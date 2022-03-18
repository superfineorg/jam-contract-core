require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const JAM_VESTING = "JamVesting";

before("Deploy JamVesting contract", async () => {
  // Prepare parameters
  const [
    deployer,
    operator,
    participant1,
    participant2,
    participant3,
    participant4,
    participant5,
    participant6,
    participant7,
    participant8
  ] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.operator = operator;
  this.participants = [
    participant1,
    participant2,
    participant3,
    participant4,
    participant5,
    participant6,
    participant7,
    participant8
  ];
  this.purchaseAmounts = [
    "20000000000000000000",
    "30000000000000000000",
    "40000000000000000000",
    "50000000000000000000",
    "35000000000000000000",
    "50000000000000000000",
    "17000000000000000000",
    "102000000000000000000"
  ];
  this.unlockDistance = 3;

  // Deploy JamVesting
  this.vestingFactory = await hre.ethers.getContractFactory(JAM_VESTING);
  this.vestingContract = await this.vestingFactory.deploy();
  await this.vestingContract.deployed();
});

describe("Test JamVesting contract", () => {
  it("Set up JamVesting contract", async () => {
    await this.vestingFactory
      .connect(this.deployer)
      .attach(this.vestingContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Check claimable amount", async () => {
    let claimableAmount = await this.vestingContract.getClaimableAmount(this.participants[2].address, 3);
    expect(claimableAmount.toString()).to.equal("0");
  });

  it("Check total claimable amount", async () => {
    let totalClaimableAmount = await this.vestingContract.getTotalClaimableAmount(this.participants[3].address);
    expect(totalClaimableAmount.toString()).to.equal("0");
  });

  it("Create 8 vesting programs", async () => {
    let start = Math.floor(Date.now() / 1000) - 5;
    let end = Math.floor(Date.now() / 1000) + 30;
    let unlockMoment = Math.floor(Date.now() / 1000) + 3;
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .createPrograms(
        Math.floor(Date.now() / 1000) - 1,
        [
          "https Seed Sale",
          "https Strategic Partners Sales",
          "https Community Growth",
          "https Ecosystem Growth",
          "https Team",
          "https Advisor",
          "https Treasury",
          "https IDO"
        ],
        [start, start, start, start, start, start, start, start],
        [end, end, end, end, end, end, end, end],
        [
          "50000000000000000000000000",
          "150000000000000000000000000",
          "230000000000000000000000000",
          "220000000000000000000000000",
          "180000000000000000000000000",
          "20000000000000000000000000",
          "120000000000000000000000000",
          "30000000000000000000000000"
        ],
        [0, 0, 3200, 2000, 0, 0, 1000, 10000],
        [
          unlockMoment,
          unlockMoment,
          unlockMoment,
          unlockMoment,
          unlockMoment,
          unlockMoment,
          unlockMoment,
          unlockMoment
        ],
        [
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance,
          this.unlockDistance
        ],
        [2000, 2000, 1360, 1600, 2000, 1250, 1125, 0]
      );
    let numPrograms = await this.vestingContract.numPrograms();
    let allPrograms = await this.vestingContract.getProgramsInfo();
    expect(numPrograms.toString()).to.equal("8");
    expect(allPrograms.length).to.equal(8);
    expect(allPrograms[6]?.metadata).to.equal("https Treasury");
  });

  it("Register 8 participants", async () => {
    for (let i = 0; i < 8; i++)
      await this.vestingFactory
        .connect(this.operator)
        .attach(this.vestingContract.address)
        .registerParticipant(this.participants[i].address, i, i !== 3, { value: this.purchaseAmounts[i] });
    let vestingAmounts = [];
    for (let i = 0; i < 8; i++)
      vestingAmounts.push(await this.vestingContract.getVestingAmount(this.participants[i].address, i));
    let allPrograms = await this.vestingContract.getProgramsInfo();
    for (let i = 0; i < 8; i++)
      expect(vestingAmounts[i].toString()).to.equal(this.purchaseAmounts[i]);
    expect(allPrograms.length).to.equal(8);
    expect(allPrograms[6].participants.length).to.equal(1);
    expect(allPrograms[6].participants[0]).to.equal(this.participants[6].address);
  });

  it("Check total vesting amount", async () => {
    let totalVestingAmount = await this.vestingContract.getTotalVestingAmount(this.participants[7].address);
    expect(totalVestingAmount.toString()).to.equal(this.purchaseAmounts[7]);
  });

  it("Update metadata", async () => {
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .updateMetadata(5, "new https Advisor");
    let allPrograms = await this.vestingContract.getProgramsInfo();
    expect(allPrograms.length).to.equal(8);
    expect(allPrograms[5].metadata).to.equal("new https Advisor");
  });

  it("Try to remove an investor from a program", async () => {
    await expect(
      this.vestingFactory
        .connect(this.operator)
        .attach(this.vestingContract.address)
        .removeParticipant(this.participants[4].address, 4)
    ).to.be.revertedWith("Cannot remove an investor");
  });

  it("Remove a non-investor participant", async () => {
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .removeParticipant(this.participants[3].address, 3);
    await expect(
      this.vestingFactory
        .connect(this.operator)
        .attach(this.vestingContract.address)
        .removeParticipant(this.participants[3].address, 3)
    ).to.be.revertedWith("Participant already removed");
    let allPrograms = await this.vestingContract.getProgramsInfo();
    expect(allPrograms[3].participants.length).to.equal(0);
  });

  it("Add more participants and remove them", async () => {
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .registerParticipant(this.participants[2].address, 6, false, { value: this.purchaseAmounts[2] });
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .registerParticipant(this.participants[5].address, 6, false, { value: this.purchaseAmounts[5] });
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .registerParticipant(this.participants[2].address, 6, false, { value: this.purchaseAmounts[2] });
    let allPrograms = await this.vestingContract.getProgramsInfo();
    expect(allPrograms[6].participants.length).to.equal(3);
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .removeParticipant(this.participants[2].address, 6);
    allPrograms = await this.vestingContract.getProgramsInfo();
    expect(allPrograms[6].participants.length).to.equal(2);
    expect(allPrograms[6].participants[1]).to.equal(this.participants[5].address);
  });

  it("Claim vesting tokens", async () => {
    await sleep(12000);
    await this.vestingFactory
      .connect(this.participants[0])
      .attach(this.vestingContract.address)
      .claimTokens(0);
    let claimedAmount = await this.vestingContract.getClaimedAmount(this.participants[0].address, 0);
    expect(claimedAmount.toString()).not.to.equal("0");
  });

  it("Claim all vesting tokens", async () => {
    await this.vestingFactory
      .connect(this.participants[2])
      .attach(this.vestingContract.address)
      .claimAllTokens();
    let claimedAmount = await this.vestingContract.getTotalClaimedAmount(this.participants[2].address);
    expect(claimedAmount.toString()).to.not.equal("0");
  });
});

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

// Run: npx hardhat test test/test-vesting-program.js
require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const VESTING = "Vesting";

before("Deploy Vesting contract", async () => {
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
  this.block = 3;

  // Deploy Vesting
  this.vestingFactory = await hre.ethers.getContractFactory(VESTING);
  this.vestingContract = await this.vestingFactory.deploy(this.block * 2);
  await this.vestingContract.deployed();
});

describe("Test Vesting contract", () => {
  it("Set up Vesting contract", async () => {
    await this.vestingFactory
      .connect(this.deployer)
      .attach(this.vestingContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Set block length", async () => {
    await this.vestingFactory
      .connect(this.operator)
      .attach(this.vestingContract.address)
      .setBlockLength(this.block);
  });

  it("Check claimable amount", async () => {
    let claimableAmount = await this.vestingContract.getClaimableAmount(this.participants[2].address);
    expect(claimableAmount.toString()).to.equal("0");
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
          "Seed Sale",
          "Strategic Partners Sales",
          "Community Growth",
          "Ecosystem Growth",
          "Team",
          "Advisor",
          "Treasury",
          "IDO"
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
        [2000, 2000, 1360, 1600, 2000, 1250, 1125, 0]
      );
    let numPrograms = await this.vestingContract.allProgramsLength();
    let treasuryInfo = await this.vestingContract.allPrograms(6);
    expect(numPrograms.toString()).to.equal("8");
    expect(treasuryInfo?.name).to.equal("Treasury");
  });

  it("Register 8 participants", async () => {
    for (let i = 0; i < 7; i++)
      await this.vestingFactory
        .connect(this.operator)
        .attach(this.vestingContract.address)
        .registerParticipant(this.participants[i].address, i, { value: this.purchaseAmounts[i] });
    let vestingAmounts = [];
    for (let i = 0; i < 7; i++)
      vestingAmounts.push(await this.vestingContract.getVestingAmount(this.participants[i].address, i));
    for (let i = 0; i < 7; i++)
      expect(vestingAmounts[i].toString()).to.equal(this.purchaseAmounts[i]);
  });

  it("Claim vesting tokens", async () => {
    await sleep(12000);
    await this.vestingFactory
      .connect(this.participants[0])
      .attach(this.vestingContract.address)
      .claimTokens();
    let claimedAmount = await this.vestingContract.getClaimedAmount(this.participants[0].address);
    expect(claimedAmount.toString()).not.to.equal("0");
  });
});

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

// Run: npx hardhat test test/test-vesting-program.js
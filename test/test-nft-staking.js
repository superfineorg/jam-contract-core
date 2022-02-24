require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const NFT_STAKING = "NFTStaking";
const SIMPLE_NFT_721 = "SimpleERC721";
const SIMPLE_NFT_1155 = "SimpleERC1155";

before("Deploy NFTStaking contract and a simple NFT contracts", async () => {
  // Prepare parameters
  const [deployer, operator, participant] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.operator = operator;
  this.participant = participant;

  // Deploy NFTStaking contract
  this.nftStakingFactory = await hre.ethers.getContractFactory(NFT_STAKING);
  this.nftStakingContract = await this.nftStakingFactory.deploy(4000, 150000);
  await this.nftStakingContract.deployed();

  // Deploy a simple NFT 721 contract
  this.nft721Factory = await hre.ethers.getContractFactory(SIMPLE_NFT_721);
  this.nft721Contract = await this.nft721Factory.deploy(
    deployer.address,
    "Gamejam Awesome NFT",
    "JamNFT",
    "https://xxx.com/"
  );
  await this.nft721Contract.deployed();

  // Deploy a simple NFT 1155 contract
  this.nft1155Factory = await hre.ethers.getContractFactory(SIMPLE_NFT_1155);
  this.nft1155Contract = await this.nft1155Factory.deploy();
  await this.nft1155Contract.deployed();
});

describe("Test NFT staking program", () => {
  it("Setup operator role", async () => {
    await this.nftStakingFactory
      .connect(this.deployer)
      .attach(this.nftStakingContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Transfer some initial fund to the contract", async () => {
    await this.deployer.sendTransaction({
      to: this.nftStakingContract.address,
      value: hre.ethers.utils.parseEther("1.0")
    });
  });

  it("Set lock duration", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .setLockDuration(5000);
    let lockDuration = await this.nftStakingContract.lockDuration();
    expect(lockDuration.toString()).to.equal("5000");
  });

  it("Set total reward per day", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .setRewardPerDay(100000);
    let rewardPerDay = await this.nftStakingContract.rewardPerDay();
    expect(rewardPerDay.toString()).to.equal("100000");
  });

  it("Whitelist the previously deployed NFT721 token", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .whitelistNFT(
        [this.nft721Contract.address, this.nft1155Contract.address],
        [0, 1],
        [true, true]
      );
    let nft721Whitelist = await this.nftStakingContract.nftWhitelist(0);
    let nft1155Whitelist = await this.nftStakingContract.nftWhitelist(1);
    expect(nft721Whitelist).to.equal(this.nft721Contract.address);
    expect(nft1155Whitelist).to.equal(this.nft1155Contract.address);
  });

  it("Remove this NFT721 token from the whitelist", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .whitelistNFT([this.nft721Contract.address], [0], [false]);
    let nft1155Whitelist = await this.nftStakingContract.nftWhitelist(0);
    expect(nft1155Whitelist).to.equal(this.nft1155Contract.address);
    await expect(this.nftStakingContract.nftWhitelist(1)).to.be.reverted;
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .whitelistNFT([this.nft721Contract.address], [0], [true]);    // Whitelist it again for testing later
  });

  it("Stake a new NFT721", async () => {
    await this.nft721Factory
      .connect(this.deployer)
      .attach(this.nft721Contract.address)
      .awardItem(this.participant.address);
    await this.nft721Factory
      .connect(this.participant)
      .attach(this.nft721Contract.address)
      .approve(this.nftStakingContract.address, 1);
    let owner = await this.nft721Contract.ownerOf(1);
    let spender = await this.nft721Contract.getApproved(1);
    expect(owner).to.equal(this.participant.address);
    expect(spender).to.equal(this.nftStakingContract.address);
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .stake(this.nft721Contract.address, 1, 1);
    let numStakedNFTs = await this.nftStakingContract.getNumStakedNFTs(this.participant.address);
    let stakedNFTTokenIds = await this.nftStakingContract.getStakedNFTTokenIds(
      this.participant.address,
      this.nft721Contract.address
    );
    let stakedQuantity = await this.nftStakingContract.getStakedQuantity(
      this.participant.address,
      this.nft721Contract.address,
      1
    );
    let stakingMoment = await this.nftStakingContract.getStakingMoment(
      this.participant.address,
      this.nft721Contract.address,
      1
    );
    expect(numStakedNFTs.toString()).to.equal("1");
    expect(stakedNFTTokenIds.length).to.equal(1);
    expect(stakedNFTTokenIds[0].toString()).to.equal("1");
    expect(stakedQuantity.toString()).to.equal("1");
    expect(stakingMoment.toString()).not.to.equal("0");
  });

  it("Claim the reward", async () => {
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .claimReward();
  });

  it("Unstake the staked NFT above", async () => {
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .unstake(this.nft721Contract.address, 1, 1);
    let numStakedNFTs = await this.nftStakingContract.getNumStakedNFTs(this.participant.address);
    let stakedNFTTokenIds = await this.nftStakingContract.getStakedNFTTokenIds(
      this.participant.address,
      this.nft721Contract.address
    );
    let stakedQuantity = await this.nftStakingContract.getStakedQuantity(
      this.participant.address,
      this.nft721Contract.address,
      1
    );
    let stakingMoment = await this.nftStakingContract.getStakingMoment(
      this.participant.address,
      this.nft721Contract.address,
      1
    );
    expect(numStakedNFTs.toString()).to.equal("0");
    expect(stakedNFTTokenIds.length).to.equal(0);
    expect(stakedQuantity.toString()).to.equal("0");
    expect(stakingMoment.toString()).to.equal("0");
  });

  it("Stake a new NFT 1155", async () => {
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .mintTo(this.participant.address, 20);
    await this.nft1155Factory
      .connect(this.participant)
      .attach(this.nft1155Contract.address)
      .setApprovalForAll(this.nftStakingContract.address, true);
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .stake(this.nft1155Contract.address, 1, 15);
    let numStakedNFTs = await this.nftStakingContract.getNumStakedNFTs(this.participant.address);
    let stakedNFTTokenIds = await this.nftStakingContract.getStakedNFTTokenIds(
      this.participant.address,
      this.nft1155Contract.address
    );
    let stakedQuantity = await this.nftStakingContract.getStakedQuantity(
      this.participant.address,
      this.nft1155Contract.address,
      1
    );
    let stakingMoment = await this.nftStakingContract.getStakingMoment(
      this.participant.address,
      this.nft1155Contract.address,
      1
    );
    expect(numStakedNFTs.toString()).to.equal("15");
    expect(stakedNFTTokenIds.length).to.equal(1);
    expect(stakedNFTTokenIds[0].toString()).to.equal("1");
    expect(stakedQuantity.toString()).to.equal("15");
    expect(stakingMoment.toString()).not.to.equal("0");
  });

  it("Unstake the staked NFTs above", async () => {
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .unstake(this.nft1155Contract.address, 1, 8);
    let numStakedNFTs = await this.nftStakingContract.getNumStakedNFTs(this.participant.address);
    let stakedNFTTokenIds = await this.nftStakingContract.getStakedNFTTokenIds(
      this.participant.address,
      this.nft1155Contract.address
    );
    let stakedQuantity = await this.nftStakingContract.getStakedQuantity(
      this.participant.address,
      this.nft1155Contract.address,
      1
    );
    let stakingMoment = await this.nftStakingContract.getStakingMoment(
      this.participant.address,
      this.nft1155Contract.address,
      1
    );
    expect(numStakedNFTs.toString()).to.equal("7");
    expect(stakedNFTTokenIds.length).to.equal(1);
    expect(stakedQuantity.toString()).to.equal("7");
    expect(stakingMoment.toString()).not.to.equal("0");
  });

  it("Pause the contract", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .pause();
    let isPaused = await this.nftStakingContract.paused();
    expect(isPaused).to.equal(true);
  });

  it("Emergency withdraw", async () => {
    await this.nftStakingFactory
      .connect(this.deployer)
      .attach(this.nftStakingContract.address)
      .emergencyWithdraw(this.deployer.address);
  });

  it("Unpause the contract", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .unpause();
    let isPaused = await this.nftStakingContract.paused();
    expect(isPaused).to.equal(false);
  });
});

// Run: npx hardhat test test/test-nft-staking.js
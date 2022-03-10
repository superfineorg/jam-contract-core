require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const NFT_STAKING = "NFTStaking";
const SIMPLE_ERC_20 = "SimpleERC20";
const SIMPLE_ERC_721 = "SimpleERC721";
const SIMPLE_ERC_1155 = "SimpleERC1155";
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

before("Deploy NFTStaking contract and simple NFT contracts", async () => {
  // Prepare parameters
  const [deployer, operator, participant] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.operator = operator;
  this.participant = participant;

  // Deploy NFTStaking contract
  this.nftStakingFactory = await hre.ethers.getContractFactory(NFT_STAKING);
  this.nftStakingContract = await this.nftStakingFactory.deploy(4, 150000, ZERO_ADDRESS);
  await this.nftStakingContract.deployed();

  // Deploy a simple ERC20 token to award
  this.erc20Factory = await hre.ethers.getContractFactory(SIMPLE_ERC_20);
  this.erc20Contract = await this.erc20Factory.deploy(
    deployer.address,
    "10000000000000000000000000",
    "10000000000000000000000000",
    18,
    "Reward Token",
    "RT"
  );
  await this.erc20Contract.deployed();

  // Deploy a simple NFT 721 contract
  this.nft721Factory = await hre.ethers.getContractFactory(SIMPLE_ERC_721);
  this.nft721Contract = await this.nft721Factory.deploy(
    deployer.address,
    "Gamejam Awesome NFT",
    "JamNFT",
    "https://xxx.com/"
  );
  await this.nft721Contract.deployed();

  // Deploy a simple NFT 1155 contract
  this.nft1155Factory = await hre.ethers.getContractFactory(SIMPLE_ERC_1155);
  this.nft1155Contract = await this.nft1155Factory.deploy();
  await this.nft1155Contract.deployed();
});

describe("Test NFT staking program", () => {
  it("Check initial values", async () => {
    let lockDuration = await this.nftStakingContract.lockDuration();
    let rewardPerDay = await this.nftStakingContract.rewardPerDay();
    let rewardToken = await this.nftStakingContract.rewardToken();
    expect(lockDuration.toString()).to.equal("4");
    expect(rewardPerDay.toString()).to.equal("150000");
    expect(rewardToken).to.equal(ZERO_ADDRESS);
  });

  it("Setup operator role", async () => {
    await this.nftStakingFactory
      .connect(this.deployer)
      .attach(this.nftStakingContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Transfer some initial fund to the contract", async () => {
    await this.deployer.sendTransaction({
      to: this.nftStakingContract.address,
      value: hre.ethers.utils.parseEther("80")
    });
    await this.erc20Factory
      .connect(this.deployer)
      .attach(this.erc20Contract.address)
      .transfer(this.nftStakingContract.address, "9000000000000000000000000");
  });

  it("Set lock duration", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .setLockDuration(8);
    let lockDuration = await this.nftStakingContract.lockDuration();
    expect(lockDuration.toString()).to.equal("8");
  });

  it("Set total reward per day", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .setRewardPerDay(100000);
    let rewardPerDay = await this.nftStakingContract.rewardPerDay();
    expect(rewardPerDay.toString()).to.equal("100000");
  });

  it("Whitelist the previously deployed ERC721 and ERC1155 tokens", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .whitelistNFT(
        [
          this.nft721Contract.address,
          this.nft1155Contract.address,
          this.nft1155Contract.address,
          this.nft1155Contract.address
        ],
        [0, 1, 1, 1],
        [0, 1, 3, 4]
      );
  });

  it("Whitelist more NFT1155 token ids", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .whitelistNFT([this.nft1155Contract.address], [1], [6]);
  });

  it("Mint some ERC721 and ERC1155 NFTs to a participant", async () => {
    await this.nft721Factory
      .connect(this.deployer)
      .attach(this.nft721Contract.address)
      .awardItem(this.participant.address);   // id = 1
    await this.nft721Factory
      .connect(this.deployer)
      .attach(this.nft721Contract.address)
      .awardItem(this.participant.address);   // id = 2
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .mintTo(this.participant.address, 20);  // id = 1
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .mintTo(this.participant.address, 30);  // id = 2
    let unstakedNFTs = await this.nftStakingContract.getUnstakedNFTs(this.participant.address);
    expect(unstakedNFTs.length).to.equal(3);
    expect(unstakedNFTs[0].nftType).to.equal(0);
    expect(unstakedNFTs[0].nftAddress).to.equal(this.nft721Contract.address);
    expect(unstakedNFTs[0].tokenId.toString()).to.equal("1");
    expect(unstakedNFTs[0].quantity.toString()).to.equal("1");
    expect(unstakedNFTs[0].stakingMoment.toString()).to.equal("0");
    expect(unstakedNFTs[1].nftType).to.equal(0);
    expect(unstakedNFTs[1].nftAddress).to.equal(this.nft721Contract.address);
    expect(unstakedNFTs[1].tokenId.toString()).to.equal("2");
    expect(unstakedNFTs[1].quantity.toString()).to.equal("1");
    expect(unstakedNFTs[1].stakingMoment.toString()).to.equal("0");
    expect(unstakedNFTs[2].nftType).to.equal(1);
    expect(unstakedNFTs[2].nftAddress).to.equal(this.nft1155Contract.address);
    expect(unstakedNFTs[2].tokenId.toString()).to.equal("1");
    expect(unstakedNFTs[2].quantity.toString()).to.equal("20");
    expect(unstakedNFTs[2].stakingMoment.toString()).to.equal("0");
  });

  it("Stake a new ERC721 NFT", async () => {
    await this.nft721Factory
      .connect(this.participant)
      .attach(this.nft721Contract.address)
      .approve(this.nftStakingContract.address, 1);
    let owner = await this.nft721Contract.ownerOf(1);
    let spender = await this.nft721Contract.getApproved(1);
    expect(owner).to.equal(this.participant.address);
    expect(spender).to.equal(this.nftStakingContract.address);
    await expect(
      this.nftStakingFactory
        .connect(this.participant)
        .attach(this.nftStakingContract.address)
        .stake([this.nft721Contract.address], [1], [3])
    ).to.be.revertedWith("NFTStaking: cannot stake more than 1 ERC721 NFT at a time");
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .stake([this.nft721Contract.address], [1], [1]);
    let stakedNFTs = await this.nftStakingContract.getStakedNFTs(this.participant.address);
    expect(stakedNFTs.length).to.equal(1);
    expect(stakedNFTs[0].nftType).to.equal(0);
    expect(stakedNFTs[0].nftAddress).to.equal(this.nft721Contract.address);
    expect(stakedNFTs[0].tokenId.toString()).to.equal("1");
    expect(stakedNFTs[0].quantity.toString()).to.equal("1");
    expect(stakedNFTs[0].stakingMoment.toString()).not.to.equal("0");
  });

  it("Claim the reward", async () => {
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .claimReward();
  });

  it("Stake some new ERC1155 NFTs", async () => {
    await this.nft1155Factory
      .connect(this.participant)
      .attach(this.nft1155Contract.address)
      .setApprovalForAll(this.nftStakingContract.address, true);
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .stake([this.nft1155Contract.address], [1], [15]);
    let stakedNFTs = await this.nftStakingContract.getStakedNFTs(this.participant.address);
    expect(stakedNFTs.length).to.equal(2);
    expect(stakedNFTs[1].nftType).to.equal(1);
    expect(stakedNFTs[1].nftAddress).to.equal(this.nft1155Contract.address);
    expect(stakedNFTs[1].tokenId.toString()).to.equal("1");
    expect(stakedNFTs[1].quantity.toString()).to.equal("15");
    expect(stakedNFTs[1].stakingMoment.toString()).not.to.equal("0");
  });

  it("Stake some non-supported ERC1155 NFTs", async () => {
    await expect(
      this.nftStakingFactory
        .connect(this.participant)
        .attach(this.nftStakingContract.address)
        .stake([this.nft1155Contract.address], [2], [25])
    ).to.be.revertedWith("NFTStaking: this NFT is not supported");
  });

  it("Unstake the staked ERC721 NFT above", async () => {
    await expect(
      this.nftStakingFactory
        .connect(this.deployer)
        .attach(this.nftStakingContract.address)
        .unstake([this.nft721Contract.address], [1], [1])
    ).to.be.revertedWith("NFTStaking: only owner can unstake");
    await expect(
      this.nftStakingFactory
        .connect(this.participant)
        .attach(this.nftStakingContract.address)
        .unstake([this.nft721Contract.address], [1], [1])
    ).to.be.revertedWith("NFTStaking: NFT not unlocked yet");
    await sleep(9000);
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .unstake([this.nft721Contract.address], [1], [1]);
    let stakedNFTs = await this.nftStakingContract.getStakedNFTs(this.participant.address);
    expect(stakedNFTs.length).to.equal(1);
    expect(stakedNFTs[0].nftType).to.equal(1);
    expect(stakedNFTs[0].nftAddress).to.equal(this.nft1155Contract.address);
    expect(stakedNFTs[0].tokenId.toString()).to.equal("1");
    expect(stakedNFTs[0].quantity.toString()).to.equal("15");
    expect(stakedNFTs[0].stakingMoment.toString()).not.to.equal("0");
  });

  it("Admin changes the reward token", async () => {
    await this.nftStakingFactory
      .connect(this.operator)
      .attach(this.nftStakingContract.address)
      .setRewardToken(this.erc20Contract.address);
    let rewardToken = await this.nftStakingContract.rewardToken();
    expect(rewardToken).to.equal(this.erc20Contract.address);
  });

  it("Unstake the staked ERC1155 NFTs above", async () => {
    await expect(
      this.nftStakingFactory
        .connect(this.deployer)
        .attach(this.nftStakingContract.address)
        .unstake([this.nft1155Contract.address], [1], [8])
    ).to.be.revertedWith("NFTStaking: only owner can unstake");
    await expect(
      this.nftStakingFactory
        .connect(this.participant)
        .attach(this.nftStakingContract.address)
        .unstake([this.nft1155Contract.address], [1], [16])
    ).to.be.revertedWith("NFTStaking: not enough NFTs to unstake");
    await this.nftStakingFactory
      .connect(this.participant)
      .attach(this.nftStakingContract.address)
      .unstake([this.nft1155Contract.address], [1], [8]);
    let stakedNFTs = await this.nftStakingContract.getStakedNFTs(this.participant.address);
    expect(stakedNFTs.length).to.equal(1);
    expect(stakedNFTs[0].nftType).to.equal(1);
    expect(stakedNFTs[0].nftAddress).to.equal(this.nft1155Contract.address);
    expect(stakedNFTs[0].tokenId.toString()).to.equal("1");
    expect(stakedNFTs[0].quantity.toString()).to.equal("7");
    expect(stakedNFTs[0].stakingMoment.toString()).not.to.equal("0");
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

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

// Run: npx hardhat test test/test-nft-staking.js
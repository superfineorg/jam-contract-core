require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const { soliditySha3 } = require('web3-utils');
const ERC20 = "GameToken";
const ERC721 = "JamNFT721Airdrop";
const ERC1155 = "JamNFT1155Airdrop";
const AIRDROP = "PlaylinkAirdrop";

before("Deploy contracts", async () => {
  // Prepare parameters
  const [deployer, operator, creator, user1, user2] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.operator = operator;
  this.creator = creator;
  this.user1 = user1;
  this.user2 = user2;
  this.MAX_BATCH_SIZE = 3;
  this.FEE_PER_BATCH = hre.ethers.utils.parseEther("0.01");

  // Deploy ERC20 contract
  this.erc20Factory = await hre.ethers.getContractFactory(ERC20);
  this.erc20Contract = await this.erc20Factory.deploy();
  await this.erc20Contract.deployed();

  // Deploy ERC721 contract
  this.erc721Factory = await hre.ethers.getContractFactory(ERC721);
  this.erc721Contract = await this.erc721Factory.deploy("JamNFT721", "JNFT", "https://gamejam.com/nft721/");
  await this.erc721Contract.deployed();

  // Deploy ERC1155 contract
  this.erc1155Factory = await hre.ethers.getContractFactory(ERC1155);
  this.erc1155Contract = await this.erc1155Factory.deploy("https://gamejam.com/nft1155/");
  await this.erc1155Contract.deployed();

  // Deploy JamAidrop contract
  this.airdropFactory = await hre.ethers.getContractFactory(AIRDROP);
  this.airdropContract = await this.airdropFactory.deploy(this.MAX_BATCH_SIZE, this.FEE_PER_BATCH);
  await this.airdropContract.deployed();
});

describe("Test PlaylinkAirdrop", () => {
  it("Some initial preparation", async () => {
    // Mint some initial ERC20 tokens
    let minterRole = await this.erc20Contract.MINTER_ROLE();
    await this.erc20Factory
      .connect(this.deployer)
      .attach(this.erc20Contract.address)
      .grantRole(minterRole, this.deployer.address);
    await this.erc20Factory
      .connect(this.deployer)
      .attach(this.erc20Contract.address)
      .mint(this.creator.address, hre.ethers.utils.parseEther("100"));
    let erc20Balance = await this.erc20Contract.balanceOf(this.creator.address);
    expect(erc20Balance.toString()).to.equal(hre.ethers.utils.parseEther("100").toString());

    // Mint some initial ERC721 tokens
    await this.erc721Factory
      .connect(this.deployer)
      .attach(this.erc721Contract.address)
      .grantRole(minterRole, this.deployer.address);
    for (let i = 0; i < 100; i++)
      await this.erc721Factory
        .connect(this.deployer)
        .attach(this.erc721Contract.address)
        .mint(this.creator.address);
    let erc721Balance = await this.erc721Contract.balanceOf(this.creator.address);
    expect(erc721Balance.toString()).to.equal("100");

    // Mint some initial ERC1155 tokens
    await this.erc1155Factory
      .connect(this.deployer)
      .attach(this.erc1155Contract.address)
      .grantRole(minterRole, this.deployer.address);
    await this.erc1155Factory
      .connect(this.deployer)
      .attach(this.erc1155Contract.address)
      .mint(this.creator.address, 999, 100, soliditySha3("Mint ERC1155 tokens"));
    let erc1155Balance = await this.erc1155Contract.balanceOf(this.creator.address, 999);
    expect(erc1155Balance.toString()).to.equal("100");
  });

  it("Set airdrop operator", async () => {
    await this.airdropFactory
      .connect(this.deployer)
      .attach(this.airdropContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Check airdrop fee calculation", async () => {
    let airdropFee = await this.airdropContract.estimateAirdropFee(10);
    expect(airdropFee.toString()).to.equal(hre.ethers.utils.parseEther("0.04"));
  });

  it("The creator creates a new campaign and approve assets", async () => {
    let currentTime = await now();
    await this.airdropFactory
      .connect(this.creator)
      .attach(this.airdropContract.address)
      .createAirdropCampaign(
        "01BX5ZZKBKACTAV9WEVGEMMVRY",
        [
          {
            assetType: 0,
            assetAddress: this.erc20Contract.address,
            assetId: 0,
            availableAmount: hre.ethers.utils.parseEther("40").toString()
          },
          {
            assetType: 1,
            assetAddress: this.erc721Contract.address,
            assetId: 7,
            availableAmount: 1
          },
          {
            assetType: 2,
            assetAddress: this.erc1155Contract.address,
            assetId: 999,
            availableAmount: 70
          },
          {
            assetType: 0,
            assetAddress: this.erc20Contract.address,
            assetId: 0,
            availableAmount: hre.ethers.utils.parseEther("45").toString()
          }
        ],
        currentTime + 30 * 60, // This campaign will start 30m later
        { value: hre.ethers.utils.parseEther("0.024") }
      );
    await this.erc20Factory
      .connect(this.creator)
      .attach(this.erc20Contract.address)
      .approve(this.airdropContract.address, hre.ethers.utils.parseEther("85"));
    await this.erc721Factory
      .connect(this.creator)
      .attach(this.erc721Contract.address)
      .approve(this.airdropContract.address, 7);
    await this.erc1155Factory
      .connect(this.creator)
      .attach(this.erc1155Contract.address)
      .setApprovalForAll(this.airdropContract.address, true);
    let campaign = await this.airdropContract.getCampaignById("01BX5ZZKBKACTAV9WEVGEMMVRY");
    expect(campaign.campaignId).to.equal("01BX5ZZKBKACTAV9WEVGEMMVRY");
    expect(campaign.creator).to.equal(this.creator.address);
    expect(campaign.assets.length).to.equal(4);
    expect(campaign.maxBatchSize.toString()).to.equal("3");
    expect(campaign.startingTime.toString()).to.equal((currentTime + 30 * 60).toString());
    expect(campaign.totalAvailableAssets.toString()).to.equal("85000000000000000071");
    expect(campaign.airdropFee.toString()).to.equal(hre.ethers.utils.parseEther("0.02"));
  });

  it("The campaign creator update assets", async () => {
    let currentTime = await now();
    await this.airdropFactory
      .connect(this.creator)
      .attach(this.airdropContract.address)
      .updateCampaign(
        "01BX5ZZKBKACTAV9WEVGEMMVRY",
        [
          {
            assetType: 0,
            assetAddress: this.erc20Contract.address,
            assetId: 0,
            availableAmount: hre.ethers.utils.parseEther("40").toString()
          },
          {
            assetType: 1,
            assetAddress: this.erc721Contract.address,
            assetId: 7,
            availableAmount: 1
          },
          {
            assetType: 2,
            assetAddress: this.erc1155Contract.address,
            assetId: 999,
            availableAmount: 65
          },
          {
            assetType: 0,
            assetAddress: this.erc20Contract.address,
            assetId: 0,
            availableAmount: hre.ethers.utils.parseEther("45").toString()
          },
          {
            assetType: 1,
            assetAddress: this.erc721Contract.address,
            assetId: 15,
            availableAmount: 1
          }
        ],
        currentTime + 15 * 60
      );
    let campaign = await this.airdropContract.getCampaignById("01BX5ZZKBKACTAV9WEVGEMMVRY");
    expect(campaign.campaignId).to.equal("01BX5ZZKBKACTAV9WEVGEMMVRY");
    expect(campaign.creator).to.equal(this.creator.address);
    expect(campaign.assets.length).to.equal(5);
    expect(campaign.maxBatchSize.toString()).to.equal("3");
    expect(campaign.totalAvailableAssets.toString()).to.equal("85000000000000000067");
    expect(campaign.airdropFee.toString()).to.equal(hre.ethers.utils.parseEther("0.02"));
  });

  it("The operator starts to airdrop", async () => {
    await expect(
      this.airdropFactory
        .connect(this.operator)
        .attach(this.airdropContract.address)
        .airdrop("01BX5ZZKBKACTAV9WEVGEMMVRY", [0], [this.user2.address])
    ).to.be.revertedWith("PlaylinkAirdrop: campaign not start yet");

    // 16 minutes later...
    await hre.network.provider.request({ method: "evm_increaseTime", params: [16 * 60] });
    await hre.network.provider.request({ method: "evm_mine", params: [] });

    await this.airdropFactory
      .connect(this.operator)
      .attach(this.airdropContract.address)
      .airdrop("01BX5ZZKBKACTAV9WEVGEMMVRY", [0, 3, 1], [this.user1.address, this.user2.address, this.user1.address]);
    await this.airdropFactory
      .connect(this.operator)
      .attach(this.airdropContract.address)
      .airdrop("01BX5ZZKBKACTAV9WEVGEMMVRY", [2], [this.user1.address]);
    let erc20Balance1 = await this.erc20Contract.balanceOf(this.user1.address);
    let erc20Balance2 = await this.erc20Contract.balanceOf(this.user2.address);
    let erc721Owner = await this.erc721Contract.ownerOf(7);
    let erc1155Balance = await this.erc1155Contract.balanceOf(this.user1.address, 999);
    expect(erc20Balance1.toString()).to.equal(hre.ethers.utils.parseEther("40"));
    expect(erc20Balance2.toString()).to.equal(hre.ethers.utils.parseEther("45"));
    expect(erc721Owner).to.equal(this.user1.address);
    expect(erc1155Balance.toString()).to.equal("65");
  });

  it("The owner withdraw all campaign fees", async () => {
    await this.airdropFactory
      .connect(this.deployer)
      .attach(this.airdropContract.address)
      .withdrawAirdropFee(this.deployer.address);
  });
});

let now = async () => {
  let blockNumber = await hre.network.provider.request({ method: "eth_blockNumber", params: [] });
  let block = await hre.network.provider.request({ method: "eth_getBlockByNumber", params: [blockNumber, false] });
  return parseInt(block.timestamp || 0, 16);
};
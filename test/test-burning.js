require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const { soliditySha3 } = require('web3-utils');

const GAME_NFT_721 = "GameNFT721";
const GAME_NFT_1155 = "GameNFT1155";
const GAME_NFT_BURNING = "GameNFTBurning";

before("Deploy GameNFT721, GameNFT1155 and GameNFTBurning contracts", async () => {
  // Prepare parameters
  const [deployer, minter, burningOperator, user] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.minter = minter;
  this.burningOperator = burningOperator;
  this.user = user;

  // Deploy GameNFT721 contract
  this.nft721Factory = await hre.ethers.getContractFactory(GAME_NFT_721);
  this.nft721Contract = await this.nft721Factory.deploy(
    "Gamejam NFT 721",
    "GNFT721",
    "https://gamejam.com/nft721/"
  );
  await this.nft721Contract.deployed();

  // Deploy GameNFT1155 contract
  this.nft1155Factory = await hre.ethers.getContractFactory(GAME_NFT_1155);
  this.nft1155Contract = await this.nft1155Factory.deploy("https://gamejam.com/nft1155/");
  await this.nft1155Contract.deployed();

  // Deploy GameNFTBurning contract
  this.nftBurningFactory = await hre.ethers.getContractFactory(GAME_NFT_BURNING);
  this.nftBurningContract = await this.nftBurningFactory.deploy();
  await this.nftBurningContract.deployed();
});

describe("Test burning contract", () => {
  it("Set up minter role", async () => {
    let minterRole = await this.nft721Contract.MINTER_ROLE();
    await this.nft721Factory
      .connect(this.deployer)
      .attach(this.nft721Contract.address)
      .grantRole(minterRole, this.minter.address);
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole721 = await this.nft721Contract.hasRole(minterRole, this.minter.address);
    let checkRole1155 = await this.nft1155Contract.hasRole(minterRole, this.minter.address);
    expect(checkRole721).to.equal(true);
    expect(checkRole1155).to.equal(true);
  });

  it("Set up burner role", async () => {
    await this.nftBurningFactory
      .connect(this.deployer)
      .attach(this.nftBurningContract.address)
      .setOperators([this.burningOperator.address], [true]);
  });

  it("Mint an NFT721 to user", async () => {
    await this.nft721Factory
      .connect(this.minter)
      .attach(this.nft721Contract.address)
      .mint(this.user.address);
    let owner = await this.nft721Contract.ownerOf(0);
    let tokenUri = await this.nft721Contract.tokenURI(0);
    expect(owner).to.equal(this.user.address);
    expect(tokenUri).to.equal("https://gamejam.com/nft721/0.json");
  });

  it("Mint some NFT1155s to user", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 0, 12, soliditySha3("Mint #0"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 0);
    let tokenUri = await this.nft1155Contract.uri(0);
    expect(balance.toString()).to.equal("12");
    expect(tokenUri).to.equal("https://gamejam.com/nft1155/0.json");
  });

  it("Mint more NFT1155s to user", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 1, 30, soliditySha3("Mint #1"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 1);
    let tokenUri = await this.nft1155Contract.uri(1);
    expect(balance.toString()).to.equal("30");
    expect(tokenUri).to.equal("https://gamejam.com/nft1155/1.json");
  });

  it("Change the ERC721 base token URI", async () => {
    await expect(
      this.nft721Factory
        .connect(this.minter)
        .attach(this.nft721Contract.address)
        .setBaseTokenURI("https://example.com/")
    ).to.be.revertedWith("GameNFT721: must have admin role to set");
    await this.nft721Factory
      .connect(this.deployer)
      .attach(this.nft721Contract.address)
      .setBaseTokenURI("https://gamejam.com/nft721/example/");
    let tokenUri = await this.nft721Contract.tokenURI(0);
    expect(tokenUri).to.equal("https://gamejam.com/nft721/example/0.json");
  });

  it("User approves his ERC721 NFT to the burning contract and the system burns it", async () => {
    await this.nft721Factory
      .connect(this.user)
      .attach(this.nft721Contract.address)
      .approve(this.nftBurningContract.address, 0);
    await this.nftBurningFactory
      .connect(this.burningOperator)
      .attach(this.nftBurningContract.address)
      .burnErc721IntoGames([this.nft721Contract.address], [0]);
    await expect(this.nft721Contract.ownerOf(0)).to.be.revertedWith("ERC721: owner query for nonexistent token");
  });

  it("User approves his ERC1155 NFTs to the burning contract and the system burns them", async () => {
    await this.nft1155Factory
      .connect(this.user)
      .attach(this.nft1155Contract.address)
      .setApprovalForAll(this.nftBurningContract.address, true);
    await this.nftBurningFactory
      .connect(this.burningOperator)
      .attach(this.nftBurningContract.address)
      .burnErc1155IntoGames(
        [this.user.address],
        [this.nft1155Contract.address],
        [[0, 1]],
        [[9, 25]]
      );
    let remainingQuantity0 = await this.nft1155Contract.balanceOf(this.user.address, 0);
    let remainingQuantity1 = await this.nft1155Contract.balanceOf(this.user.address, 1);
    expect(remainingQuantity0.toString()).to.equal("3");
    expect(remainingQuantity1.toString()).to.equal("5");
  });
});
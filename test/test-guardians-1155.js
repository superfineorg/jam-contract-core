require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { soliditySha3 } = require("web3-utils");
const { expect } = require("chai");
const GAMEJAM_NFT_1155 = "GamejamNFT1155";

before("Deploy NFTStaking contract and simple NFT contracts", async () => {
  // Prepare parameters
  const [deployer, minter, user] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.minter = minter;
  this.user = user;

  // Deploy the GamejamNFT1155 contract
  this.nft1155Factory = await hre.ethers.getContractFactory(GAMEJAM_NFT_1155);
  this.nft1155Contract = await this.nft1155Factory.deploy("https://gamejam.com/nft1155/");
  await this.nft1155Contract.deployed();
});

describe("Test NFT staking program", () => {
  it("Grant minter role", async () => {
    let minterRole = await this.nft1155Contract.MINTER_ROLE();
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole = await this.nft1155Contract.hasRole(minterRole, this.minter.address);
    expect(checkRole).to.equal(true);
  });

  it("Mint some first NFTs", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 7, 20, soliditySha3("First minting time"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 7);
    let uri = await this.nft1155Contract.uri(7);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(balance.toString()).to.equal("20");
    expect(uri).to.equal("https://gamejam.com/nft1155/7.json");
    expect(ownedNFTs.length).to.equal(1);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[0].quantity.toString()).to.equal("20");
    expect(ownedNFTs[0].uri).to.equal("https://gamejam.com/nft1155/7.json");
  });

  it("Mint some more NFTs", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 15, 6, soliditySha3("Second minting time"));
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 98, 41, soliditySha3("Third minting time"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 15);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(balance.toString()).to.equal("6");
    expect(ownedNFTs.length).to.equal(3);
    expect(ownedNFTs[1].tokenId.toString()).to.equal("15");
    expect(ownedNFTs[1].quantity.toString()).to.equal("6");
    expect(ownedNFTs[1].uri).to.equal("https://gamejam.com/nft1155/15.json");
    expect(ownedNFTs[2].tokenId.toString()).to.equal("98");
    expect(ownedNFTs[2].quantity.toString()).to.equal("41");
    expect(ownedNFTs[2].uri).to.equal("https://gamejam.com/nft1155/98.json");
  });

  it("Set token base URI", async () => {
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .setBaseTokenURI("https://nft-gamejam.com/1155/");
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(ownedNFTs.length).to.equal(3);
    expect(ownedNFTs[1].uri).to.equal("https://nft-gamejam.com/1155/15.json");
  });

  it("Check URI of non-existent NFTs", async () => {
    await expect(this.nft1155Contract.uri(100)).to.be.revertedWith("GamejamNFT1155: URI query for non-existent token");
  });

  it("Burn some NFTs", async () => {
    await this.nft1155Factory
      .connect(this.user)
      .attach(this.nft1155Contract.address)
      .setApprovalForAll(this.minter.address, true);
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .burn(this.user.address, 15, 4);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(ownedNFTs.length).to.equal(3);
    expect(ownedNFTs[1].quantity.toString()).to.equal("2");
  });

  it("Mint batch", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .mintBatch(
        this.user.address,
        [12, 23, 34],
        [43, 32, 21],
        soliditySha3("Forth minting time")
      );
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(ownedNFTs.length).to.equal(6);
    expect(ownedNFTs[3].tokenId.toString()).to.equal("12");
    expect(ownedNFTs[3].quantity.toString()).to.equal("43");
    expect(ownedNFTs[3].uri).to.equal("https://nft-gamejam.com/1155/12.json");
    expect(ownedNFTs[4].tokenId.toString()).to.equal("23");
    expect(ownedNFTs[4].quantity.toString()).to.equal("32");
    expect(ownedNFTs[4].uri).to.equal("https://nft-gamejam.com/1155/23.json");
    expect(ownedNFTs[5].tokenId.toString()).to.equal("34");
    expect(ownedNFTs[5].quantity.toString()).to.equal("21");
    expect(ownedNFTs[5].uri).to.equal("https://nft-gamejam.com/1155/34.json");
  });

  it("Burn all NFTs of an id", async () => {
    await this.nft1155Factory
      .connect(this.minter)
      .attach(this.nft1155Contract.address)
      .burn(this.user.address, 12, 43);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address);
    expect(ownedNFTs.length).to.equal(5);
    expect(ownedNFTs[3].tokenId.toString()).to.equal("23");
    expect(ownedNFTs[3].quantity.toString()).to.equal("32");
    expect(ownedNFTs[3].uri).to.equal("https://nft-gamejam.com/1155/23.json");
    expect(ownedNFTs[4].tokenId.toString()).to.equal("34");
    expect(ownedNFTs[4].quantity.toString()).to.equal("21");
    expect(ownedNFTs[4].uri).to.equal("https://nft-gamejam.com/1155/34.json");
  });
});
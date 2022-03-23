require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { soliditySha3 } = require("web3-utils");
const { expect } = require("chai");
const GAMEJAM_NFT_1155 = "JamNFT1155";
const PROXY_REGISTRY = "contracts/tokens/ERC1155/ERC1155Tradable.sol:ProxyRegistry";

before("Deploy JamNFT1155 contract", async () => {
  // Prepare parameters
  const [deployer, creator, user] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.creator = creator;
  this.user = user;

  // Deploy ProxyRegistry
  this.proxyFactory = await hre.ethers.getContractFactory(PROXY_REGISTRY);
  this.proxyContract = await this.proxyFactory.deploy();
  await this.proxyContract.deployed();

  // Deploy the JamNFT1155 contract
  this.nft1155Factory = await hre.ethers.getContractFactory(GAMEJAM_NFT_1155);
  this.nft1155Contract = await this.nft1155Factory.deploy(
    "Guardians of Glory NFT1155",
    "JamNFT1155",
    "https://gamejam.com/nft1155/",
    this.proxyContract.address
  );
  await this.nft1155Contract.deployed();
});

describe("Test JamNFT1155 contract", () => {
  it("Create a new token id", async () => {
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .create(this.user.address, 12, 333, "https://test.com/", soliditySha3("Create #12"));
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .create(this.user.address, 7, 222, "", soliditySha3("Create 7"));
    let creator1 = await this.nft1155Contract.creators(12);
    let creator2 = await this.nft1155Contract.creators(7);
    expect(creator1).to.equal(this.deployer.address);
    expect(creator2).to.equal(this.deployer.address);
  });

  it("Transfer creator role", async () => {
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .setCreator(this.creator.address, [7, 12]);
    let creator1 = await this.nft1155Contract.creators(12);
    let creator2 = await this.nft1155Contract.creators(7);
    expect(creator1).to.equal(this.creator.address);
    expect(creator2).to.equal(this.creator.address);
  });

  it("Mint some first NFTs", async () => {
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 7, 20, soliditySha3("First minting time"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 7);
    let uri = await this.nft1155Contract.uri(7);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 8);
    expect(balance.toString()).to.equal("242");
    expect(uri).to.equal("https://gamejam.com/nft1155/7.json");
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("12");
    expect(ownedNFTs[0].quantity.toString()).to.equal("333");
    expect(ownedNFTs[0].uri).to.equal("https://test.com/");
    expect(ownedNFTs[1].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[1].quantity.toString()).to.equal("242");
    expect(ownedNFTs[1].uri).to.equal("https://gamejam.com/nft1155/7.json");
  });

  it("Mint some more NFTs", async () => {
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 12, 6, soliditySha3("Second minting time"));
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 7, 41, soliditySha3("Third minting time"));
    let balance = await this.nft1155Contract.balanceOf(this.user.address, 12);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 5);
    expect(balance.toString()).to.equal("339");
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("12");
    expect(ownedNFTs[0].quantity.toString()).to.equal("339");
    expect(ownedNFTs[0].uri).to.equal("https://test.com/");
    expect(ownedNFTs[1].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[1].quantity.toString()).to.equal("283");
    expect(ownedNFTs[1].uri).to.equal("https://gamejam.com/nft1155/7.json");
  });

  it("Set token base URI", async () => {
    await this.nft1155Factory
      .connect(this.deployer)
      .attach(this.nft1155Contract.address)
      .setBaseTokenURI("https://nft-gamejam.com/1155/");
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 4);
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[1].uri).to.equal("https://nft-gamejam.com/1155/7.json");
  });

  it("Check URI of non-existent NFTs", async () => {
    await expect(this.nft1155Contract.uri(100)).to.be.revertedWith("JamNFT1155: URI query for non-existent token");
  });

  it("Burn some NFTs", async () => {
    await this.nft1155Factory
      .connect(this.user)
      .attach(this.nft1155Contract.address)
      .setApprovalForAll(this.creator.address, true);
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .burn(this.user.address, 7, 4);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 9);
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[1].quantity.toString()).to.equal("279");
  });

  it("Mint batch", async () => {
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .batchMint(
        this.user.address,
        [12, 7, 12],
        [43, 32, 21],
        soliditySha3("Forth minting time")
      );
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 6);
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("12");
    expect(ownedNFTs[0].quantity.toString()).to.equal("403");
    expect(ownedNFTs[0].uri).to.equal("https://test.com/");
    expect(ownedNFTs[1].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[1].quantity.toString()).to.equal("311");
    expect(ownedNFTs[1].uri).to.equal("https://nft-gamejam.com/1155/7.json");
  });

  it("Burn all NFTs of an id", async () => {
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .burn(this.user.address, 12, 403);
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 12);
    expect(ownedNFTs.length).to.equal(1);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[0].quantity.toString()).to.equal("311");
    expect(ownedNFTs[0].uri).to.equal("https://nft-gamejam.com/1155/7.json");
  });

  it("Mint again to check uri", async () => {
    await this.nft1155Factory
      .connect(this.creator)
      .attach(this.nft1155Contract.address)
      .mint(this.user.address, 12, 999, soliditySha3("Fifth minting time"));
    let ownedNFTs = await this.nft1155Contract.getOwnedTokens(this.user.address, 0, 33);
    expect(ownedNFTs.length).to.equal(2);
    expect(ownedNFTs[0].tokenId.toString()).to.equal("12");
    expect(ownedNFTs[0].quantity.toString()).to.equal("999");
    expect(ownedNFTs[0].uri).to.equal("https://nft-gamejam.com/1155/12.json");
    expect(ownedNFTs[1].tokenId.toString()).to.equal("7");
    expect(ownedNFTs[1].quantity.toString()).to.equal("311");
    expect(ownedNFTs[1].uri).to.equal("https://nft-gamejam.com/1155/7.json");
  });
});
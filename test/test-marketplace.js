require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const JAM_MARKETPLACE_HUB = "JamMarketplaceHub";
const JAM_MARKETPLACE = "JamMarketplace";
const JAM_CLOCK_AUCTION = "JamClockAuction";
const JAM_TRADITIONAL_AUCTION = "JamTraditionalAuction";
const JAM_P2P_TRADING = "JamP2PTrading";
const ERC20 = "JamToken";
const ERC721 = "JamNFT721";
const PROXY_REGISTRY = "contracts/tokens/ERC721/ERC721Tradable.sol:ProxyRegistry";
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

before("Deploy all contracts", async () => {
  // Prepare parameters
  const [deployer, minter, seller, buyer1, buyer2] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.minter = minter;
  this.seller = seller;
  this.buyer1 = buyer1;
  this.buyer2 = buyer2;

  // Deploy ERC20 contract as payment
  this.erc20Factory = await hre.ethers.getContractFactory(ERC20);
  this.erc20Contract = await this.erc20Factory.deploy("JAM Token", "JAM", hre.ethers.utils.parseEther("10000000000"));
  await this.erc20Contract.deployed();

  // Deploy ProxyRegistry contract to support ERC721
  this.proxyRegistryFactory = await hre.ethers.getContractFactory(PROXY_REGISTRY);
  this.proxyRegistryContract = await this.proxyRegistryFactory.deploy();
  await this.proxyRegistryContract.deployed();

  // Deploy ERC721 contract as trading assets
  this.erc721Factory = await hre.ethers.getContractFactory(ERC721);
  this.erc721Contract = await this.erc721Factory.deploy(
    "Gamejam Main NFT",
    "JamNFT",
    "https://gamejam.com/nft721/",
    this.proxyRegistryContract.address,
    100
  );
  await this.erc721Contract.deployed();

  // Deploy JamMarketplaceHub
  this.hubFactory = await hre.ethers.getContractFactory(JAM_MARKETPLACE_HUB);
  this.hubContract = await this.hubFactory.deploy();
  await this.hubContract.deployed();

  // Deploy JamMarketplace
  this.jamMarketplaceFactory = await hre.ethers.getContractFactory(JAM_MARKETPLACE);
  this.jamMarketplaceContract = await this.jamMarketplaceFactory.deploy(this.hubContract.address, 2000);
  await this.jamMarketplaceContract.deployed();

  // Deploy JamClockAuction
  this.jamClockAuctionFactory = await hre.ethers.getContractFactory(JAM_CLOCK_AUCTION);
  this.jamClockAuctionContract = await this.jamClockAuctionFactory.deploy(this.hubContract.address, 2000);
  await this.jamClockAuctionContract.deployed();

  // Deploy JamTraditionalAuction
  this.jamTraditionalAuctionFactory = await hre.ethers.getContractFactory(JAM_TRADITIONAL_AUCTION);
  this.jamTraditonalAuctionContract = await this.jamClockAuctionFactory.deploy(this.hubContract.address, 2000);
  await this.jamTraditonalAuctionContract.deployed();

  // Deploy JamP2PTrading
  this.jamP2PTradingFactory = await hre.ethers.getContractFactory(JAM_P2P_TRADING);
  this.jamP2PTradingContract = await this.jamP2PTradingFactory.deploy(this.hubContract.address, 2000);
  await this.jamP2PTradingContract.deployed();
});

describe("Set all marketplace contracts up", () => {
  it("Set JamMarketplace up", async () => {
    await this.jamMarketplaceFactory
      .connect(this.deployer)
      .attach(this.jamMarketplaceContract.address)
      .registerWithHub();
    let checkAddress = await this.hubContract.isMarketplace(this.jamMarketplaceContract.address);
    let id = await this.jamMarketplaceContract.marketplaceId();
    let addr = await this.hubContract.getMarketplace(id);
    expect(checkAddress).to.equal(true);
    expect(addr).to.equal(this.jamMarketplaceContract.address);
  });

  it("Set JamClockAuction up", async () => {
    await this.jamClockAuctionFactory
      .connect(this.deployer)
      .attach(this.jamClockAuctionContract.address)
      .registerWithHub();
    let checkAddress = await this.hubContract.isMarketplace(this.jamClockAuctionContract.address);
    let id = await this.jamClockAuctionContract.marketplaceId();
    let addr = await this.hubContract.getMarketplace(id);
    expect(checkAddress).to.equal(true);
    expect(addr).to.equal(this.jamClockAuctionContract.address);
  });

  it("Set JamTraditionalMarketplace up", async () => {
    await this.jamTraditionalAuctionFactory
      .connect(this.deployer)
      .attach(this.jamTraditonalAuctionContract.address)
      .registerWithHub();
    let checkAddress = await this.hubContract.isMarketplace(this.jamTraditonalAuctionContract.address);
    let id = await this.jamTraditonalAuctionContract.marketplaceId();
    let addr = await this.hubContract.getMarketplace(id);
    expect(checkAddress).to.equal(true);
    expect(addr).to.equal(this.jamTraditonalAuctionContract.address);
  });

  it("Set JamP2PTrading up", async () => {
    await this.jamP2PTradingFactory
      .connect(this.deployer)
      .attach(this.jamP2PTradingContract.address)
      .registerWithHub();
    let checkAddress = await this.hubContract.isMarketplace(this.jamP2PTradingContract.address);
    let id = await this.jamP2PTradingContract.marketplaceId();
    let addr = await this.hubContract.getMarketplace(id);
    expect(checkAddress).to.equal(true);
    expect(addr).to.equal(this.jamP2PTradingContract.address);
  });
});

describe("Mint users some NFTs to sell and some ERC20 tokens to buy", () => {
  it("Set ERC20 tokens minter role", async () => {
    let minterRole = await this.erc20Contract.MINTER_ROLE();
    await this.erc20Factory
      .connect(this.deployer)
      .attach(this.erc20Contract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole = await this.erc20Contract.hasRole(minterRole, this.minter.address);
    expect(checkRole).to.equal(true);
  });

  it("Set NFT minter role", async () => {
    let minterRole = await this.erc721Contract.MINTER_ROLE();
    await this.erc721Factory
      .connect(this.deployer)
      .attach(this.erc721Contract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole = await this.erc721Contract.hasRole(minterRole, this.minter.address);
    expect(checkRole).to.equal(true);
  });

  it("Mint some ERC20 tokens", async () => {
    await this.erc20Factory
      .connect(this.minter)
      .attach(this.erc20Contract.address)
      .mint(this.buyer1.address, hre.ethers.utils.parseEther("10000"));
    await this.erc20Factory
      .connect(this.minter)
      .attach(this.erc20Contract.address)
      .mint(this.buyer2.address, hre.ethers.utils.parseEther("20000"));
    let balance1 = await this.erc20Contract.balanceOf(this.buyer1.address);
    let balance2 = await this.erc20Contract.balanceOf(this.buyer2.address);
    expect(balance1.toString()).to.equal(hre.ethers.utils.parseEther("10000"));
    expect(balance2.toString()).to.equal(hre.ethers.utils.parseEther("20000"));
  });

  it("Mint some NFTs to sell", async () => {
    for (let tokenId = 0; tokenId < 8; tokenId++)
      await this.erc721Factory
        .connect(this.minter)
        .attach(this.erc721Contract.address)
        .mint(this.seller.address, tokenId, `https://gamejam.com/NFT${tokenId}.json`);
    let balance = await this.erc721Contract.balanceOf(this.seller.address);
    expect(balance.toString()).to.equal("8");
  });
});

describe("Test JamMarketplace", () => {
  //
});

describe("Test JamClockAuction", () => {
  it("Create new auction with native token as payment", async () => {
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamClockAuctionContract.address, 2);
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        2,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("50"),
        hre.ethers.utils.parseEther("20"),
        60
      );
    let currentOwner = await this.erc721Contract.ownerOf(2);
    let auction = await this.jamClockAuctionContract.getAuction(this.erc721Contract.address, 2);
    expect(currentOwner).to.equal(this.jamClockAuctionContract.address);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.startingPrice.toString()).to.equal(hre.ethers.utils.parseEther("50"));
    expect(auction.endingPrice.toString()).to.equal(hre.ethers.utils.parseEther("20"));
    expect(auction.duration.toString()).to.equal("60");
    expect(auction.startedAt.toString()).not.to.equal("0");
  });

  it("Update this auction's information", async () => {
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .updateAuction(
        this.erc721Contract.address,
        2,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("70"),
        hre.ethers.utils.parseEther("40"),
        120
      );
    let auction = await this.jamClockAuctionContract.getAuction(this.erc721Contract.address, 2);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.startingPrice.toString()).to.equal(hre.ethers.utils.parseEther("70"));
    expect(auction.endingPrice.toString()).to.equal(hre.ethers.utils.parseEther("40"));
    expect(auction.duration.toString()).to.equal("120");
    expect(auction.startedAt.toString()).not.to.equal("0");
  });

  it("A buyer bids this auction using erc20 tokens", async () => {
    let price = await this.jamClockAuctionContract.getCurrentPrice(this.erc721Contract.address, 2);
    await this.erc20Factory
      .connect(this.buyer1)
      .attach(this.erc20Contract.address)
      .approve(this.jamClockAuctionContract.address, price);
    await this.jamClockAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamClockAuctionContract.address)
      .bid(this.erc721Contract.address, 2, price);
    let currentOwner = await this.erc721Contract.ownerOf(2);
    expect(currentOwner).to.equal(this.buyer1.address);
  });

  it("Create another auction with native tokens as payment", async () => {
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamClockAuctionContract.address, 3);
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        3,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        hre.ethers.utils.parseEther("25"),
        60
      );
    let auction = await this.jamClockAuctionContract.getAuction(this.erc721Contract.address, 3);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.startingPrice.toString()).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endingPrice.toString()).to.equal(hre.ethers.utils.parseEther("25"));
    expect(auction.duration.toString()).to.equal("60");
    expect(auction.startedAt.toString()).not.to.equal("0");
  });

  it("Cancel this auction", async () => {
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .cancelAuction(this.erc721Contract.address, 3, this.seller.address);
    let currentOwner = await this.erc721Contract.ownerOf(3);
    expect(currentOwner).to.equal(this.seller.address);
  });

  it("Create an auction again and a seller buys it with native token", async () => {
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamClockAuctionContract.address, 3);
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        3,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        hre.ethers.utils.parseEther("15"),
        60
      );
    let price = await this.jamClockAuctionContract.getCurrentPrice(this.erc721Contract.address, 3);
    await expect(
      this.jamClockAuctionFactory
        .connect(this.buyer1)
        .attach(this.jamClockAuctionContract.address)
        .bid(this.erc721Contract.address, 3, 0, { value: price })
    ).to.be.revertedWith("JamClockAuction: bid amount info mismatch");
    await this.jamClockAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamClockAuctionContract.address)
      .bid(this.erc721Contract.address, 3, price, { value: price });
    let currentOwner = await this.erc721Contract.ownerOf(3);
    expect(currentOwner).to.equal(this.buyer1.address);
  });
});

describe("Test JamTraditionalAuction", () => {
  //
});

describe("Test JamP2PTrading", () => {
  //
});

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};
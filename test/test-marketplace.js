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
  this.jamTraditionalAuctionContract = await this.jamTraditionalAuctionFactory.deploy(this.hubContract.address, 2000);
  await this.jamTraditionalAuctionContract.deployed();

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
      .attach(this.jamTraditionalAuctionContract.address)
      .registerWithHub();
    let checkAddress = await this.hubContract.isMarketplace(this.jamTraditionalAuctionContract.address);
    let id = await this.jamTraditionalAuctionContract.marketplaceId();
    let addr = await this.hubContract.getMarketplace(id);
    expect(checkAddress).to.equal(true);
    expect(addr).to.equal(this.jamTraditionalAuctionContract.address);
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
    for (let tokenId = 0; tokenId < 9; tokenId++)
      await this.erc721Factory
        .connect(this.minter)
        .attach(this.erc721Contract.address)
        .mint(this.seller.address, tokenId, `https://gamejam.com/NFT${tokenId}.json`);
    let balance = await this.erc721Contract.balanceOf(this.seller.address);
    expect(balance.toString()).to.equal("9");
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
      .cancelAuction(this.erc721Contract.address, 3);
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
  it("Create new auction with native token as payment", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 3000);
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamTraditionalAuctionContract.address, 4);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        4,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        auctionEndsAt
      );
    let auction = await this.jamTraditionalAuctionContract.getAuction(this.erc721Contract.address, 4);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.highestBidder).to.equal(ZERO_ADDRESS);
    expect(auction.highestBidAmount.toString()).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endAt.toString()).to.equal(auctionEndsAt.toString());
  });

  it("Cancel this auction", async () => {
    let isAuctionCancelable = await this.jamTraditionalAuctionContract.isAuctionCancelable(this.erc721Contract.address, 4);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .cancelAuction(this.erc721Contract.address, 4);
    let currentOwner = await this.erc721Contract.ownerOf(4);
    expect(isAuctionCancelable).to.equal(true);
    expect(currentOwner).to.equal(this.seller.address);
  });

  it("Create this auction again an update its information", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 3000);
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamTraditionalAuctionContract.address, 4);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        4,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        auctionEndsAt
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .updateAuction(
        this.erc721Contract.address,
        4,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("10"),
        auctionEndsAt - 2300
      );
    let auction = await this.jamTraditionalAuctionContract.getAuction(this.erc721Contract.address, 4);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.highestBidder).to.equal(ZERO_ADDRESS);
    expect(auction.highestBidAmount.toString()).to.equal(hre.ethers.utils.parseEther("10"));
    expect(auction.endAt.toString()).to.equal((auctionEndsAt - 2300).toString());
  });

  it("The first buyer bids 25", async () => {
    await this.erc20Factory
      .connect(this.buyer1)
      .attach(this.erc20Contract.address)
      .approve(
        this.jamTraditionalAuctionContract.address,
        hre.ethers.utils.parseEther("25")
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(this.erc721Contract.address, 4, hre.ethers.utils.parseEther("25"));
    let isAuctionCancelable = await this.jamTraditionalAuctionContract.isAuctionCancelable(this.erc721Contract.address, 4);
    let auction = await this.jamTraditionalAuctionContract.getAuction(this.erc721Contract.address, 4);
    expect(isAuctionCancelable).to.equal(false);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.highestBidder).to.equal(this.buyer1.address);
    expect(auction.highestBidAmount).to.equal(hre.ethers.utils.parseEther("25"));
    expect(auction.endAt).not.to.equal("0");
  });

  it("The second buyer bids 30", async () => {
    await this.erc20Factory
      .connect(this.buyer2)
      .attach(this.erc20Contract.address)
      .approve(
        this.jamTraditionalAuctionContract.address,
        hre.ethers.utils.parseEther("30")
      );
    await expect(
      this.jamTraditionalAuctionFactory
        .connect(this.buyer2)
        .attach(this.jamTraditionalAuctionContract.address)
        .bid(this.erc721Contract.address, 4, hre.ethers.utils.parseEther("25"))
    ).to.be.revertedWith("JamTraditionalAuction: currently has higher bid");
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer2)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(this.erc721Contract.address, 4, hre.ethers.utils.parseEther("30"));
    let isAuctionCancelable = await this.jamTraditionalAuctionContract.isAuctionCancelable(this.erc721Contract.address, 4);
    let auction = await this.jamTraditionalAuctionContract.getAuction(this.erc721Contract.address, 4);
    expect(isAuctionCancelable).to.equal(false);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.highestBidder).to.equal(this.buyer2.address);
    expect(auction.highestBidAmount).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endAt).not.to.equal("0");
  });

  // it("The winner claims the asset", async () => {
  //   await expect(
  //     this.jamTraditionalAuctionFactory
  //       .connect(this.buyer2)
  //       .attach(this.jamTraditionalAuctionContract.address)
  //       .claimAsset(this.erc721Contract.address, 4)
  //   ).to.be.revertedWith("JamTraditionalAuction: auction not ends yet");
  //   await sleep(20000);
  //   await this.jamTraditionalAuctionFactory
  //     .connect(this.buyer2)
  //     .attach(this.jamTraditionalAuctionContract.address)
  //     .claimAsset(this.erc721Contract.address, 4);
  //   let currentOwner = await this.erc721Contract.ownerOf(4);
  //   expect(currentOwner).to.equal(this.buyer2.address);
  // });

  it("Create another auction using native token", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 700);
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamTraditionalAuctionContract.address, 5);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        5,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("5"),
        auctionEndsAt
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(
        this.erc721Contract.address,
        5,
        hre.ethers.utils.parseEther("6"),
        { value: hre.ethers.utils.parseEther("6") }
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer2)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(
        this.erc721Contract.address,
        5,
        hre.ethers.utils.parseEther("7"),
        { value: hre.ethers.utils.parseEther("7") }
      );
    let isAuctionCancelable = await this.jamTraditionalAuctionContract.isAuctionCancelable(this.erc721Contract.address, 5);
    let auction = await this.jamTraditionalAuctionContract.getAuction(this.erc721Contract.address, 5);
    expect(isAuctionCancelable).to.equal(false);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.highestBidder).to.equal(this.buyer2.address);
    expect(auction.highestBidAmount).to.equal(hre.ethers.utils.parseEther("7"));
    expect(auction.endAt).not.to.equal("0");
  });
});

describe("Test JamP2PTrading", () => {
  it("Make a new offer for an existent NFT", async () => {
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc721Contract.address,
        6,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("12"),
        { value: hre.ethers.utils.parseEther("12") }
      );
    let offersForNFT = await this.jamP2PTradingContract.getOffersFor(this.erc721Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc721Contract.address,
      6
    );
    expect(offer.offeror).to.equal(this.buyer1.address);
    expect(offer.nftAddress).to.equal(this.erc721Contract.address);
    expect(offer.tokenId.toString()).to.equal("6");
    expect(offer.currency).to.equal(ZERO_ADDRESS);
    expect(offer.amount.toString()).to.equal(hre.ethers.utils.parseEther("12").toString());
    expect(offersForNFT.length).to.equal(1);
    expect(offersForNFT[0].toString()).to.equal(offer.toString());
    expect(offersOfOfferer.length).to.equal(1);
    expect(offersOfOfferer[0].toString()).to.equal(offer.toString());
  });

  it("Update this offer", async () => {
    await this.erc20Factory
      .connect(this.buyer1)
      .attach(this.erc20Contract.address)
      .approve(
        this.jamP2PTradingContract.address,
        hre.ethers.utils.parseEther("15")
      );
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .updateOffer(
        this.erc721Contract.address,
        6,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("15")
      );
    let offersForNFT = await this.jamP2PTradingContract.getOffersFor(this.erc721Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc721Contract.address,
      6
    );
    expect(offer.offeror).to.equal(this.buyer1.address);
    expect(offer.nftAddress).to.equal(this.erc721Contract.address);
    expect(offer.tokenId.toString()).to.equal("6");
    expect(offer.currency).to.equal(this.erc20Contract.address);
    expect(offer.amount.toString()).to.equal(hre.ethers.utils.parseEther("15").toString());
    expect(offersForNFT.length).to.equal(1);
    expect(offersForNFT[0].toString()).to.equal(offer.toString());
    expect(offersOfOfferer.length).to.equal(1);
    expect(offersOfOfferer[0].toString()).to.equal(offer.toString());
  });

  it("Cancel this offer", async () => {
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .cancelOffer(this.erc721Contract.address, 6);
    let offersForNFT = await this.jamP2PTradingContract.getOffersFor(this.erc721Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc721Contract.address,
      6
    );
    expect(offer.offeror).to.equal(ZERO_ADDRESS);
    expect(offer.nftAddress).to.equal(ZERO_ADDRESS);
    expect(offer.tokenId.toString()).to.equal("0");
    expect(offer.currency).to.equal(ZERO_ADDRESS);
    expect(offer.amount.toString()).to.equal("0");
    expect(offersForNFT.length).to.equal(0);
    expect(offersOfOfferer.length).to.equal(0);
  });

  it("The offeror offers again and the owner accepts", async () => {
    await this.erc20Factory
      .connect(this.buyer2)
      .attach(this.erc20Contract.address)
      .approve(
        this.jamP2PTradingContract.address,
        hre.ethers.utils.parseEther("40")
      );
    await this.jamP2PTradingFactory
      .connect(this.buyer2)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc721Contract.address,
        6,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("40")
      );
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamP2PTradingContract.address, 6);
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer2.address, this.erc721Contract.address, 6);
    let currentOwner = await this.erc721Contract.ownerOf(6);
    expect(currentOwner).to.equal(this.buyer2.address);
  });

  it("Another trading session using native token", async () => {
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc721Contract.address,
        7,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("8"),
        { value: hre.ethers.utils.parseEther("8") }
      );
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamP2PTradingContract.address, 7);
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer1.address, this.erc721Contract.address, 7);
    let currentOwner = await this.erc721Contract.ownerOf(7);
    expect(currentOwner).to.equal(this.buyer1.address);
  });

  it("Accept an offer for an NFT which is currently on marketplace", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 700);
    await this.erc721Factory
      .connect(this.seller)
      .attach(this.erc721Contract.address)
      .approve(this.jamTraditionalAuctionContract.address, 8);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc721Contract.address,
        8,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("20"),
        auctionEndsAt
      );
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc721Contract.address,
        8,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        { value: hre.ethers.utils.parseEther("30") }
      );
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer1.address, this.erc721Contract.address, 8);
  });
});

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};
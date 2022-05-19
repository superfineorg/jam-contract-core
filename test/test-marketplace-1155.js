require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const { soliditySha3 } = require('web3-utils');
const JAM_NFT_OWNERS = "JamNFTOwners";
const JAM_MARKETPLACE_HUB = "JamMarketplaceHub";
const JAM_MARKETPLACE = "JamMarketplace1155";
const JAM_CLOCK_AUCTION = "JamClockAuction1155";
const JAM_TRADITIONAL_AUCTION = "JamTraditionalAuction1155";
const JAM_P2P_TRADING = "JamP2PTrading1155";
const ERC20 = "JamToken";
const ERC1155 = "JamNFT1155";
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

  // Deploy ERC1155 contract as trading assets
  this.erc1155Factory = await hre.ethers.getContractFactory(ERC1155);
  this.erc1155Contract = await this.erc1155Factory.deploy("https://gamejam.com/nft1155/");
  await this.erc1155Contract.deployed();

  // Deploy JamNFTOwners
  this.jamNFTOwnersFactory = await hre.ethers.getContractFactory(JAM_NFT_OWNERS);
  this.jamNFTOwnersContract = await this.jamNFTOwnersFactory.deploy();
  await this.jamNFTOwnersContract.deployed();

  // Deploy JamMarketplaceHub
  this.hubFactory = await hre.ethers.getContractFactory(JAM_MARKETPLACE_HUB);
  this.hubContract = await this.hubFactory.deploy(this.jamNFTOwnersContract.address);
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
    let minterRole = await this.erc1155Contract.MINTER_ROLE();
    await this.erc1155Factory
      .connect(this.deployer)
      .attach(this.erc1155Contract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole = await this.erc1155Contract.hasRole(minterRole, this.minter.address);
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
    for (let tokenId = 0; tokenId < 9; tokenId++) {
      await this.erc1155Factory
        .connect(this.minter)
        .attach(this.erc1155Contract.address)
        .mint(this.seller.address, tokenId, 1000, soliditySha3(`Mint #${tokenId}`));
      let balance = await this.erc1155Contract.balanceOf(this.seller.address, tokenId);
      expect(balance.toString()).to.equal("1000");
    }
  });
});

describe("Test JamMarketplace", () => {
  //
});

describe("Test JamClockAuction", () => {
  it("Create new auction with native token as payment", async () => {
    await this.erc1155Factory
      .connect(this.seller)
      .attach(this.erc1155Contract.address)
      .setApprovalForAll(this.jamClockAuctionContract.address, true);
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        2,
        20,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("50"),
        hre.ethers.utils.parseEther("20"),
        60
      );
    let marketplaceBalance = await this.erc1155Contract.balanceOf(this.jamClockAuctionContract.address, 2);
    let auction = await this.jamClockAuctionContract.getAuction(0);
    expect(marketplaceBalance.toString()).to.equal("20");
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.quantity).to.equal("20");
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
        0,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("70"),
        hre.ethers.utils.parseEther("40"),
        120
      );
    let auction = await this.jamClockAuctionContract.getAuction(0);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("2");
    expect(auction.quantity.toString()).to.equal("20");
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.startingPrice.toString()).to.equal(hre.ethers.utils.parseEther("70"));
    expect(auction.endingPrice.toString()).to.equal(hre.ethers.utils.parseEther("40"));
    expect(auction.duration.toString()).to.equal("120");
    expect(auction.startedAt.toString()).not.to.equal("0");
  });

  it("A buyer bids this auction using erc20 tokens", async () => {
    let price = await this.jamClockAuctionContract.getCurrentTotalPrice(0);
    await this.erc20Factory
      .connect(this.buyer1)
      .attach(this.erc20Contract.address)
      .approve(this.jamClockAuctionContract.address, price);
    await this.jamClockAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamClockAuctionContract.address)
      .bid(0, price);
    let buyer1Balance = await this.erc1155Contract.balanceOf(this.buyer1.address, 2);
    expect(buyer1Balance.toString()).to.equal("20");
  });

  it("Create another auction with native tokens as payment", async () => {
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        3,
        30,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        hre.ethers.utils.parseEther("25"),
        60
      );
    let auction = await this.jamClockAuctionContract.getAuction(1);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("3");
    expect(auction.quantity.toString()).to.equal("30");
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.startingPrice.toString()).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endingPrice.toString()).to.equal(hre.ethers.utils.parseEther("25"));
    expect(auction.duration.toString()).to.equal("60");
    expect(auction.startedAt.toString()).not.to.equal("0");
  });

  it("Cancel this auction", async () => {
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .cancelAuction1155(1);
    let currentBalance = await this.erc1155Contract.balanceOf(this.seller.address, 3);
    expect(currentBalance.toString()).to.equal("1000");
  });

  it("Create an auction again and a seller buys it with native token", async () => {
    await this.jamClockAuctionFactory
      .connect(this.seller)
      .attach(this.jamClockAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        3,
        35,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        hre.ethers.utils.parseEther("15"),
        60
      );
    let price = await this.jamClockAuctionContract.getCurrentTotalPrice(2);
    await expect(
      this.jamClockAuctionFactory
        .connect(this.buyer1)
        .attach(this.jamClockAuctionContract.address)
        .bid(2, 0, { value: price })
    ).to.be.revertedWith("JamClockAuction1155: bid amount info mismatch");
    await this.jamClockAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamClockAuctionContract.address)
      .bid(2, price, { value: price });
    let buyer1Balance = await this.erc1155Contract.balanceOf(this.buyer1.address, 3);
    expect(buyer1Balance.toString()).to.equal("35");
  });
});

describe("Test JamTraditionalAuction", () => {
  it("Create new auction with native token as payment", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 3000);
    await this.erc1155Factory
      .connect(this.seller)
      .attach(this.erc1155Contract.address)
      .setApprovalForAll(this.jamTraditionalAuctionContract.address, true);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        4,
        40,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        auctionEndsAt
      );
    let auction = await this.jamTraditionalAuctionContract.getAuction(0);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("4");
    expect(auction.quantity.toString()).to.equal("40");
    expect(auction.currency).to.equal(ZERO_ADDRESS);
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.highestBidder).to.equal(ZERO_ADDRESS);
    expect(auction.highestBidAmount.toString()).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endAt.toString()).to.equal(auctionEndsAt.toString());
  });

  it("Cancel this auction", async () => {
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .cancelAuction1155(0);
    let currentBalance = await this.erc1155Contract.balanceOf(this.seller.address, 4);
    expect(currentBalance.toString()).to.equal("1000");
  });

  it("Create this auction again an update its information", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 3000);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        4,
        45,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        auctionEndsAt
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .updateAuction(
        1,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("10"),
        auctionEndsAt - 2300
      );
    let auction = await this.jamTraditionalAuctionContract.getAuction(1);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("4");
    expect(auction.quantity.toString()).to.equal("45");
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.seller).to.equal(this.seller.address);
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
      .bid(1, hre.ethers.utils.parseEther("25"));
    let auction = await this.jamTraditionalAuctionContract.getAuction(1);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("4");
    expect(auction.quantity.toString()).to.equal("45");
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
        .bid(1, hre.ethers.utils.parseEther("25"))
    ).to.be.revertedWith("JamTraditionalAuction1155: currently has higher bid");
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer2)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(1, hre.ethers.utils.parseEther("30"));
    let auction = await this.jamTraditionalAuctionContract.getAuction(1);
    expect(auction.nftAddress).to.equal(this.erc1155Contract.address);
    expect(auction.tokenId.toString()).to.equal("4");
    expect(auction.quantity.toString()).to.equal("45");
    expect(auction.seller).to.equal(this.seller.address);
    expect(auction.currency).to.equal(this.erc20Contract.address);
    expect(auction.highestBidder).to.equal(this.buyer2.address);
    expect(auction.highestBidAmount).to.equal(hre.ethers.utils.parseEther("30"));
    expect(auction.endAt).not.to.equal("0");
  });

  // // it("The winner claims the asset", async () => {
  // //   await expect(
  // //     this.jamTraditionalAuctionFactory
  // //       .connect(this.buyer2)
  // //       .attach(this.jamTraditionalAuctionContract.address)
  // //       .finalizeAuction(this.erc721Contract.address, 4)
  // //   ).to.be.revertedWith("JamTraditionalAuction721: auction not ends yet");
  // //   await sleep(20000);
  // //   await this.jamTraditionalAuctionFactory
  // //     .connect(this.buyer2)
  // //     .attach(this.jamTraditionalAuctionContract.address)
  // //     .finalizeAuction(this.erc721Contract.address, 4);
  // //   let currentOwner = await this.erc721Contract.ownerOf(4);
  // //   expect(currentOwner).to.equal(this.buyer2.address);
  // // });

  it("Create another auction using native token", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 700);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        5,
        50,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("5"),
        auctionEndsAt
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer1)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(
        2,
        hre.ethers.utils.parseEther("6"),
        { value: hre.ethers.utils.parseEther("6") }
      );
    await this.jamTraditionalAuctionFactory
      .connect(this.buyer2)
      .attach(this.jamTraditionalAuctionContract.address)
      .bid(
        2,
        hre.ethers.utils.parseEther("7"),
        { value: hre.ethers.utils.parseEther("7") }
      );
    let auction = await this.jamTraditionalAuctionContract.getAuction(2);
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
        this.erc1155Contract.address,
        6,
        60,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("12"),
        { value: hre.ethers.utils.parseEther("12") }
      );
    let offersForNFT = await this.jamP2PTradingContract.getAllOffersFor(this.erc1155Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc1155Contract.address,
      6
    );
    expect(offer.offeror).to.equal(this.buyer1.address);
    expect(offer.nftAddress).to.equal(this.erc1155Contract.address);
    expect(offer.tokenId.toString()).to.equal("6");
    expect(offer.quantity.toString()).to.equal("60");
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
        this.erc1155Contract.address,
        6,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("15")
      );
    let offersForNFT = await this.jamP2PTradingContract.getAllOffersFor(this.erc1155Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc1155Contract.address,
      6
    );
    expect(offer.offeror).to.equal(this.buyer1.address);
    expect(offer.nftAddress).to.equal(this.erc1155Contract.address);
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
      .cancelOffer(this.erc1155Contract.address, 6);
    let offersForNFT = await this.jamP2PTradingContract.getAllOffersFor(this.erc1155Contract.address, 6);
    let offersOfOfferer = await this.jamP2PTradingContract.getOffersOf(this.buyer1.address);
    let offer = await this.jamP2PTradingContract.getSpecificOffer(
      this.buyer1.address,
      this.erc1155Contract.address,
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
        this.erc1155Contract.address,
        6,
        65,
        this.erc20Contract.address,
        hre.ethers.utils.parseEther("40")
      );
    await this.erc1155Factory
      .connect(this.seller)
      .attach(this.erc1155Contract.address)
      .setApprovalForAll(this.jamP2PTradingContract.address, true);
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer2.address, this.erc1155Contract.address, 6);
    let offerorBalance = await this.erc1155Contract.balanceOf(this.buyer2.address, 6);
    expect(offerorBalance.toString()).to.equal("65");
  });

  it("Another trading session using native token", async () => {
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc1155Contract.address,
        7,
        70,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("8"),
        { value: hre.ethers.utils.parseEther("8") }
      );
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer1.address, this.erc1155Contract.address, 7);
    let offerorBalance = await this.erc1155Contract.balanceOf(this.buyer1.address, 7);
    expect(offerorBalance.toString()).to.equal("70");
  });

  it("Accept an offer for an NFT which is currently on marketplace", async () => {
    let auctionEndsAt = Math.floor(Date.now() / 1000 + 700);
    await this.jamTraditionalAuctionFactory
      .connect(this.seller)
      .attach(this.jamTraditionalAuctionContract.address)
      .createAuction(
        this.erc1155Contract.address,
        8,
        80,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("20"),
        auctionEndsAt
      );
    await this.jamP2PTradingFactory
      .connect(this.buyer1)
      .attach(this.jamP2PTradingContract.address)
      .makeOffer(
        this.erc1155Contract.address,
        8,
        250,
        ZERO_ADDRESS,
        hre.ethers.utils.parseEther("30"),
        { value: hre.ethers.utils.parseEther("30") }
      );
    await this.jamP2PTradingFactory
      .connect(this.seller)
      .attach(this.jamP2PTradingContract.address)
      .acceptOffer(this.buyer1.address, this.erc1155Contract.address, 8);
    let offerorBalance = await this.erc1155Contract.balanceOf(this.buyer1.address, 8);
    expect(offerorBalance.toString()).to.equal("250");
  });
});
require('@nomiclabs/hardhat-ethers');

const { expect } = require('chai');
const hre = require('hardhat');
const JAM_OG_PASS = "JamOGPass";
const JAM_SUPER_HAPPY_FRENS = "JamSuperHappyFrens";
const MINTING = "JamOGPassMinting";

before("Deploy contracts", async () => {
  // Prepare parameters
  const [deployer, user] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.user = user;

  // Deploy JamOGPass contract
  this.ogPassFactory = await hre.ethers.getContractFactory(JAM_OG_PASS);
  this.ogPassContract = await this.ogPassFactory.deploy(
    "OGPass NFT",
    "OGP",
    "https://gamejam.com/og-pass/",
    "0x0000000000000000000000000000000000000022"
  );
  await this.ogPassContract.deployed();

  // Deploy JamSuperHappyFrens contract
  this.shfFactory = await hre.ethers.getContractFactory(JAM_SUPER_HAPPY_FRENS);
  this.shfContract = await this.shfFactory.deploy(
    "Super Happy Frens NFT",
    "SHF",
    "https://gamejam.com/super-happy-frens/",
    "0x0000000000000000000000000000000000000034",
    2000
  );
  await this.shfContract.deployed();

  // Deploy JamOGPassMinting contract
  this.mintingFactory = await hre.ethers.getContractFactory(MINTING);
  this.mintingContract = await this.mintingFactory.deploy(
    this.ogPassContract.address,
    this.shfContract.address,
    1234,
    3333,
    2222,
    10,
    10
  );
  await this.mintingContract.deployed();
});

describe("Test OGPass and Super Happy Frens", () => {
  it("Set minter role to the minting contract", async () => {
    let minterRole = await this.ogPassContract.MINTER_ROLE();
    await this.ogPassFactory
      .connect(this.deployer)
      .attach(this.ogPassContract.address)
      .grantRole(minterRole, this.mintingContract.address);
    await this.shfFactory
      .connect(this.deployer)
      .attach(this.shfContract.address)
      .grantRole(minterRole, this.mintingContract.address);
    let ogPassMinterRole = await this.ogPassContract.hasRole(minterRole, this.mintingContract.address);
    let shfMinterRole = await this.shfContract.hasRole(minterRole, this.mintingContract.address);
    expect(ogPassMinterRole).to.equal(true);
    expect(shfMinterRole).to.equal(true);
  });

  it("User buys some SuperHappyFrens NFTs with normal price", async () => {
    await expect(
      this.mintingFactory
        .connect(this.user)
        .attach(this.mintingContract.address)
        .purchaseSuperHappyFrens(2, { value: 6600 })
    ).to.be.revertedWith("JamOGPassMinting: not enough money");
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .purchaseSuperHappyFrens(2, { value: 7000 });
    for (let i = 0; i <= 1; i++) {
      let currentOwner = await this.shfContract.ownerOf(i);
      expect(currentOwner).to.equal(this.user.address);
    }
  });

  it("User mints himself an OGPass", async () => {
    await expect(
      this.mintingFactory
        .connect(this.user)
        .attach(this.mintingContract.address)
        .mintOGPass(2, { value: 2466 })
    ).to.be.revertedWith("JamOGPassMinting: not enough fee to mint");
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .mintOGPass(2, { value: 2468 });
    let currentOwner = await this.ogPassContract.ownerOf(0);
    expect(currentOwner).to.equal(this.user.address);
  });

  it("User buys some SuperHappyFrens NFTs with discount price", async () => {
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .purchaseSuperHappyFrens(3, { value: 7000 });
    for (let i = 2; i <= 4; i++) {
      let currentOwner = await this.shfContract.ownerOf(i);
      expect(currentOwner).to.equal(this.user.address);
    }
  });

  it("Mint more OGPass to test", async () => {
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .mintOGPass(8, { value: 2222 * 8 });
    let ownedTokens = await this.ogPassContract.getOwnedTokens(this.user.address);
    expect(ownedTokens.length).to.equal(10);
    for (let i = 0; i < 10; i++) {
      expect(ownedTokens[i].tokenId.toString()).to.equal(i.toString());
      expect(ownedTokens[i].tokenURI).to.equal(`https://gamejam.com/og-pass/${i}.json`);
    }
  });

  it("Set the new values", async () => {
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setMintFee(9999);
    await expect(
      this.mintingFactory
        .connect(this.deployer)
        .attach(this.mintingContract.address)
        .setSuperHappyFrensPrice(1111)
    ).to.be.revertedWith("JamOGPassMinting: normal price must be greater than discount price");
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setSuperHappyFrensPrice(4444);
    await expect(
      this.mintingFactory
        .connect(this.deployer)
        .attach(this.mintingContract.address)
        .setSuperHappyFrensDiscountPrice(5555)
    ).to.be.revertedWith("JamOGPassMinting: normal price must be greater than discount price");
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setSuperHappyFrensDiscountPrice(3333);
    await expect(
      this.mintingFactory
        .connect(this.deployer)
        .attach(this.mintingContract.address)
        .setMintLimit(0)
    ).to.be.revertedWith("JamOGPassMinting: mint limit must be greater than 0");
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setMintLimit(20);
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setPurchaseLimit(18);
    let mintFee = await this.mintingContract.mintFee();
    let price = await this.mintingContract.jamSuperHappyFrensPrice();
    let discountPrice = await this.mintingContract.jamSuperHappyFrensDiscountPrice();
    let mintLimit = await this.mintingContract.mintLimit();
    let purchaseLimit = await this.mintingContract.purchaseLimit();
    expect(mintFee.toString()).to.equal("9999");
    expect(price.toString()).to.equal("4444");
    expect(discountPrice.toString()).to.equal("3333");
    expect(mintLimit.toString()).to.equal("20");
    expect(purchaseLimit.toString()).to.equal("18");
  });

  it("Reclaim money", async () => {
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .reclaimEther(this.deployer.address);
  });
});
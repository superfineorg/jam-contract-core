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
    "0x0000000000000000000000000000000000000022",
    2000
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
    2222
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

  it("User mints himself an OGPass", async () => {
    await expect(
      this.mintingFactory
        .connect(this.user)
        .attach(this.mintingContract.address)
        .mintOGPass({ value: 1233 })
    ).to.be.revertedWith("JamOGPassMinting: not enough fee to mint");
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .mintOGPass({ value: 1235 });
    let currentOwner = await this.ogPassContract.ownerOf(0);
    expect(currentOwner).to.equal(this.user.address);
  });

  it("User buys a SuperHappyFrens NFT with discount price", async () => {
    await this.ogPassFactory
      .connect(this.user)
      .attach(this.ogPassContract.address)
      .approve(this.mintingContract.address, 0);
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .buySuperHappyFrens(0, { value: 3000 });
    let currentOwner = await this.shfContract.ownerOf(0);
    expect(currentOwner).to.equal(this.user.address);
  });

  it("User buys another SuperHappyFrens NFT with normal price", async () => {
    await expect(
      this.mintingFactory
        .connect(this.user)
        .attach(this.mintingContract.address)
        .buySuperHappyFrens(1, { value: 3000 })
    ).to.be.revertedWith("JamOGPassMinting: not enough money");
    await this.mintingFactory
      .connect(this.user)
      .attach(this.mintingContract.address)
      .buySuperHappyFrens(1, { value: 4000 });
    let currentOwner = await this.shfContract.ownerOf(1);
    expect(currentOwner).to.equal(this.user.address);
  });

  it("Mint more OGPass to test", async () => {
    for (let i = 0; i < 8; i++)
      await this.mintingFactory
        .connect(this.user)
        .attach(this.mintingContract.address)
        .mintOGPass({ value: 2222 });
    let ownedTokens = await this.ogPassContract.getOwnedTokens(this.user.address);
    expect(ownedTokens.length).to.equal(8);
    for (let i = 0; i < 8; i++) {
      expect(ownedTokens[i].tokenId.toString()).to.equal((i + 1).toString());
      expect(ownedTokens[i].tokenURI).to.equal(`https://gamejam.com/og-pass/${i + 1}.json`);
    }
  });
});
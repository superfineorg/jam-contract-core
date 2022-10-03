require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const CURRENCY = "GameToken";
const PLATFORM_TOPUP = "JamPlatformTopup";

before("Deploy contracts", async () => {
  // Prepare parameters
  const [deployer, user, platform] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.user = user;
  this.platform = platform;

  // Deploy ERC20 contract
  this.currencyFactory = await hre.ethers.getContractFactory(CURRENCY);
  this.currencyContract = await this.currencyFactory.deploy();
  await this.currencyContract.deployed();

  // Deploy topup contract
  this.topupFactory = await hre.ethers.getContractFactory(PLATFORM_TOPUP);
  this.topupContract = await this.topupFactory.deploy(platform.address);
  await this.topupContract.deployed();
});

describe("Test JamAirdrop", () => {
  it("Mint some initial money", async () => {
    let minterRole = await this.currencyContract.MINTER_ROLE();
    await this.currencyFactory
      .connect(this.deployer)
      .attach(this.currencyContract.address)
      .grantRole(minterRole, this.deployer.address);
    await this.currencyFactory
      .connect(this.deployer)
      .attach(this.currencyContract.address)
      .mint(this.user.address, hre.ethers.utils.parseEther("1000"));
    let balance = await this.currencyContract.balanceOf(this.user.address);
    expect(balance.toString()).to.equal(hre.ethers.utils.parseEther("1000").toString());
  });

  it("Whitelist this currency", async () => {
    await this.topupFactory
      .connect(this.deployer)
      .attach(this.topupContract.address)
      .whitelistCurrencies([this.currencyContract.address], [true]);
    let whitelistedCurrencies = await this.topupContract.getWhitelistedCurrencies();
    expect(whitelistedCurrencies.length).to.equal(1);
    expect(whitelistedCurrencies[0]).to.equal(this.currencyContract.address);
  });

  it("Top up money to platform", async () => {
    await this.currencyFactory
      .connect(this.user)
      .attach(this.currencyContract.address)
      .approve(this.topupContract.address, hre.ethers.utils.parseEther("25"));
    await this.topupFactory
      .connect(this.user)
      .attach(this.topupContract.address)
      .topup("1234", this.currencyContract.address, hre.ethers.utils.parseEther("25"));
    let platformBalance = await this.currencyContract.balanceOf(this.platform.address);
    expect(platformBalance.toString()).to.equal(hre.ethers.utils.parseEther("25"));
  });
});
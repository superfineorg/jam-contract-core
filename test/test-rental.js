require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const ORACLE = "SimpleOracle";
const RENTAL = "Rental";

before("Deploy Vesting contract", async () => {
  // Prepare parameters
  const [
    deployer
  ] = await hre.ethers.getSigners();
  this.deployer = deployer;

  this.oracleFactory = await hre.ethers.getContractFactory(ORACLE);
  this.oracleContract = await this.oracleFactory.deploy();
  await this.oracleContract.deployed();

  this.rentalFactory = await hre.ethers.getContractFactory(RENTAL);
  this.rentalContract = await this.rentalFactory.deploy(deployer.address);
  await this.oracleContract.deployed();
});

describe("Test Vesting contract", () => {
  it("", async () => {
    await this.rentalFactory
      .connect(this.deployer)
      .attach(this.rentalContract.address)
      .setOracleContract(this.oracleContract.address);
    await this.oracleFactory
      .connect(this.deployer)
      .attach(this.oracleContract.address)
      .setRate(["0x0000000000000000000000000000000000000000"], ["1000000000000"]);
    await this.rentalFactory
      .connect(this.deployer)
      .attach(this.rentalContract.address)
      .updateRentPrice(1, "0x0000000000000000000000000000000000000001", 1, "1000000");
    await this.rentalFactory
      .connect(this.deployer)
      .attach(this.rentalContract.address)
      .rentNFT(1, "0x0000000000000000000000000000000000000001", 1, 1, "0x0000000000000000000000000000000000000000", "1000000000000000000", { value: hre.ethers.utils.parseEther("1.0") });
  });
});
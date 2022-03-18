require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const JAM_ORACLE = "JamOracle";
const JAM_RENTAL = "JamRental";

before("Deploy JamRental contract", async () => {
  // Prepare parameters
  const [
    deployer
  ] = await hre.ethers.getSigners();
  this.deployer = deployer;

  this.oracleFactory = await hre.ethers.getContractFactory(JAM_ORACLE);
  this.oracleContract = await this.oracleFactory.deploy();
  await this.oracleContract.deployed();

  this.rentalFactory = await hre.ethers.getContractFactory(JAM_RENTAL);
  this.rentalContract = await this.rentalFactory.deploy(deployer.address);
  await this.oracleContract.deployed();
});

describe("Test JamRental contract", () => {
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
require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const JAM_FAUCET = "JamFaucet";

before("Deploy JamFaucet", async () => {
  // Prepare parameters
  const [deployer, client1, client2] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.client1 = client1;
  this.client2 = client2;
  this.FAUCET_AMOUNT = hre.ethers.utils.parseEther("1");

  // Deploy JamFaucet contract
  this.faucetFactory = await hre.ethers.getContractFactory(JAM_FAUCET);
  this.faucetContract = await this.faucetFactory.deploy();
  await this.faucetContract.deployed();
});

describe("Test JamFaucet", () => {
  it("Transfer some initial money to the faucet", async () => {
    let initialAmount = hre.ethers.utils.parseEther("15");
    await this.deployer.sendTransaction({
      to: this.faucetContract.address,
      value: initialAmount
    });
    let faucetBalance = await hre.waffle.provider.getBalance(this.faucetContract.address);
    expect(faucetBalance.toString()).to.equal(initialAmount);
  });

  it("Set initial parameters", async () => {
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .setFaucetAmount(this.FAUCET_AMOUNT);
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .setFaucetInterval(3);
    let faucetAmount = await this.faucetContract.faucetAmount();
    let faucetInterval = await this.faucetContract.faucetInterval();
    expect(faucetAmount.toString()).to.equal(this.FAUCET_AMOUNT);
    expect(faucetInterval.toString()).to.equal("3");
  });

  it("Faucet some money to 2 clients", async () => {
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .faucet([this.client1.address, this.client2.address]);
    let client1Balance = await this.client1.getBalance();
    let client2Balance = await this.client2.getBalance();
    expect(client1Balance.toString()).to.equal("10001000000000000000000");
    expect(client2Balance.toString()).to.equal("10001000000000000000000");
  });

  it("Pause faucet", async () => {
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .pauseFaucet();
  });

  it("Unpause faucet", async () => {
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .unpauseFaucet();
    await sleep(4000);
    await this.faucetFactory
      .connect(this.deployer)
      .attach(this.faucetContract.address)
      .faucet([this.client1.address]);
    let client1Balance = await this.client1.getBalance();
    expect(client1Balance.toString()).to.equal("10002000000000000000000");
  });
});

let sleep = ms => {
  return new Promise(resolve => setTimeout(resolve, ms));
};
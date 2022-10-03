require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");
const { soliditySha3 } = require('web3-utils');
const OPENSEA_PROXY = "contracts/tokens/ERC721/ERC721Tradable.sol:ProxyRegistry";
const WASD = "WeAllSurvivedDeath";
const WASD_MINTING = "WASDMinting";

before("Deploy WASD and WASDMinting contract", async () => {
  // Prepare parameters
  const accounts = await hre.ethers.getSigners();
  this.deployer = accounts[0];
  this.operator = accounts[1];
  this.participants = accounts.slice(2);

  // Deploy ProxyRegistry
  this.proxyFactory = await hre.ethers.getContractFactory(OPENSEA_PROXY);
  this.proxyContract = await this.proxyFactory.deploy();
  await this.proxyContract.deployed();

  // Deploy WASD
  this.wasdFactory = await hre.ethers.getContractFactory(WASD);
  this.wasdContract = await this.wasdFactory.deploy(
    "WeAllSurvivedDeath",
    "WASD",
    "https://gamejam.com/wasd/",
    this.proxyContract.address
  );
  await this.wasdContract.deployed();

  // Deploy WASD minting
  let deployedTime = await now();
  this.mintingFactory = await hre.ethers.getContractFactory(WASD_MINTING);
  this.mintingContract = await this.mintingFactory.deploy(
    this.wasdContract.address,
    10,
    deployedTime + 60 * 60,
    ["Phase 1", "Phase 2", "Phase 3"],
    [30 * 60, 60 * 60, 90 * 60],
    [4, 2, 1]
  );
  await this.mintingContract.deployed();
});

describe("Test WASD minting", () => {
  it("Check phase information", async () => {
    let phaseInfo = await this.mintingContract.getPhaseInfo();
    expect(phaseInfo.length).to.equal(3);
    expect(phaseInfo[0].metadata).to.equal("Phase 1");
    expect(phaseInfo[0].duration.toString()).to.equal("1800");
    expect(phaseInfo[0].mintLimit.toString()).to.equal("4");
    expect(phaseInfo[1].metadata).to.equal("Phase 2");
    expect(phaseInfo[1].duration.toString()).to.equal("3600");
    expect(phaseInfo[1].mintLimit.toString()).to.equal("2");
    expect(phaseInfo[2].metadata).to.equal("Phase 3");
    expect(phaseInfo[2].duration.toString()).to.equal("5400");
    expect(phaseInfo[2].mintLimit.toString()).to.equal("1");
  });

  it("Set up minter role and operators", async () => {
    // Set up the WASD minter role
    let minterRole = await this.wasdContract.MINTER_ROLE();
    await this.wasdFactory
      .connect(this.deployer)
      .attach(this.wasdContract.address)
      .grantRole(minterRole, this.mintingContract.address);
    let checkRole = await this.wasdContract.hasRole(minterRole, this.mintingContract.address);
    expect(checkRole).to.equal(true);

    // Set up operators
    await this.mintingFactory
      .connect(this.deployer)
      .attach(this.mintingContract.address)
      .setOperators([this.operator.address], [true]);
  });

  it("Operator creates minting roles and adds them to phases", async () => {
    await this.mintingFactory
      .connect(this.operator)
      .attach(this.mintingContract.address)
      .createMintingRoles(["OG", "WL"], [4, 2], [2, 1]);
    await this.mintingFactory
      .connect(this.operator)
      .attach(this.mintingContract.address)
      .addRolesToPhases([1, 2, 3], [[1, 2], [1], [0, 1, 2]]);
  });

  it("Mint before the first phase starts", async () => {
    let signature = await this.operator.signMessage(Buffer.from(soliditySha3(this.participants[1].address, 1).replaceAll("0x", ""), "hex"));
    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256,bytes,uint256)"](1, signature, 1)
    ).to.be.revertedWith("WASDMinting: personal mint limit at this phase reached");
  });

  it("Wait for the first phase to start", async () => {
    // 61 minutes later...
    await hre.network.provider.request({ method: "evm_increaseTime", params: [61 * 60] });
    await hre.network.provider.request({ method: "evm_mine", params: [] });

    // Mint without the operator's signature
    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256)"](1)
    ).to.be.revertedWith("WASDMinting: first mint - undetected role");

    // First mint with the operator's signature
    let signature = await this.operator.signMessage(Buffer.from(soliditySha3(this.participants[1].address, 1).replaceAll("0x", ""), "hex"));
    await this.mintingFactory
      .connect(this.participants[1])
      .attach(this.mintingContract.address)
    ["mintWASD(uint256,bytes,uint256)"](1, signature, 2);

    // Check the information
    let currentPhase = await this.mintingContract.getCurrentPhase();
    let participantInfo = await this.mintingContract.getParticipantInfo(this.participants[1].address);
    expect(currentPhase.toString()).to.equal("1");
    expect(participantInfo.availableMintCount.toString()).to.equal("2");
  });

  it("Participant #1 mint an NFT for himself with the signature", async () => {
    let signature = await this.operator.signMessage(Buffer.from(soliditySha3(this.participants[1].address, 1).replaceAll("0x", ""), "hex"));
    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256,bytes,uint256)"](1, signature, 2)
    ).to.be.revertedWith("WASDMinting: participant already granted minting role");
    await expect(
      this.mintingFactory
        .connect(this.participants[0])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256)"](1)
    ).to.be.revertedWith("WASDMinting: first mint - undetected role");
    await this.mintingFactory
      .connect(this.participants[1])
      .attach(this.mintingContract.address)
    ["mintWASD(uint256)"](1);
    let owner = await this.wasdContract.ownerOf(1);
    let balance = await this.wasdContract.balanceOf(this.participants[1].address);
    expect(owner).to.equal(this.participants[1].address);
    expect(balance.toString()).to.equal("3");
  });

  it("Participant #1 mint another NFT", async () => {
    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256)"](2)
    ).to.be.revertedWith("WASDMinting: personal mint limit at this phase reached");
    await this.mintingFactory
      .connect(this.participants[1])
      .attach(this.mintingContract.address)
    ["mintWASD(uint256)"](1);
    let balance = await this.wasdContract.balanceOf(this.participants[1].address);
    expect(balance.toString()).to.equal("4");
  });

  it("Community role waits to mint", async () => {
    let signature = await this.operator.signMessage(Buffer.from(soliditySha3(this.participants[12].address, 0).replaceAll("0x", ""), "hex"));
    await expect(
      this.mintingFactory
        .connect(this.participants[12])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256,bytes,uint256)"](0, signature, 1)
    ).to.be.revertedWith("WASDMinting: personal mint limit at this phase reached");

    // 91 minutes later...
    await hre.network.provider.request({ method: "evm_increaseTime", params: [91 * 60] });
    await hre.network.provider.request({ method: "evm_mine", params: [] });

    await expect(
      this.mintingFactory
        .connect(this.participants[12])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256,bytes,uint256)"](0, signature, 2)
    ).to.be.revertedWith("WASDMinting: personal mint limit at this phase reached");
    await this.mintingFactory
      .connect(this.participants[12])
      .attach(this.mintingContract.address)
    ["mintWASD(uint256,bytes,uint256)"](0, signature, 1);
    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256)"](2)
    ).to.be.revertedWith("WASDMinting: personal mint limit at this phase reached");
    let owner = await this.wasdContract.ownerOf(5);
    let balance = await this.wasdContract.balanceOf(this.participants[12].address);
    expect(owner).to.equal(this.participants[12].address);
    expect(balance.toString()).to.equal("1");
  });

  it("Minting time is over", async () => {
    // 91 minutes later...
    await hre.network.provider.request({ method: "evm_increaseTime", params: [91 * 60] });
    await hre.network.provider.request({ method: "evm_mine", params: [] });

    await expect(
      this.mintingFactory
        .connect(this.participants[1])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256)"](1)
    ).to.be.revertedWith("WASDMinting: minting time is over");

    let signature = await this.operator.signMessage(Buffer.from(soliditySha3(this.participants[14].address, 0).replaceAll("0x", ""), "hex"));
    await expect(
      this.mintingFactory
        .connect(this.participants[14])
        .attach(this.mintingContract.address)
      ["mintWASD(uint256,bytes,uint256)"](0, signature, 1)
    ).to.be.revertedWith("WASDMinting: minting time is over");
  });
});

let now = async () => {
  let blockNumber = await hre.network.provider.request({ method: "eth_blockNumber", params: [] });
  let block = await hre.network.provider.request({ method: "eth_getBlockByNumber", params: [blockNumber, false] });
  return parseInt(block.timestamp || 0, 16);
};
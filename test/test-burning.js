require('@nomiclabs/hardhat-ethers');

const hre = require('hardhat');
const { expect } = require("chai");

const GAME_NFT = "GameNFT";
const GAME_NFT_BURNING = "GameNFTBurning";

before("Deploy GameNFT and GameNFTBurning contracts", async () => {
  // Prepare parameters
  const [deployer, minter, user] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.minter = minter;
  this.user = user;

  // Deploy GameNFT contract
  this.nftFactory = await hre.ethers.getContractFactory(GAME_NFT);
  this.nftContract = await this.nftFactory.deploy(
    "Gamejam NFT",
    "GNFT",
    "https://gamejam.com/"
  );
  await this.nftContract.deployed();

  // Deploy GameNFTBurning contract
  this.nftBurningFactory = await hre.ethers.getContractFactory(GAME_NFT_BURNING);
  this.nftBurningContract = await this.nftBurningFactory.deploy();
  await this.nftBurningContract.deployed();
});

describe("Test burning contract", () => {
  it("Set up minter role", async () => {
    let minterRole = await this.nftContract.MINTER_ROLE();
    await this.nftFactory
      .connect(this.deployer)
      .attach(this.nftContract.address)
      .grantRole(minterRole, this.minter.address);
    let checkRole = await this.nftContract.hasRole(minterRole, this.minter.address);
    expect(checkRole).to.equal(true);
  });

  it("Mint an NFT to user", async () => {
    await this.nftFactory
      .connect(this.minter)
      .attach(this.nftContract.address)
      .mint(this.user.address);
    let owner = await this.nftContract.ownerOf(0);
    let tokenUri = await this.nftContract.tokenURI(0);
    expect(owner).to.equal(this.user.address);
    expect(tokenUri).to.equal("https://gamejam.com/0.json");
  });

  it("Change the base token URI", async () => {
    await expect(
      this.nftFactory
        .connect(this.minter)
        .attach(this.nftContract.address)
        .setBaseTokenURI("https://example.com/")
    ).to.be.revertedWith("GameNFT: must have admin role to set");
    await this.nftFactory
      .connect(this.deployer)
      .attach(this.nftContract.address)
      .setBaseTokenURI("https://gamejam.com/nft721/");
    let tokenUri = await this.nftContract.tokenURI(0);
    expect(tokenUri).to.equal("https://gamejam.com/nft721/0.json");
  });

  it("User approves his NFT to the burning contract and system burns it", async () => {
    await this.nftFactory
      .connect(this.user)
      .attach(this.nftContract.address)
      .approve(this.nftBurningContract.address, 0);
    await this.nftBurningFactory
      .connect(this.deployer)
      .attach(this.nftBurningContract.address)
      .burnIntoGames(this.nftContract.address, 0);
    await expect(this.nftContract.ownerOf(0)).to.be.revertedWith("ERC721: owner query for nonexistent token");
  });
});
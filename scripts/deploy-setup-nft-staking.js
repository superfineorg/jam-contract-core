const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const JAM_NFT_STAKING = "JamNFTStaking";
const ERC20 = "SimpleERC20";
const ERC721 = "SimpleERC721";
const ERC1155 = "JamNFT1155";
const TESTER_ADDR = "0x7871aa48fc61A25f444e4B3F53125FBca5AF437B";
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

let deploy = async () => {
  const [deployer] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy ERC20
  console.log(`Deploying ${ERC20} with parameters: "${deployer.address}" "100000000000000000000000000" "100000000000000000000000000" "18" "Test-JAM" "TJAM"`);
  this.erc20Factory = await hre.ethers.getContractFactory(ERC20);
  this.erc20Contract = await this.erc20Factory.deploy(
    deployer.address,
    "100000000000000000000000000",
    "100000000000000000000000000",
    18,
    "Test-JAM",
    "TJAM"
  );
  await this.erc20Contract.deployed();
  deployInfo[this.networkName][ERC20] = this.erc20Contract.address;

  // Deploy ERC721
  console.log(`Deploying ${ERC721} with parameters: "${TESTER_ADDR}" "Gamejam-Awesome-NFT" "JamNFT" "https://gamejam.com/nft721/"`);
  this.erc721Factory = await hre.ethers.getContractFactory(ERC721);
  this.erc721Contract = await this.erc721Factory.deploy(
    TESTER_ADDR,
    "Gamejam-Awesome-NFT",
    "JamNFT",
    "https://gamejam.com/nft721/"
  );
  await this.erc721Contract.deployed();
  deployInfo[this.networkName][ERC721] = this.erc721Contract.address;

  // Deploy ERC1155
  console.log(`Deploying ${ERC1155} with parameters: "https://gamejam.com/nft1155/`);
  this.erc1155Factory = await hre.ethers.getContractFactory(ERC1155);
  this.erc1155Contract = await this.erc1155Factory.deploy("https://gamejam.com/nft1155/");
  await this.erc1155Contract.deployed();
  deployInfo[this.networkName][ERC1155] = this.erc1155Contract.address;

  // Deploy JamNFTStaking
  console.log(`Deploying ${JAM_NFT_STAKING} with parameters: "${300}" "${"2000000000000000000"}" "${ZERO_ADDRESS}"`);
  this.nftStakingFactory = await hre.ethers.getContractFactory(JAM_NFT_STAKING);
  this.nftStakingContract = await this.nftStakingFactory.deploy(300, "2000000000000000000", ZERO_ADDRESS);
  await this.nftStakingContract.deployed();
  deployInfo[this.networkName][JAM_NFT_STAKING] = this.nftStakingContract.address;

  // Write the result to deploy.json
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
};

let setup = async () => {
  // Set operator role for tester
  console.log(`Setting operator role for ${TESTER_ADDR}...`);
  await this.nftStakingFactory
    .connect(this.deployer)
    .attach(this.nftStakingContract.address)
    .setOperators([TESTER_ADDR], [true]);

  // Transfer some initial fund to the contract
  console.log("Transferring some initial fund to the JamNFTStaking contract...");
  await this.deployer.sendTransaction({
    to: this.nftStakingContract.address,
    value: hre.ethers.utils.parseEther("80")
  });
  await this.erc20Factory
    .connect(this.deployer)
    .attach(this.erc20Contract.address)
    .transfer(this.nftStakingContract.address, "50000000000000000000000000");

  // Whitelist the ERC721 and ERC1155 tokens
  console.log(`Whitelisting ${this.erc721Contract.address} and ${this.erc1155Contract.address}...`);
  await this.nftStakingFactory
    .connect(this.deployer)
    .attach(this.nftStakingContract.address)
    .whitelistNFT([this.erc721Contract.address, this.erc1155Contract.address], [0, 1]);

  // Set minter role of JamNFT1155 to tester
  console.log(`Setting minter role of JamNFT1155 to ${TESTER_ADDR}`);
  let minterRole = await this.erc1155Contract.MINTER_ROLE();
  await this.erc1155Factory
    .connect(this.deployer)
    .attach(this.erc1155Contract.address)
    .grantRole(minterRole, TESTER_ADDR);
};

let main = async () => {
  await deploy();
  await setup();
};

main();
const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const NFT_STAKING = "NFTStaking";
const ERC721 = "SimpleERC721";
const ERC1155 = "SimpleERC1155";
const TESTER_ADDR = "0x7871aa48fc61A25f444e4B3F53125FBca5AF437B";

let deploy = async () => {
  const [deployer] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy ERC721
  console.log(`Deploying ${ERC721} with parameters: "${TESTER_ADDR}" "Gamejam Awesome NFT" "JamNFT" "https://xxx.com/"`);
  this.erc721Factory = await hre.ethers.getContractFactory(ERC721);
  this.erc721Contract = await this.erc721Factory.deploy(
    TESTER_ADDR,
    "Gamejam Awesome NFT",
    "JamNFT",
    "https://xxx.com/"
  );
  await this.erc721Contract.deployed();
  deployInfo[this.networkName][ERC721] = this.erc721Contract.address;

  // Deploy ERC1155
  console.log(`Deploying ${ERC1155} with no parameter`);
  this.erc1155Factory = await hre.ethers.getContractFactory(ERC1155);
  this.erc1155Contract = await this.erc1155Factory.deploy();
  await this.erc1155Contract.deployed();
  deployInfo[this.networkName][ERC1155] = this.erc1155Contract.address;

  // Deploy NFTStaking
  console.log(`Deploying ${NFT_STAKING} with parameters: "${300}" "${"2000000000000000000"}"`);
  this.nftStakingFactory = await hre.ethers.getContractFactory(NFT_STAKING);
  this.nftStakingContract = await this.nftStakingFactory.deploy(300, "2000000000000000000");
  await this.nftStakingContract.deployed();
  deployInfo[this.networkName][NFT_STAKING] = this.nftStakingContract.address;

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
  console.log("Transferring some initial fund to the NFTStaking contract...");
  await this.deployer.sendTransaction({
    to: this.nftStakingContract.address,
    value: hre.ethers.utils.parseEther("80")
  });

  // Whitelist the ERC721 and ERC1155 tokens
  console.log(`Whitelisting ${this.erc721Contract.address} and ${this.erc1155Contract.address}...`);
  await this.nftStakingFactory
    .connect(this.deployer)
    .attach(this.nftStakingContract.address)
    .whitelistNFT(
      [
        this.erc721Contract.address,
        this.erc1155Contract.address,
        this.erc1155Contract.address,
        this.erc1155Contract.address,
        this.erc1155Contract.address
      ],
      [0, 1, 1, 1, 1],
      [0, 1, 3, 4, 6]
    );

  // Transfer ownership of ERC1155 contract to tester
  console.log(`Transferring ownership of ERC1155 token to ${TESTER_ADDR}`);
  await this.erc1155Factory
    .connect(this.deployer)
    .attach(this.erc1155Contract.address)
    .transferOwnership(TESTER_ADDR);
};

let main = async () => {
  await deploy();
  await setup();
};

main();
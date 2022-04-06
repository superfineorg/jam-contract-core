const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const ERC721 = "SimpleERC721";

let deploy = async () => {
  const [deployer] = await hre.ethers.getSigners();
  this.deployer = deployer;
  this.networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy Simple ERC721
  console.log(`Deploying ${ERC721} with parameters: "${deployer.address}" "Gamejam-Awesome-NFT" "JamNFT" "https://gamejam.com/nft721/"`);
  this.erc721Factory = await hre.ethers.getContractFactory(ERC721);
  this.erc721Contract = await this.erc721Factory.deploy(
    deployer.address,
    "Gamejam-Awesome-NFT",
    "JamNFT",
    "https://gamejam.com/nft721/"
  );
  await this.erc721Contract.deployed();
  deployInfo[this.networkName][ERC721] = this.erc721Contract.address;

  // Write the result to deploy.json
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
};

deploy();
const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const GAME_NFT_721 = "GameNFT721";
const GAME_NFT_1155 = "GameNFT1155";
const GAME_NFT_BURNING = "GameNFTBurning";

async function deploy() {
  // Info
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy GameNFT721 contract
  console.log(`${GAME_NFT_721}: "Gamejam NFT" "GNFT" "https://gamejam.com/nft721/"`);
  let nft721Factory = await hre.ethers.getContractFactory(GAME_NFT_721);
  let nft721Contract = await nft721Factory.deploy(
    "Gamejam NFT",
    "GNFT",
    "https://gamejam.com/nft721/"
  );
  await nft721Contract.deployed();

  // Deploy GameNFT1155 contract
  console.log(`${GAME_NFT_1155}: "https://gamejam.com/nft1155/"`);
  let nft1155Factory = await hre.ethers.getContractFactory(GAME_NFT_1155);
  let nft1155Contract = await nft1155Factory.deploy("https://gamejam.com/nft1155/");
  await nft1155Contract.deployed();

  // Deploy GameNFTBurning contract
  console.log(`${GAME_NFT_BURNING}`);
  let nftBurningFactory = await hre.ethers.getContractFactory(GAME_NFT_BURNING);
  let nftBurningContract = await nftBurningFactory.deploy();
  await nftBurningContract.deployed();

  // Write the result to deploy.json
  deployInfo[networkName][GAME_NFT_721] = nft721Contract.address;
  deployInfo[networkName][GAME_NFT_1155] = nft1155Contract.address;
  deployInfo[networkName][GAME_NFT_BURNING] = nftBurningContract.address;
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
}

deploy();
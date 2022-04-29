const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const GAME_NFT = "GameNFT";
const GAME_NFT_BURNING = "GameNFTBurning";

async function deploy() {
  // Info
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy GameNFT contract
  let nftFactory = await hre.ethers.getContractFactory(GAME_NFT);
  let nftContract = await nftFactory.deploy(
    "Gamejam NFT",
    "GNFT",
    "https://gamejam.com/"
  );
  await nftContract.deployed();

  // Deploy GameNFTBurning contract
  let nftBurningFactory = await hre.ethers.getContractFactory(GAME_NFT_BURNING);
  let nftBurningContract = await nftBurningFactory.deploy();
  await nftBurningContract.deployed();

  // Write the result to deploy.json
  deployInfo[networkName][GAME_NFT] = nftContract.address;
  deployInfo[networkName][GAME_NFT_BURNING] = nftBurningContract.address;
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
}

deploy();
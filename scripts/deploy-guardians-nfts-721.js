const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const CONTRACT_NAME = "GamejamNFT721";
const PROXY_REGISTRY_ADDRESS = {
  "rinkeby": "0xf57b2c51ded3a29e6891aba85459d600256cf317",
  "eth": "0xa5409ec958C83C3f309868babACA7c86DCB077c1",
  "jamchaintestnet": "0x0000000000000000000000000000000000000001"
};

async function deploy() {
  // Deploy
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());
  const factory = await hre.ethers.getContractFactory(CONTRACT_NAME);
  console.log(`Deploying ${CONTRACT_NAME} with parameters: "Gamejam Main NFT" "JamNFT" "https://gamejam.com/nft/" "${PROXY_REGISTRY_ADDRESS[networkName]}" "${20}"`);
  const contract = await factory.deploy(
    "Gamejam Main NFT",
    "JamNFT",
    "https://gamejam.com/nft/",
    PROXY_REGISTRY_ADDRESS[networkName],
    20
  );
  await contract.deployed();
  console.log(`${CONTRACT_NAME} deployed address: ${contract.address}`);

  // Write the result to deploy.json
  deployInfo[networkName][CONTRACT_NAME] = contract.address;
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
}

deploy();
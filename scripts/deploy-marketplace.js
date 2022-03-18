const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const CONTRACT_NAME = "JamMarketplace";

async function deploy() {
  // Deploy
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());
  const factory = await hre.ethers.getContractFactory(CONTRACT_NAME);
  console.log(`Deploying ${CONTRACT_NAME} with parameters: "300"`);
  const contract = await factory.deploy(300);
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
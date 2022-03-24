const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const CONTRACT_NAME = "GuardianOfGloryItems";
const OWNER = "0x4d50Edd2273cB1f666d53f7a15785f4F79ea0EAA";
const MINTER = "0x03B301BEA83eef1eB74136a57Ed59F9724387F79";

async function deploy() {
  // Deploy
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());
  const factory = await hre.ethers.getContractFactory(CONTRACT_NAME);
  console.log(`Deploying ${CONTRACT_NAME} with parameters: "${OWNER}" "${MINTER}" "https://assets.gamejam.co/platform/GuardianGlory/nft/items/"`);
  const contract = await factory.deploy(OWNER, MINTER, "https://assets.gamejam.co/platform/GuardianGlory/nft/items/");
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
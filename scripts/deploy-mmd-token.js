const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const TOKEN = "MMDToken";

async function deploy() {
  const networkName = hre.network.name;
  let factory = await hre.ethers.getContractFactory(TOKEN);
  let contract = await factory.deploy("MMDToken", "MMCoin", "1000000000000000000000000000000");
  await contract.deployed();

  // Write the result to deploy.json
  deployInfo[networkName][TOKEN] = contract.address;
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
}

deploy();
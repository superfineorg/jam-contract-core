const hre = require("hardhat");

const CONTRACT_NAME = "Namfungible";

let deploy = async () => {
  let factory = await hre.ethers.getContractFactory(CONTRACT_NAME);
  let contract = await factory.deploy("Namfungible", "NamNFT", "https://gamejam.com/", "0xf57b2c51ded3a29e6891aba85459d600256cf317");
  await contract.deployed();
  console.log(contract.address);
};

deploy();
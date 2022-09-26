const hre = require("hardhat");
const deployInfo = require("../../deploy.json");

const SEND_BATCH = "PlaylinkAirdrop";
const ERC20 = "GameToken";
const ERC721 = "GameNFT721";
const ERC1155 = "GameNFT1155";

let setup = async () => {
  // Info
  const [deployer] = await hre.ethers.getSigners();
  const OPERATOR = "0x28Dcc538F79525d713DFb3566C8507cE6BaA9e70";
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Prepare contracts
  let batchSendFactory = await hre.ethers.getContractFactory(SEND_BATCH);
  let batchSendContract = batchSendFactory.attach(deployInfo[networkName][SEND_BATCH]);
  let erc20Factory = await hre.ethers.getContractFactory(ERC20);
  let erc20Contract = erc20Factory.attach(deployInfo[networkName][ERC20]);
  let erc721Factory = await hre.ethers.getContractFactory(ERC721);
  let erc721Contract = erc721Factory.attach(deployInfo[networkName][ERC721]);
  let erc1155Factory = await hre.ethers.getContractFactory(ERC1155);
  let erc1155Contract = erc1155Factory.attach(deployInfo[networkName][ERC1155]);

  // Set minter role
  console.log("Set minter roles...");
  let minterRole = await erc20Contract.MINTER_ROLE();
  await erc20Factory
    .connect(deployer)
    .attach(erc20Contract.address)
    .grantRole(minterRole, OPERATOR);
  await erc721Factory
    .connect(deployer)
    .attach(erc721Contract.address)
    .grantRole(minterRole, OPERATOR);
  await erc1155Factory
    .connect(deployer)
    .attach(erc1155Contract.address)
    .grantRole(minterRole, OPERATOR);

  // Set up operators
  console.log("Set up operators...");
  await batchSendFactory
    .connect(deployer)
    .attach(batchSendContract.address)
    .setOperators([OPERATOR], [true]);
};

setup();
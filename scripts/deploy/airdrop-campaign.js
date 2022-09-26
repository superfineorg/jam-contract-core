const hre = require("hardhat");
const fs = require("fs");
const deployInfo = require("../../deploy.json");

const AIRDROP = "PlaylinkAirdrop";
const ERC20 = "GameToken";
const ERC721 = "GameNFT721";
const ERC1155 = "GameNFT1155";

const CONTRACTS_INFO_PATH = "./contracts-info";

let updateAddresses = () => {
  return [
    {
      name: AIRDROP,
      constructorArgs: [20, hre.ethers.utils.parseEther("0.01").toString()]
    },
    {
      name: ERC20,
      constructorArgs: []
    },
    {
      name: ERC721,
      constructorArgs: ["Qwerty", "QWERTY", "ipfs://QmQXLbHuubm832k7vr1DfznAuxaqUsoiA7TSX2GfjF46AJ/"]
    },
    {
      name: ERC1155,
      constructorArgs: ["ipfs://QmQXLbHuubm832k7vr1DfznAuxaqUsoiA7TSX2GfjF46AJ/"]
    }
  ];
};

async function deploy() {
  // Info
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy and compute encoded constructor arguments
  let encodedArgs = {};
  let contracts = updateAddresses(deployInfo[networkName]);
  for (let i = 0; i < contracts.length; i++) {
    // Deploy
    console.log(`Deploying ${contracts[i].name}: ${contracts[i].constructorArgs.map(arg => `"${arg}"`).join(" ")}`);
    let factory = await hre.ethers.getContractFactory(contracts[i].name);
    let contract = await factory.deploy(...contracts[i].constructorArgs);
    await contract.deployed();
    deployInfo[networkName][contracts[i].name] = contract.address;
    contracts = updateAddresses(deployInfo[networkName]);

    // Compute encoded constructor arguments
    encodedArgs[contracts[i].name] = factory
      .getDeployTransaction(...contracts[i].constructorArgs)
      .data
      .replace(factory.bytecode, "");
  }

  // Save the results
  if (!fs.existsSync(CONTRACTS_INFO_PATH))
    fs.mkdirSync(CONTRACTS_INFO_PATH);
  contracts.forEach(contract => {
    let infoFolder = `${CONTRACTS_INFO_PATH}/${getFolderName(contract.name)}`;
    if (!fs.existsSync(infoFolder))
      fs.mkdirSync(infoFolder);
    fs.writeFileSync(`${infoFolder}/constructor-args.txt`, encodedArgs[contract.name]);
  });
  fs.writeFileSync("deploy.json", JSON.stringify(deployInfo, null, "\t"));

  console.log("Finish!");
}

let getFolderName = contractName => {
  let splits = contractName.split("/");
  return splits[splits.length - 1].split(".")[0];
};

deploy();
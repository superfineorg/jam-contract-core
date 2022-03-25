const hre = require("hardhat");
const FileSystem = require("fs");
const deployInfo = require("../deploy.json");

const OG_PASS = "JamOGPass";
const SHF = "JamSuperHappyFrens";
const MINTING = "JamOGPassMinting";

async function deploy() {
  // Info
  const [deployer] = await hre.ethers.getSigners();
  const networkName = hre.network.name;
  console.log("Deployer:", deployer.address);
  console.log("Balance:", (await deployer.getBalance()).toString());

  // Deploy OGPass
  console.log(`Deploying ${OG_PASS} with parameters: "OGPass NFT" "OGP" "https://gamejam.com/og-pass/" "0x0000000000000000000000000000000000000022" "2000"`);
  let ogPassFactory = await hre.ethers.getContractFactory(OG_PASS);
  let ogPassContract = await ogPassFactory.deploy(
    "OGPass NFT",
    "OGP",
    "https://gamejam.com/og-pass/",
    "0x0000000000000000000000000000000000000022",
    2000
  );
  await ogPassContract.deployed();
  console.log(`${OG_PASS} has been deployed at: ${ogPassContract.address}`);

  // Deploy SuperHapprFrens
  console.log(`Deploying ${SHF} with parameters: "Super Happy Frens NFT" "SHF" "https://gamejam.com/super-happy-frens/" "0x0000000000000000000000000000000000000034" "2000"`);
  let shfFactory = await hre.ethers.getContractFactory(SHF);
  let shfContract = await shfFactory.deploy(
    "Super Happy Frens NFT",
    "SHF",
    "https://gamejam.com/super-happy-frens/",
    "0x0000000000000000000000000000000000000034",
    2000
  );
  await shfContract.deployed();
  console.log(`${SHF} has been deployed at ${shfContract.address}`);

  // Deploy Minting
  console.log(`Deploying ${MINTING} with parameters: "${ogPassContract.address}" "${shfContract.address}" "1234" "3333" "2222"`);
  let mintingFactory = await hre.ethers.getContractFactory(MINTING);
  let mintingContract = await mintingFactory.deploy(
    ogPassContract.address,
    shfContract.address,
    1234,
    3333,
    2222
  );
  await mintingContract.deployed();
  console.log(`${MINTING} has been deployed at ${mintingContract.address}`);

  // Write the result to deploy.json
  deployInfo[networkName][OG_PASS] = ogPassContract.address;
  deployInfo[networkName][MINTING] = mintingContract.address;
  deployInfo[networkName][SHF] = shfContract.address;
  FileSystem.writeFile("deploy.json", JSON.stringify(deployInfo, null, "\t"), err => {
    if (err)
      console.log("Error when trying to write to deploy.json!", err);
    else
      console.log("Information has been written to deploy.json!");
  });
}

deploy();
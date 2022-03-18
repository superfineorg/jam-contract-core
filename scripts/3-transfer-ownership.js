const hre = require("hardhat");
const deployInfo = require("../deploy.json");

const JAM_NFT_STAKING = "JamNFTStaking";
const JAM_VESTING = "JamVesting";

async function transferOwnership() {
  // Declaration
  const [deployer] = await hre.ethers.getSigners();
  const nftStakingNewOwner = "0x7871aa48fc61A25f444e4B3F53125FBca5AF437B";
  const vestingNewOwner = "0xEd0bC1D60a1a58630dfab759150F9b35E2f8e6aC";
  let nftStakingFactory = await hre.ethers.getContractFactory(JAM_NFT_STAKING);
  let vestingFactory = await hre.ethers.getContractFactory(JAM_VESTING);
  const networkName = hre.network.name;

  // Set operators
  await nftStakingFactory
    .connect(deployer)
    .attach(deployInfo[networkName]?.JamNFTStaking)
    .setOperators([nftStakingNewOwner], [true]);
  await vestingFactory
    .connect(deployer)
    .attach(deployInfo[networkName]?.JamVesting)
    .setOperators([vestingNewOwner], [true]);

  // Transfer ownership
  await nftStakingFactory
    .connect(deployer)
    .attach(deployInfo[networkName]?.JamNFTStaking)
    .transferOwnership(nftStakingNewOwner);
  await vestingFactory
    .connect(deployer)
    .attach(deployInfo[networkName]?.JamVesting)
    .transferOwnership(vestingNewOwner);
}

transferOwnership();
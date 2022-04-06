const hre = require("hardhat");
const deployInfo = require("../deploy.json");
const { soliditySha3 } = require('web3-utils');

const ERC721 = "SimpleERC721";
const ERC1155 = "JamNFT1155";
const TESTERS = [
  "0xE393F47642359536F05446Ae722C13dAA0119b4e",
  "0xEd6922a7065Ad1E7aAA34baA80828796cA011C3d",
  "0x2e39990D3B9FA56aee206f38cbE12eef02651366",
  "0x681dDF7E581bd61c85111b8e76A788692bBBaA32",
  "0xC3aBA23f1377C764b503CAF7EE3100ADfe669137",
  "0x0797fACB87725022f22bE906E917e3a872D55541",
  "0x831C2C057762B6633A2010e695D2d0E3B7b9A199",
  "0x2460593cB6e6D0fF490fD2EA1F212236c8c9a7B6",
  "0x4aF36Ca461987a498F269663EBD376D1832Fe0F1",
  "0x8dD4478C0f033A72751084b476138b01fAAD24B3",
  "0xec4E6478A0739b3c3700b1Df04A75f36BC964a48",
  "0xf0CA075fd1E2F37d6c1a6204079489e9f438865c",
  "0xC74EC9C1C6461554Acd21622aE8a228c3B17826F",
  "0x9E7B187F052a101f138F15f722B828C9fcFDd3d8",
  "0xa2d4FeC0Fc927865Ad89f8e59A66CA9F348D12C2",
  "0xD31041c4D8611C023Bb16E1da3C6582DAEF0dB65",
  "0xeBC7d806cdD76bd3792e901B7f9CC84f3Ea147dB"
];

let main = async () => {
  const [deployer] = await hre.ethers.getSigners();

  let erc721Factory = await hre.ethers.getContractFactory(ERC721);
  let erc1155Factory = await hre.ethers.getContractFactory(ERC1155);

  for (let i = 0; i < TESTERS.length; i++) {
    // Send some JAM tokens
    console.log(`Sending 1 JAM to ${TESTERS[i]}...`);
    await deployer.sendTransaction({ to: TESTERS[i], value: hre.ethers.utils.parseEther("1.0") });

    // Mint some ERC721 NFTs
    console.log(`Minting an ERC721 NFT to ${TESTERS[i]}...`);
    await erc721Factory
      .connect(deployer)
      .attach(deployInfo.jamchaintestnet[ERC721])
      .awardItem(TESTERS[i]);

    // Mint some ERC1155 NFTs
    console.log(`Minting an ERC1155 NFT to ${TESTERS[i]}...`);
    await erc1155Factory
      .connect(deployer)
      .attach(deployInfo.jamchaintestnet[ERC1155])
      .mint(TESTERS[i], 10000 + i, 100, soliditySha3(`#${10000 + i}`));
  }
};

main();
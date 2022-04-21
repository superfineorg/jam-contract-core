require("dotenv").config();
require("@nomiclabs/hardhat-waffle");
require('@nomiclabs/hardhat-ethers');
require("@nomiclabs/hardhat-etherscan");
require('@openzeppelin/hardhat-upgrades');
// require('@openzeppelin/hardhat-defender');
require("@nomiclabs/hardhat-web3");

const ETHEREUM_PROVIDER = process.env.ETHEREUM_PROVIDER;
const BSC_PROVIDER = process.env.BSC_PROVIDER;
const JAMCHAIN_PROVIDER = process.env.JAMCHAIN_PROVIDER;
const RINKEBY_PROVIDER = process.env.RINKEBY_PROVIDER;
const GOERLI_PROVIDER = process.env.GOERLI_PROVIDER;
const BSC_TESTNET_PROVIDER = process.env.BSC_TESTNET_PROVIDER;
const JAMCHAIN_TESTNET_PROVIDER = process.env.JAMCHAIN_TESTNET_PROVIDER;
const POLYGON_TESTNET_PROVIDER = process.env.POLYGON_TESTNET_PROVIDER;
const BSC_API_KEY = process.env.BSC_API_KEY;
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY;
const POLYGON_API_KEY = process.env.POLYGON_API_KEY;
const ADDRESS = process.env.ADDRESS;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const MNEMONIC = process.env.MNEMONIC;

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();
  for (const account of accounts) {
    console.log(account.address);
  }
});

module.exports = {
  defaultNetwork: "hardhat",
  networks: {
    eth: {
      url: ETHEREUM_PROVIDER,
      chainId: 1,
      gas: 8000000,
      accounts: [PRIVATE_KEY]
    },
    bsc: {
      url: BSC_PROVIDER,
      chainId: 56,
      gas: 8000000,
      accounts: [PRIVATE_KEY]
    },
    jamchain: {
      url: JAMCHAIN_PROVIDER,
      chainId: 2077,
      gas: 5500000,
      accounts: [PRIVATE_KEY]
    },
    rinkeby: {
      url: RINKEBY_PROVIDER,
      chainId: 4,
      gas: 5500000,
      accounts: [PRIVATE_KEY]
    },
    goerli: {
      url: GOERLI_PROVIDER,
      chainId: 5,
      gas: 5500000,
      accounts: [PRIVATE_KEY]
    },
    bsctestnet: {
      url: BSC_TESTNET_PROVIDER,
      chainId: 97,
      gas: 8000000,
      accounts: [PRIVATE_KEY],
      networkCheckTimeout: 1000000000,
      timeoutBlocks: 200000000,
      skipDryRun: true,
    },
    jamchaintestnet: {
      url: JAMCHAIN_TESTNET_PROVIDER,
      chainId: 2710,
      gas: 5500000,
      accounts: [PRIVATE_KEY]
    },
    polygontestnet: {
      url: POLYGON_TESTNET_PROVIDER,
      chainId: 80001,
      gas: 5500000,
      accounts: [PRIVATE_KEY]
    },
    hardhat: {},
    localhost: {
      url: "http://127.0.0.1:7545",
      accounts: [PRIVATE_KEY],
      gasLimit: 6000000000,
      defaultBalanceEther: 10,
    }
  },
  etherscan: {
    apiKey: ETHERSCAN_API_KEY
  },
  solidity: {
    version: "0.8.6",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  }
};

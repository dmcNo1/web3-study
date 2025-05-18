require("@nomicfoundation/hardhat-toolbox");
// require("dotenv").config();
require("@chainlink/env-enc").config();
require("hardhat-deploy")
// require("./tasks/deploy-fundme")
// require("./tasks/interact-fundme")
require("./tasks")

const SEPOLIA_URL = process.env.SEPOLIA_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const PRIVATE_KEY_2 = process.env.PRIVATE_KEY_2;
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY;

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  defaultNetwork: "hardhat",
  networks: {
    sepolia: {
      url: SEPOLIA_URL,
      accounts: [
        PRIVATE_KEY,
        PRIVATE_KEY_2
      ],
      chainId: 11155111,
      timeout: 1000000
    },
    hardhat: {
      timeout: 1000000
    }
  },
  etherscan: {
    apiKey: {
      sepolia: ETHERSCAN_API_KEY
    }
  },
  namedAccounts: {
    firstAccount: {
      default: 0  // 获取accounts数组的第0号元素
    },
    secondAccount: {
      default: 1
    }
  }
};

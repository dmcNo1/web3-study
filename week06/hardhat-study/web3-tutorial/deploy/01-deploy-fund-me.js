// function deployFunction() {
//     console.log("this is a deploy function");
// }

const { network } = require("hardhat");
const { developmentChains, networkConfig, CONFIRMATIONS } = require("../helper-hardhat-config");

// module.exports.default = deployFunction;

// module.exports = async(hre) => {
//     const getNamedAccounts = hre.getNamedAccounts;
//     const deployments = hre.deployments;
//     console.log("this is a deploy function");
// }

module.exports = async ({ getNamedAccounts, deployments }) => {
    console.log("this is a deploy function")
    const firstAccount = (await getNamedAccounts()).firstAccount;

    let mockDataFeedAddr;
    let confirmations;
    if (developmentChains.includes(network.name)) {
        const mockV3Aggregator = await deployments.get("MockV3Aggregator");
        mockDataFeedAddr = mockV3Aggregator.address;
        confirmations = 0;
    } else {
        mockDataFeedAddr = "";
        confirmations = CONFIRMATIONS;
    }

    
    const { deploy } = deployments;
    const fundMe = await deploy("FundMe", {
        from: firstAccount,
        args: [180, mockDataFeedAddr],
        log: true,
        waitConfirmations: confirmations,
    });

    if (hre.network.config.chainId === 11155111 && process.env.ETHERSCAN_API_KEY) {
        // verify
    } else {
        console.log("network is not sepolia, verify is skipped");
    }
}

module.exports.tags = ["all", "fundMe"];
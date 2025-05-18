// function deployFunction() {
//     console.log("this is a deploy function");
// }

const { network } = require("hardhat");
const { developmentChains, networkConfig } = require("../helper-hardhat-config");

// module.exports.default = deployFunction;

// module.exports = async(hre) => {
//     const getNamedAccounts = hre.getNamedAccounts;
//     const deployments = hre.deployments;
//     console.log("this is a deploy function");
// }

module.exports = async ({ getNamedAccounts, deployments }) => {
    const firstAccount = (await getNamedAccounts()).firstAccount;
    const { deploy } = deployments;
    let dataFeedAddr;
    if (developmentChains.includes(network.name)) {
        const mockV3Aggregator = await deployments.get("MockV3Aggregator")
        dataFeedAddr = mockV3Aggregator.address;
    } else {
        dataFeedAddr = networkConfig[network.config.chainId].ethUsdDataFeed;
    }
    // 部署合约
    await deploy("FundMe", {
        from: firstAccount,
        args: [180, mockDataAddress.address],
        log: true
    })
    console.log(`first account is ${firstAccount}`);
    console.log("this is a deploy function");
}

module.exports.tags = ["all", "fundMe"];
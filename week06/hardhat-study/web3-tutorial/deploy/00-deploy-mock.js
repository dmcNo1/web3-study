const { getNamedAccounts, deployments } = require("hardhat");
const { DECIMAL, INITIAL_ANSWER, developmentChains} = require("../helper-hardhat-config");

module.exports = async({getNamedAccounts, deployments}) => {
    const {firstAccount} = await getNamedAccounts();
    const {deploy} = deployments;
    if (developmentChains.includes(network.name)) {
        await deploy("MockV3Aggregator", {
            from: firstAccount,
            args: [DECIMAL, INITIAL_ANSWER],
            log: true
        })
    } else {
        console.log("environment is not local, mock skip");
    }
}

// 如果是npx hardhat deploy --tags mock或者npx hardhat deploy --tags all的话，就会执行这个部署脚本
module.exports.tags = ["all", "mock"]
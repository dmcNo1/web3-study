const { task } = require("hardhat/config");

// 定义一个task
task("deploy-fundme").setAction(async(taskArgs, hre) => {
    // 创建一个合约工厂
    const fundMeFactory = await ethers.getContractFactory("FundMe");
    // 部署合约，如果没有配置，默认是运行在本地环境
    const fundMe = await fundMeFactory.deploy(1200);
    await fundMe.waitForDeployment();
    console.log("contract has been deployed successfully, contract address: ", fundMe.target);

    // verify fundMe
    // 如果失败了，vpn记得打开tun模式
    console.log(hre.network.config.chainId);
    if (hre.network.config.chainId == 11155111 && process.env.ETHERSCAN_API_KEY) {
        console.log("Waiting for 5 blocks to verify...");
        // 等待五个区块
        await fundMe.deploymentTransaction().wait(5);
        console.log("Verifying contract...");
        // await verifyFundMe(fundMe.target, [3000]);
    } else {
        console.log("Etherscan verification is not needed");
    }
})

async function verifyFundMe(fundMeAddress, args) {
    await hre.run("verify:verify", {
        address: fundMeAddress,
        constructorArguments: args,
    });
}

module.exports = {}
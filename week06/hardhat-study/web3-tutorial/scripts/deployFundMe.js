const {ethers} = require("hardhat");    // 引入ether.js依赖

// 创建一个main方法
async function main() {
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

    // init 2 accounts，从配置中读取accounts
    const [firstAccount, secondAccount] = await ethers.getSigners();

    // fund contract with first account
    const fundTx = await fundMe.fund({value: ethers.parseEther("0.5")});
    await fundTx.wait()

    // check balance of contract
    const balanceOfContract = await ethers.provider.getBalance(fundMe.target);
    console.log(`Balance of contract: ${balanceOfContract}`);

    // fund contract with second account，如果不加上connect，默认是用第一个账户
    const fundTx2 = await fundMe.connect(secondAccount).fund({value: ethers.parseEther("0.2")});
    await fundTx2.wait();

    // check balance of contract
    const balanceOfContract2 = await ethers.provider.getBalance(fundMe.target);
    console.log(`Balance of contract: ${balanceOfContract2}`);

    // check mapping
    const accountBalance = await fundMe.fundersToAmount(firstAccount.address);
    const accountBalance2 = await fundMe.fundersToAmount(secondAccount.address);
    console.log(`Account balance: ${firstAccount.address}-${accountBalance}, ${secondAccount.address}-${accountBalance2}`);
}

async function verifyFundMe(fundMeAddress, args) {
    await hre.run("verify:verify", {
        address: fundMeAddress,
        constructorArguments: args,
    });
}

// 调用main方法
main().then().catch((error) => {
    console.error(error);
    process.exit(0);
});
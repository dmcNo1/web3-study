const { task } = require("hardhat/config");

task("interact-fundme")
    .addParam("address", "fundme contract address")
    .setAction(async(taskArgs, hre) => {
        // 获取合约
        const fundmeFactory = await ethers.getContractFactory("FundMe");
        const fundme = await fundmeFactory.attach(taskArgs.address);

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
})

module.exports = {}
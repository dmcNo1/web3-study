const {ethers} = require("hardhat");

async function main() {
    // create factory
    const fundMeFactory = await ethers.getContractFactory("FundMe");
    // deploy contract from factory
    const fundMe = await fundMeFactory.deploy(300);
    await fundMe.waitForDeployment();
    console.log("contract has been deployed successfully, contract address: ", fundMe.target);

    // verify fundMe
    // 如果失败了，vpn记得打开tun模式
    console.log(hre.network.config.chainId);
    if (hre.network.config.chainId == 11155111 && process.env.ETHERSCAN_API_KEY) {
        console.log("Waiting for 5 blocks to verify...");
        await fundMe.deploymentTransaction().wait(5);
        await verifyFundMe(fundMe.target, [300]);
    } else {
        console.log("Etherscan verification is not needed");
    }

    // init 2 accounts，从配置中读取accounts
    const [firstAccount, secondAccount] = await ethers.getSigners();

    // fund contract with first account
    const fundTx = await fundMe.fund({value: ethers.parseEther("0.001")});
    await fundTx.wait();

    // check balance of contract
    const balanceOfContract = ethers.provider.getBalance(fundMe.target);
    console.log(`Balance of contract: ${balanceOfContract}`);

    // fund contract with second account，如果不加上connect，默认是用第一个账户
    const fundTx2 = await fundMe.connect(secondAccount).fund({value: ethers.parseEther("0.0001")});
    await fundTx2.wait();

    // check balance of contract
    const balanceOfContract2 = ethers.provider.getBalance(fundMe.target);
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

main().then().catch((error) => {
    console.error(error);
    process.exit(1);
})
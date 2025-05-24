const { ethers, deployments, getNamedAccounts } = require('hardhat');
const { assert, expect } = require('chai');
const helpers = require('@nomicfoundation/hardhat-network-helpers');
const { developmentChains } = require('../../helper-hardhat-config');

// 集成测试，只有在部署到测试网络才会执行
developmentChains.includes(network.name)
    ? describe.skip
    : describe("test FundMe contract", async () => {
        let fundMe;
        let firstAccount;

        beforeEach(async () => {
            await deployments.fixture(['all']);
            firstAccount = (await getNamedAccounts()).firstAccount;
            const fundMeDeployment = await deployments.get('FundMe');
            fundMe = await ethers.getContractAt("FundMe", fundMeDeployment.address);
        });

        it("fund and getFund successfully", async () => {
            // make sure target reached
            await fundMe.fund({ value: ethers.parseEther("0.5") });
            // 等待时间窗口的结束（不是本地，所以没法用mine来快速处理）
            await new Promise(resolve => setTimeout(resolve, 181 * 1000))
            const getFundTx = await fundMe.getFund();
            // 等待交易成功写入链上，获取到回执
            const getFundReceipt = await getFundTx.wait();
            expect(getFundReceipt)
                .to.be.emit(fundMe, "FundWithdrawByOwner")
                .withArgs(firstAccount, ethers.parseEther("0.5"));
        });

        it("fund and reFund successfully", async () => {
            await fundMe.fund({ value: ethers.parseEther("0.1") });
            await new Promise(resolve => setTimeout(resolve, 181 * 1000));
            const reFundTx = await fundMe.reFund();
            const reFundReceipt = await reFundTx.wait();
            expect(getFundReceipt)
                .to.be.emit(fundMe, "ReFundByFunder")
                .withArgs(firstAccount, ethers.parseEther("0.1"));
        })
    })
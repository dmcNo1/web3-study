const { assert, expect } = require("chai");
const { ethers, deployments } = require("hardhat");
const helpers = require("@nomicfoundation/hardhat-network-helpers")
const { developmentChains } = require("../../helper-hardhat-config");

!developmentChains.includes(network.name)
    ? describe.skip
    : describe("test fundMe contract", async function () {
        let firstAccount;
        let secondAccount;
        let fundMe;
        let fundMeSecondAccount;
        let mockV3Aggregator;

        // 每个test执行前都会执行这个方法
        beforeEach(async function () {
            await deployments.fixture(["all"]);
            firstAccount = (await getNamedAccounts()).firstAccount;
            secondAccount = (await getNamedAccounts()).secondAccount;
            const fundMeDeployment = await deployments.get("FundMe");
            mockV3Aggregator = await deployments.get("MockV3Aggregator");
            fundMe = await ethers.getContractAt("FundMe", fundMeDeployment.address);
            fundMeSecondAccount = await ethers.getContract("FundMe", secondAccount);
        })

        // 验证合约owner
        it("test if the owner is msg.sender", async function () {
            await fundMe.waitForDeployment();
            assert.equal(await fundMe.owner(), firstAccount);
        })

        // 验证合约dataFeed
        it("test if the dataFeed is assigned correctly", async function () {
            await fundMe.waitForDeployment();
            assert.equal(await fundMe.dataFeed(), mockV3Aggregator.address);
        })

        // fund, getFund, refund
        // unit test for fund
        // window open, value greater than minimum value, funder balance
        it("window closed, value greater than minimum, fund failed", async function () {
            // 确保时间窗口关闭
            await helpers.time.increase(200);
            await helpers.mine();
            expect(fundMe.fund({ value: ethers.parseEther("0.1") }))
                .to.be.revertedWith("window is closed");
        })

        it("window open, value less than minimum, fund failed", async function () {
            expect(fundMe.fund({ value: ethers.parseEther("0.001") }))
                .to.be.revertedWith("Send more ETH");
        })

        it("window open, value is greater than minimum, fund success", async function () {
            await fundMe.fund({ value: ethers.parseEther("0.1") });
            const balance = await fundMe.fundersToAmount(firstAccount);
            expect(balance).to.equal(ethers.parseEther("0.1"));
        })

        // getFund
        // onlyOwner, windowClosed, target reached
        it("not owner, window closed, target reached, getFund failed", async function () {
            // make sure target is reached
            await fundMe.fund({ value: ethers.parseEther("1") });

            // 确保时间窗口关闭
            await helpers.time.increase(200);
            await helpers.mine();

            expect(fundMeSecondAccount.getFund())
                .to.be.revertedWith("this function can only be called by owner")
        })

        it("window open, target reached", async () => {
            await fundMe.fund({ value: ethers.parseEther("1") });
            await expect(fundMe.getFund())
                .to.be.revertedWith("window is not closed")
        })

        it("window closed, target not reached, getFund failed", async () => {
            await fundMe.fund({ value: ethers.parseEther("0.1") });
            // 确保时间窗口关闭
            await helpers.time.increase(200);
            await helpers.mine();
            await expect(fundMe.getFund())
                .to.be.revertedWith("Target is not reached")
        })

        it("window closed, target reached, getFund success", async () => {
            await fundMe.fund({ value: ethers.parseEther("1") });
            await helpers.time.increase(200);
            await helpers.mine();
            await expect(fundMe.getFund())
                // 通过返回的event来校验交易的结果
                .to.emit(fundMe, "FundWithdrawByOwner").withArgs(ethers.parseEther("1"))
        })

        // refund
        // window closed, target not reached, funder has balance
        it("window open, target not reached, funder has balance", async () => {
            await fundMe.fund({ value: ethers.parseEther("0.1") });
            await expect(fundMe.refund()).to.be.revertedWith("window is not closed")
        })

        it("window closed, target reach, funder has balance",
            async function () {
                await fundMe.fund({ value: ethers.parseEther("1") })
                // make sure the window is closed
                await helpers.time.increase(200)
                await helpers.mine()
                await expect(fundMe.refund())
                    .to.be.revertedWith("Target is reached");
            }
        )

        it("window closed, target not reach, funder does not has balance",
            async function () {
                await fundMe.fund({ value: ethers.parseEther("0.1") })
                // make sure the window is closed
                await helpers.time.increase(200)
                await helpers.mine()
                await expect(fundMeSecondAccount.refund())
                    .to.be.revertedWith("there is no fund for you");
            }
        )

        it("window closed, target not reached, funder has balance",
            async function () {
                await fundMe.fund({ value: ethers.parseEther("0.1") })
                // make sure the window is closed
                await helpers.time.increase(200)
                await helpers.mine()
                await expect(fundMe.refund())
                    .to.emit(fundMe, "RefundByFunder")
                    .withArgs(firstAccount, ethers.parseEther("0.1"))
            }
        )
    })
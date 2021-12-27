const JamDistribute = artifacts.require("JamDistribute");
const SimpleCoin = artifacts.require("SimpleERC20Token")

contract("JamDistribute", (accounts) => {
    it("Test distribution flow", async () => {
        var jamdis = await JamDistribute.deployed()
        var erc20 = await SimpleCoin.deployed()
        
        // Add some fund to Distribute Contract
        const INIT_CONTRACT_FUND = 400
        await erc20.transfer(jamdis.address, INIT_CONTRACT_FUND)
        balance = (await erc20.balanceOf(jamdis.address)).toNumber()
        assert.equal(balance, INIT_CONTRACT_FUND, "Expecting 400 JAM Token")
        
        // Add Distributors
        await jamdis.updateDistributors(erc20.address, accounts[0], true)

        // Add Rewards
        await jamdis.addRewards(erc20.address, [accounts[1],accounts[2], accounts[3]], [20,30,40])
        account1Reward = await jamdis.viewReward(erc20.address, accounts[1])
        account2Reward = await jamdis.viewReward(erc20.address, accounts[2])
        account3Reward = await jamdis.viewReward(erc20.address, accounts[3])
        assert.equal(account1Reward, 20, "Reward ammount not matching")
        assert.equal(account2Reward, 30, "Reward ammount not matching")
        assert.equal(account3Reward, 40, "Reward ammount not matching")
        
        // Get Reward
        reward = await jamdis.getReward(erc20.address, {from: accounts[1]})
        balance = (await erc20.balanceOf(accounts[1])).toNumber()
        assert.equal(balance, 20, "User do not receive the correct reward")

        // Add Rewards
        await jamdis.addRewards(erc20.address, [accounts[1],accounts[2], accounts[3]], [20,30,40])
        account1Reward = await jamdis.viewReward(erc20.address, accounts[1])
        account2Reward = await jamdis.viewReward(erc20.address, accounts[2])
        account3Reward = await jamdis.viewReward(erc20.address, accounts[3])
        assert.equal(account1Reward, 20, "Reward ammount not matching")
        assert.equal(account2Reward, 60, "Reward ammount not matching")
        assert.equal(account3Reward, 80, "Reward ammount not matching")
    })
})
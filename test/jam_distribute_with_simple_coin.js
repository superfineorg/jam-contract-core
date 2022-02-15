const truffleAssert = require("truffle-assertions")
const JamDistribute = artifacts.require("JamDistribute")
const SimpleCoin = artifacts.require("SimpleERC20Token")

contract("JamDistribute", async (accounts) => {
    const INIT_CONTRACT_FUND = 400
    var jamdis
    var erc20

    before("Deployed contract", async () => {
        jamdis = await JamDistribute.deployed()
        erc20 = await SimpleCoin.deployed()
    })

    it("Contract should be able to receive anykind of ERC20 token", async () => {
        await erc20.transfer(jamdis.address, INIT_CONTRACT_FUND)
        balance = (await erc20.balanceOf(jamdis.address)).toNumber()
        assert.equal(balance, INIT_CONTRACT_FUND, "Expecting 400 JAM Token")
    })

    it("Should be fail to add distributors before support that token", async () => {
        await truffleAssert.reverts(jamdis.updateDistributors(erc20.address, accounts[0], true))
    })
    
    it("Should be able to add erc20 token to supported token and emit an AddToken event", async() => {
        await truffleAssert.eventEmitted(
            await jamdis.updateSupportedToken(erc20.address, true), 
            'AddToken', 
            (ev) => { return ev.tokenAddr = erc20.address }, 
            "updateSupportedToken should emit AddToken with tokenAddr of the supported token"
        )
    })
    
    it("Now it should be able to add distributors for supported token", async () => {
        await truffleAssert.eventEmitted(
            await jamdis.updateDistributors(erc20.address, accounts[0], true),
            'AddDistributor',
            (ev) => { return ev.tokenAddr === erc20.address && ev.distributor === accounts[0] },
            "updateDistributors should emit AddDistributor event with tokenAddr of supported token, and address of added distributor"
        )
    })
        
    it("Should be able to add reward one time a day by the distributor", async () => {
        await jamdis.addRewards(erc20.address, [accounts[1], accounts[2], accounts[3]], [20,30,40])
        await truffleAssert.fails(jamdis.addRewards(erc20.address, [accounts[1], accounts[2], accounts[3]], [20,30,40]))
        account1Reward = await jamdis.viewReward(erc20.address, accounts[1])
        account2Reward = await jamdis.viewReward(erc20.address, accounts[2])
        account3Reward = await jamdis.viewReward(erc20.address, accounts[3])
        assert.equal(account1Reward, 20, "Reward ammount not matching")
        assert.equal(account2Reward, 30, "Reward ammount not matching")
        assert.equal(account3Reward, 40, "Reward ammount not matching")
    })

    it("User should be able to get reward if there is", async () => {
        reward = await jamdis.getReward(erc20.address, {from: accounts[1]})
        balance = (await erc20.balanceOf(accounts[1])).toNumber()
        assert.equal(balance, 20, "User do not receive the correct reward")
    })
})

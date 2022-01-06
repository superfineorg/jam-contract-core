const truffleAssert = require("truffle-assertions")
const truffleFlattener = require("truffle-flattener")
const JamFaucet = artifacts.require("JamFaucet")

contract("JamFaucet", async(accounts) => {
    
    var faucet

    before("Deployed contract", async() => {
        faucet = await JamFaucet.deployed()
        web3.eth.defaultAccount = accounts[0]
        web3.eth.sendTransaction({
            from: web3.eth.defaultAccount,
            to: faucet.address,
            value: web3.utils.toWei('1', 'ether')
        })
    })

    it("Contract should have 1 ether", async () => {
        balance = await web3.eth.getBalance(faucet.address)
        assert.equal(balance, web3.utils.toWei('1', 'ether'))
    }) 

    it("Set faucet wei should be fail if not owner", async () => {
        await truffleAssert.fails(
            faucet.setFaucetWei(10, {from :accounts[2]}), 
            truffleAssert.ErrorType.REVERT,
            "Ownable: caller is not the owner"
        ) 
    })

    it("Set faucet wei should be success if it is owner", async () => {
        await faucet.setFaucetWei(10)
        faucetWei = await faucet.faucetWei()
        assert.equal(faucetWei, 10)        
    })

    it("Set interval to 10 minutes", async () => {
        await faucet.setFaucetInterval(600)
        faucetInterval = await faucet.faucetInterval()
        assert.equal(faucetInterval, 600)
    })

    it("Faucet should be allow, user get 10 wei from it", async() => {
        previousBalance = await web3.eth.getBalance(accounts[1])
        await faucet.faucet(accounts[1])
        afterBalance = await web3.eth.getBalance(accounts[1])
        addedAmount= web3.utils.toBN("10")
        assert.equal(web3.utils.toBN(previousBalance).add(addedAmount).toString(), web3.utils.toBN(afterBalance).toString())
    }) 


    it("Faucet should not be allow, since user already faucet it before and still in 10 minute frame", async() => {
        await truffleAssert.reverts(faucet.faucet(accounts[1]))
    }) 
})
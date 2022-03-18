const truffleAssert = require('truffle-assertions');
const Erc20Factory = artifacts.require("JamERC20Factory");

contract("ERC20Fatory", async (accounts) => {
    var fatory;

    before("Deploy contract", async () => {
        factory = await Erc20Factory.deployed();
        web3.eth.defaultAccount = accounts[0];
    });

    it("Create contract  with initial supply higher than capacity should be fail", async () => {
        truffleAssert.reverts(factory.createERC20(300, 1000, 18, "JAM", "JAM"));
    });

    it("Create contract should be able to done by any user and return correct event", async () => {
        await truffleAssert.eventEmitted(
            await factory.createERC20(3000, 1000, 18, "JAM", "JAM"),
            "CreateERC20",
            (ev) => {
                return ev.name === "JAM" && ev.symbol === "JAM";
            }
        );
    });

});
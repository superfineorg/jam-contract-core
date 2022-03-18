const truffleAssert = require('truffle-assertions');
const Erc721Factory = artifacts.require("JamERC721Factory");

contract("ERC20Fatory", async (accounts) => {
    var fatory;

    before("Deploy contract", async () => {
        factory = await Erc721Factory.deployed();
        web3.eth.defaultAccount = accounts[0];
    });

    it("Create contract", async () => {
        await truffleAssert.eventEmitted(
            await factory.createERC721("JAM721", "JAM721", "https://google.com"),
            "CreateERC721",
            (ev) => {
                return ev.template_type === "SimpleERC721-V1" && ev.symbol === "JAM721";
            }
        );
    });

});
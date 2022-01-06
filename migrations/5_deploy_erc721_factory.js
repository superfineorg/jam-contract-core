var erc721Factory = artifacts.require("ERC721Factory");

module.exports = function(deployer) {
    deployer.deploy(erc721Factory)
}
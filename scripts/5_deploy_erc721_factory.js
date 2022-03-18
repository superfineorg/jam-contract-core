var erc721Factory = artifacts.require("JamERC721Factory");

module.exports = function (deployer) {
    deployer.deploy(erc721Factory);
};
var erc20Factory = artifacts.require("JamERC20Factory");

module.exports = function (deployer) {
    deployer.deploy(erc20Factory);
};
var erc20Factory = artifacts.require("ERC20Factory");

module.exports = function(deployer) {
    deployer.deploy(erc20Factory)
}
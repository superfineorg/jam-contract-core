var JamDistribute = artifacts.require("JamDistribute");
var SimpleERC20Token = artifacts.require("SimpleERC20Token");

module.exports = function(deployer) {
   deployer.deploy(JamDistribute,"0x502A5D8BA5E39e2b314D9Cb226Aa055193a20B7e");
   deployer.deploy(SimpleERC20Token, "0x502A5D8BA5E39e2b314D9Cb226Aa055193a20B7e", 30000, 10000, 18, "JAM", "JAM");
};
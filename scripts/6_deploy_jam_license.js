var JamLicense = artifacts.require("JamLicense");
var NodeLicenseNFT = artifacts.require("JamNodeLicenseNFT");

module.exports = function (deployer) {
   jamlicense = deployer.deploy(JamLicense, "0x7345f5761553af07cb70cd786ef224ceda513a6d");
   nodenft = deployer.deploy(NodeLicenseNFT);
};
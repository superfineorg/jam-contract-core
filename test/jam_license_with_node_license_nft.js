const truffleAssert = require("truffle-assertions");
const JamLicense = artifacts.require("JamLicense");
const NodeLicenseNFT = artifacts.require("JamNodeLicenseNFT");

contract("JamLicense", async (accounts) => {
    const INIT_FIRST_PRICE = 1;
    const INIT_PRICE_STEP = 2;
    const INIT_PRICE_INCREASE = 3;
    const INIT_SLIPPAGE_TOLERANCE = 4;
    const INIT_MAX_NODE_BUY_PER_TRANSACTION = 5;
    const INIT_RATE_FOR_ETHER = "1000000000000000000";
    const ETHER_DEFINE_ADDR = "0x0000000000000000000000000000000000000000";
    var jamLicense;
    var nodeNFT;

    before("Deployed contract", async () => {
        jamLicense = await JamLicense.deployed();
        nodeNFT = await NodeLicenseNFT.deployed();
        await nodeNFT.transferOwnership(jamLicense.address);
    });

    it("Check owner of nodeNFT is jamLicense", async () => {

        owner = await nodeNFT.owner();
        assert.equal(owner, jamLicense.address, "Expecting true owner");
    });

    it("Check set jamLicense's configuration is ok", async () => {
        await jamLicense.setNodeAddress(nodeNFT.address);
        nodeAddr = await jamLicense.getNodeAddress();
        assert.equal(nodeAddr, nodeNFT.address, "Expecting true node NFT address");
    });

    it("Check set jamLicense's configuration is ok", async () => {
        await jamLicense.setNodeAddress(nodeNFT.address);
        nodeAddr = await jamLicense.getNodeAddress();
        assert.equal(nodeAddr, nodeNFT.address, "Expecting true node NFT address");
        await jamLicense.setPriceIncrease(INIT_PRICE_INCREASE);
        priceIncrease = await jamLicense.getPriceIncrease();
        assert.equal(priceIncrease, INIT_PRICE_INCREASE, "Expecting true INIT_PRICE_INCREASE");
        await jamLicense.setFirstPrice(INIT_FIRST_PRICE);
        firstPrice = await jamLicense.getFirstPrice();
        assert.equal(firstPrice, INIT_FIRST_PRICE, "Expecting true INIT_FIRST_PRICE");
        await jamLicense.setPriceStep(INIT_PRICE_STEP);
        priceStep = await jamLicense.getPriceStep();
        assert.equal(priceStep, INIT_PRICE_STEP, "Expecting true INIT_PRICE_STEP");
        await jamLicense.setSlippageTolerance(INIT_SLIPPAGE_TOLERANCE);
        slt = await jamLicense.getSlippageTolerance();
        assert.equal(slt, INIT_SLIPPAGE_TOLERANCE, "Expecting true INIT_SLIPPAGE_TOLERANCE");
        await jamLicense.setMaxNodeBuyPerTransaction(INIT_MAX_NODE_BUY_PER_TRANSACTION);
        m = await jamLicense.getMaxNodeBuyPerTransaction();
        assert.equal(m, INIT_MAX_NODE_BUY_PER_TRANSACTION, "Expecting true INIT_MAX_NODE_BUY_PER_TRANSACTION");
        await jamLicense.setRate([ETHER_DEFINE_ADDR], [INIT_RATE_FOR_ETHER]);
        rateEther = await jamLicense.getRate(ETHER_DEFINE_ADDR);
        assert.equal(rateEther.toString(), INIT_RATE_FOR_ETHER, "Expecting true rate");
    });


    //
    it("Should be return true billing amount in USD, and revert when numNodeLicense out limit", async () => {
        bill = await jamLicense.getBillAmountInUSD(3);
        assert.equal(bill.toString(), "6", "Expecting true billAmount");
        await truffleAssert.reverts(jamLicense.getBillAmountInUSD(INIT_MAX_NODE_BUY_PER_TRANSACTION + 1));
    });

    it("Should be return true billing amount with Ether, and revert when numNodeLicense out limit", async () => {
        bill = await jamLicense.getBillAmount(ETHER_DEFINE_ADDR, 3);
        assert.equal(bill.toString(), "6000000000000000000", "Expecting true billAmount");
        await truffleAssert.reverts(jamLicense.getBillAmount(ETHER_DEFINE_ADDR, INIT_MAX_NODE_BUY_PER_TRANSACTION + 1));
    });

    it("Should be return number of NodeNFT after buyNode", async () => {
        await jamLicense.buyNode(ETHER_DEFINE_ADDR, 3, "6000000000000000000", { from: accounts[1], value: "6000000000000000000" });
        balance = await nodeNFT.balanceOf(accounts[1]);
        assert.equal(balance.toString(), "3", "Expecting true  number of NodeNFT");
    });


});

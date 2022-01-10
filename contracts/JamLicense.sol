// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./HasNoEther.sol";
import "./NodeLicenseNFT.sol";

contract JamLicense is ReentrancyGuard, HasNoEther {
    using SafeMath for uint256;

    address private nodeAddress;
    uint256 private firstPrice; // first price in usdt
    uint256 private priceStep;  // number of node for each time increase price
    uint256 private priceIncrease; // increase price in usdt
    uint256 private maxNodeBuyPerTransaction;
    uint256 private slippageTolerance; // slippageTolerance 0 -> 10000 mean 0% to 100%


    mapping(address => uint256) private tokenRate; // rate token(with decimals)/USDT

    constructor(address _newOwner) {
        transferOwnership(_newOwner);
    }

    function _getNodeLicense(address _nftAddress) internal pure returns (NodeLicenseNFT) {
        NodeLicenseNFT candidateContract = NodeLicenseNFT(_nftAddress);
        return candidateContract;
    }

    function _getERC20Contract(address _erc20Address) internal pure returns (IERC20) {
        IERC20 candidateContract = IERC20(_erc20Address);
        return candidateContract;
    }

    fallback() external payable {}

    receive() external payable {}

    /* ======== SetInfo ========= */

    function setNodeAddress(address addr) external onlyOwner nonReentrant {
        nodeAddress = addr;
    }

    function setPriceIncrease(uint256 _priceIncrease) external onlyOwner nonReentrant {
        priceIncrease = _priceIncrease;
    }

    function setFirstPrice(uint256 _firstPrice) external onlyOwner nonReentrant {
        firstPrice = _firstPrice;
    }

    function setPriceStep(uint256 _priceStep) external onlyOwner nonReentrant {
        priceStep = _priceStep;
    }

    function setSlippageTolerance(uint256 _slippageTolerance) external onlyOwner nonReentrant {
        slippageTolerance = _slippageTolerance;
    }

    function setMaxNodeBuyPerTransaction(uint256 _maxNodeBuyPerTransaction) external onlyOwner nonReentrant {
        maxNodeBuyPerTransaction = _maxNodeBuyPerTransaction;
    }

    function setRate(address[] memory tokenAddrs, uint256[] memory rate) external onlyOwner nonReentrant {
        require(tokenAddrs.length == rate.length, "addrs and amount does not same length");
        for (uint i = 0; i < tokenAddrs.length; i++) {
            tokenRate[tokenAddrs[i]] = rate[i];
        }
    }

    function returnOwnerShip() external onlyOwner nonReentrant {
        NodeLicenseNFT node = _getNodeLicense(nodeAddress);
        node.transferOwnership(msg.sender);
    }


    /* ======== GetInfo ========= */
    function getNodeAddress() public view returns (address) {
        return nodeAddress;
    }

    function getPriceIncrease() public view returns (uint256) {
        return priceIncrease;
    }

    function getFirstPrice() public view returns (uint256) {
        return firstPrice;
    }

    function getPriceStep() public view returns (uint256) {
        return priceStep;
    }

    function getSlippageTolerance() public view returns (uint256) {
        return slippageTolerance;
    }

    function getMaxNodeBuyPerTransaction() public view returns (uint256) {
        return maxNodeBuyPerTransaction;
    }

    function getRate(address tokenAddr) public view returns (uint256) {
        return tokenRate[tokenAddr];
    }

    function getBillAmount(address paidToken, uint256 numNodeLicense) validNumPurchaseNodeLicense(numNodeLicense) validCurrency(paidToken) public view returns (uint256) {
        uint256 billAmount = getBillAmountInUSD(numNodeLicense);
        uint256 convertRate = tokenRate[paidToken];
        return billAmount.mul(convertRate);
    }

    function getBillAmountInUSD(uint256 numNodeLicense) validNumPurchaseNodeLicense(numNodeLicense) public view returns (uint256) {
        NodeLicenseNFT nodeContract = _getNodeLicense(nodeAddress);
        uint256 nodeCount = nodeContract.totalSupply();
        uint256 nextNodeCount = nodeCount.add(numNodeLicense)-1;
        uint256 currentBatch = nodeCount.div(priceStep);
        uint256 currentPrice = firstPrice + currentBatch * priceIncrease;
        uint256 nextBatch = nextNodeCount.div(priceStep);
        uint256 nextPrice = firstPrice + nextBatch * priceIncrease;
        uint256 billAmount;
        if (nextBatch == currentBatch) {
            return currentPrice.mul(numNodeLicense);
        }
        uint256 nextRemain = nextNodeCount.mod(priceStep);
        billAmount = currentPrice.mul(priceStep - nextRemain) + nextPrice.mul(nextRemain+1);
        for (uint256 i = 1; i < nextBatch.sub(currentBatch); i++) {
            uint256 batchPrice = currentPrice + i * priceIncrease;
            billAmount += batchPrice * priceStep;
        }
        return billAmount;
    }

    /* ======== PurchasedNode ========= */

    function buyNode(address paidToken, uint256 numNodeLicense, uint256 amount) validNumPurchaseNodeLicense(numNodeLicense) validCurrency(paidToken) external payable {
        uint256 billAmount = getBillAmount(paidToken, numNodeLicense);
        uint256 chargeAmount = billAmount;
        if (billAmount > amount) {
            // check slippageTolerance
            uint256 diff = billAmount - amount;
            require(diff.mul(10000) < billAmount.mul(slippageTolerance), "JAM_LICENSE: out of slippage tolerance");
            chargeAmount = amount;
        }
        // process transfer token if currency is ERC20
        if (paidToken != address(0)) {
            SafeERC20.safeTransferFrom(IERC20(paidToken), msg.sender, address(this), billAmount);
        } else {
            require(msg.value >= billAmount, "Not enough balance");
            if (msg.value > billAmount) {
                uint256 _billExceed = msg.value - billAmount;
                payable(msg.sender).transfer(_billExceed);
            }
        }
        // mint token
        NodeLicenseNFT nodeContract = _getNodeLicense(nodeAddress);
        for (uint256 i = 0; i < numNodeLicense; i ++) {
            nodeContract.mint(msg.sender);
        }
        emit PurchasedNode(msg.sender, paidToken, numNodeLicense, chargeAmount);
    }


    /* ======== Withdraw ======== */
    function reclaimERC20(address _erc20Address) external onlyOwner {
        IERC20 erc20Contract = _getERC20Contract(_erc20Address);
        erc20Contract.transfer(owner(), erc20Contract.balanceOf(address(this)));
    }
    /* ======== Modfier ========= */
    modifier validNumPurchaseNodeLicense(uint256 numNodeLicense) {
        require(numNodeLicense > 0, "node buy must be greather than 0");
        require(numNodeLicense <= maxNodeBuyPerTransaction, "too many node buying");
        _;
    }
    modifier validCurrency(address paidToken) {
        require(tokenRate[paidToken] > 0, "currency not support");
        _;
    }

    /* ======== Event ======== */

    event PurchasedNode(address indexed buyer, address currencyToken, uint256 numNodeLicense, uint256 billAmount);
}

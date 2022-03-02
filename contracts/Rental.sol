pragma solidity ^0.8.0;

import "./IRentalContract.sol";
import "./utils/HasNoEther.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./IOracle.sol";

contract Rental is IRentalContract, HasNoEther, ReentrancyGuard {


    address private backendAddr;

    address private oracleContract;

    // Cut owner takes on each auction, measured in basis points (1/100 of a percent).
    // Values 0-10,000 map to 0%-100%
    uint256 public ownerCut;

    uint256 public slippageTolerance; // slippageTolerance 0 -> 10000 mean 0% to 100%

    mapping(uint256 => mapping(address => mapping(uint256 => uint256))) private rentPriceUSD; // pricing Of NFT in USDT (with decimals 6)

    mapping(string => uint256) private receiptLog;
    mapping(address => uint256) private ownerProfit;
    mapping(address => mapping(string => uint256)) private rentedExpire;

    constructor(address _newOwner) {
        transferOwnership(_newOwner);
        backendAddr = _newOwner;
    }

    /* ======== Config Rental ======== */

    function setOwnerCut(uint256 _ownerCut) onlyOwner {
        require(_ownerCut <= 10000, "Owner cut cannot exceed 100%");
        ownerCut = _ownerCut;
    }

    function setBackendAddress(uint256 _backendAddr) onlyOwner {
        backendAddr = _backendAddr;
    }

    function setBackendAddress(uint256 _oracleContract) onlyOwner {
        oracleContract = _oracleContract;
    }

    function setSlippageTolerance(uint256 _slippageTolerance) external onlyOwner
    {
        slippageTolerance = _slippageTolerance;
    }

    /* ======== End Config Rental ======== */

    /* ======== Internal function ======== */

    function _getERC20Contract(address _erc20Address)
    internal
    pure
    returns (IERC20)
    {
        IERC20 candidateContract = IERC20(_erc20Address);
        return candidateContract;
    }


    function _getOracleContract()
    internal
    pure
    returns (IOracle)
    {
        IOracle candidateContract = IOracle(_nftAddress);
        return candidateContract;
    }

    function _updateExpire(address renter, uint256 chainId, address contractAddress, uint256 tokenId, uint8 rentedDay)
    internal
    nonReentrant
    pure
    {
        string memory key = abi.encodePacked(Strings.toString(chainId), "-", abi.encodePacked(contractAddress), "-" , Strings.toString(tokenId));
        uint256 currentTime = block.timestamp;
        if (currentTime > rentedExpire[renter][key]) {
            rentedExpire[renter][key] = currentTime.add(rentedDay * 1 days);
        }else{
            rentedExpire[renter][key] = rentedExpire[renter][key].add(rentedDay * 1 days);
        }
        return candidateContract;
    }

    /* ======== End Internal function ======== */

    /* ======== Business function ======== */

    /**
     * @dev update RentNFT pricing, 0 mean not public for rent
    */
    function updateRentPrice(uint256 chainId, address contractAddress, uint256 tokenId, uint256 pricingPerDay) external onlyBackend returns (bool){
        rentPriceUSD[chainId][contractAddress][tokenId] = pricingPerDay;
        // Emit event UpdateRentPrice
        emit UpdateRentPrice(chainId, contractAddress, tokenId, pricingPerDay);
        return true;
    }

    /**
      * @dev add profit to user. USDT alway have 6 decimal
      * TODO: Need oracle to convert usdt to jam
    */
    function addProfit(address owner, uint256 addingBalance, string memory receiptId) external onlyNewReceiptId(receiptId) onlyBackend returns (bool) {
        require(addingBalance > 0);
        // TODO: need convert addingBalance from usdt to jam
        IOracle oracleContract = _getOracleContract();
        uint256 jamRate = oracleContract.GetRate(address(0));
        // calculate the jamBalance
        uint256 jamBalance = jamRate.mul(addingBalance);
        ownerProfit[owner] = ownerProfit[owner].add(jamBalance);
        receiptLog[receiptId] = jamBalance;
        return true;
    }

    /**
      * @dev rent the NFT
      * TODO : need oracle to continue
     */
    function rentNFT(uint256 chainId, address contractAddress, uint256 tokenId, uint8 rentedDay, address paidToken, uint256 amount) payable external returns (bool){
        if (paidToken == address(0)) {
            require(msg.value >= amount, "not enough balance");
        }
        IOracle oracleContract = _getOracleContract();
        uint256 tokenRate = oracleContract.GetRate(paidToken);
        require(tokenRate > 0 , "paid token is not support");
        uint256 totalBalanceUSD = amount.div(tokenRate);
        uint256 pricingPerDayUSD = rentPriceUSD[chainId][contractAddress][tokenId];
        //check slippageTolerance
        uint256 billAmount = pricingPerDayUSD.mul(rentedDay);
        uint256 diff = billAmount.sub(amount);
        require(diff.mul(10000) < billAmount.mul(slippageTolerance), "out of slippage tolerance");
        // calculate fee and addingBalance
        uint256 fee = amount.mul(ownerCut).div(10000);
        uint256 addingBalance = amount.sub(fee);
        emit RentNFT(msg.sender, chainId, contractAddress, tokenId, rentedDay, paidToken, amount, addingBalance, fee);
        return true;
    }

    /**
     * @dev rent the NFT via IAP
     * TODO : need oracle to continue
    */
    function rentNFTViaIAP(uint256 chainId, address contractAddress, uint256 tokenId, address renter, address owner, uint8 rentedDay, uint256 addingBalance, string memory receiptId) external onlyBackend returns (bool){
        // TODO: need oracle to convert addingBalance form USDT to JAM
        addingBalance(owner, addingBalance, receiptId);
        emit RentNFTViaIAP(renter, owner, chainId, contractAddress, tokenId, rentedDay, receiptId, addingBalance);
        return true;
    }

    /* ======== End Business function ======== */

    /* ======== Query function ======== */

    function viewProfit()
    public
    view
    returns (uint256)
    {
        return ownerProfit[msg.sender];
    }

    function viewNFTExpire(uint256 chainId, address contractAddress, uint256 tokenId)
    public
    view
    returns (uint256)
    {
        string memory key = abi.encodePacked(Strings.toString(chainId), "-", abi.encodePacked(contractAddress), "-" , Strings.toString(tokenId));
        return rentedExpire[msg.sender][key];
    }

    /* ======== End Query function ======== */


    /* ======== Admin Withdraw ERC20 ======== */
    function reclaimERC20(address _erc20Address) external onlyOwner {
        IERC20 erc20Contract = _getERC20Contract(_erc20Address);
        erc20Contract.transfer(owner(), erc20Contract.balanceOf(address(this)));
    }

    /* ======== User Query function ======== */
    function reclaimProfit() external {
        uint256 amount = ownerProfit[msg.sender];
        (bool success,) = payable(msg.sender).call{value : amount}(
            ""
        );
        ownerProfit[msg.sender] = ownerProfit[msg.sender].sub(amount);
        require(success, "Transfer failed.");
    }

    /* ======== End User Query function ======== */

    /* ======== Modfier ========= */
    modifier onlyBackend() {
        require(backendAddr == msg.sender, "only backend call");
        _;
    }

    modifier onlyNewReceiptId(string memory receiptId) {
        require(receiptLog[receiptId] == 0, "only new receiptId accepted");
        _;
    }

}

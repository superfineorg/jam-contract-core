// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "./DateTime.sol";

contract JamDistribute is ReentrancyGuard {

    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    // Owner of the contract
    address owner;

    // Store distributor status link to specific type of token
    mapping(address => mapping(address => bool)) private allowedDistributors;

    // Store latest day that user get reward link to specific type of token
    // latestRewardDate[tokenAddress][accountAddress] = latestRewardDate
    mapping(address => mapping(address => uint256)) private latestRewardDate;

    // Store rewards of user link to specific type of token
    // reward[tokenAddress][accountAddress] = rewardAmount 
    mapping(address => mapping(address => uint256)) private rewards;

    // totalAmount of reward of a specific token, 
    //just for the sake of no looping :D
    mapping(address => uint256) private totalRewards;
    
    /* ========== CONSTRUCTOR ========== */
    constructor(
        address _owner
    ) {
        owner = _owner;
    }


    /* ======== Native Function ========= */
    function viewReward(address tokenAddr, address addr) public view returns (uint256){
        return rewards[tokenAddr][addr];
    }

    function addRewards(address tokenAddr, address[] memory addrs, uint256[] memory amount) external
        nonReentrant
        onlyDistributor(tokenAddr)
        enoughBalance(tokenAddr, amount)
    {
        require(addrs.length == amount.length, "addrs and amount does not same length");
        uint256 addedRewardAmount = _calSumAmount(amount);
        for (uint i=0; i < addrs.length; i++) {
            rewards[tokenAddr][addrs[i]] = rewards[tokenAddr][addrs[i]].add( amount[i]);
        }
        totalRewards[tokenAddr] = totalRewards[tokenAddr].add(addedRewardAmount);
    }

    function updateDistributors(address tokenAddr, address distributor, bool ok) external 
        onlyOwner 
    {
        allowedDistributors[tokenAddr][distributor] = ok;
        if (ok) {
            emit AddDistributor(tokenAddr, distributor);
        } else {
            emit RemoveDistributor(tokenAddr, distributor);
        }
    }

    function getReward(address tokenAddr) external 
        nonReentrant
        haveReward(tokenAddr)
        availableForToday(tokenAddr)
    {
        uint256 today = DateTime.toDateUnit(block.timestamp);
        uint256 rewardAmount = rewards[tokenAddr][msg.sender];
        if (tokenAddr == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: rewardAmount}(
                ""
            );
            require(success, "Transfer failed.");
        } else {
            IERC20(tokenAddr).safeTransfer(msg.sender, rewardAmount);
        }
        rewards[tokenAddr][msg.sender] = 0;
        totalRewards[tokenAddr] = totalRewards[tokenAddr] - rewardAmount;
        latestRewardDate[tokenAddr][msg.sender] = today;
    }

    function _calSumAmount(uint256[] memory amount) private pure returns (uint256) {
        uint256 total = 0;
        for(uint i = 0; i < amount.length; i++) {
            total = total.add(amount[i]);
        }
        return total;
    }
 
    fallback() external payable {}

    receive() external payable {}
    

    /* ======== Modifier ========= */
    modifier onlyOwner() {
        require(msg.sender == owner, "only onwer can take this action");
        _;
    }

    modifier onlyDistributor(address _token) {
        require(allowedDistributors[_token][msg.sender], "only distributor can take this action");
        _;
    }

    modifier enoughBalance(address tokenAddr, uint256[] memory amount) {
        uint256 total = _calSumAmount(amount);
        uint256 thisAccountBalance = 0;
        if (tokenAddr == address(0)) {
            thisAccountBalance = address(this).balance;
        } else {
            thisAccountBalance = IERC20(tokenAddr).balanceOf(address(this));
        }
        require(totalRewards[tokenAddr].add(total) <= thisAccountBalance, "balance not enough");
        _;
    }

    modifier availableForToday(address tokenAddr) {
        uint256 today = DateTime.toDateUnit(block.timestamp);
        require(latestRewardDate[tokenAddr][msg.sender] <= today, "already taken for today");
        _;
    }

    modifier haveReward(address tokenAddr) {
        require(rewards[tokenAddr][msg.sender] > 0, "you don't have any reward");
        _;
    }

    /* ======== Event ========= */
    event AddDistributor(address indexed tokenAddr, address indexed distributor);
    event RemoveDistributor(address indexed tokenAddr, address indexed distributor);
}

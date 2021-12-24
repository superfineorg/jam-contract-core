// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
// import "./HasNoEther.sol";
import "./DateTime.sol";

contract JamDistribute is ReentrancyGuard, AccessControl {
    using SafeMath for uint256;
    mapping(address => uint256) private currentDistributedDate;
    mapping(address => uint256) private currentDistribute;
    mapping(address => uint256[]) private currentDistributedAmount;
    mapping(address => address[]) private currentDistributedAddr;
    mapping(address => bool) private doneDistributed;
    mapping(address => mapping(address => bool)) private allowedDistributors;
    address private defaultDistributor;

    //// Event
    event DistributeFail(address indexed tokenAddr, address indexed receiver, uint256 amount);
    event AddDistributor(address indexed tokenAddr, address indexed distributor);
    event RemoveDistributor(address indexed tokenAddr, address indexed distributor);


    constructor(address _defaultDistributor) payable{
        defaultDistributor = _defaultDistributor;
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
    }


    function adminUpdateDistributors(address token, address _dist, bool ok) external onlyAdmin {
        allowedDistributors[token][_dist] = ok;
        if (ok) {
            emit AddDistributor(token, _dist);
        } else {
            emit RemoveDistributor(token, _dist);
        }
    }

    fallback() external payable {
    }

    receive() external payable {
    }

    //// Modifier
    function _onlyAdmin() public view
    {
        require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "sender is not admin");
    }

    modifier onlyAdmin()
    {
        _onlyAdmin();
        _;
    }

    modifier isDistributor(address token)
    {
        require((_msgSender() == defaultDistributor) || (allowedDistributors[token][_msgSender()]), "sender is not distributor");
        _;
    }


    modifier enoughBalance(address tokenAddr) 
    {
        if (tokenAddr == address(0)) {
            require(currentDistribute[tokenAddr] < address(this).balance, "Balance not enough");
        } else {
            IERC20 erc20Contract = _getERC20Contract(tokenAddr);
            require(currentDistribute[tokenAddr] <= erc20Contract.balanceOf(address(this)), "Balance not enough");
        }
        _;
    }

    modifier isNotDone(address tokenAddr) 
    {
        require(!doneDistributed[tokenAddr], "state is distributed done");
        _;
    }

    modifier isToday(address tokenAddr) 
    {
        uint256 currentDate = DateTime.toDateUnit(block.timestamp);
        require(currentDate == currentDistributedDate[tokenAddr], "current date not equal currentDistributedDate");
        _;
    }

    // modifier initDistributeCap(address tokenAddr) 
    // {
    //     uint256 memory d = currentDistributeCap[tokenAddr];
    //     uint256 memory currentDate = DateTime.toDateUnit(block.timestamp);
    //     uint256 memory lastDate =  currentDistributedDate[tokenAddr];
    //     require(currentDate >= lastDate, "error currentDate is less than currentDistributedDate");
    //     if (currentDate > lastDate) {
    //         require(distributeCapPerDay[tokenAddr] > 0, "distributeCapPerDay not set");
    //         if (lastDate == 0) {
    //             lastDate = currentDate -1 ;
    //         }
    //         if (lastDate < currentDate) {
    //             //Calculate Max cap for current distributed
    //             currentDistributeCap[tokenAddr] = distributeCapPerDay[tokenAddr] * (currentDate - lastDate);
    //             // Set currentDistributedDate = today
    //             currentDistributedDate[tokenAddr] = currentDate;
    //         }
    //     }
    //     _;
    // }

    ////

    
    function addDistribute(address tokenAddr, address[] memory addrs, uint256[] memory amount) external 
        nonReentrant
        isDistributor(tokenAddr)
        isNotDone(tokenAddr)
    {
        require(addrs.length == amount.length, "address and amount does not same length");
        uint i;
        uint256 total = 0;
        for (i = 0; i < addrs.length; i += 1) {
            currentDistributedAmount[tokenAddr].push(amount[i]);
            currentDistributedAddr[tokenAddr].push(addrs[i]);
            total += amount[i];
        }
        currentDistribute[tokenAddr] += total;
    }

    function initDistribute(address tokenAddr) public 
        nonReentrant 
        isDistributor(tokenAddr) 
    {
        uint256 currentDate = DateTime.toDateUnit(block.timestamp);
        // First init
        if (currentDistributedDate[tokenAddr] == 0) {
            doneDistributed[tokenAddr] = true;
            currentDistribute[tokenAddr] = 0;
        } 
        
        //If state is already distributed
        if (doneDistributed[tokenAddr]) {
            //Check currentTime and last Distributed
            require(currentDate > currentDistributedDate[tokenAddr], "current date less than currentDistributedDate");
            // reset all
            doneDistributed[tokenAddr] = false;
            currentDistribute[tokenAddr] = 0;
        }
        
        currentDistributedDate[tokenAddr] = currentDate;
    }

    function executor(address tokenAddr) public 
        nonReentrant
        isDistributor(tokenAddr)
        isNotDone(tokenAddr)
        enoughBalance(tokenAddr) 
    {
        uint256 tt = currentDistributedAddr[tokenAddr].length;
        for (uint256 i = 0; i < tt ; i += 1) {
            uint256 amount = currentDistributedAmount[tokenAddr][currentDistributedAmount[tokenAddr].length - 1]; 
            currentDistributedAmount[tokenAddr].pop();
            address addr = currentDistributedAddr[tokenAddr][currentDistributedAddr[tokenAddr].length - 1];
            currentDistributedAddr[tokenAddr].pop();
            bool success = false;
            if (tokenAddr == address(0)) {
                (success,) = payable(addr).call{value : amount}(
                    ""
                );
            } else {
                IERC20 erc20Contract = _getERC20Contract(tokenAddr);
                success = erc20Contract.transfer(addr, amount);
            }
            if (!success) {
                emit DistributeFail(tokenAddr, addr, amount);
            }
        }
        doneDistributed[tokenAddr] = true;
    }

    // ERC20;
    function _getERC20Contract(address _erc20Address) public pure returns (IERC20)
    {
        return IERC20(_erc20Address);
    }
}
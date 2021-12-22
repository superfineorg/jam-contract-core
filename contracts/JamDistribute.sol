pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./HasNoEther.sol";
import "./DateTime.sol";

contract JamDistribute is HasNoEther {
    using SafeMath for uint256;
    mapping(address => uint256) private currentDistributedDate;
    mapping(address => uint256) private currentDistribute;
    mapping(address => bool) private reentrancyLock;
    mapping(address => uint256[]) private currentDistributedAmount;
    mapping(address => address[]) private currentDistributedAddr;
    mapping(address => bool) private doneDistributed;
    mapping(address => address) private distributor;
    address private defaultDistributor;


    event DistributeFail(address indexed tokenAddr, address indexed receiver, uint256 amount);

    function _getERC20Contract(address _erc20Address) internal pure returns (IERC20)
    {
        IERC20 candidateContract = IERC20(_erc20Address);
        return candidateContract;
    }

    modifier nonReentrant(address addr) {
        // On the first call to nonReentrant, _notEntered will be true
        require(reentrancyLock[addr] != 1, "ReentrancyGuard: reentrant call");

        // Any calls to nonReentrant after this point will fail
        reentrancyLock[addr] = 1;

        _;

        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        reentrancyLock[addr] = 0;
    }

    modifier isDistributor(address tokenAddr) {
        require(_msgSender() == defaultDistributor || distributor[tokenAddr] == _msgSender(), "Not distributor");
        _;
    }

    modifier enoughBalance(address tokenAddr) {
        if (tokenAddr == address(0)) {
            require(currentDistribute[tokenAddr] < address(this).balance, "Balance not enough");
        } else {
            IERC20 erc20Contract = _getERC20Contract(_erc20Address);
            require(currentDistribute[tokenAddr] <= erc20Contract.balanceOf(address(this)), "Balance not enough");
        }
        _;
    }

    modifier isNotDone(address tokenAddr) {
        require(!doneDistributed[tokenAddr], "state is distributed done");
    }

    //    modifier isToday(address tokenAddr) {
    //        currentDate = DateTime.toDateUnit(block.timestamp);
    //        require(currentDistribute == currentDistributedDate[tokenAddr], "current date not equal currentDistributedDate");
    //    }
    //
    //    modifier initDistributeCap(address tokenAddr) {
    //        uint256 memory d = currentDistributeCap[tokenAddr];
    //        uint256 memory currentDate = DateTime.toDateUnit(block.timestamp);
    //        uint256 memory lastDate =  currentDistributedDate[tokenAddr];
    //        require(currentDate >= lastDate, "error currentDate is less than currentDistributedDate");
    //        if (currentDate > lastDate) {
    //            require(distributeCapPerDay[tokenAddr] > 0, "distributeCapPerDay not set");
    //            if (lastDate == 0) {
    //                lastDate = currentDate -1 ;
    //            }
    //            if (lastDate < currentDate) {
    //                //Calculate Max cap for current distributed
    //                currentDistributeCap[tokenAddr] = distributeCapPerDay[tokenAddr] * (currentDate - lastDate);
    //                // Set currentDistributedDate = today
    //                currentDistributedDate[tokenAddr] = currentDate;
    //            }
    //        }
    //        _;
    //    }

    function addDistribute(address tokenAddr, address[] addrs, uint256[] amount) public nonReentrant(tokenAddr) isDistributor(tokenAddr) isNotDone(tokenAddr) {
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

    function initDistribute(address tokenAddr) public nonReentrant(tokenAddr) isDistributor(tokenAddr) {
        currentDate = DateTime.toDateUnit(block.timestamp);
        // First init
        if (currentDistributedDate[tokenAddr] == 0) {
            doneDistributed[tokenAddr] = true;
            currentDistribute[tokenAddr] = 0;
        }

        //If state is already distributed
        if (doneDistributed[tokenAddr]) {
            //Check currentTime and last Distributed
            require(currentDistribute > currentDistributedDate[tokenAddr], "current date less than currentDistributedDate");
            // reset all
            doneDistributed[tokenAddr] = false;
            currentDistribute[tokenAddr] = 0;
        }
        currentDistributedDate[tokenAddr] = currentDate;
    }


    function executor(address tokenAddr) public nonReentrant(tokenAddr) isDistributor(tokenAddr) isNotDone(tokenAddr) enoughBalance(tokenAddr) {
        uint256 memory tt = currentDistributedAddr[tokenAddr].length;
        for (uint256 i = 0; i < tt ; i += 1) {
            uint256 amount = currentDistributedAmount[tokenAddr].pop();
            address addr = currentDistributedAddr[tokenAddr].pop();
            bool memory success = false;
            if (tokenAddr == address(0)) {
                (success,) = payable(addr).call{value : amount}(
                    ""
                );
            } else {
                IERC20 erc20Contract = _getERC20Contract(tokenAddr);
                (success,) = erc20Contract.transfer(addr, amount);
            }
            if (!success) {
                emit DistributeFail(tokenAddr, addr, amount);
            }
        }
        doneDistributed[tokenAddr] = true;
    }

}
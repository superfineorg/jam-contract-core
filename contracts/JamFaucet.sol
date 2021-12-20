pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./HasNoEther.sol";

contract JamFaucet is HasNoEther {
    using SafeMath for uint256;

    uint256 public faucetWei;

    uint256 public faucetInterval;

    mapping(address => uint256) lastFaucet;

    
    modifier canFaucet(address addr) {
        require(lastFaucet[addr] + faucetInterval < block.timestamp, "requestFaucet too fast");
        require(address(this).balance > faucetWei * 2, "out of balance");
        _;
    }

    constructor() payable {
    }

    receive() external payable {
    }

    function setFaucetWei(uint256 _faucetWei) public onlyOwner {
        faucetWei = _faucetWei;
    }
    function setFaucetInterval(uint256 _faucetInterval) public onlyOwner {
        faucetInterval = _faucetInterval;
    }

    function faucet(address addr) public canFaucet(addr) {
        (bool success, ) = payable(addr).call{value: faucetWei}(
            ""
        );
        require(success, "Faucet failed.");
        lastFaucet[addr] = block.timestamp;
    }

    function getLastFaucet(address addr) public view returns (uint256) {
        return lastFaucet[addr];
    }
}
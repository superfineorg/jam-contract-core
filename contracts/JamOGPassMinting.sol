/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./tokens/ERC721/JamOGPass.sol";
import "./tokens/ERC721/JamSuperHappyFrens.sol";

contract JamOGPassMinting is Ownable {
    JamOGPass public jamOGPassContract;
    JamSuperHappyFrens public jamSuperHappyFrensContract;
    uint256 public mintingFee;
    uint256 public jamSuperHappyFrensPrice;
    uint256 public jamSuperHappyFrensDiscountPrice;

    constructor(
        address jamOGPass,
        address jamSuperHappyFrens,
        uint256 mintingFee_,
        uint256 jamSuperHappyFrensPrice_,
        uint256 jamSuperHappyFrensDiscountPrice_
    ) {
        require(
            jamSuperHappyFrensPrice_ > jamSuperHappyFrensDiscountPrice_,
            "JamOGPassMinting: normal price must be greater than discount price"
        );
        jamOGPassContract = JamOGPass(jamOGPass);
        jamSuperHappyFrensContract = JamSuperHappyFrens(jamSuperHappyFrens);
        mintingFee = mintingFee_;
        jamSuperHappyFrensPrice = jamSuperHappyFrensPrice_;
        jamSuperHappyFrensDiscountPrice = jamSuperHappyFrensDiscountPrice_;
    }

    function mintOGPass() external payable {
        require(
            msg.value >= mintingFee,
            "JamOGPassMinting: not enough fee to mint"
        );
        jamOGPassContract.mintTo(msg.sender);
        (bool success, ) = payable(msg.sender).call{
            value: msg.value - mintingFee
        }("");
        require(success, "JamOGPassMinting: return fee failed");
    }

    function buySuperHappyFrens(uint256 quantity) external payable {
        require(quantity > 0, "JamOGPassMinting: cannot mint zero NFT");
        uint256 price = jamSuperHappyFrensPrice;
        if (jamOGPassContract.balanceOf(msg.sender) > 0)
            price = jamSuperHappyFrensDiscountPrice;
        require(
            msg.value >= price * quantity,
            "JamOGPassMinting: not enough money"
        );
        for (uint256 i = 0; i < quantity; i++)
            jamSuperHappyFrensContract.mintTo(msg.sender);
        (bool success, ) = payable(msg.sender).call{
            value: msg.value - price * quantity
        }("");
        require(success, "JamOGPassMinting: return money failed");
    }

    function reclaimEther(address payable recipient) external onlyOwner {
        require(
            recipient != address(0),
            "JamOGPassMinting: cannot reclaim to zero address"
        );
        (bool success, ) = recipient.call{value: address(this).balance}("");
        require(success, "JamOGPassMinting: reclaim failed");
    }
}

/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./tokens/ERC721/JamOGPass.sol";
import "./tokens/ERC721/JamSuperHappyFrens.sol";

contract JamOGPassMinting is Ownable {
    JamOGPass public jamOGPassContract;
    JamSuperHappyFrens public jamSuperHappyFrensContract;
    uint256 public mintFee;
    uint256 public jamSuperHappyFrensPrice;
    uint256 public jamSuperHappyFrensDiscountPrice;
    uint256 public mintLimit;
    uint256 public purchaseLimit;

    constructor(
        address jamOGPass,
        address jamSuperHappyFrens,
        uint256 mintFee_,
        uint256 jamSuperHappyFrensPrice_,
        uint256 jamSuperHappyFrensDiscountPrice_,
        uint256 mintLimit_,
        uint256 purchaseLimit_
    ) {
        require(
            jamSuperHappyFrensPrice_ > jamSuperHappyFrensDiscountPrice_,
            "JamOGPassMinting: normal price must be greater than discount price"
        );
        require(
            mintLimit_ > 0,
            "JamOGPassMinting: mint limit must be greater than 0"
        );
        require(
            purchaseLimit_ > 0,
            "JamOGPassMinting: purchase limit must be greater than 0"
        );
        jamOGPassContract = JamOGPass(jamOGPass);
        jamSuperHappyFrensContract = JamSuperHappyFrens(jamSuperHappyFrens);
        mintFee = mintFee_;
        jamSuperHappyFrensPrice = jamSuperHappyFrensPrice_;
        jamSuperHappyFrensDiscountPrice = jamSuperHappyFrensDiscountPrice_;
        mintLimit = mintLimit_;
        purchaseLimit = purchaseLimit_;
    }

    function setMintFee(uint256 newFee) external onlyOwner {
        mintFee = newFee;
    }

    function setSuperHappyFrensPrice(uint256 newPrice) external onlyOwner {
        require(
            newPrice > jamSuperHappyFrensDiscountPrice,
            "JamOGPassMinting: normal price must be greater than discount price"
        );
        jamSuperHappyFrensPrice = newPrice;
    }

    function setSuperHappyFrensDiscountPrice(uint256 newPrice)
        external
        onlyOwner
    {
        require(
            newPrice < jamSuperHappyFrensPrice,
            "JamOGPassMinting: normal price must be greater than discount price"
        );
        jamSuperHappyFrensDiscountPrice = newPrice;
    }

    function setMintLimit(uint256 newLimit) external onlyOwner {
        require(
            newLimit > 0,
            "JamOGPassMinting: mint limit must be greater than 0"
        );
        mintLimit = newLimit;
    }

    function setPurchaseLimit(uint256 newLimit) external onlyOwner {
        require(
            newLimit > 0,
            "JamOGPassMinting: purchase limit must be greater than 0"
        );
        purchaseLimit = newLimit;
    }

    function mintOGPass(uint256 quantity) external payable {
        require(quantity > 0, "JamOGPassMinting: cannot mint zero OGPass");
        require(
            quantity <= mintLimit,
            "JamOGPassMinting: quantity exceeds limit"
        );
        require(
            msg.value >= mintFee * quantity,
            "JamOGPassMinting: not enough fee to mint"
        );
        for (uint256 i = 0; i < quantity; i++)
            jamOGPassContract.mintTo(msg.sender);
        (bool success, ) = payable(msg.sender).call{
            value: msg.value - mintFee * quantity
        }("");
        require(success, "JamOGPassMinting: return fee failed");
    }

    function purchaseSuperHappyFrens(uint256 quantity) external payable {
        require(quantity > 0, "JamOGPassMinting: cannot purchase zero NFT");
        require(
            quantity <= purchaseLimit,
            "JamOGPassMinting: quantity exceeds limit"
        );
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

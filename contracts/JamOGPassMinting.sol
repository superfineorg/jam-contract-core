/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.0;

import "./tokens/ERC721/JamOGPass.sol";
import "./tokens/ERC721/JamSuperHappyFrens.sol";

contract JamOGPassMinting {
    JamOGPass public jamOGPassContract;
    JamSuperHappyFrens public jamSuperHappyFrensContract;
    uint256 public mintingFee;

    constructor(
        address jamOGPass,
        address jamSuperHappyFrens,
        uint256 mintingFee_
    ) {
        jamOGPassContract = JamOGPass(jamOGPass);
        jamSuperHappyFrensContract = JamSuperHappyFrens(jamSuperHappyFrens);
        mintingFee = mintingFee_;
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

    function exchangeOGPass(uint256 id) external {
        jamOGPassContract.burn(id);
        jamSuperHappyFrensContract.mint(msg.sender, id, "");
    }
}

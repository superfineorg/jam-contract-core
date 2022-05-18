/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./JamMarketplaceHelpers.sol";
import "./JamNFTOwners.sol";

contract JamMarketplaceHub is Ownable {
    address private _jamNFTOwners;
    mapping(bytes32 => address) private _marketplaceIdToAddress;
    mapping(address => bytes32) private _marketplaceAddressToId;

    constructor(address nftOwnersContract) {
        _jamNFTOwners = nftOwnersContract;
    }

    function isMarketplace(address addr) external view returns (bool) {
        return _marketplaceAddressToId[addr] != 0x0;
    }

    function getMarketplace(bytes32 id) external view returns (address) {
        return _marketplaceIdToAddress[id];
    }

    function registerMarketplace(bytes32 id) external {
        require(
            Ownable(msg.sender).owner() == owner(),
            "JamMarketplaceHub: invalid caller contract"
        );
        require(id != 0x0, "JamMarketplaceHub: invalid marketplace id");
        _marketplaceIdToAddress[id] = msg.sender;
        _marketplaceAddressToId[msg.sender] = id;
    }

    function unregisterMarketplace(bytes32 id) external onlyOwner {
        delete _marketplaceAddressToId[_marketplaceIdToAddress[id]];
        delete _marketplaceIdToAddress[id];
    }

    function unregisterMarketplace(address addr) external onlyOwner {
        delete _marketplaceIdToAddress[_marketplaceAddressToId[addr]];
        delete _marketplaceAddressToId[addr];
    }

    function setRoyaltyFee(
        address nftAddress,
        bytes32[] memory marketplaceIds,
        address[] memory recipients,
        uint256[] memory percentages
    ) external {
        require(
            msg.sender == JamNFTOwners(_jamNFTOwners).getNFTOwner(nftAddress),
            "JamMarketplaceHub: caller is not NFT owner"
        );
        require(
            marketplaceIds.length == recipients.length &&
                marketplaceIds.length == percentages.length,
            "JamMarketplaceHub: lengths mismatch"
        );
        for (uint256 i = 0; i < marketplaceIds.length; i++) {
            address marketplaceAddr = _marketplaceIdToAddress[
                marketplaceIds[i]
            ];
            require(
                _marketplaceAddressToId[marketplaceAddr] != 0x0,
                "JamMarketplaceHub: invalid marketplace id"
            );
            JamMarketplaceHelpers(marketplaceAddr).setRoyaltyFee(
                nftAddress,
                recipients[i],
                percentages[i]
            );
        }
    }
}

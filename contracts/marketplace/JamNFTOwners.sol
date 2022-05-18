/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";

contract JamNFTOwners is Ownable {
    mapping(address => bool) private _operators;
    mapping(address => address) private _ownerOf;

    constructor() Ownable() {}

    modifier onlyOperators() {
        require(_operators[msg.sender], "JamNFTOwners: caller is not operator");
        _;
    }

    function getNFTOwner(address nftAddress) external view returns (address) {
        return _ownerOf[nftAddress];
    }

    function setOperators(address[] memory operators, bool[] memory isOperators)
        external
        onlyOwner
    {
        require(
            operators.length == isOperators.length,
            "JamNFTOwners: lengths mismatch"
        );
        for (uint256 i = 0; i < operators.length; i++)
            _operators[operators[i]] = isOperators[i];
    }

    function setNFTOwners(
        address[] memory nftAddresses,
        address[] memory nftOwners
    ) external onlyOperators {
        require(
            nftAddresses.length == nftOwners.length,
            "JamNFTOwners: lengths mismatch"
        );
        for (uint256 i = 0; i < nftAddresses.length; i++)
            _ownerOf[nftAddresses[i]] = nftOwners[i];
    }
}

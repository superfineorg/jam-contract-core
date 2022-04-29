/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract GameNFTBurning is Ownable {
    event BurnedIntoGames(address indexed nftAddress, uint256 indexed tokenId);

    function burnIntoGames(address nftAddress, uint256 tokenId)
        external
        onlyOwner
    {
        bool success = ERC165Checker.supportsERC165(nftAddress);
        if (success)
            success = IERC165(nftAddress).supportsInterface(
                type(IERC721).interfaceId
            );
        require(success, "GameNFTBurning: invalid NFT address");
        ERC721Burnable(nftAddress).burn(tokenId);
        emit BurnedIntoGames(nftAddress, tokenId);
    }
}

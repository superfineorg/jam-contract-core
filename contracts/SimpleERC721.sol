pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "./Ownable.sol";

contract UniqueItem is ERC721Enumerable, Ownable, IERC2981 {
    uint256 _tokenIds;
    string public tokenURIPrefix = "https://asset.gamejam.co/gamejam-nft/erc/721/card/";
    string public tokenURISuffix = ".json";
    address public nftAddress;

    /**
      * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event AwardItem(address indexed player, uint256 indexed tokenId);

    constructor(address _owner, string memory name, string memory symbol) ERC721(name, symbol) payable {
        owner = payable(_owner);
        nftAddress = address(this);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC165, IERC165) returns (bool) {
        return
        interfaceId == type(IERC721).interfaceId ||
    interfaceId == type(IERC721).interfaceId ||
    interfaceId == type(IERC721Metadata).interfaceId ||
    super.supportsInterface(interfaceId);
    }

    function setTokenURIPrefix(string calldata _tokenURIPrefix) public onlyOwner returns (string calldata) {
        tokenURIPrefix = _tokenURIPrefix;
        return _tokenURIPrefix;
    }

    function setTokenURISuffix(string calldata _tokenURISuffix) public onlyOwner returns (string calldata) {
        tokenURISuffix = _tokenURISuffix;
        return _tokenURISuffix;
    }

    function awardItem(address player) public onlyOwner returns (uint256) {
        _tokenIds += 1;
        _mint(player, _tokenIds);
        // _setTokenURI(_tokenIds, tokenURI);
        emit AwardItem(player, _tokenIds);
        return _tokenIds;
    }

    function tokenURI(uint256 _tokenId) override public view returns (string memory) {
        bytes memory _tokenURIPrefixBytes = bytes(tokenURIPrefix);
        bytes memory _tokenURISuffixBytes = bytes(tokenURISuffix);
        uint256 _tmpTokenId = _tokenId;
        uint256 _length;

        do {
            _length++;
            _tmpTokenId /= 10;
        }
        while (_tmpTokenId > 0);

        bytes memory _tokenURIBytes = new bytes(_tokenURIPrefixBytes.length + _length + 5);
        uint256 _i = _tokenURIBytes.length - 6;

        _tmpTokenId = _tokenId;

        do {
            _tokenURIBytes[_i--] = bytes1(uint8(48 + _tmpTokenId % 10));
            _tmpTokenId /= 10;
        }
        while (_tmpTokenId > 0);

        for (_i = 0; _i < _tokenURIPrefixBytes.length; _i++) {
            _tokenURIBytes[_i] = _tokenURIPrefixBytes[_i];
        }

        for (_i = 0; _i < _tokenURISuffixBytes.length; _i++) {
            _tokenURIBytes[_tokenURIBytes.length + _i - 5] = _tokenURISuffixBytes[_i];
        }

        return string(_tokenURIBytes);
    }

    function tokenOfOwner(address owner) public view returns (uint256[] memory) {
        uint256 totalToken = ERC721.balanceOf(owner);
        uint256[] memory tokenIds = new uint256[](totalToken);
        for (uint256 i = 0; i < totalToken; i += 1) {
            uint256 tmp = ERC721Enumerable.tokenOfOwnerByIndex(owner, i);
            tokenIds[i] = tmp;
            //tokenIds.push(tmp);
        }
        return tokenIds;
    }
}
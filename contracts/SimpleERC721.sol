pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./Ownable.sol";

contract UniqueItem is ERC721Enumerable, Ownable, IERC2981 {
    using SafeMath for uint256;
    uint256 _tokenIds;
    string public tokenURIPrefix = "https://asset.gamejam.co/gamejam-nft/erc/721/card/";
    string public tokenURISuffix = ".json";
    address public nftAddress;
    
    mapping(uint256 => address) public firstOwner;

    // royalty fee takes on each auction, measured in basis points (1/100 of a percent).
   // Values 0-10,00 map to 0%-10%
    uint64 public royaltyFee;


    /**
      * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event AwardItem(address indexed player, uint256 indexed tokenId);
    /**
      * @dev Emitted when owner of contract update royaltyFee from `previous_fee` to `current_fee`
     */
    event UpdateRoyaltyFee(uint64 previous_fee, uint64 current_fee);

    constructor(address _owner, string memory name, string memory symbol) ERC721(name, symbol) payable {
        owner = payable(_owner);
        nftAddress = address(this);
        royaltyFee = 0;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721Enumerable, IERC165) returns (bool) {
        bool b = interfaceId == type(IERC721).interfaceId ||
    interfaceId == type(IERC721Metadata).interfaceId ||
    (interfaceId == type(IERC2981).interfaceId && royaltyFee > 0 ) ||
    super.supportsInterface(interfaceId);
        return b;
    }

    function royaltyInfo(uint256 tokenId, uint256 salePrice)
    external
    view
    virtual override(IERC2981)
    returns (address receiver, uint256 royaltyAmount) {
        receiver = firstOwner[tokenId];
        royaltyAmount = salePrice.mul(royaltyFee).div(10000);
        return (receiver, royaltyAmount);
    }

    function setRoyaltyFee(uint64 _royaltyFee) public onlyOwner {
        require(_royaltyFee >= 0 && _royaltyFee <= 1000);
        uint64 previousFee = royaltyFee;
        royaltyFee = _royaltyFee;
        emit UpdateRoyaltyFee(previousFee, royaltyFee);
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
        uint256 _ttId = _tokenIds += 1;
        _safeMint(player, _ttId);
        emit AwardItem(player, _ttId);
        firstOwner[_ttId] = player;
        return _ttId;
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

    function formatStartIndex(uint256 _from_index) internal pure returns (uint256) {
        return _from_index > 0 ? _from_index : 0; 
    }
    
    function formatEndIndex(uint256 _end_index, uint256 _max) internal pure returns (uint256) {
        return _end_index > _max ? _max : _end_index; 
    }

    function tokenOfOwner(address owner, uint256 _from_index, uint256 _end_index) public view returns (uint256[] memory) {
        uint256 totalToken = ERC721.balanceOf(owner);
        _from_index = formatStartIndex (_from_index);
        _end_index = formatEndIndex(_end_index, totalToken);
        uint256[] memory tokenIds = new uint256[](_end_index-_from_index);
        for (uint256 i = _from_index; i < _end_index; i += 1) {
            uint256 tmp = ERC721Enumerable.tokenOfOwnerByIndex(owner, i);
            tokenIds[i-_from_index] = tmp;
        }
        return tokenIds;
    }
}
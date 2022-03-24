// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/interfaces/IERC165.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../utils/HasNoEther.sol";

contract JamMarketplace is HasNoEther, Pausable, ReentrancyGuard {
    using SafeMath for uint256;

    // Represents an auction on an NFT
    struct Auction {
        address seller;
        uint128 price;
        address erc20Address; // erc20Address, address(0) if place a price with (wei)
        address nftAddress;
        uint256 tokenId;
        uint64 startedAt;
    }

    uint256 withdrawDuration = 14 days;

    // Cut owner takes on each auction, measured in basis points (1/100 of a percent).
    // Values 0-10,000 map to 0%-100%
    uint256 public ownerCut;

    // The total amount of royalty cut which cannot be reclaimed by the owner of the contract
    mapping(address => uint256) private _totalRoyaltyCut;

    // Map from an NFT to its corresponding auction.
    mapping(address => mapping(uint256 => Auction)) public auctions;

    // mapping royaltyCut amount for each (address, erc20Address) pair,
    // erc20Address = address(0) mean it value is wei
    mapping(address => mapping(address => uint256)) private royaltyCuts;
    mapping(address => mapping(address => uint256)) private lastWithdraws;

    event AuctionCreated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 price,
        address seller,
        address erc20Address
    );

    event AuctionSuccessful(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 price,
        address winner
    );

    event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId);

    // Modifiers to check that inputs can be safely stored with a certain
    // number of bits. We use constants and multiple modifiers to save gas.
    modifier canBeStoredWith64Bits(uint256 _value) {
        require(
            _value <= type(uint64).max,
            "JamMarketplace: cannot be stored with 64 bits"
        );
        _;
    }

    modifier canBeStoredWith128Bits(uint256 _value) {
        require(
            _value < type(uint128).max,
            "JamMarketplace: cannot be stored with 128 bits"
        );
        _;
    }

    constructor(uint256 _ownerCut) {
        require(
            _ownerCut <= 10000,
            "JamMarketplace: owner cut cannot exceed 100%"
        );
        ownerCut = _ownerCut;
    }

    function updateOwnerCut(uint256 _ownerCut) public onlyOwner {
        require(
            _ownerCut <= 10000,
            "JamMarketplace: owner cut cannot exceed 100%"
        );
        ownerCut = _ownerCut;
    }

    function _computeCut(uint256 _price) internal view returns (uint256) {
        // NOTE: We don't use SafeMath (or similar) in this function because
        //  all of our entry functions carefully cap the maximum values for
        //  currency (at 128-bits), and ownerCut <= 10000 (see the require()
        //  statement in the ClockAuction constructor). The result of this
        //  function is always guaranteed to be <= _price.
        return (_price * ownerCut) / 10000;
    }

    function _isOnAuction(Auction storage _auction)
        internal
        view
        returns (bool)
    {
        return (_auction.startedAt > 0);
    }

    function _getNftContract(address _nftAddress)
        internal
        pure
        returns (IERC721)
    {
        IERC721 candidateContract = IERC721(_nftAddress);
        return candidateContract;
    }

    function _supportIERC2981(address _nftAddress)
        internal
        view
        returns (bool)
    {
        bool success;
        success = ERC165Checker.supportsERC165(_nftAddress);
        if (success) {
            success = IERC165(_nftAddress).supportsInterface(
                type(IERC2981).interfaceId
            );
        }
        return success;
    }

    function _getERC2981(address _nftAddress) internal pure returns (IERC2981) {
        IERC2981 candidateContract = IERC2981(_nftAddress);
        return candidateContract;
    }

    function _getERC20Contract(address _erc20Address)
        internal
        pure
        returns (IERC20)
    {
        IERC20 candidateContract = IERC20(_erc20Address);
        return candidateContract;
    }

    function _owns(
        address _nftAddress,
        address _claimant,
        uint256 _tokenId
    ) internal view returns (bool) {
        IERC721 _nftContract = _getNftContract(_nftAddress);
        return (_nftContract.ownerOf(_tokenId) == _claimant);
    }

    function _transfer(
        address _nftAddress,
        address _receiver,
        uint256 _tokenId
    ) internal {
        IERC721 _nftContract = _getNftContract(_nftAddress);

        // It will throw if transfer fails
        _nftContract.transferFrom(address(this), _receiver, _tokenId);
    }

    function _addAuction(
        address _nftAddress,
        uint256 _tokenId,
        Auction memory _auction,
        address _seller
    ) internal {
        auctions[_nftAddress][_tokenId] = _auction;

        emit AuctionCreated(
            _nftAddress,
            _tokenId,
            uint256(_auction.price),
            _seller,
            address(_auction.erc20Address)
        );
    }

    function _escrow(
        address _nftAddress,
        address _owner,
        uint256 _tokenId
    ) internal {
        IERC721 _nftContract = _getNftContract(_nftAddress);

        // It will throw if transfer fails
        _nftContract.transferFrom(_owner, address(this), _tokenId);
    }

    /// @dev Removes an auction from the list of open auctions.
    /// @param _tokenId - ID of NFT on auction.
    function _removeAuction(address _nftAddress, uint256 _tokenId) internal {
        delete auctions[_nftAddress][_tokenId];
    }

    /// @dev Cancels an auction unconditionally.
    function _cancelAuction(
        address _nftAddress,
        uint256 _tokenId,
        address _seller
    ) internal {
        _removeAuction(_nftAddress, _tokenId);
        _transfer(_nftAddress, _seller, _tokenId);
        emit AuctionCancelled(_nftAddress, _tokenId);
    }

    function getAuction(address _nftAddress, uint256 _tokenId)
        external
        view
        returns (
            address seller,
            uint256 price,
            address erc20Address,
            uint256 startedAt
        )
    {
        Auction storage _auction = auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace: not on auction");
        return (
            _auction.seller,
            _auction.price,
            _auction.erc20Address,
            _auction.startedAt
        );
    }

    function getWithdrawInfo(address[] memory _contractAddress)
        external
        view
        returns (
            address[] memory contractAddress,
            uint256[] memory balances,
            uint256[] memory lastWithdraw
        )
    {
        uint256[] memory _balances = new uint256[](_contractAddress.length);
        uint256[] memory _lastWithdraw = new uint256[](_contractAddress.length);
        for (uint256 i = 0; i < _contractAddress.length; i += 1) {
            address ad = _contractAddress[i];
            uint256 balance = royaltyCuts[msg.sender][ad];
            _balances[i] = balance;
            uint256 lw = lastWithdraws[msg.sender][ad];
            _lastWithdraw[i] = lw;
        }
        return (_contractAddress, balances, lastWithdraw);
    }

    /// @dev Create an auction.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT.
    /// @param _price - Price of erc20 address.
    /// @param _erc20Address - Address of price need to be list on.
    function createAuction(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _price,
        address _erc20Address
    ) external whenNotPaused canBeStoredWith128Bits(_price) {
        address _seller = msg.sender;
        require(
            _owns(_nftAddress, _seller, _tokenId),
            "JamMarketplace: only owner can create auction"
        );
        _escrow(_nftAddress, _seller, _tokenId);

        Auction memory _auction = Auction(
            _seller,
            uint128(_price),
            _erc20Address,
            _nftAddress,
            _tokenId,
            uint64(block.timestamp)
        );
        _addAuction(_nftAddress, _tokenId, _auction, _seller);
    }

    /// @dev Cancels an auction.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT on auction to cancel.
    function cancelAuction(address _nftAddress, uint256 _tokenId) external {
        Auction storage _auction = auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace: not on auction");
        require(
            msg.sender == _auction.seller,
            "JamMarketplace: only seller can cancel auction"
        );
        _cancelAuction(_nftAddress, _tokenId, _auction.seller);
    }

    /// @dev Cancels an auction when the contract is paused.
    ///  Only the owner may do this, and NFTs are returned to
    ///  the seller. This should only be used in emergencies.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT on auction to cancel.
    function cancelAuctionWhenPaused(address _nftAddress, uint256 _tokenId)
        external
        whenPaused
        onlyOwner
    {
        Auction storage _auction = auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace: not on auction");
        _cancelAuction(_nftAddress, _tokenId, _auction.seller);
    }

    function buy(address _nftAddress, uint256 _tokenId)
        external
        payable
        whenNotPaused
    {
        Auction memory auction = auctions[_nftAddress][_tokenId];
        require(
            auction.erc20Address == address(0),
            "JamMarketplace: cannot buy this item with native token"
        );
        // _bid will throw if the bid or funds transfer fails
        _buy(_nftAddress, _tokenId, msg.value);
        _transfer(_nftAddress, msg.sender, _tokenId);
    }

    function buyByERC20(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _maxPrice
    ) external whenNotPaused nonReentrant {
        Auction storage _auction = auctions[_nftAddress][_tokenId];
        require(
            _auction.erc20Address != address(0),
            "JamMarketplace: can only be paid with native token"
        );
        require(_isOnAuction(_auction), "JamMarketplace: not on auction");
        uint256 _price = _auction.price;
        require(_price <= _maxPrice, "JamMarketplace: item price is too high");
        uint256 tokenId = _tokenId;
        address _erc20Address = _auction.erc20Address;

        IERC20 erc20Contract = _getERC20Contract(_erc20Address);

        uint256 _allowance = erc20Contract.allowance(msg.sender, address(this));

        require(_allowance >= _price, "JamMarketplace: not enough allowance");

        address _seller = _auction.seller;
        uint256 _auctioneerCut = _computeCut(_price);
        uint256 _sellerProceeds = _price - _auctioneerCut;
        bool success = erc20Contract.transferFrom(
            msg.sender,
            address(this),
            _price
        );
        require(success, "JamMarketplace: not enough balance");
        erc20Contract.transfer(_seller, _sellerProceeds);
        if (_supportIERC2981(_nftAddress)) {
            IERC2981 royaltyContract = _getERC2981(_nftAddress);
            (address firstOwner, uint256 amount) = royaltyContract.royaltyInfo(
                tokenId,
                _auctioneerCut
            );
            require(
                amount < _auctioneerCut,
                "JamMarketplace: royalty amount must be less than auctioneer cut"
            );
            _totalRoyaltyCut[_erc20Address] = _totalRoyaltyCut[_erc20Address]
                .add(amount);
            royaltyCuts[firstOwner][_erc20Address] = royaltyCuts[firstOwner][
                _erc20Address
            ].add(amount);
        }
        _transfer(_nftAddress, msg.sender, _tokenId);
        _removeAuction(_nftAddress, _tokenId);
    }

    /// @dev Computes the price and transfers winnings.
    /// Does NOT transfer ownership of token.
    function _buy(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _bidAmount
    ) internal returns (uint256) {
        // Get a reference to the auction struct
        Auction storage _auction = auctions[_nftAddress][_tokenId];

        require(_isOnAuction(_auction), "JamMarketplace: not on auction");
        uint256 _price = _auction.price;

        require(
            _bidAmount >= _price,
            "JamMarketplace: bid amount cannot be less than price"
        );

        address _seller = _auction.seller;

        _removeAuction(_nftAddress, _tokenId);

        if (_price > 0) {
            //  Calculate the auctioneer's cut.
            // (NOTE: _computeCut() is guaranteed to return a
            //  value <= price, so this subtraction can't go negative.)
            uint256 _auctioneerCut = _computeCut(_price);
            uint256 _sellerProceeds = _price - _auctioneerCut;
            payable(_seller).transfer(_sellerProceeds);
            if (_supportIERC2981(_nftAddress)) {
                IERC2981 royaltyContract = _getERC2981(_nftAddress);
                (address firstOwner, uint256 amount) = royaltyContract
                    .royaltyInfo(_tokenId, _auctioneerCut);
                require(
                    amount < _auctioneerCut,
                    "JamMarketplace: royalty amount must be less than auctioneer cut"
                );
                _totalRoyaltyCut[address(0)] = _totalRoyaltyCut[address(0)].add(
                    amount
                );
                royaltyCuts[firstOwner][address(0)] = royaltyCuts[firstOwner][
                    address(0)
                ].add(amount);
            }
        }

        if (_bidAmount > _price) {
            // Calculate any excess funds included with the bid. If the excess
            // is anything worth worrying about, transfer it back to bidder.
            // NOTE: We checked above that the bid amount is greater than or
            // equal to the price so this cannot underflow.
            uint256 _bidExcess = _bidAmount - _price;

            // Return the funds. Similar to the previous transfer, this is
            // not susceptible to a re-entry attack because the auction is
            // removed before any transfers occur.
            payable(msg.sender).transfer(_bidExcess);
        }

        // Tell the world!
        emit AuctionSuccessful(_nftAddress, _tokenId, _price, msg.sender);

        return _price;
    }

    function withdraw(address _erc20Address) public {
        uint256 lastWithdraw = lastWithdraws[msg.sender][_erc20Address];
        require(
            lastWithdraw + withdrawDuration <= block.timestamp,
            "JamMarketplace: only withdraw after 14 days before previous withdraw"
        );
        uint256 royaltyCut = royaltyCuts[msg.sender][_erc20Address];
        require(royaltyCut > 0, "JamMarketplace: no royalty cut to withdraw");
        royaltyCuts[msg.sender][_erc20Address] = 0;
        _totalRoyaltyCut[_erc20Address] = _totalRoyaltyCut[_erc20Address].sub(
            royaltyCut
        );
        lastWithdraws[msg.sender][_erc20Address] = block.timestamp;
        if (_erc20Address == address(0)) {
            payable(msg.sender).transfer(royaltyCut);
        } else {
            IERC20 erc20Contract = _getERC20Contract(_erc20Address);
            bool success = erc20Contract.transfer(msg.sender, royaltyCut);
            require(success, "JamMarketplace: transfer failed");
        }
    }

    // function reclaimEther() external override onlyOwner {
    //     (bool success, ) = payable(owner()).call{
    //         value: address(this).balance.sub(_totalRoyaltyCut[address(0)])
    //     }("");
    //     require(success, "JamMarketplace: reclaim Ether failed");
    // }

    function reclaimERC20(address _erc20Address) external onlyOwner {
        IERC20 erc20Contract = _getERC20Contract(_erc20Address);
        erc20Contract.transfer(
            owner(),
            erc20Contract.balanceOf(address(this)).sub(
                _totalRoyaltyCut[_erc20Address]
            )
        );
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }
}

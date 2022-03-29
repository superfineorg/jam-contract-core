/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./JamMarketplaceHelpers.sol";

contract JamClockAuction is JamMarketplaceHelpers, ReentrancyGuard {
    // The information of an auction
    struct Auction {
        address seller;
        uint128 startingPrice;
        uint128 endingPrice;
        uint64 duration;
        uint64 startedAt;
    }

    // Mapping from (NFT address + token ID) to the NFT's auction.
    mapping(address => mapping(uint256 => Auction)) public auctions;

    event AuctionCreated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration,
        address seller
    );

    event AuctionUpdated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration,
        address seller
    );

    event AuctionSuccessful(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 totalPrice,
        address winner
    );

    event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId);

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_CLOCK_AUCTION");
    }

    receive() external payable {}

    // Modifiers to check that inputs can be safely stored with a certain number of bits
    modifier canBeStoredWith64Bits(uint256 value) {
        require(
            value <= 18446744073709551615,
            "JamClockAuction: value is longer than 64 bit"
        );
        _;
    }

    modifier canBeStoredWith128Bits(uint256 value) {
        require(
            value < 340282366920938463463374607431768211455,
            "JamClockAuction: value is longer than 128 bit"
        );
        _;
    }

    /**
     * @dev Returns auction info for an NFT currently on auction.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of NFT on auction.
     */
    function getAuction(address nftAddress, uint256 tokenId)
        external
        view
        returns (
            address seller,
            uint256 startingPrice,
            uint256 endingPrice,
            uint256 duration,
            uint256 startedAt
        )
    {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        return (
            auction.seller,
            auction.startingPrice,
            auction.endingPrice,
            auction.duration,
            auction.startedAt
        );
    }

    /**
     * @dev Returns the current price of an auction.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the token price we are checking.
     */
    function getCurrentPrice(address nftAddress, uint256 tokenId)
        external
        view
        returns (uint256)
    {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        return _getCurrentPrice(auction);
    }

    /**
     * @dev Check if this auction is currently cancelable
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction
     */
    function isAuctionCancelable(address nftAddress, uint256 tokenId)
        external
        pure
        override
        returns (bool)
    {
        return true;
    }

    /**
     * @dev Creates and begins a new auction.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endingPrice - Price of item (in wei) at end of auction.
     * @param duration - Length of time to move between starting price and ending price (in seconds).
     */
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration
    )
        external
        whenNotPaused
        canBeStoredWith128Bits(startingPrice)
        canBeStoredWith128Bits(endingPrice)
        canBeStoredWith64Bits(duration)
    {
        address seller = msg.sender;
        require(
            _owns(nftAddress, seller, tokenId),
            "JamClockAuction: sender is not owner of NFT"
        );
        _escrow(nftAddress, seller, tokenId);
        Auction memory auction = Auction(
            seller,
            uint128(startingPrice),
            uint128(endingPrice),
            uint64(duration),
            uint64(block.timestamp)
        );
        _addAuction(nftAddress, tokenId, auction, seller);
    }

    /**
     * @dev Bids on an open auction, completing the auction and transferring ownership of the NFT if enough Ether is supplied.
     * @param nftAddress - address of a deployed contract implementing the Nonfungible Interface.
     * @param tokenId - ID of token to bid on.
     */
    function bid(address nftAddress, uint256 tokenId)
        external
        payable
        whenNotPaused
        nonReentrant
    {
        _bid(nftAddress, tokenId, msg.value);
        _transfer(nftAddress, msg.sender, tokenId);
    }

    /**
     * @dev Updates an existent auction's information.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endingPrice - Price of item (in wei) at end of auction.
     * @param duration - Length of time to move between starting price and ending price (in seconds).
     */
    function updateAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration
    )
        external
        whenNotPaused
        canBeStoredWith128Bits(startingPrice)
        canBeStoredWith128Bits(endingPrice)
        canBeStoredWith64Bits(duration)
    {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        require(
            msg.sender == auction.seller,
            "JamClockAuction: only seller can update auction"
        );
        auction.startingPrice = uint128(startingPrice);
        auction.endingPrice = uint128(endingPrice);
        auction.duration = uint64(duration);
        emit AuctionUpdated(
            nftAddress,
            tokenId,
            startingPrice,
            endingPrice,
            duration,
            msg.sender
        );
    }

    /**
     * @dev Cancels an auction that hasn't been won yet.
     * Returns the NFT to original owner.
     * @notice This is a state-modifying function that can be called while the contract is paused.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of token on auction
     */
    function cancelAuction(address nftAddress, uint256 tokenId)
        external
        override
    {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        require(
            msg.sender == auction.seller,
            "JamClockAuction: only seller can cancel auction"
        );
        _cancelAuction(nftAddress, tokenId, auction.seller);
    }

    /**
     * @dev Cancels an auction when the contract is paused.
     * Only the owner may do this, and NFTs are returned to
     * the seller. This should only be used in emergencies.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the NFT on auction to cancel.
     */
    function cancelAuctionWhenPaused(address nftAddress, uint256 tokenId)
        external
        whenPaused
        onlyOwner
    {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        _cancelAuction(nftAddress, tokenId, auction.seller);
    }

    /**
     * @dev Returns true if the NFT is on auction.
     * @param auction - Auction to check.
     */
    function _isOnAuction(Auction storage auction)
        internal
        view
        returns (bool)
    {
        return (auction.startedAt > 0);
    }

    /**
     * @dev Gets the NFT object from an address, validating that implementsERC721 is true.
     * @param nftAddress - Address of the NFT.
     */
    function _getNftContract(address nftAddress)
        internal
        pure
        returns (IERC721)
    {
        IERC721 candidateContract = IERC721(nftAddress);
        // require(candidateContract.implementsERC721());
        return candidateContract;
    }

    /**
     * @dev Returns current price of an NFT on auction. Broken into two
     *  functions (this one, that computes the duration from the auction
     *  structure, and the other that does the price computation) so we
     *  can easily test that the price computation works correctly.
     */
    function _getCurrentPrice(Auction storage auction)
        internal
        view
        returns (uint256)
    {
        uint256 secondsPassed = block.timestamp - auction.startedAt;
        return
            _computeCurrentPrice(
                auction.startingPrice,
                auction.endingPrice,
                auction.duration,
                secondsPassed
            );
    }

    /**
     * @dev Computes the current price of an auction. Factored out
     *  from _currentPrice so we can run extensive unit tests.
     *  When testing, make this function external and turn on
     *  `Current price computation` test suite.
     */
    function _computeCurrentPrice(
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration,
        uint256 secondsPassed
    ) internal pure returns (uint256) {
        if (secondsPassed >= duration) return endingPrice;
        else {
            int256 totalPriceChange = int256(endingPrice) -
                int256(startingPrice);
            int256 currentPriceChange = (totalPriceChange *
                int256(secondsPassed)) / int256(duration);
            int256 currentPrice = int256(startingPrice) + currentPriceChange;
            return uint256(currentPrice);
        }
    }

    /**
     * @dev Returns true if the claimant owns the token.
     * @param nftAddress - The address of the NFT.
     * @param claimant - Address claiming to own the token.
     * @param tokenId - ID of token whose ownership to verify.
     */
    function _owns(
        address nftAddress,
        address claimant,
        uint256 tokenId
    ) internal view returns (bool) {
        IERC721 nftContract = _getNftContract(nftAddress);
        return (nftContract.ownerOf(tokenId) == claimant);
    }

    /**
     * @dev Adds an auction to the list of open auctions. Also fires the
     * AuctionCreated event.
     * @param tokenId The ID of the token to be put on auction.
     * @param auction Auction to add.
     */
    function _addAuction(
        address nftAddress,
        uint256 tokenId,
        Auction memory auction,
        address seller
    ) internal {
        require(
            auction.duration >= 1 minutes,
            "JamClockAuction: too short auction"
        );
        auctions[nftAddress][tokenId] = auction;
        emit AuctionCreated(
            nftAddress,
            tokenId,
            uint256(auction.startingPrice),
            uint256(auction.endingPrice),
            uint256(auction.duration),
            seller
        );
    }

    /**
     * @dev Removes an auction from the list of open auctions.
     * @param tokenId - ID of NFT on auction.
     */
    function _removeAuction(address nftAddress, uint256 tokenId) internal {
        delete auctions[nftAddress][tokenId];
    }

    /** @dev Cancels an auction unconditionally. */
    function _cancelAuction(
        address nftAddress,
        uint256 tokenId,
        address seller
    ) internal {
        _removeAuction(nftAddress, tokenId);
        _transfer(nftAddress, seller, tokenId);
        emit AuctionCancelled(nftAddress, tokenId);
    }

    /**
     * @dev Escrows the NFT, assigning ownership to this contract.
     * Throws if the escrow fails.
     * @param nftAddress - The address of the NFT.
     * @param owner - Current owner address of token to escrow.
     * @param tokenId - ID of token whose approval to verify.
     */
    function _escrow(
        address nftAddress,
        address owner,
        uint256 tokenId
    ) internal {
        IERC721 nftContract = _getNftContract(nftAddress);
        nftContract.transferFrom(owner, address(this), tokenId);
    }

    /**
     * @dev Transfers an NFT owned by this contract to another address.
     * Returns true if the transfer succeeds.
     * @param nftAddress - The address of the NFT.
     * @param receiver - Address to transfer NFT to.
     * @param tokenId - ID of token to transfer.
     */
    function _transfer(
        address nftAddress,
        address receiver,
        uint256 tokenId
    ) internal {
        IERC721 nftContract = _getNftContract(nftAddress);
        nftContract.transferFrom(address(this), receiver, tokenId);
    }

    /**
     * @dev Computes owner's cut of a sale.
     * @param price - Sale price of NFT.
     */
    function _computeCut(uint256 price) internal view returns (uint256) {
        return (price * ownerCut) / 10000;
    }

    /**
     * @dev Computes the price and transfers winnings.
     * Does NOT transfer ownership of token.
     */
    function _bid(
        address nftAddress,
        uint256 tokenId,
        uint256 bidAmount
    ) internal returns (uint256) {
        Auction storage auction = auctions[nftAddress][tokenId];
        require(_isOnAuction(auction), "JamClockAuction: auction not exists");
        uint256 price = _getCurrentPrice(auction);
        require(bidAmount >= price, "JamClockAuction: insufficient bid amount");
        address seller = auction.seller;
        _removeAuction(nftAddress, tokenId);
        if (price > 0) {
            uint256 auctioneerCut = _computeCut(price);
            uint256 sellerProceeds = price - auctioneerCut;
            (bool success, ) = payable(seller).call{value: sellerProceeds}("");
            require(success, "JamClockAuction: transfer proceeds failed");
            if (_supportIERC2981(nftAddress)) {
                (address recipient, uint256 amount) = IERC2981(nftAddress)
                    .royaltyInfo(tokenId, auctioneerCut);
                require(
                    amount < auctioneerCut,
                    "JamClockAuction: royalty amount must be less than auctioneer cut"
                );
                _totalRoyaltyCut[address(0)] += amount;
                _royaltyCuts[recipient][address(0)] += amount;
            }
        }
        if (bidAmount > price) {
            uint256 bidExcess = bidAmount - price;
            (bool success, ) = payable(msg.sender).call{value: bidExcess}("");
            require(success, "JamClockAuction: return excess failed");
        }
        emit AuctionSuccessful(nftAddress, tokenId, price, msg.sender);
        return price;
    }
}

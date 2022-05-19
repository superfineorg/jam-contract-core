/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "../JamMarketplaceHelpers.sol";

contract JamClockAuction1155 is JamMarketplaceHelpers, ERC1155Holder {
    using Counters for Counters.Counter;

    // The information of an auction
    struct Auction {
        uint256 auctionId;
        address nftAddress;
        uint256 tokenId;
        uint256 quantity;
        address seller;
        address currency;
        uint128 startingPrice;
        uint128 endingPrice;
        uint64 duration;
        uint64 startedAt;
    }

    // The current ID of auction
    Counters.Counter private _currentAuctionId;

    // Mapping from auction ID to the auction info
    mapping(uint256 => Auction) private _auctions;

    event AuctionCreated(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration,
        address seller
    );

    event AuctionUpdated(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 startingPrice,
        uint256 endingPrice,
        uint256 duration,
        address seller
    );

    event AuctionSuccessful(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 finalPrice,
        address winner
    );

    event AuctionCancelled(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_CLOCK_AUCTION_1155");
    }

    receive() external payable {}

    // Modifiers to check that inputs can be safely stored with a certain number of bits
    modifier canBeStoredWith64Bits(uint256 value) {
        require(
            value <= type(uint64).max,
            "JamClockAuction1155: value is longer than 64 bit"
        );
        _;
    }

    modifier canBeStoredWith128Bits(uint256 value) {
        require(
            value < type(uint128).max,
            "JamClockAuction1155: value is longer than 128 bit"
        );
        _;
    }

    /**
     * @dev Returns auction info for an NFT currently on auction.
     * @param auctionId - ID of the auction.
     */
    function getAuction(uint256 auctionId)
        external
        view
        returns (
            address nftAddress,
            uint256 tokenId,
            uint256 quantity,
            address seller,
            address currency,
            uint256 startingPrice,
            uint256 endingPrice,
            uint256 duration,
            uint256 startedAt
        )
    {
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        return (
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.seller,
            auction.currency,
            auction.startingPrice,
            auction.endingPrice,
            auction.duration,
            auction.startedAt
        );
    }

    /**
     * @dev Returns the current price of an auction.
     * @param auctionId - ID of the auction.
     */
    function getCurrentTotalPrice(uint256 auctionId)
        external
        view
        returns (uint256)
    {
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        return _getCurrentTotalPrice(auction);
    }

    /**
     * @dev Creates and begins a new auction.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param quantity - The number of NFTs to sell.
     * @param currency - The address of ERC20 or native token used as payment.
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endingPrice - Price of item (in wei) at end of auction.
     * @param duration - Length of time to move between starting price and ending price (in seconds).
     */
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
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
            _ownsEnough(nftAddress, seller, tokenId, quantity),
            "JamClockAuction1155: sender does not have enough NFTs"
        );
        _escrow(nftAddress, seller, tokenId, quantity);
        uint256 auctionId = _currentAuctionId.current();
        Auction memory auction = Auction(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            seller,
            currency,
            uint128(startingPrice),
            uint128(endingPrice),
            uint64(duration),
            uint64(block.timestamp)
        );
        _addAuction(auction);
        _currentAuctionId.increment();
    }

    /**
     * @dev Bids on an open auction, completing the auction and transferring ownership of the NFT if enough Ether is supplied.
     * @param auctionId - The ID of the auction which a bidder wants to bid
     * @param bidAmount - The amount of money a bidder is willing to bid
     */
    function bid(uint256 auctionId, uint256 bidAmount)
        external
        payable
        whenNotPaused
        nonReentrant
    {
        Auction memory auction = _auctions[auctionId];
        if (auction.currency == address(0))
            require(
                bidAmount == msg.value,
                "JamClockAuction1155: bid amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamClockAuction1155: return money failed");
        }
        _bid(auctionId, bidAmount);
        _transfer(
            auction.nftAddress,
            msg.sender,
            auction.tokenId,
            auction.quantity
        );
    }

    /**
     * @dev Updates an existent auction's information.
     * @param auctionId - The ID of the auction to update
     * @param currency - The address of ERC20 or native token used as payment.
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endingPrice - Price of item (in wei) at end of auction.
     * @param duration - Length of time to move between starting price and ending price (in seconds).
     */
    function updateAuction(
        uint256 auctionId,
        address currency,
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
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        require(
            msg.sender == auction.seller,
            "JamClockAuction1155: only seller can update auction"
        );
        auction.currency = currency;
        auction.startingPrice = uint128(startingPrice);
        auction.endingPrice = uint128(endingPrice);
        auction.duration = uint64(duration);
        emit AuctionUpdated(
            auctionId,
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            currency,
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
     * @param auctionId - The ID of the auction to cancel
     */
    function cancelAuction1155(uint256 auctionId) external override {
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        require(
            msg.sender == auction.seller ||
                msg.sender ==
                JamMarketplaceHub(_marketplaceHub).getMarketplace(
                    keccak256("JAM_P2P_TRADING_1155")
                ),
            "JamClockAuction1155: only seller can cancel auction"
        );
        _cancelAuction(auctionId, msg.sender);
    }

    /**
     * @dev Cancels an auction when the contract is paused.
     * Only the owner may do this, and NFTs are returned to
     * the seller. This should only be used in emergencies.
     * @param auctionId - The ID of the auction to cancel
     */
    function cancelAuctionWhenPaused(uint256 auctionId)
        external
        whenPaused
        onlyOwner
    {
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        _cancelAuction(auctionId, auction.seller);
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
        returns (IERC1155)
    {
        IERC1155 candidateContract = IERC1155(nftAddress);
        return candidateContract;
    }

    /**
     * @dev Returns current price of an NFT on auction. Broken into two
     *  functions (this one, that computes the duration from the auction
     *  structure, and the other that does the price computation) so we
     *  can easily test that the price computation works correctly.
     */
    function _getCurrentTotalPrice(Auction storage auction)
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
    function _ownsEnough(
        address nftAddress,
        address claimant,
        uint256 tokenId,
        uint256 quantity
    ) internal view returns (bool) {
        IERC1155 nftContract = _getNftContract(nftAddress);
        return (nftContract.balanceOf(claimant, tokenId) >= quantity);
    }

    /**
     * @dev Adds an auction to the list of open auctions. Also fires the
     * AuctionCreated event.
     * @param auction Auction to add.
     */
    function _addAuction(Auction memory auction) internal {
        require(
            auction.duration >= 1 minutes,
            "JamClockAuction1155: too short auction"
        );
        _auctions[auction.auctionId] = auction;
        emit AuctionCreated(
            auction.auctionId,
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.currency,
            uint256(auction.startingPrice),
            uint256(auction.endingPrice),
            uint256(auction.duration),
            auction.seller
        );
    }

    /**
     * @dev Removes an auction from the list of open auctions.
     * @param auctionId - The ID of the auction to remove
     */
    function _removeAuction(uint256 auctionId) internal {
        delete _auctions[auctionId];
    }

    /** @dev Cancels an auction unconditionally. */
    function _cancelAuction(uint256 auctionId, address recipient) internal {
        Auction storage auction = _auctions[auctionId];
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        uint256 quantity = auction.quantity;
        _removeAuction(auctionId);
        _transfer(nftAddress, recipient, tokenId, quantity);
        emit AuctionCancelled(auctionId, nftAddress, tokenId, quantity);
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
        uint256 tokenId,
        uint256 quantity
    ) internal {
        IERC1155 nftContract = _getNftContract(nftAddress);
        nftContract.safeTransferFrom(
            owner,
            address(this),
            tokenId,
            quantity,
            abi.encodePacked("Transfer")
        );
    }

    /**
     * @dev Transfers an NFT owned by this contract to another address.
     * Returns true if the transfer succeeds.
     * @param nftAddress - The address of the NFT.
     * @param receiver - Address to transfer NFT to.
     * @param tokenId - ID of token to transfer.
     * @param quantity - The number of NFTs to transfer
     */
    function _transfer(
        address nftAddress,
        address receiver,
        uint256 tokenId,
        uint256 quantity
    ) internal {
        IERC1155 nftContract = _getNftContract(nftAddress);
        nftContract.safeTransferFrom(
            address(this),
            receiver,
            tokenId,
            quantity,
            abi.encodePacked("Transfer")
        );
    }

    /**
     * @dev Computes the price and transfers winnings.
     * Does NOT transfer ownership of token.
     */
    function _bid(uint256 auctionId, uint256 bidAmount)
        internal
        returns (uint256)
    {
        Auction storage auction = _auctions[auctionId];
        require(
            _isOnAuction(auction),
            "JamClockAuction1155: auction not exists"
        );
        uint256 price = _getCurrentTotalPrice(auction);
        require(
            bidAmount >= price,
            "JamClockAuction1155: insufficient bid amount"
        );
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        uint256 quantity = auction.quantity;
        address seller = auction.seller;
        address currency = auction.currency;
        _removeAuction(auctionId);
        if (currency != address(0))
            IERC20(currency).transferFrom(msg.sender, address(this), bidAmount);
        if (price > 0)
            _computeFeesAndPaySeller(nftAddress, seller, currency, price);
        if (bidAmount > price) {
            uint256 bidExcess = bidAmount - price;
            if (currency == address(0)) {
                (bool success, ) = payable(msg.sender).call{value: bidExcess}(
                    ""
                );
                require(success, "JamClockAuction1155: return excess failed");
            } else {
                bool success = IERC20(currency).transfer(msg.sender, bidExcess);
                require(success, "JamClockAuction1155: return excess failed");
            }
        }
        emit AuctionSuccessful(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            currency,
            price,
            msg.sender
        );
        return price;
    }
}

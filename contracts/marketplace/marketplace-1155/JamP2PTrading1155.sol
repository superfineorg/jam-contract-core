/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "../JamMarketplaceHelpers.sol";

contract JamP2PTrading1155 is JamMarketplaceHelpers {
    struct Offer {
        address offeror;
        address nftAddress;
        uint256 tokenId;
        uint256 quantity;
        address currency;
        uint256 amount;
    }

    // Mapping from an NFT type (NFT address + token ID) to its pending offers
    mapping(address => mapping(uint256 => Offer[])) private _offersFor;

    // Mapping from a user address to the offers he has made so far
    mapping(address => Offer[]) private _offersOf;

    event OfferCreated(
        address offeror,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 amount
    );

    event OfferUpdated(
        address offeror,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 amount
    );

    event OfferAccepted(
        address offeror,
        address accepter,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 amount
    );

    event OfferCancelled(
        address offeror,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_P2P_TRADING_1155");
    }

    function getAllOffersFor(address nftAddress, uint256 tokenId)
        external
        view
        returns (Offer[] memory)
    {
        return _offersFor[nftAddress][tokenId];
    }

    function getAcceptableOffersFor(
        address nftAddress,
        uint256 tokenId,
        address user
    ) external view returns (Offer[] memory) {
        uint256 offerCount = 0;
        uint256 numAcceptableOffers = 0;
        uint256 balance = IERC1155(nftAddress).balanceOf(user, tokenId);
        Offer[] memory allOffers = _offersFor[nftAddress][tokenId];
        for (uint256 i = 0; i < allOffers.length; i++)
            if (allOffers[i].quantity <= balance) numAcceptableOffers++;
        Offer[] memory acceptableOffers = new Offer[](numAcceptableOffers);
        for (uint256 i = 0; i < allOffers.length; i++)
            if (allOffers[i].quantity <= balance) {
                acceptableOffers[offerCount] = allOffers[i];
                offerCount++;
            }
        return acceptableOffers;
    }

    function getOffersOf(address offeror)
        external
        view
        returns (Offer[] memory)
    {
        return _offersOf[offeror];
    }

    function getSpecificOffer(
        address offeror,
        address nftAddress,
        uint256 tokenId
    ) public view returns (Offer memory) {
        Offer memory offer;
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++) {
            offer = _offersFor[nftAddress][tokenId][i];
            if (offer.offeror == offeror) return offer;
        }
        return offer;
    }

    function makeOffer(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 offerAmount
    ) external payable {
        // Check offering conditions
        require(
            getSpecificOffer(msg.sender, nftAddress, tokenId).offeror ==
                address(0),
            "JamP2PTrading1155: already offered before"
        );
        if (currency == address(0))
            require(
                offerAmount == msg.value,
                "JamP2PTrading1155: offer amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamP2PTrading1155: return money failed");
        }

        // Save the offer's information
        Offer memory offer = Offer(
            msg.sender,
            nftAddress,
            tokenId,
            quantity,
            currency,
            offerAmount
        );
        _offersFor[nftAddress][tokenId].push(offer);
        _offersOf[msg.sender].push(offer);

        // Lock offeror's money
        if (currency != address(0))
            IERC20(currency).transferFrom(
                msg.sender,
                address(this),
                offerAmount
            );

        emit OfferCreated(
            msg.sender,
            nftAddress,
            tokenId,
            quantity,
            currency,
            offerAmount
        );
    }

    function updateOffer(
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 offerAmount
    ) external payable nonReentrant {
        Offer memory offer = getSpecificOffer(msg.sender, nftAddress, tokenId);

        // Check updating conditions
        require(
            offer.offeror == msg.sender,
            "JamP2PTrading1155: no offer found"
        );
        if (currency == address(0))
            require(
                offerAmount == msg.value,
                "JamP2PTrading1155: offer amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamP2PTrading1155: return money failed");
        }

        // Save the new information
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++)
            if (_offersFor[nftAddress][tokenId][i].offeror == msg.sender) {
                _offersFor[nftAddress][tokenId][i].currency = currency;
                _offersFor[nftAddress][tokenId][i].amount = offerAmount;
                break;
            }
        for (uint256 i = 0; i < _offersOf[msg.sender].length; i++)
            if (
                _offersOf[msg.sender][i].nftAddress == nftAddress &&
                _offersOf[msg.sender][i].tokenId == tokenId
            ) {
                _offersOf[msg.sender][i].currency = currency;
                _offersOf[msg.sender][i].amount = offerAmount;
                break;
            }

        // Return old locked offer amount
        if (offer.currency == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: offer.amount}(
                ""
            );
            require(success, "JamP2PTrading1155: return money failed");
        } else {
            bool success = IERC20(offer.currency).transfer(
                msg.sender,
                offer.amount
            );
            require(success, "JamP2PTrading1155: return money failed");
        }

        // Lock new offer amount
        if (currency != address(0))
            IERC20(currency).transferFrom(
                msg.sender,
                address(this),
                offerAmount
            );

        emit OfferUpdated(
            msg.sender,
            nftAddress,
            tokenId,
            offer.quantity,
            currency,
            offerAmount
        );
    }

    function cancelOffer(address nftAddress, uint256 tokenId)
        external
        nonReentrant
    {
        Offer memory offer = getSpecificOffer(msg.sender, nftAddress, tokenId);

        // Check cancelling conditions
        require(
            offer.offeror == msg.sender,
            "JamP2PTrading1155: sender is not offeror"
        );

        // Delete offer's information
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++)
            if (_offersFor[nftAddress][tokenId][i].offeror == msg.sender) {
                uint256 length = _offersFor[nftAddress][tokenId].length;
                _offersFor[nftAddress][tokenId][i] = _offersFor[nftAddress][
                    tokenId
                ][length - 1];
                _offersFor[nftAddress][tokenId].pop();
                break;
            }
        for (uint256 i = 0; i < _offersOf[msg.sender].length; i++)
            if (
                _offersOf[msg.sender][i].nftAddress == nftAddress &&
                _offersOf[msg.sender][i].tokenId == tokenId
            ) {
                uint256 length = _offersOf[msg.sender].length;
                _offersOf[msg.sender][i] = _offersOf[msg.sender][length - 1];
                _offersOf[msg.sender].pop();
                break;
            }

        // Return locked offer amount
        if (offer.currency == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: offer.amount}(
                ""
            );
            require(success, "JamP2PTrading1155: return money failed");
        } else {
            bool success = IERC20(offer.currency).transfer(
                msg.sender,
                offer.amount
            );
            require(success, "JamP2PTrading1155: return money failed");
        }

        emit OfferCancelled(msg.sender, nftAddress, tokenId, offer.quantity);
    }

    /**
     * @dev The owner of the NFT accepts the offer even when the NFT is on marketplace.
     * @param offeror The person who offered this NFT so far.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the offered token.
     * @notice If the NFT is currently not on marketplace, the owner must approve it for this contract
     */
    function acceptOffer(
        address offeror,
        address nftAddress,
        uint256 tokenId
    ) external nonReentrant {
        // Check accepting conditions
        Offer memory offer = getSpecificOffer(offeror, nftAddress, tokenId);
        require(offer.offeror == offeror, "JamP2PTrading1155: no offer found");

        require(
            IERC1155(nftAddress).balanceOf(msg.sender, tokenId) >=
                offer.quantity,
            "JamP2PTrading1155: not enough NFTs to accept"
        );

        // Sell it to the offeror
        IERC1155(nftAddress).safeTransferFrom(
            msg.sender,
            offeror,
            tokenId,
            offer.quantity,
            abi.encodePacked("Accept offer")
        );

        // Compute auctioneer cut and royalty cut then return proceeds to seller
        if (offer.amount > 0)
            _computeFeesAndPaySeller(
                nftAddress,
                msg.sender,
                offer.currency,
                offer.amount
            );

        emit OfferAccepted(
            offeror,
            msg.sender,
            nftAddress,
            tokenId,
            offer.quantity,
            offer.currency,
            offer.amount
        );
    }
}

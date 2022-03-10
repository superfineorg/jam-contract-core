// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract NFTStaking is
    ERC721Holder,
    ERC1155Holder,
    Ownable,
    Pausable,
    ReentrancyGuard
{
    enum NFTType {
        ERC721,
        ERC1155
    }

    struct StakingInfo {
        uint256 numStakedNFTs;
        uint256 lastClaimMoment;
        // Mapping from an NFT address to a list of staked tokenIds
        mapping(address => uint256[]) stakedTokenIds;
        // Mapping from an NFT (NFT address + tokenID) to its staked quantity
        mapping(address => mapping(uint256 => uint256)) stakedQuantityOf;
        // Mapping from an NFT (NFT address + tokenID) to a moment when that NFT is staked
        mapping(address => mapping(uint256 => uint256)) stakingMomentOf;
    }

    struct NFTInfo {
        NFTType nftType;
        address nftAddress;
        uint256 tokenId;
        uint256 quantity;
        uint256 stakingMoment;
    }

    uint256 public lockDuration;
    uint256 public rewardPerDay;
    address public rewardToken;
    uint256 private _totalStakedNFTs;
    mapping(address => NFTType) private _typeOf;
    mapping(address => StakingInfo) private _stakingInfoOf;
    mapping(address => mapping(uint256 => address)) private _ownerOf;
    mapping(address => bool) private _operators;

    // Whitelist
    address[] private _erc721Whitelist;
    address[] private _erc1155Whitelist;
    mapping(address => uint256[]) private _erc1155TokenIdWhitelist;
    mapping(address => bool) private _isERC721Whitelisted;
    mapping(address => mapping(uint256 => bool))
        private _isERC1155TokenIdWhitelisted;

    event NFTStaked(
        address participant,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );
    event NFTUnstaked(
        address participant,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );
    event RewardClaimed(address participant, uint256 rewardAmount);
    event EmergencyWithdrawn(address recipient, uint256 amount);

    receive() external payable {}

    constructor(
        uint256 lockDuration_,
        uint256 rewardPerDay_,
        address rewardToken_
    ) {
        lockDuration = lockDuration_;
        rewardPerDay = rewardPerDay_;
        rewardToken = rewardToken_;
        _operators[msg.sender] = true;
    }

    modifier onlyOperator() {
        require(_operators[msg.sender], "NFTStaking: caller is not operator");
        _;
    }

    function getCurrentReward(address participant)
        public
        view
        returns (uint256)
    {
        StakingInfo storage stakingInfo = _stakingInfoOf[participant];
        uint256 elapsedTime = block.timestamp - stakingInfo.lastClaimMoment;
        if (_totalStakedNFTs == 0) return 0;
        return
            (elapsedTime * rewardPerDay * stakingInfo.numStakedNFTs) /
            _totalStakedNFTs /
            1 days;
    }

    function getStakedNFTs(address participant)
        external
        view
        returns (NFTInfo[] memory)
    {
        uint256 numStakedNFTs = 0;

        // Get the number of staked ERC721 NFTs
        for (uint256 i = 0; i < _erc721Whitelist.length; i++)
            numStakedNFTs += _stakingInfoOf[participant]
                .stakedTokenIds[_erc721Whitelist[i]]
                .length;

        // Get the number of staked ERC1155 NFTs
        for (uint256 j = 0; j < _erc1155Whitelist.length; j++)
            numStakedNFTs += _stakingInfoOf[participant]
                .stakedTokenIds[_erc1155Whitelist[j]]
                .length;

        NFTInfo[] memory stakedNFTs = new NFTInfo[](numStakedNFTs);
        uint256 nftCount = 0;

        // Get staked ERC721 NFTs
        for (uint256 i = 0; i < _erc721Whitelist.length; i++) {
            address nftAddress = _erc721Whitelist[i];
            uint256[] memory stakedTokenIds = _stakingInfoOf[participant]
                .stakedTokenIds[nftAddress];
            for (uint256 j = 0; j < stakedTokenIds.length; j++) {
                uint256 tokenId = stakedTokenIds[j];
                uint256 stakingMoment = _stakingInfoOf[participant]
                    .stakingMomentOf[nftAddress][tokenId];
                stakedNFTs[nftCount] = NFTInfo(
                    NFTType.ERC721,
                    nftAddress,
                    tokenId,
                    1,
                    stakingMoment
                );
                nftCount++;
            }
        }

        // Get staked ERC1155 NFTs
        for (uint256 k = 0; k < _erc1155Whitelist.length; k++) {
            address nftAddress = _erc1155Whitelist[k];
            uint256[] memory stakedTokenIds = _stakingInfoOf[participant]
                .stakedTokenIds[nftAddress];
            for (uint256 l = 0; l < stakedTokenIds.length; l++) {
                uint256 tokenId = stakedTokenIds[l];
                uint256 quantity = _stakingInfoOf[participant].stakedQuantityOf[
                    nftAddress
                ][tokenId];
                uint256 stakingMoment = _stakingInfoOf[participant]
                    .stakingMomentOf[nftAddress][tokenId];
                stakedNFTs[nftCount] = NFTInfo(
                    NFTType.ERC1155,
                    nftAddress,
                    tokenId,
                    quantity,
                    stakingMoment
                );
                nftCount++;
            }
        }

        return stakedNFTs;
    }

    function getUnstakedNFTs(address participant)
        external
        view
        returns (NFTInfo[] memory)
    {
        uint256 numUnstakedNFTs = 0;

        // Get the number of owned ERC721 NFTs
        for (uint256 i = 0; i < _erc721Whitelist.length; i++)
            numUnstakedNFTs += ERC721Enumerable(_erc721Whitelist[i]).balanceOf(
                participant
            );

        // Get the number of owned ERC1155 NFTs
        for (uint256 j = 0; j < _erc1155Whitelist.length; j++) {
            address nftAddress = _erc1155Whitelist[j];
            uint256[] memory whitelistedTokenIds = _erc1155TokenIdWhitelist[
                nftAddress
            ];
            for (uint256 i = 0; i < whitelistedTokenIds.length; i++)
                if (
                    IERC1155(nftAddress).balanceOf(
                        participant,
                        whitelistedTokenIds[i]
                    ) > 0
                ) numUnstakedNFTs++;
        }

        NFTInfo[] memory unstakedNFTs = new NFTInfo[](numUnstakedNFTs);
        uint256 nftCount = 0;

        // Get unstaked ERC721 NFTs
        for (uint256 i = 0; i < _erc721Whitelist.length; i++) {
            address nftAddress = _erc721Whitelist[i];
            ERC721Enumerable nftContract = ERC721Enumerable(nftAddress);
            uint256 numOwnedNFTs = nftContract.balanceOf(participant);
            for (uint256 index = 0; index < numOwnedNFTs; index++) {
                uint256 tokenId = nftContract.tokenOfOwnerByIndex(
                    participant,
                    index
                );
                unstakedNFTs[nftCount] = NFTInfo(
                    NFTType.ERC721,
                    nftAddress,
                    tokenId,
                    1,
                    0
                );
                nftCount++;
            }
        }

        // Get unstaked ERC1155 NFTs
        for (uint256 i = 0; i < _erc1155Whitelist.length; i++) {
            address nftAddress = _erc1155Whitelist[i];
            uint256[] memory whitelistedTokenIds = _erc1155TokenIdWhitelist[
                nftAddress
            ];
            for (uint256 j = 0; j < whitelistedTokenIds.length; j++) {
                uint256 quantity = IERC1155(nftAddress).balanceOf(
                    participant,
                    whitelistedTokenIds[j]
                );
                if (quantity > 0) {
                    unstakedNFTs[nftCount] = NFTInfo(
                        NFTType.ERC1155,
                        nftAddress,
                        whitelistedTokenIds[j],
                        quantity,
                        0
                    );
                    nftCount++;
                }
            }
        }

        return unstakedNFTs;
    }

    function setOperators(address[] memory operators, bool[] memory isOperators)
        external
        onlyOwner
    {
        require(
            operators.length == isOperators.length,
            "NFTStaking: lengths mismatch"
        );
        for (uint256 i = 0; i < operators.length; i++)
            _operators[operators[i]] = isOperators[i];
    }

    function setLockDuration(uint256 lockDuration_) external onlyOperator {
        lockDuration = lockDuration_;
    }

    function setRewardPerDay(uint256 rewardPerDay_) external onlyOperator {
        rewardPerDay = rewardPerDay_;
    }

    function setRewardToken(address rewardToken_) external onlyOperator {
        rewardToken = rewardToken_;
    }

    function whitelistNFT(
        address[] calldata nftAddresses,
        NFTType[] calldata types,
        uint256[] calldata tokenIds
    ) external onlyOperator {
        require(
            nftAddresses.length == types.length,
            "NFTStaking: lengths mismatch"
        );
        require(
            nftAddresses.length == tokenIds.length,
            "NFTStaking: lengths mismatch"
        );

        for (uint256 i = 0; i < nftAddresses.length; i++) {
            address nftAddress = nftAddresses[i];
            _typeOf[nftAddress] = types[i];
            if (types[i] == NFTType.ERC721) {
                _isERC721Whitelisted[nftAddress] = true;
                bool addedBefore = false;
                for (uint256 j = 0; j < _erc721Whitelist.length; j++)
                    if (_erc721Whitelist[j] == nftAddress) {
                        addedBefore = true;
                        break;
                    }
                if (!addedBefore) _erc721Whitelist.push(nftAddress);
            } else if (types[i] == NFTType.ERC1155) {
                _isERC1155TokenIdWhitelisted[nftAddress][tokenIds[i]] = true;
                bool nftAddedBefore = false;
                for (uint256 k = 0; k < _erc1155Whitelist.length; k++)
                    if (_erc1155Whitelist[k] == nftAddress) {
                        nftAddedBefore = true;
                        break;
                    }
                if (!nftAddedBefore) _erc1155Whitelist.push(nftAddress);
                bool tokenIdAddedBefore = false;
                for (
                    uint256 l = 0;
                    l < _erc1155TokenIdWhitelist[nftAddress].length;
                    l++
                )
                    if (
                        _erc1155TokenIdWhitelist[nftAddress][l] == tokenIds[i]
                    ) {
                        tokenIdAddedBefore = true;
                        break;
                    }
                if (!tokenIdAddedBefore)
                    _erc1155TokenIdWhitelist[nftAddress].push(tokenIds[i]);
            }
        }
    }

    function stake(
        address[] memory nftAddresses,
        uint256[] memory tokenIds,
        uint256[] memory quantities
    ) external whenNotPaused nonReentrant {
        require(
            nftAddresses.length == tokenIds.length,
            "NFTStaking: lengths mismatch"
        );
        require(
            nftAddresses.length == quantities.length,
            "NFTStaking: lengths mismatch"
        );
        for (uint256 i = 0; i < nftAddresses.length; i++)
            _stake(nftAddresses[i], tokenIds[i], quantities[i]);
    }

    function unstake(
        address[] memory nftAddresses,
        uint256[] memory tokenIds,
        uint256[] memory quantities
    ) external whenNotPaused nonReentrant {
        require(
            nftAddresses.length == tokenIds.length,
            "NFTStaking: lengths mismatch"
        );
        require(
            nftAddresses.length == quantities.length,
            "NFTStaking: lengths mismatch"
        );
        for (uint256 i = 0; i < nftAddresses.length; i++)
            _unstake(nftAddresses[i], tokenIds[i], quantities[i]);
    }

    function _stake(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    ) private {
        require(
            _isERC721Whitelisted[nftAddress] ||
                _isERC1155TokenIdWhitelisted[nftAddress][tokenId],
            "NFTStaking: this NFT is not supported"
        );
        require(quantity > 0, "NFTStaking: stake nothing");
        _settle(msg.sender);
        StakingInfo storage stakingInfo = _stakingInfoOf[msg.sender];
        if (_typeOf[nftAddress] == NFTType.ERC721) {
            require(
                quantity == 1,
                "NFTStaking: cannot stake more than 1 ERC721 NFT at a time"
            );
            ERC721Enumerable(nftAddress).safeTransferFrom(
                msg.sender,
                address(this),
                tokenId
            );
            stakingInfo.stakedTokenIds[nftAddress].push(tokenId);
            stakingInfo.stakedQuantityOf[nftAddress][tokenId] = 1;
            stakingInfo.stakingMomentOf[nftAddress][tokenId] = block.timestamp;
        } else if (_typeOf[nftAddress] == NFTType.ERC1155) {
            IERC1155(nftAddress).safeTransferFrom(
                msg.sender,
                address(this),
                tokenId,
                quantity,
                ""
            );
            if (stakingInfo.stakedQuantityOf[nftAddress][tokenId] == 0)
                stakingInfo.stakedTokenIds[nftAddress].push(tokenId);
            stakingInfo.stakedQuantityOf[nftAddress][tokenId] += quantity;
            stakingInfo.stakingMomentOf[nftAddress][tokenId] = block.timestamp;
        }
        stakingInfo.lastClaimMoment = block.timestamp;
        stakingInfo.numStakedNFTs += quantity;
        _totalStakedNFTs += quantity;
        _ownerOf[nftAddress][tokenId] = msg.sender;
        emit NFTStaked(msg.sender, nftAddress, tokenId, quantity);
    }

    function _unstake(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    ) private {
        require(
            _isERC721Whitelisted[nftAddress] ||
                _isERC1155TokenIdWhitelisted[nftAddress][tokenId],
            "NFTStaking: this NFT is not supported"
        );
        require(
            _ownerOf[nftAddress][tokenId] == msg.sender,
            "NFTStaking: only owner can unstake"
        );
        require(quantity > 0, "NFTStaking: unstake nothing");
        _settle(msg.sender);
        StakingInfo storage stakingInfo = _stakingInfoOf[msg.sender];
        require(
            block.timestamp >=
                stakingInfo.stakingMomentOf[nftAddress][tokenId] + lockDuration,
            "NFTStaking: NFT not unlocked yet"
        );
        if (_typeOf[nftAddress] == NFTType.ERC721) {
            require(
                quantity == 1,
                "NFTStaking: cannot unstake more than 1 ERC721 NFT at a time"
            );
            require(
                stakingInfo.stakedQuantityOf[nftAddress][tokenId] == 1,
                "NFTStaking: NFT not found"
            );
            ERC721Enumerable(nftAddress).safeTransferFrom(
                address(this),
                msg.sender,
                tokenId
            );
            for (
                uint256 i = 0;
                i < stakingInfo.stakedTokenIds[nftAddress].length;
                i++
            )
                if (stakingInfo.stakedTokenIds[nftAddress][i] == tokenId) {
                    stakingInfo.stakedTokenIds[nftAddress][i] = stakingInfo
                        .stakedTokenIds[nftAddress][
                            stakingInfo.stakedTokenIds[nftAddress].length - 1
                        ];
                    stakingInfo.stakedTokenIds[nftAddress].pop();
                    break;
                }
            stakingInfo.stakedQuantityOf[nftAddress][tokenId] = 0;
            stakingInfo.stakingMomentOf[nftAddress][tokenId] = 0;
        } else if (_typeOf[nftAddress] == NFTType.ERC1155) {
            require(
                stakingInfo.stakedQuantityOf[nftAddress][tokenId] >= quantity,
                "NFTStaking: not enough NFTs to unstake"
            );
            IERC1155(nftAddress).safeTransferFrom(
                address(this),
                msg.sender,
                tokenId,
                quantity,
                ""
            );
            stakingInfo.stakedQuantityOf[nftAddress][tokenId] -= quantity;
            if (stakingInfo.stakedQuantityOf[nftAddress][tokenId] == 0) {
                stakingInfo.stakingMomentOf[nftAddress][tokenId] = 0;
                for (
                    uint256 i = 0;
                    i < stakingInfo.stakedTokenIds[nftAddress].length;
                    i++
                )
                    if (stakingInfo.stakedTokenIds[nftAddress][i] == tokenId) {
                        stakingInfo.stakedTokenIds[nftAddress][i] = stakingInfo
                            .stakedTokenIds[nftAddress][
                                stakingInfo.stakedTokenIds[nftAddress].length -
                                    1
                            ];
                        stakingInfo.stakedTokenIds[nftAddress].pop();
                        break;
                    }
            }
        }
        stakingInfo.lastClaimMoment = block.timestamp;
        stakingInfo.numStakedNFTs -= quantity;
        _totalStakedNFTs -= quantity;
        delete _ownerOf[nftAddress][tokenId];
        emit NFTUnstaked(msg.sender, nftAddress, tokenId, quantity);
    }

    function claimReward() external whenNotPaused {
        uint256 rewardAmount = getCurrentReward(msg.sender);
        _settle(msg.sender);
        emit RewardClaimed(msg.sender, rewardAmount);
    }

    function _settle(address participant) private {
        uint256 reward = getCurrentReward(participant);
        if (reward == 0) return;
        _stakingInfoOf[participant].lastClaimMoment = block.timestamp;
        if (rewardToken == address(0)) {
            (bool success, ) = payable(participant).call{value: reward}("");
            require(success, "NFTStaking: native token settle failed");
        } else {
            bool success = IERC20(rewardToken).transfer(participant, reward);
            require(success, "NFTStaking: ERC20 token settle failed");
        }
    }

    function pause() external onlyOperator {
        _pause();
    }

    function unpause() external onlyOperator {
        _unpause();
    }

    function emergencyWithdraw(address payable recipient)
        external
        onlyOwner
        whenPaused
    {
        uint256 amount = address(this).balance;
        (bool success, ) = recipient.call{value: amount}("");
        require(success, "NFTStaking: emergency withdraw failed");
        emit EmergencyWithdrawn(recipient, amount);
    }
}

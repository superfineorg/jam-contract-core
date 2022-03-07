// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

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
        // Mapping from an NFT to its staked quantity
        mapping(address => mapping(uint256 => uint256)) stakedQuantityOf;
        // Mapping from (NFT address + tokenID) to a moment when that NFT is staked
        mapping(address => mapping(uint256 => uint256)) stakingMomentOf;
    }

    uint256 public lockDuration;
    uint256 public rewardPerDay;
    address[] public nftWhitelist;
    uint256 private _numStakedNFTs;
    mapping(address => bool) private _isNFTWhitelisted;
    mapping(address => NFTType) private _typeOf;
    mapping(address => StakingInfo) private _stakingInfoOf;
    mapping(address => bool) private _operators;

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

    constructor(uint256 lockDuration_, uint256 rewardPerDay_) {
        lockDuration = lockDuration_;
        rewardPerDay = rewardPerDay_;
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
        if (_numStakedNFTs == 0) return 0;
        return
            (elapsedTime * rewardPerDay * stakingInfo.numStakedNFTs) /
            _numStakedNFTs /
            1 days;
    }

    function getNumStakedNFTs(address participant)
        external
        view
        returns (uint256)
    {
        return _stakingInfoOf[participant].numStakedNFTs;
    }

    function getStakedNFTTokenIds(address participant, address nftAddress)
        external
        view
        returns (uint256[] memory)
    {
        return _stakingInfoOf[participant].stakedTokenIds[nftAddress];
    }

    function getStakedQuantity(
        address participant,
        address nftAddress,
        uint256 tokenId
    ) external view returns (uint256) {
        return
            _stakingInfoOf[participant].stakedQuantityOf[nftAddress][tokenId];
    }

    function getStakingMoment(
        address participant,
        address nftAddress,
        uint256 tokenId
    ) external view returns (uint256) {
        return _stakingInfoOf[participant].stakingMomentOf[nftAddress][tokenId];
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

    function whitelistNFT(
        address[] calldata nftAddresses,
        NFTType[] calldata types,
        bool[] calldata statuses
    ) external onlyOperator {
        require(
            nftAddresses.length == types.length,
            "NFTStaking: lengths mismatch"
        );
        require(
            nftAddresses.length == statuses.length,
            "NFTStaking: lengths mismatch"
        );
        for (uint256 i = 0; i < nftAddresses.length; i++) {
            _typeOf[nftAddresses[i]] = types[i];
            _isNFTWhitelisted[nftAddresses[i]] = statuses[i];
            if (statuses[i]) nftWhitelist.push(nftAddresses[i]);
            else {
                for (uint256 j = 0; j < nftWhitelist.length; j++)
                    if (nftWhitelist[j] == nftAddresses[i]) {
                        nftWhitelist[j] = nftWhitelist[nftWhitelist.length - 1];
                        nftWhitelist.pop();
                        break;
                    }
            }
        }
    }

    function stake(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    ) external whenNotPaused nonReentrant {
        require(
            _isNFTWhitelisted[nftAddress],
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
            IERC721(nftAddress).safeTransferFrom(
                msg.sender,
                address(this),
                tokenId
            );
            stakingInfo.stakedTokenIds[nftAddress].push(tokenId);
            stakingInfo.stakedQuantityOf[nftAddress][tokenId] = 1;
            stakingInfo.stakingMomentOf[nftAddress][tokenId] = block.timestamp;
        } else {
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
        _numStakedNFTs += quantity;
        emit NFTStaked(msg.sender, nftAddress, tokenId, quantity);
    }

    function unstake(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    ) external whenNotPaused nonReentrant {
        require(
            _isNFTWhitelisted[nftAddress],
            "NFTStaking: this NFT is not supported"
        );
        require(quantity > 0, "NFTStaking: unstake nothing");
        _settle(msg.sender);
        StakingInfo storage stakingInfo = _stakingInfoOf[msg.sender];
        if (_typeOf[nftAddress] == NFTType.ERC721) {
            require(
                quantity == 1,
                "NFTStaking: cannot unstake more than 1 ERC721 NFT at a time"
            );
            require(
                stakingInfo.stakedQuantityOf[nftAddress][tokenId] == 1,
                "NFTStaking: NFT not found"
            );
            require(
                block.timestamp >=
                    stakingInfo.stakingMomentOf[nftAddress][tokenId] +
                        lockDuration,
                "NFTStaking: NFT not unlocked yet"
            );
            IERC721(nftAddress).safeTransferFrom(
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
        } else {
            require(
                stakingInfo.stakedQuantityOf[nftAddress][tokenId] >= quantity,
                "NFTStaking: not enough NFTs to unstake"
            );
            require(
                block.timestamp >=
                    stakingInfo.stakingMomentOf[nftAddress][tokenId] +
                        lockDuration,
                "NFTStaking: NFT not unlocked yet"
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
        _numStakedNFTs -= quantity;
        emit NFTUnstaked(msg.sender, nftAddress, tokenId, quantity);
    }

    function claimReward() external whenNotPaused {
        uint256 rewardAmount = getCurrentReward(msg.sender);
        _settle(msg.sender);
        emit RewardClaimed(msg.sender, rewardAmount);
    }

    function _settle(address participant) private {
        uint256 rewardAmount = getCurrentReward(participant);
        if (rewardAmount == 0) return;
        _stakingInfoOf[participant].lastClaimMoment = block.timestamp;
        (bool success, ) = payable(participant).call{value: rewardAmount}("");
        require(success, "NFTStaking: settle failed");
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

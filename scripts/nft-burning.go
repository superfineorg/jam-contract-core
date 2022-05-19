package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gamejam-contract-core/go/contracts"
)

const (
	Erc721ApprovalEvent       string = "Approval(address,address,uint256)"
	Erc1155ApprovalEvent      string = "ApprovalForAll(address,address,bool)"
	Erc721ContractAddress     string = "81a6AA39BB15a97d2f0231872F7f939E5d986070"
	Erc1155ContractAddress    string = "D5aB6238E3dA51138A4727736A3Fb23fc3eAaE04"
	NFTBurningContractAddress string = "7A25b44999E9181BFAaF8D2c38f9E58002474ab9"
	RPCProvider               string = "https://crypto.gamejam.com/testnet/rpc/"
	DeploymentBlock           int    = 2641604
	PrivateKey                string = "289972fdba8eafe076b5751454a00b43505cdec1b3e97618bba5636cd44343ab"
)

func buildQuery(contract common.Address, eventSelector common.Hash, startBlock *big.Int, endBlock *big.Int) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contract},
		Topics:    [][]common.Hash{{eventSelector}},
	}
	return query
}

func main() {
	// burnErc721()
	burnErc1155(
		common.HexToAddress("Ed6922a7065Ad1E7aAA34baA80828796cA011C3d"),
		[]*big.Int{big.NewInt(123)},
		[]*big.Int{big.NewInt(10)},
	)
}

func burnErc721() {
	c, _ := ethclient.Dial(RPCProvider)
	currentBlock := DeploymentBlock

	// Initialize general information
	caller, err := contracts.NewGameNFT721Caller(common.HexToAddress(Erc721ContractAddress), c)
	if err != nil {
		fmt.Println("Failed to initialize the NFT721 contract")
		return
	}
	burningContract, err := contracts.NewGameNFTBurning(common.HexToAddress(NFTBurningContractAddress), c)
	if err != nil {
		fmt.Println("Failed to initialize the burning contract")
		return
	}
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		fmt.Println("Failed to initialize the private key")
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(2710))
	if err != nil {
		fmt.Println("Failed to authenticate the private key")
		return
	}

	for {
		latestBlock, err := c.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("Failed to get latest block")
		} else if currentBlock <= int(latestBlock) {
			fmt.Printf("Scanning from block %d to block %d...\n", currentBlock, latestBlock)
			var logs []types.Log
			var err error
			for ok := true; ok; ok = (err != nil) {
				logs, err = c.FilterLogs(context.Background(), buildQuery(
					common.HexToAddress(Erc721ContractAddress),
					common.HexToHash(hex.EncodeToString(crypto.Keccak256([]byte(Erc721ApprovalEvent)))),
					big.NewInt(int64(currentBlock)),
					big.NewInt(int64(latestBlock)),
				))
			}
			tokenIdsToBurn := []*big.Int{}
			for i := 0; i < len(logs); i++ {
				// Get information from the event
				approved := common.BytesToAddress(logs[i].Topics[2].Bytes()).String()
				tokenId, success := new(big.Int).SetString(hex.EncodeToString(logs[i].Topics[3].Bytes()), 16)
				if !success {
					fmt.Println("Error when parsing approved tokenId")
					break
				}
				fmt.Printf("Token #%d has been approved to %s\n", tokenId.Int64(), approved)

				// List all token IDs to prepare to burn
				if strings.TrimPrefix(approved, "0x") == NFTBurningContractAddress {
					_, err = caller.OwnerOf(&bind.CallOpts{}, tokenId)
					if err != nil {
						fmt.Println("NFT does not exist or burned already")
					} else {
						tokenIdsToBurn = append(tokenIdsToBurn, tokenId)
					}
				}
			}

			// Burn all listed token IDs
			if len(tokenIdsToBurn) > 0 {
				fmt.Printf("Burning the token #%v...\n", tokenIdsToBurn)
				nftAddresses := []common.Address{}
				for i := 0; i < len(tokenIdsToBurn); i++ {
					nftAddresses = append(nftAddresses, common.HexToAddress(Erc721ContractAddress))
				}
				_, err = burningContract.BurnErc721IntoGames(auth, nftAddresses, tokenIdsToBurn)
				if err != nil {
					fmt.Printf("Failed to burn the tokens #%v\n", tokenIdsToBurn)
				} else {
					fmt.Printf("Tokens #%s has been burned successfully!\n", tokenIdsToBurn)
				}
			}
			currentBlock = int(latestBlock) + 1
		}
		time.Sleep(10 * time.Second)
	}
}

func burnErc1155(owner common.Address, tokenIds []*big.Int, quantities []*big.Int) {
	c, _ := ethclient.Dial(RPCProvider)
	burningContract, err := contracts.NewGameNFTBurning(common.HexToAddress(NFTBurningContractAddress), c)
	if err != nil {
		fmt.Println("Failed to initialize the burning contract")
		return
	}
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		fmt.Println("Failed to initialize the private key")
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(2710))
	if err != nil {
		fmt.Println("Failed to authenticate the private key")
		return
	}

	if len(tokenIds) != len(quantities) {
		fmt.Println("Lenghts mismatch")
		return
	}

	// Start burning
	fmt.Println("Burning:")
	for i := 0; i < len(tokenIds); i++ {
		fmt.Printf("\t- %d NFT #%d of 0x%s\n", quantities[i].Int64(), tokenIds[i].Int64(), hex.EncodeToString(owner.Bytes()))
	}
	_, err = burningContract.BurnErc1155IntoGames(
		auth,
		[]common.Address{owner},
		[]common.Address{common.HexToAddress(Erc1155ContractAddress)},
		[][]*big.Int{tokenIds},
		[][]*big.Int{quantities},
	)
	if err != nil {
		fmt.Println("Failed to burn ERC1155 NFTs", err)
	} else {
		fmt.Println("Burn successfully!")
	}
}

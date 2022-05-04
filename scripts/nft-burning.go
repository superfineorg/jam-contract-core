package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gamejam-contract-core/go/contracts"
)

const (
	// Keccak256("Approval(address,address,uint256)")
	ApprovalEventSelector     string = "8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	NFTContractAddress        string = "Ab18DA4F945E61BECF4b55aA8ec960493EbB4D94"
	NFTBurningContractAddress string = "8546fAA4F858cAF34Fd605865B6Bbf3A813D960b"
	RinkebyProvider           string = "https://crypto.gamejam.com/testnet/rpc/"
	DeploymentBlock           int    = 2346062
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
	c, _ := ethclient.Dial(RinkebyProvider)
	currentBlock := DeploymentBlock

	// Initialize general information
	caller, err := contracts.NewGameNFTCaller(common.HexToAddress(NFTContractAddress), c)
	if err != nil {
		fmt.Println("Failed to initialize the NFT contract")
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
			logs, err := c.FilterLogs(context.Background(), buildQuery(
				common.HexToAddress(NFTContractAddress),
				common.HexToHash(ApprovalEventSelector),
				big.NewInt(int64(currentBlock)),
				big.NewInt(int64(latestBlock)),
			))
			if err != nil {
				fmt.Print("Unable to filter logs", err)
			} else {
				for i := 0; i < len(logs); i++ {
					// Get information from the event
					approved := common.BytesToAddress(logs[i].Topics[2].Bytes()).String()
					tokenId, success := new(big.Int).SetString(hex.EncodeToString(logs[i].Topics[3].Bytes()), 16)
					if !success {
						fmt.Println("Error when parsing approved tokenId")
						break
					}
					fmt.Printf("Token #%d has been approved to %s\n", tokenId.Int64(), approved)

					// Burn the approved NFTs
					if approved[2:] == NFTBurningContractAddress {
						_, err = caller.OwnerOf(&bind.CallOpts{}, tokenId)
						if err != nil {
							fmt.Println("NFT does not exist or burned already")
						} else {
							fmt.Printf("Burning the token #%d...\n", tokenId)
							_, err = burningContract.BurnIntoGames(auth, common.HexToAddress(NFTContractAddress), tokenId)
							if err != nil {
								fmt.Printf("Failed to burn the token #%d\n", tokenId.Int64())
							} else {
								fmt.Printf("Token #%d has been burned successfully!\n", tokenId.Int64())
							}
						}
					}
				}
				currentBlock = int(latestBlock) + 1
			}
		}
		time.Sleep(10 * time.Second)
	}
}

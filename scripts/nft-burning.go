package main

import (
	"bytes"
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
	Erc721ContractAddress     string = "Bb0C52212D96706d17F7231B30EF34AA086f2275"
	Erc1155ContractAddress    string = "4Cf7962B32754C4D7B4976A615a7f57dAd419054"
	NFTBurningContractAddress string = "44077399AEaD781Db0728b99c744F3b905908Bde"
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
	c, _ := ethclient.Dial(RPCProvider)
	currentBlock := DeploymentBlock

	// Initialize general information
	caller721, err := contracts.NewGameNFT721Caller(common.HexToAddress(Erc721ContractAddress), c)
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
			scanErc721ApprovalEvents(
				c,
				int64(currentBlock),
				int64(latestBlock),
				caller721,
				auth,
				burningContract,
			)
			scanErc1155ApprovalEvents(
				c,
				int64(currentBlock),
				int64(latestBlock),
				auth,
				burningContract,
			)
			currentBlock = int(latestBlock) + 1
		}
		time.Sleep(10 * time.Second)
	}
}

func scanErc721ApprovalEvents(
	c *ethclient.Client,
	fromBlock int64,
	toBlock int64,
	caller721 *contracts.GameNFT721Caller,
	auth *bind.TransactOpts,
	burningContract *contracts.GameNFTBurning,
) {
	var logs []types.Log
	var err error
	for ok := true; ok; ok = (err != nil) {
		logs, err = c.FilterLogs(context.Background(), buildQuery(
			common.HexToAddress(Erc721ContractAddress),
			common.HexToHash(hex.EncodeToString(crypto.Keccak256([]byte(Erc721ApprovalEvent)))),
			big.NewInt(fromBlock),
			big.NewInt(toBlock),
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
			_, err = caller721.OwnerOf(&bind.CallOpts{}, tokenId)
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
		_, err = burningContract.BurnErc721IntoGames(auth, []common.Address{common.HexToAddress(Erc721ContractAddress)}, tokenIdsToBurn)
		if err != nil {
			fmt.Printf("Failed to burn the tokens #%v\n", tokenIdsToBurn)
		} else {
			fmt.Printf("Tokens #%s has been burned successfully!\n", tokenIdsToBurn)
		}
	}
}

func scanErc1155ApprovalEvents(
	c *ethclient.Client,
	fromBlock int64,
	toBlock int64,
	auth *bind.TransactOpts,
	burningContract *contracts.GameNFTBurning,
) {
	var logs []types.Log
	var err error
	for ok := true; ok; ok = (err != nil) {
		logs, err = c.FilterLogs(context.Background(), buildQuery(
			common.HexToAddress(Erc1155ContractAddress),
			common.HexToHash(hex.EncodeToString(crypto.Keccak256([]byte(Erc1155ApprovalEvent)))),
			big.NewInt(fromBlock),
			big.NewInt(toBlock),
		))
	}
	for i := 0; i < len(logs); i++ {
		owner := common.BytesToAddress(logs[i].Topics[1].Bytes()).String()
		operator := common.BytesToAddress(logs[i].Topics[2].Bytes()).String()
		if !bytes.Equal(logs[i].Data, make([]byte, 32)) {
			fmt.Printf("%s has approved all ERC1155 NFTs to %s\n", owner, operator)

			// Burn all approved NFTs
			if strings.TrimPrefix(operator, "0x") == NFTBurningContractAddress {
				fmt.Printf("Burning all ERC1155 tokens of %s...\n", owner)
				_, err = burningContract.BurnErc1155IntoGames(auth, []common.Address{common.HexToAddress(owner)}, []common.Address{common.HexToAddress(Erc1155ContractAddress)})
				if err != nil {
					fmt.Printf("Failed to burn ERC1155 NFTs of %s\n", owner)
				} else {
					fmt.Printf("All ERC1155 NFTs of %s are burned successfully!\n", owner)
				}
			}
		}
	}
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GameNFT721MetaData contains all meta data concerning the GameNFT721 contract.
var GameNFT721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002644380380620026448339810160408190526200003491620002a3565b8251839083906200004d90600090602085019062000146565b5080516200006390600190602084019062000146565b505081516200007b9150600b90602084019062000146565b506200008960003362000092565b50505062000387565b6200009e8282620000a2565b5050565b6000828152600a602090815260408083206001600160a01b038516845290915290205460ff166200009e576000828152600a602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620001023390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b828054620001549062000334565b90600052602060002090601f016020900481019282620001785760008555620001c3565b82601f106200019357805160ff1916838001178555620001c3565b82800160010185558215620001c3579182015b82811115620001c3578251825591602001919060010190620001a6565b50620001d1929150620001d5565b5090565b5b80821115620001d15760008155600101620001d6565b600082601f830112620001fe57600080fd5b81516001600160401b03808211156200021b576200021b62000371565b604051601f8301601f19908116603f0116810190828211818310171562000246576200024662000371565b816040528381526020925086838588010111156200026357600080fd5b600091505b8382101562000287578582018301518183018401529082019062000268565b83821115620002995760008385830101525b9695505050505050565b600080600060608486031215620002b957600080fd5b83516001600160401b0380821115620002d157600080fd5b620002df87838801620001ec565b94506020860151915080821115620002f657600080fd5b6200030487838801620001ec565b935060408601519150808211156200031b57600080fd5b506200032a86828701620001ec565b9150509250925092565b600181811c908216806200034957607f821691505b602082108114156200036b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6122ad80620003976000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80634f6ccce7116100f9578063a22cb46511610097578063d539139311610071578063d53913931461038d578063d547741f146103b4578063d547cfb7146103c7578063e985e9c5146103cf57600080fd5b8063a22cb46514610354578063b88d4fde14610367578063c87b56dd1461037a57600080fd5b806370a08231116100d357806370a082311461031e57806391d148541461033157806395d89b4114610344578063a217fddf1461034c57600080fd5b80634f6ccce7146102e55780636352211e146102f85780636a6278421461030b57600080fd5b8063248a9ca31161016657806330176e131161014057806330176e131461029957806336568abe146102ac57806342842e0e146102bf57806342966c68146102d257600080fd5b8063248a9ca3146102505780632f2ff15d146102735780632f745c591461028657600080fd5b806301ffc9a7146101ae57806306fdde03146101d6578063081812fc146101eb578063095ea7b31461021657806318160ddd1461022b57806323b872dd1461023d575b600080fd5b6101c16101bc366004611df6565b61040b565b60405190151581526020015b60405180910390f35b6101de61042b565b6040516101cd919061202e565b6101fe6101f9366004611dba565b6104bd565b6040516001600160a01b0390911681526020016101cd565b610229610224366004611d90565b610557565b005b6008545b6040519081526020016101cd565b61022961024b366004611c9c565b61066d565b61022f61025e366004611dba565b6000908152600a602052604090206001015490565b610229610281366004611dd3565b61069f565b61022f610294366004611d90565b6106c5565b6102296102a7366004611e30565b61075b565b6102296102ba366004611dd3565b6107d9565b6102296102cd366004611c9c565b610853565b6102296102e0366004611dba565b61086e565b61022f6102f3366004611dba565b6108e8565b6101fe610306366004611dba565b61097b565b610229610319366004611c4e565b6109f2565b61022f61032c366004611c4e565b610a9a565b6101c161033f366004611dd3565b610b21565b6101de610b4c565b61022f600081565b610229610362366004611d54565b610b5b565b610229610375366004611cd8565b610b66565b6101de610388366004611dba565b610b9e565b61022f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102296103c2366004611dd3565b610c37565b6101de610c5d565b6101c16103dd366004611c69565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b600061041682610ceb565b80610425575061042582610d0c565b92915050565b60606000805461043a90612189565b80601f016020809104026020016040519081016040528092919081815260200182805461046690612189565b80156104b35780601f10610488576101008083540402835291602001916104b3565b820191906000526020600020905b81548152906001019060200180831161049657829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b031661053b5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b60006105628261097b565b9050806001600160a01b0316836001600160a01b031614156105d05760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610532565b336001600160a01b03821614806105ec57506105ec81336103dd565b61065e5760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610532565b6106688383610d31565b505050565b610678335b82610d9f565b6106945760405162461bcd60e51b815260040161053290612093565b610668838383610e96565b6000828152600a60205260409020600101546106bb813361103d565b61066883836110a1565b60006106d083610a9a565b82106107325760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610532565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b610766600033610b21565b6107c25760405162461bcd60e51b815260206004820152602760248201527f47616d654e46543732313a206d75737420686176652061646d696e20726f6c65604482015266081d1bc81cd95d60ca1b6064820152608401610532565b80516107d590600b906020840190611b23565b5050565b6001600160a01b03811633146108495760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610532565b6107d58282611127565b61066883838360405180602001604052806000815250610b66565b61087733610672565b6108dc5760405162461bcd60e51b815260206004820152603060248201527f4552433732314275726e61626c653a2063616c6c6572206973206e6f74206f7760448201526f1b995c881b9bdc88185c1c1c9bdd995960821b6064820152608401610532565b6108e58161118e565b50565b60006108f360085490565b82106109565760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610532565b6008828154811061096957610969612235565b90600052602060002001549050919050565b6000818152600260205260408120546001600160a01b0316806104255760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610532565b610a1c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633610b21565b610a7a5760405162461bcd60e51b815260206004820152602960248201527f47616d654e46543732313a206d7573742068617665206d696e74657220726f6c60448201526819481d1bc81b5a5b9d60ba1b6064820152608401610532565b610a8c81610a87600c5490565b611235565b6108e5600c80546001019055565b60006001600160a01b038216610b055760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610532565b506001600160a01b031660009081526003602052604090205490565b6000918252600a602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606001805461043a90612189565b6107d533838361124f565b610b703383610d9f565b610b8c5760405162461bcd60e51b815260040161053290612093565b610b988484848461131e565b50505050565b6000818152600260205260409020546060906001600160a01b0316610c055760405162461bcd60e51b815260206004820181905260248201527f47616d654e46543732313a20746f6b656e20646f6573206e6f742065786973746044820152606401610532565b600b610c1083611351565b604051602001610c21929190611ec1565b6040516020818303038152906040529050919050565b6000828152600a6020526040902060010154610c53813361103d565b6106688383611127565b600b8054610c6a90612189565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9690612189565b8015610ce35780601f10610cb857610100808354040283529160200191610ce3565b820191906000526020600020905b815481529060010190602001808311610cc657829003601f168201915b505050505081565b60006001600160e01b03198216637965db0b60e01b14806104255750610425825b60006001600160e01b0319821663780e9d6360e01b148061042557506104258261144f565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610d668261097b565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000818152600260205260408120546001600160a01b0316610e185760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610532565b6000610e238361097b565b9050806001600160a01b0316846001600160a01b03161480610e5e5750836001600160a01b0316610e53846104bd565b6001600160a01b0316145b80610e8e57506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b0316610ea98261097b565b6001600160a01b031614610f0d5760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610532565b6001600160a01b038216610f6f5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610532565b610f7a83838361149f565b610f85600082610d31565b6001600160a01b0383166000908152600360205260408120805460019290610fae90849061212f565b90915550506001600160a01b0382166000908152600360205260408120805460019290610fdc9084906120e4565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6110478282610b21565b6107d55761105f816001600160a01b031660146114aa565b61106a8360206114aa565b60405160200161107b929190611f7c565b60408051601f198184030181529082905262461bcd60e51b82526105329160040161202e565b6110ab8282610b21565b6107d5576000828152600a602090815260408083206001600160a01b03851684529091529020805460ff191660011790556110e33390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6111318282610b21565b156107d5576000828152600a602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60006111998261097b565b90506111a78160008461149f565b6111b2600083610d31565b6001600160a01b03811660009081526003602052604081208054600192906111db90849061212f565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6107d582826040518060200160405280600081525061164d565b816001600160a01b0316836001600160a01b031614156112b15760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610532565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b611329848484610e96565b61133584848484611680565b610b985760405162461bcd60e51b815260040161053290612041565b6060816113755750506040805180820190915260018152600360fc1b602082015290565b8160005b811561139f5780611389816121c4565b91506113989050600a836120fc565b9150611379565b60008167ffffffffffffffff8111156113ba576113ba61224b565b6040519080825280601f01601f1916602001820160405280156113e4576020820181803683370190505b5090505b8415610e8e576113f960018361212f565b9150611406600a866121df565b6114119060306120e4565b60f81b81838151811061142657611426612235565b60200101906001600160f81b031916908160001a905350611448600a866120fc565b94506113e8565b60006001600160e01b031982166380ac58cd60e01b148061148057506001600160e01b03198216635b5e139f60e01b145b8061042557506301ffc9a760e01b6001600160e01b0319831614610425565b61066883838361178d565b606060006114b9836002612110565b6114c49060026120e4565b67ffffffffffffffff8111156114dc576114dc61224b565b6040519080825280601f01601f191660200182016040528015611506576020820181803683370190505b509050600360fc1b8160008151811061152157611521612235565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061155057611550612235565b60200101906001600160f81b031916908160001a9053506000611574846002612110565b61157f9060016120e4565b90505b60018111156115f7576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106115b3576115b3612235565b1a60f81b8282815181106115c9576115c9612235565b60200101906001600160f81b031916908160001a90535060049490941c936115f081612172565b9050611582565b5083156116465760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610532565b9392505050565b6116578383611845565b6116646000848484611680565b6106685760405162461bcd60e51b815260040161053290612041565b60006001600160a01b0384163b1561178257604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906116c4903390899088908890600401611ff1565b602060405180830381600087803b1580156116de57600080fd5b505af192505050801561170e575060408051601f3d908101601f1916820190925261170b91810190611e13565b60015b611768573d80801561173c576040519150601f19603f3d011682016040523d82523d6000602084013e611741565b606091505b5080516117605760405162461bcd60e51b815260040161053290612041565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610e8e565b506001949350505050565b6001600160a01b0383166117e8576117e381600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b61180b565b816001600160a01b0316836001600160a01b03161461180b5761180b8382611993565b6001600160a01b0382166118225761066881611a30565b826001600160a01b0316826001600160a01b031614610668576106688282611adf565b6001600160a01b03821661189b5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610532565b6000818152600260205260409020546001600160a01b0316156119005760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610532565b61190c6000838361149f565b6001600160a01b03821660009081526003602052604081208054600192906119359084906120e4565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600060016119a084610a9a565b6119aa919061212f565b6000838152600760205260409020549091508082146119fd576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b600854600090611a429060019061212f565b60008381526009602052604081205460088054939450909284908110611a6a57611a6a612235565b906000526020600020015490508060088381548110611a8b57611a8b612235565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611ac357611ac361221f565b6001900381819060005260206000200160009055905550505050565b6000611aea83610a9a565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b828054611b2f90612189565b90600052602060002090601f016020900481019282611b515760008555611b97565b82601f10611b6a57805160ff1916838001178555611b97565b82800160010185558215611b97579182015b82811115611b97578251825591602001919060010190611b7c565b50611ba3929150611ba7565b5090565b5b80821115611ba35760008155600101611ba8565b600067ffffffffffffffff80841115611bd757611bd761224b565b604051601f8501601f19908116603f01168101908282118183101715611bff57611bff61224b565b81604052809350858152868686011115611c1857600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b0381168114611c4957600080fd5b919050565b600060208284031215611c6057600080fd5b61164682611c32565b60008060408385031215611c7c57600080fd5b611c8583611c32565b9150611c9360208401611c32565b90509250929050565b600080600060608486031215611cb157600080fd5b611cba84611c32565b9250611cc860208501611c32565b9150604084013590509250925092565b60008060008060808587031215611cee57600080fd5b611cf785611c32565b9350611d0560208601611c32565b925060408501359150606085013567ffffffffffffffff811115611d2857600080fd5b8501601f81018713611d3957600080fd5b611d4887823560208401611bbc565b91505092959194509250565b60008060408385031215611d6757600080fd5b611d7083611c32565b915060208301358015158114611d8557600080fd5b809150509250929050565b60008060408385031215611da357600080fd5b611dac83611c32565b946020939093013593505050565b600060208284031215611dcc57600080fd5b5035919050565b60008060408385031215611de657600080fd5b82359150611c9360208401611c32565b600060208284031215611e0857600080fd5b813561164681612261565b600060208284031215611e2557600080fd5b815161164681612261565b600060208284031215611e4257600080fd5b813567ffffffffffffffff811115611e5957600080fd5b8201601f81018413611e6a57600080fd5b610e8e84823560208401611bbc565b60008151808452611e91816020860160208601612146565b601f01601f19169290920160200192915050565b60008151611eb7818560208601612146565b9290920192915050565b600080845481600182811c915080831680611edd57607f831692505b6020808410821415611efd57634e487b7160e01b86526022600452602486fd5b818015611f115760018114611f2257611f4f565b60ff19861689528489019650611f4f565b60008b81526020902060005b86811015611f475781548b820152908501908301611f2e565b505084890196505b505050505050611f73611f628286611ea5565b64173539b7b760d91b815260050190565b95945050505050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611fb4816017850160208801612146565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611fe5816028840160208801612146565b01602801949350505050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061202490830184611e79565b9695505050505050565b6020815260006116466020830184611e79565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b600082198211156120f7576120f76121f3565b500190565b60008261210b5761210b612209565b500490565b600081600019048311821515161561212a5761212a6121f3565b500290565b600082821015612141576121416121f3565b500390565b60005b83811015612161578181015183820152602001612149565b83811115610b985750506000910152565b600081612181576121816121f3565b506000190190565b600181811c9082168061219d57607f821691505b602082108114156121be57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156121d8576121d86121f3565b5060010190565b6000826121ee576121ee612209565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b0319811681146108e557600080fdfea264697066735822122053af62b059a1cf14713a6f0717d6a7724943dd76432a2aef4101de509ea0d61e64736f6c63430008060033",
}

// GameNFT721ABI is the input ABI used to generate the binding from.
// Deprecated: Use GameNFT721MetaData.ABI instead.
var GameNFT721ABI = GameNFT721MetaData.ABI

// GameNFT721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameNFT721MetaData.Bin instead.
var GameNFT721Bin = GameNFT721MetaData.Bin

// DeployGameNFT721 deploys a new Ethereum contract, binding an instance of GameNFT721 to it.
func DeployGameNFT721(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, baseTokenURI_ string) (common.Address, *types.Transaction, *GameNFT721, error) {
	parsed, err := GameNFT721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameNFT721Bin), backend, name, symbol, baseTokenURI_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameNFT721{GameNFT721Caller: GameNFT721Caller{contract: contract}, GameNFT721Transactor: GameNFT721Transactor{contract: contract}, GameNFT721Filterer: GameNFT721Filterer{contract: contract}}, nil
}

// GameNFT721 is an auto generated Go binding around an Ethereum contract.
type GameNFT721 struct {
	GameNFT721Caller     // Read-only binding to the contract
	GameNFT721Transactor // Write-only binding to the contract
	GameNFT721Filterer   // Log filterer for contract events
}

// GameNFT721Caller is an auto generated read-only Go binding around an Ethereum contract.
type GameNFT721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFT721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GameNFT721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFT721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameNFT721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFT721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameNFT721Session struct {
	Contract     *GameNFT721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameNFT721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameNFT721CallerSession struct {
	Contract *GameNFT721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GameNFT721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameNFT721TransactorSession struct {
	Contract     *GameNFT721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GameNFT721Raw is an auto generated low-level Go binding around an Ethereum contract.
type GameNFT721Raw struct {
	Contract *GameNFT721 // Generic contract binding to access the raw methods on
}

// GameNFT721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameNFT721CallerRaw struct {
	Contract *GameNFT721Caller // Generic read-only contract binding to access the raw methods on
}

// GameNFT721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameNFT721TransactorRaw struct {
	Contract *GameNFT721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGameNFT721 creates a new instance of GameNFT721, bound to a specific deployed contract.
func NewGameNFT721(address common.Address, backend bind.ContractBackend) (*GameNFT721, error) {
	contract, err := bindGameNFT721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameNFT721{GameNFT721Caller: GameNFT721Caller{contract: contract}, GameNFT721Transactor: GameNFT721Transactor{contract: contract}, GameNFT721Filterer: GameNFT721Filterer{contract: contract}}, nil
}

// NewGameNFT721Caller creates a new read-only instance of GameNFT721, bound to a specific deployed contract.
func NewGameNFT721Caller(address common.Address, caller bind.ContractCaller) (*GameNFT721Caller, error) {
	contract, err := bindGameNFT721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFT721Caller{contract: contract}, nil
}

// NewGameNFT721Transactor creates a new write-only instance of GameNFT721, bound to a specific deployed contract.
func NewGameNFT721Transactor(address common.Address, transactor bind.ContractTransactor) (*GameNFT721Transactor, error) {
	contract, err := bindGameNFT721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFT721Transactor{contract: contract}, nil
}

// NewGameNFT721Filterer creates a new log filterer instance of GameNFT721, bound to a specific deployed contract.
func NewGameNFT721Filterer(address common.Address, filterer bind.ContractFilterer) (*GameNFT721Filterer, error) {
	contract, err := bindGameNFT721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameNFT721Filterer{contract: contract}, nil
}

// bindGameNFT721 binds a generic wrapper to an already deployed contract.
func bindGameNFT721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameNFT721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFT721 *GameNFT721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFT721.Contract.GameNFT721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFT721 *GameNFT721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFT721.Contract.GameNFT721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFT721 *GameNFT721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFT721.Contract.GameNFT721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFT721 *GameNFT721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFT721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFT721 *GameNFT721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFT721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFT721 *GameNFT721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFT721.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameNFT721.Contract.DEFAULTADMINROLE(&_GameNFT721.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameNFT721.Contract.DEFAULTADMINROLE(&_GameNFT721.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721Session) MINTERROLE() ([32]byte, error) {
	return _GameNFT721.Contract.MINTERROLE(&_GameNFT721.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT721 *GameNFT721CallerSession) MINTERROLE() ([32]byte, error) {
	return _GameNFT721.Contract.MINTERROLE(&_GameNFT721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT721 *GameNFT721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT721 *GameNFT721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNFT721.Contract.BalanceOf(&_GameNFT721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT721 *GameNFT721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNFT721.Contract.BalanceOf(&_GameNFT721.CallOpts, owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT721 *GameNFT721Caller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT721 *GameNFT721Session) BaseTokenURI() (string, error) {
	return _GameNFT721.Contract.BaseTokenURI(&_GameNFT721.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT721 *GameNFT721CallerSession) BaseTokenURI() (string, error) {
	return _GameNFT721.Contract.BaseTokenURI(&_GameNFT721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNFT721.Contract.GetApproved(&_GameNFT721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNFT721.Contract.GetApproved(&_GameNFT721.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT721 *GameNFT721Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT721 *GameNFT721Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameNFT721.Contract.GetRoleAdmin(&_GameNFT721.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT721 *GameNFT721CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameNFT721.Contract.GetRoleAdmin(&_GameNFT721.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT721 *GameNFT721Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT721 *GameNFT721Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameNFT721.Contract.HasRole(&_GameNFT721.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT721 *GameNFT721CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameNFT721.Contract.HasRole(&_GameNFT721.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT721 *GameNFT721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT721 *GameNFT721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNFT721.Contract.IsApprovedForAll(&_GameNFT721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT721 *GameNFT721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNFT721.Contract.IsApprovedForAll(&_GameNFT721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT721 *GameNFT721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT721 *GameNFT721Session) Name() (string, error) {
	return _GameNFT721.Contract.Name(&_GameNFT721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT721 *GameNFT721CallerSession) Name() (string, error) {
	return _GameNFT721.Contract.Name(&_GameNFT721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNFT721.Contract.OwnerOf(&_GameNFT721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT721 *GameNFT721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNFT721.Contract.OwnerOf(&_GameNFT721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT721 *GameNFT721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT721 *GameNFT721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNFT721.Contract.SupportsInterface(&_GameNFT721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT721 *GameNFT721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNFT721.Contract.SupportsInterface(&_GameNFT721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT721 *GameNFT721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT721 *GameNFT721Session) Symbol() (string, error) {
	return _GameNFT721.Contract.Symbol(&_GameNFT721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT721 *GameNFT721CallerSession) Symbol() (string, error) {
	return _GameNFT721.Contract.Symbol(&_GameNFT721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GameNFT721.Contract.TokenByIndex(&_GameNFT721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GameNFT721.Contract.TokenByIndex(&_GameNFT721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GameNFT721.Contract.TokenOfOwnerByIndex(&_GameNFT721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT721 *GameNFT721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GameNFT721.Contract.TokenOfOwnerByIndex(&_GameNFT721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT721 *GameNFT721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT721 *GameNFT721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNFT721.Contract.TokenURI(&_GameNFT721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT721 *GameNFT721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNFT721.Contract.TokenURI(&_GameNFT721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT721 *GameNFT721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT721 *GameNFT721Session) TotalSupply() (*big.Int, error) {
	return _GameNFT721.Contract.TotalSupply(&_GameNFT721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT721 *GameNFT721CallerSession) TotalSupply() (*big.Int, error) {
	return _GameNFT721.Contract.TotalSupply(&_GameNFT721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.Approve(&_GameNFT721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.Approve(&_GameNFT721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.Burn(&_GameNFT721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.Burn(&_GameNFT721.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.GrantRole(&_GameNFT721.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.GrantRole(&_GameNFT721.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT721 *GameNFT721Transactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT721 *GameNFT721Session) Mint(to common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.Mint(&_GameNFT721.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT721 *GameNFT721TransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.Mint(&_GameNFT721.TransactOpts, to)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.RenounceRole(&_GameNFT721.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.RenounceRole(&_GameNFT721.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.RevokeRole(&_GameNFT721.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT721 *GameNFT721TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT721.Contract.RevokeRole(&_GameNFT721.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.SafeTransferFrom(&_GameNFT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.SafeTransferFrom(&_GameNFT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT721 *GameNFT721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT721 *GameNFT721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT721.Contract.SafeTransferFrom0(&_GameNFT721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT721 *GameNFT721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT721.Contract.SafeTransferFrom0(&_GameNFT721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT721 *GameNFT721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT721 *GameNFT721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT721.Contract.SetApprovalForAll(&_GameNFT721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT721 *GameNFT721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT721.Contract.SetApprovalForAll(&_GameNFT721.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT721 *GameNFT721Transactor) SetBaseTokenURI(opts *bind.TransactOpts, baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "setBaseTokenURI", baseTokenURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT721 *GameNFT721Session) SetBaseTokenURI(baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT721.Contract.SetBaseTokenURI(&_GameNFT721.TransactOpts, baseTokenURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT721 *GameNFT721TransactorSession) SetBaseTokenURI(baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT721.Contract.SetBaseTokenURI(&_GameNFT721.TransactOpts, baseTokenURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.TransferFrom(&_GameNFT721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT721 *GameNFT721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT721.Contract.TransferFrom(&_GameNFT721.TransactOpts, from, to, tokenId)
}

// GameNFT721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GameNFT721 contract.
type GameNFT721ApprovalIterator struct {
	Event *GameNFT721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721Approval represents a Approval event raised by the GameNFT721 contract.
type GameNFT721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GameNFT721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721ApprovalIterator{contract: _GameNFT721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GameNFT721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721Approval)
				if err := _GameNFT721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) ParseApproval(log types.Log) (*GameNFT721Approval, error) {
	event := new(GameNFT721Approval)
	if err := _GameNFT721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFT721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GameNFT721 contract.
type GameNFT721ApprovalForAllIterator struct {
	Event *GameNFT721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721ApprovalForAll represents a ApprovalForAll event raised by the GameNFT721 contract.
type GameNFT721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNFT721 *GameNFT721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GameNFT721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721ApprovalForAllIterator{contract: _GameNFT721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNFT721 *GameNFT721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GameNFT721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721ApprovalForAll)
				if err := _GameNFT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNFT721 *GameNFT721Filterer) ParseApprovalForAll(log types.Log) (*GameNFT721ApprovalForAll, error) {
	event := new(GameNFT721ApprovalForAll)
	if err := _GameNFT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFT721RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GameNFT721 contract.
type GameNFT721RoleAdminChangedIterator struct {
	Event *GameNFT721RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721RoleAdminChanged represents a RoleAdminChanged event raised by the GameNFT721 contract.
type GameNFT721RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameNFT721 *GameNFT721Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GameNFT721RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721RoleAdminChangedIterator{contract: _GameNFT721.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameNFT721 *GameNFT721Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GameNFT721RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721RoleAdminChanged)
				if err := _GameNFT721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameNFT721 *GameNFT721Filterer) ParseRoleAdminChanged(log types.Log) (*GameNFT721RoleAdminChanged, error) {
	event := new(GameNFT721RoleAdminChanged)
	if err := _GameNFT721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFT721RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GameNFT721 contract.
type GameNFT721RoleGrantedIterator struct {
	Event *GameNFT721RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721RoleGranted represents a RoleGranted event raised by the GameNFT721 contract.
type GameNFT721RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameNFT721RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721RoleGrantedIterator{contract: _GameNFT721.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GameNFT721RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721RoleGranted)
				if err := _GameNFT721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) ParseRoleGranted(log types.Log) (*GameNFT721RoleGranted, error) {
	event := new(GameNFT721RoleGranted)
	if err := _GameNFT721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFT721RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GameNFT721 contract.
type GameNFT721RoleRevokedIterator struct {
	Event *GameNFT721RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721RoleRevoked represents a RoleRevoked event raised by the GameNFT721 contract.
type GameNFT721RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameNFT721RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721RoleRevokedIterator{contract: _GameNFT721.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GameNFT721RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721RoleRevoked)
				if err := _GameNFT721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT721 *GameNFT721Filterer) ParseRoleRevoked(log types.Log) (*GameNFT721RoleRevoked, error) {
	event := new(GameNFT721RoleRevoked)
	if err := _GameNFT721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFT721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GameNFT721 contract.
type GameNFT721TransferIterator struct {
	Event *GameNFT721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GameNFT721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFT721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GameNFT721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GameNFT721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFT721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFT721Transfer represents a Transfer event raised by the GameNFT721 contract.
type GameNFT721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GameNFT721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFT721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFT721TransferIterator{contract: _GameNFT721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GameNFT721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFT721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFT721Transfer)
				if err := _GameNFT721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNFT721 *GameNFT721Filterer) ParseTransfer(log types.Log) (*GameNFT721Transfer, error) {
	event := new(GameNFT721Transfer)
	if err := _GameNFT721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

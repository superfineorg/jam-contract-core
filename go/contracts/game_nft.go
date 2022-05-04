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

// GameNFTMetaData contains all meta data concerning the GameNFT contract.
var GameNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200263d3803806200263d8339810160408190526200003491620002a3565b8251839083906200004d90600090602085019062000146565b5080516200006390600190602084019062000146565b505081516200007b9150600b90602084019062000146565b506200008960003362000092565b50505062000387565b6200009e8282620000a2565b5050565b6000828152600a602090815260408083206001600160a01b038516845290915290205460ff166200009e576000828152600a602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620001023390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b828054620001549062000334565b90600052602060002090601f016020900481019282620001785760008555620001c3565b82601f106200019357805160ff1916838001178555620001c3565b82800160010185558215620001c3579182015b82811115620001c3578251825591602001919060010190620001a6565b50620001d1929150620001d5565b5090565b5b80821115620001d15760008155600101620001d6565b600082601f830112620001fe57600080fd5b81516001600160401b03808211156200021b576200021b62000371565b604051601f8301601f19908116603f0116810190828211818310171562000246576200024662000371565b816040528381526020925086838588010111156200026357600080fd5b600091505b8382101562000287578582018301518183018401529082019062000268565b83821115620002995760008385830101525b9695505050505050565b600080600060608486031215620002b957600080fd5b83516001600160401b0380821115620002d157600080fd5b620002df87838801620001ec565b94506020860151915080821115620002f657600080fd5b6200030487838801620001ec565b935060408601519150808211156200031b57600080fd5b506200032a86828701620001ec565b9150509250925092565b600181811c908216806200034957607f821691505b602082108114156200036b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6122a680620003976000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80634f6ccce7116100f9578063a22cb46511610097578063d539139311610071578063d53913931461038d578063d547741f146103b4578063d547cfb7146103c7578063e985e9c5146103cf57600080fd5b8063a22cb46514610354578063b88d4fde14610367578063c87b56dd1461037a57600080fd5b806370a08231116100d357806370a082311461031e57806391d148541461033157806395d89b4114610344578063a217fddf1461034c57600080fd5b80634f6ccce7146102e55780636352211e146102f85780636a6278421461030b57600080fd5b8063248a9ca31161016657806330176e131161014057806330176e131461029957806336568abe146102ac57806342842e0e146102bf57806342966c68146102d257600080fd5b8063248a9ca3146102505780632f2ff15d146102735780632f745c591461028657600080fd5b806301ffc9a7146101ae57806306fdde03146101d6578063081812fc146101eb578063095ea7b31461021657806318160ddd1461022b57806323b872dd1461023d575b600080fd5b6101c16101bc366004611def565b61040b565b60405190151581526020015b60405180910390f35b6101de61042b565b6040516101cd9190612027565b6101fe6101f9366004611db3565b6104bd565b6040516001600160a01b0390911681526020016101cd565b610229610224366004611d89565b610557565b005b6008545b6040519081526020016101cd565b61022961024b366004611c95565b61066d565b61022f61025e366004611db3565b6000908152600a602052604090206001015490565b610229610281366004611dcc565b61069f565b61022f610294366004611d89565b6106c5565b6102296102a7366004611e29565b61075b565b6102296102ba366004611dcc565b6107d5565b6102296102cd366004611c95565b61084f565b6102296102e0366004611db3565b61086a565b61022f6102f3366004611db3565b6108e4565b6101fe610306366004611db3565b610977565b610229610319366004611c47565b6109ee565b61022f61032c366004611c47565b610a93565b6101c161033f366004611dcc565b610b1a565b6101de610b45565b61022f600081565b610229610362366004611d4d565b610b54565b610229610375366004611cd1565b610b5f565b6101de610388366004611db3565b610b97565b61022f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102296103c2366004611dcc565b610c30565b6101de610c56565b6101c16103dd366004611c62565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b600061041682610ce4565b80610425575061042582610d05565b92915050565b60606000805461043a90612182565b80601f016020809104026020016040519081016040528092919081815260200182805461046690612182565b80156104b35780601f10610488576101008083540402835291602001916104b3565b820191906000526020600020905b81548152906001019060200180831161049657829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b031661053b5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b600061056282610977565b9050806001600160a01b0316836001600160a01b031614156105d05760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610532565b336001600160a01b03821614806105ec57506105ec81336103dd565b61065e5760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610532565b6106688383610d2a565b505050565b610678335b82610d98565b6106945760405162461bcd60e51b81526004016105329061208c565b610668838383610e8f565b6000828152600a60205260409020600101546106bb8133611036565b610668838361109a565b60006106d083610a93565b82106107325760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610532565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b610766600033610b1a565b6107be5760405162461bcd60e51b8152602060048201526024808201527f47616d654e46543a206d75737420686176652061646d696e20726f6c6520746f604482015263081cd95d60e21b6064820152608401610532565b80516107d190600b906020840190611b1c565b5050565b6001600160a01b03811633146108455760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610532565b6107d18282611120565b61066883838360405180602001604052806000815250610b5f565b61087333610672565b6108d85760405162461bcd60e51b815260206004820152603060248201527f4552433732314275726e61626c653a2063616c6c6572206973206e6f74206f7760448201526f1b995c881b9bdc88185c1c1c9bdd995960821b6064820152608401610532565b6108e181611187565b50565b60006108ef60085490565b82106109525760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610532565b600882815481106109655761096561222e565b90600052602060002001549050919050565b6000818152600260205260408120546001600160a01b0316806104255760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610532565b610a187f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633610b1a565b610a735760405162461bcd60e51b815260206004820152602660248201527f47616d654e46543a206d7573742068617665206d696e74657220726f6c6520746044820152651bc81b5a5b9d60d21b6064820152608401610532565b610a8581610a80600c5490565b61122e565b6108e1600c80546001019055565b60006001600160a01b038216610afe5760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610532565b506001600160a01b031660009081526003602052604090205490565b6000918252600a602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606001805461043a90612182565b6107d1338383611248565b610b693383610d98565b610b855760405162461bcd60e51b81526004016105329061208c565b610b9184848484611317565b50505050565b6000818152600260205260409020546060906001600160a01b0316610bfe5760405162461bcd60e51b815260206004820152601d60248201527f47616d654e46543a20746f6b656e20646f6573206e6f742065786973740000006044820152606401610532565b600b610c098361134a565b604051602001610c1a929190611eba565b6040516020818303038152906040529050919050565b6000828152600a6020526040902060010154610c4c8133611036565b6106688383611120565b600b8054610c6390612182565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8f90612182565b8015610cdc5780601f10610cb157610100808354040283529160200191610cdc565b820191906000526020600020905b815481529060010190602001808311610cbf57829003601f168201915b505050505081565b60006001600160e01b03198216637965db0b60e01b14806104255750610425825b60006001600160e01b0319821663780e9d6360e01b1480610425575061042582611448565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610d5f82610977565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000818152600260205260408120546001600160a01b0316610e115760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610532565b6000610e1c83610977565b9050806001600160a01b0316846001600160a01b03161480610e575750836001600160a01b0316610e4c846104bd565b6001600160a01b0316145b80610e8757506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b0316610ea282610977565b6001600160a01b031614610f065760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610532565b6001600160a01b038216610f685760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610532565b610f73838383611498565b610f7e600082610d2a565b6001600160a01b0383166000908152600360205260408120805460019290610fa7908490612128565b90915550506001600160a01b0382166000908152600360205260408120805460019290610fd59084906120dd565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6110408282610b1a565b6107d157611058816001600160a01b031660146114a3565b6110638360206114a3565b604051602001611074929190611f75565b60408051601f198184030181529082905262461bcd60e51b825261053291600401612027565b6110a48282610b1a565b6107d1576000828152600a602090815260408083206001600160a01b03851684529091529020805460ff191660011790556110dc3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61112a8282610b1a565b156107d1576000828152600a602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600061119282610977565b90506111a081600084611498565b6111ab600083610d2a565b6001600160a01b03811660009081526003602052604081208054600192906111d4908490612128565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6107d1828260405180602001604052806000815250611646565b816001600160a01b0316836001600160a01b031614156112aa5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610532565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b611322848484610e8f565b61132e84848484611679565b610b915760405162461bcd60e51b81526004016105329061203a565b60608161136e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156113985780611382816121bd565b91506113919050600a836120f5565b9150611372565b60008167ffffffffffffffff8111156113b3576113b3612244565b6040519080825280601f01601f1916602001820160405280156113dd576020820181803683370190505b5090505b8415610e87576113f2600183612128565b91506113ff600a866121d8565b61140a9060306120dd565b60f81b81838151811061141f5761141f61222e565b60200101906001600160f81b031916908160001a905350611441600a866120f5565b94506113e1565b60006001600160e01b031982166380ac58cd60e01b148061147957506001600160e01b03198216635b5e139f60e01b145b8061042557506301ffc9a760e01b6001600160e01b0319831614610425565b610668838383611786565b606060006114b2836002612109565b6114bd9060026120dd565b67ffffffffffffffff8111156114d5576114d5612244565b6040519080825280601f01601f1916602001820160405280156114ff576020820181803683370190505b509050600360fc1b8160008151811061151a5761151a61222e565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106115495761154961222e565b60200101906001600160f81b031916908160001a905350600061156d846002612109565b6115789060016120dd565b90505b60018111156115f0576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106115ac576115ac61222e565b1a60f81b8282815181106115c2576115c261222e565b60200101906001600160f81b031916908160001a90535060049490941c936115e98161216b565b905061157b565b50831561163f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610532565b9392505050565b611650838361183e565b61165d6000848484611679565b6106685760405162461bcd60e51b81526004016105329061203a565b60006001600160a01b0384163b1561177b57604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906116bd903390899088908890600401611fea565b602060405180830381600087803b1580156116d757600080fd5b505af1925050508015611707575060408051601f3d908101601f1916820190925261170491810190611e0c565b60015b611761573d808015611735576040519150601f19603f3d011682016040523d82523d6000602084013e61173a565b606091505b5080516117595760405162461bcd60e51b81526004016105329061203a565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610e87565b506001949350505050565b6001600160a01b0383166117e1576117dc81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b611804565b816001600160a01b0316836001600160a01b03161461180457611804838261198c565b6001600160a01b03821661181b5761066881611a29565b826001600160a01b0316826001600160a01b031614610668576106688282611ad8565b6001600160a01b0382166118945760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610532565b6000818152600260205260409020546001600160a01b0316156118f95760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610532565b61190560008383611498565b6001600160a01b038216600090815260036020526040812080546001929061192e9084906120dd565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000600161199984610a93565b6119a39190612128565b6000838152600760205260409020549091508082146119f6576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b600854600090611a3b90600190612128565b60008381526009602052604081205460088054939450909284908110611a6357611a6361222e565b906000526020600020015490508060088381548110611a8457611a8461222e565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611abc57611abc612218565b6001900381819060005260206000200160009055905550505050565b6000611ae383610a93565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b828054611b2890612182565b90600052602060002090601f016020900481019282611b4a5760008555611b90565b82601f10611b6357805160ff1916838001178555611b90565b82800160010185558215611b90579182015b82811115611b90578251825591602001919060010190611b75565b50611b9c929150611ba0565b5090565b5b80821115611b9c5760008155600101611ba1565b600067ffffffffffffffff80841115611bd057611bd0612244565b604051601f8501601f19908116603f01168101908282118183101715611bf857611bf8612244565b81604052809350858152868686011115611c1157600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b0381168114611c4257600080fd5b919050565b600060208284031215611c5957600080fd5b61163f82611c2b565b60008060408385031215611c7557600080fd5b611c7e83611c2b565b9150611c8c60208401611c2b565b90509250929050565b600080600060608486031215611caa57600080fd5b611cb384611c2b565b9250611cc160208501611c2b565b9150604084013590509250925092565b60008060008060808587031215611ce757600080fd5b611cf085611c2b565b9350611cfe60208601611c2b565b925060408501359150606085013567ffffffffffffffff811115611d2157600080fd5b8501601f81018713611d3257600080fd5b611d4187823560208401611bb5565b91505092959194509250565b60008060408385031215611d6057600080fd5b611d6983611c2b565b915060208301358015158114611d7e57600080fd5b809150509250929050565b60008060408385031215611d9c57600080fd5b611da583611c2b565b946020939093013593505050565b600060208284031215611dc557600080fd5b5035919050565b60008060408385031215611ddf57600080fd5b82359150611c8c60208401611c2b565b600060208284031215611e0157600080fd5b813561163f8161225a565b600060208284031215611e1e57600080fd5b815161163f8161225a565b600060208284031215611e3b57600080fd5b813567ffffffffffffffff811115611e5257600080fd5b8201601f81018413611e6357600080fd5b610e8784823560208401611bb5565b60008151808452611e8a81602086016020860161213f565b601f01601f19169290920160200192915050565b60008151611eb081856020860161213f565b9290920192915050565b600080845481600182811c915080831680611ed657607f831692505b6020808410821415611ef657634e487b7160e01b86526022600452602486fd5b818015611f0a5760018114611f1b57611f48565b60ff19861689528489019650611f48565b60008b81526020902060005b86811015611f405781548b820152908501908301611f27565b505084890196505b505050505050611f6c611f5b8286611e9e565b64173539b7b760d91b815260050190565b95945050505050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611fad81601785016020880161213f565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611fde81602884016020880161213f565b01602801949350505050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061201d90830184611e72565b9695505050505050565b60208152600061163f6020830184611e72565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b600082198211156120f0576120f06121ec565b500190565b60008261210457612104612202565b500490565b6000816000190483118215151615612123576121236121ec565b500290565b60008282101561213a5761213a6121ec565b500390565b60005b8381101561215a578181015183820152602001612142565b83811115610b915750506000910152565b60008161217a5761217a6121ec565b506000190190565b600181811c9082168061219657607f821691505b602082108114156121b757634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156121d1576121d16121ec565b5060010190565b6000826121e7576121e7612202565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b0319811681146108e157600080fdfea26469706673582212206ba8ea8499eba111e30985e6d8fae4d4d7a7dc33dd22ee79a6f3496513e284ac64736f6c63430008060033",
}

// GameNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use GameNFTMetaData.ABI instead.
var GameNFTABI = GameNFTMetaData.ABI

// GameNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameNFTMetaData.Bin instead.
var GameNFTBin = GameNFTMetaData.Bin

// DeployGameNFT deploys a new Ethereum contract, binding an instance of GameNFT to it.
func DeployGameNFT(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, baseTokenURI_ string) (common.Address, *types.Transaction, *GameNFT, error) {
	parsed, err := GameNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameNFTBin), backend, name, symbol, baseTokenURI_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameNFT{GameNFTCaller: GameNFTCaller{contract: contract}, GameNFTTransactor: GameNFTTransactor{contract: contract}, GameNFTFilterer: GameNFTFilterer{contract: contract}}, nil
}

// GameNFT is an auto generated Go binding around an Ethereum contract.
type GameNFT struct {
	GameNFTCaller     // Read-only binding to the contract
	GameNFTTransactor // Write-only binding to the contract
	GameNFTFilterer   // Log filterer for contract events
}

// GameNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameNFTSession struct {
	Contract     *GameNFT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameNFTCallerSession struct {
	Contract *GameNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GameNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameNFTTransactorSession struct {
	Contract     *GameNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GameNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameNFTRaw struct {
	Contract *GameNFT // Generic contract binding to access the raw methods on
}

// GameNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameNFTCallerRaw struct {
	Contract *GameNFTCaller // Generic read-only contract binding to access the raw methods on
}

// GameNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameNFTTransactorRaw struct {
	Contract *GameNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameNFT creates a new instance of GameNFT, bound to a specific deployed contract.
func NewGameNFT(address common.Address, backend bind.ContractBackend) (*GameNFT, error) {
	contract, err := bindGameNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameNFT{GameNFTCaller: GameNFTCaller{contract: contract}, GameNFTTransactor: GameNFTTransactor{contract: contract}, GameNFTFilterer: GameNFTFilterer{contract: contract}}, nil
}

// NewGameNFTCaller creates a new read-only instance of GameNFT, bound to a specific deployed contract.
func NewGameNFTCaller(address common.Address, caller bind.ContractCaller) (*GameNFTCaller, error) {
	contract, err := bindGameNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFTCaller{contract: contract}, nil
}

// NewGameNFTTransactor creates a new write-only instance of GameNFT, bound to a specific deployed contract.
func NewGameNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*GameNFTTransactor, error) {
	contract, err := bindGameNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFTTransactor{contract: contract}, nil
}

// NewGameNFTFilterer creates a new log filterer instance of GameNFT, bound to a specific deployed contract.
func NewGameNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*GameNFTFilterer, error) {
	contract, err := bindGameNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameNFTFilterer{contract: contract}, nil
}

// bindGameNFT binds a generic wrapper to an already deployed contract.
func bindGameNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFT *GameNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFT.Contract.GameNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFT *GameNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFT.Contract.GameNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFT *GameNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFT.Contract.GameNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFT *GameNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFT *GameNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFT *GameNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameNFT.Contract.DEFAULTADMINROLE(&_GameNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameNFT.Contract.DEFAULTADMINROLE(&_GameNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTSession) MINTERROLE() ([32]byte, error) {
	return _GameNFT.Contract.MINTERROLE(&_GameNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_GameNFT *GameNFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _GameNFT.Contract.MINTERROLE(&_GameNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT *GameNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT *GameNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNFT.Contract.BalanceOf(&_GameNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNFT *GameNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNFT.Contract.BalanceOf(&_GameNFT.CallOpts, owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT *GameNFTCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT *GameNFTSession) BaseTokenURI() (string, error) {
	return _GameNFT.Contract.BaseTokenURI(&_GameNFT.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_GameNFT *GameNFTCallerSession) BaseTokenURI() (string, error) {
	return _GameNFT.Contract.BaseTokenURI(&_GameNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNFT.Contract.GetApproved(&_GameNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNFT.Contract.GetApproved(&_GameNFT.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT *GameNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT *GameNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameNFT.Contract.GetRoleAdmin(&_GameNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameNFT *GameNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameNFT.Contract.GetRoleAdmin(&_GameNFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT *GameNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT *GameNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameNFT.Contract.HasRole(&_GameNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameNFT *GameNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameNFT.Contract.HasRole(&_GameNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT *GameNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT *GameNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNFT.Contract.IsApprovedForAll(&_GameNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNFT *GameNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNFT.Contract.IsApprovedForAll(&_GameNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT *GameNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT *GameNFTSession) Name() (string, error) {
	return _GameNFT.Contract.Name(&_GameNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNFT *GameNFTCallerSession) Name() (string, error) {
	return _GameNFT.Contract.Name(&_GameNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNFT.Contract.OwnerOf(&_GameNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNFT *GameNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNFT.Contract.OwnerOf(&_GameNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT *GameNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT *GameNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNFT.Contract.SupportsInterface(&_GameNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNFT *GameNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNFT.Contract.SupportsInterface(&_GameNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT *GameNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT *GameNFTSession) Symbol() (string, error) {
	return _GameNFT.Contract.Symbol(&_GameNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNFT *GameNFTCallerSession) Symbol() (string, error) {
	return _GameNFT.Contract.Symbol(&_GameNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT *GameNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT *GameNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GameNFT.Contract.TokenByIndex(&_GameNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_GameNFT *GameNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _GameNFT.Contract.TokenByIndex(&_GameNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT *GameNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT *GameNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GameNFT.Contract.TokenOfOwnerByIndex(&_GameNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_GameNFT *GameNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _GameNFT.Contract.TokenOfOwnerByIndex(&_GameNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT *GameNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT *GameNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNFT.Contract.TokenURI(&_GameNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNFT *GameNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNFT.Contract.TokenURI(&_GameNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT *GameNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GameNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT *GameNFTSession) TotalSupply() (*big.Int, error) {
	return _GameNFT.Contract.TotalSupply(&_GameNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GameNFT *GameNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _GameNFT.Contract.TotalSupply(&_GameNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.Approve(&_GameNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.Approve(&_GameNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT *GameNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.Burn(&_GameNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.Burn(&_GameNFT.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.GrantRole(&_GameNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.GrantRole(&_GameNFT.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT *GameNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT *GameNFTSession) Mint(to common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.Mint(&_GameNFT.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_GameNFT *GameNFTTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.Mint(&_GameNFT.TransactOpts, to)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.RenounceRole(&_GameNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.RenounceRole(&_GameNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.RevokeRole(&_GameNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameNFT *GameNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameNFT.Contract.RevokeRole(&_GameNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.SafeTransferFrom(&_GameNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.SafeTransferFrom(&_GameNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT *GameNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT *GameNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT.Contract.SafeTransferFrom0(&_GameNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_GameNFT *GameNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _GameNFT.Contract.SafeTransferFrom0(&_GameNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT *GameNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT *GameNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT.Contract.SetApprovalForAll(&_GameNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNFT *GameNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNFT.Contract.SetApprovalForAll(&_GameNFT.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT *GameNFTTransactor) SetBaseTokenURI(opts *bind.TransactOpts, baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "setBaseTokenURI", baseTokenURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT *GameNFTSession) SetBaseTokenURI(baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT.Contract.SetBaseTokenURI(&_GameNFT.TransactOpts, baseTokenURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseTokenURI_) returns()
func (_GameNFT *GameNFTTransactorSession) SetBaseTokenURI(baseTokenURI_ string) (*types.Transaction, error) {
	return _GameNFT.Contract.SetBaseTokenURI(&_GameNFT.TransactOpts, baseTokenURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.TransferFrom(&_GameNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNFT *GameNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFT.Contract.TransferFrom(&_GameNFT.TransactOpts, from, to, tokenId)
}

// GameNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GameNFT contract.
type GameNFTApprovalIterator struct {
	Event *GameNFTApproval // Event containing the contract specifics and raw log

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
func (it *GameNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTApproval)
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
		it.Event = new(GameNFTApproval)
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
func (it *GameNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTApproval represents a Approval event raised by the GameNFT contract.
type GameNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNFT *GameNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GameNFTApprovalIterator, error) {

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

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTApprovalIterator{contract: _GameNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNFT *GameNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GameNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTApproval)
				if err := _GameNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseApproval(log types.Log) (*GameNFTApproval, error) {
	event := new(GameNFTApproval)
	if err := _GameNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GameNFT contract.
type GameNFTApprovalForAllIterator struct {
	Event *GameNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GameNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTApprovalForAll)
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
		it.Event = new(GameNFTApprovalForAll)
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
func (it *GameNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTApprovalForAll represents a ApprovalForAll event raised by the GameNFT contract.
type GameNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNFT *GameNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GameNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTApprovalForAllIterator{contract: _GameNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNFT *GameNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GameNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTApprovalForAll)
				if err := _GameNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseApprovalForAll(log types.Log) (*GameNFTApprovalForAll, error) {
	event := new(GameNFTApprovalForAll)
	if err := _GameNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GameNFT contract.
type GameNFTRoleAdminChangedIterator struct {
	Event *GameNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GameNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTRoleAdminChanged)
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
		it.Event = new(GameNFTRoleAdminChanged)
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
func (it *GameNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTRoleAdminChanged represents a RoleAdminChanged event raised by the GameNFT contract.
type GameNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameNFT *GameNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GameNFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTRoleAdminChangedIterator{contract: _GameNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameNFT *GameNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GameNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTRoleAdminChanged)
				if err := _GameNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseRoleAdminChanged(log types.Log) (*GameNFTRoleAdminChanged, error) {
	event := new(GameNFTRoleAdminChanged)
	if err := _GameNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GameNFT contract.
type GameNFTRoleGrantedIterator struct {
	Event *GameNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *GameNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTRoleGranted)
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
		it.Event = new(GameNFTRoleGranted)
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
func (it *GameNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTRoleGranted represents a RoleGranted event raised by the GameNFT contract.
type GameNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT *GameNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameNFTRoleGrantedIterator, error) {

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

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTRoleGrantedIterator{contract: _GameNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT *GameNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GameNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTRoleGranted)
				if err := _GameNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseRoleGranted(log types.Log) (*GameNFTRoleGranted, error) {
	event := new(GameNFTRoleGranted)
	if err := _GameNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GameNFT contract.
type GameNFTRoleRevokedIterator struct {
	Event *GameNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GameNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTRoleRevoked)
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
		it.Event = new(GameNFTRoleRevoked)
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
func (it *GameNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTRoleRevoked represents a RoleRevoked event raised by the GameNFT contract.
type GameNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT *GameNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameNFTRoleRevokedIterator, error) {

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

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTRoleRevokedIterator{contract: _GameNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameNFT *GameNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GameNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTRoleRevoked)
				if err := _GameNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseRoleRevoked(log types.Log) (*GameNFTRoleRevoked, error) {
	event := new(GameNFTRoleRevoked)
	if err := _GameNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GameNFT contract.
type GameNFTTransferIterator struct {
	Event *GameNFTTransfer // Event containing the contract specifics and raw log

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
func (it *GameNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTTransfer)
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
		it.Event = new(GameNFTTransfer)
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
func (it *GameNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTTransfer represents a Transfer event raised by the GameNFT contract.
type GameNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNFT *GameNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GameNFTTransferIterator, error) {

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

	logs, sub, err := _GameNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTTransferIterator{contract: _GameNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNFT *GameNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GameNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GameNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTTransfer)
				if err := _GameNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GameNFT *GameNFTFilterer) ParseTransfer(log types.Log) (*GameNFTTransfer, error) {
	event := new(GameNFTTransfer)
	if err := _GameNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

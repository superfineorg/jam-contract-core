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

// GameNFTBurningMetaData contains all meta data concerning the GameNFTBurning contract.
var GameNFTBurningMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"}],\"name\":\"Erc1155BurnedIntoGames\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Erc721BurnedIntoGames\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"nftAddresses\",\"type\":\"address[]\"}],\"name\":\"burnErc1155IntoGames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nftAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnErc721IntoGames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isOperators\",\"type\":\"bool[]\"}],\"name\":\"setOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361003c565b336000908152600160208190526040909120805460ff1916909117905561008c565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61127a8061009b6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631a2b08d714610067578063552faf9f1461007c578063715018a61461008f5780638da5cb5b14610097578063f2fde38b146100b6578063f4305035146100c9575b600080fd5b61007a610075366004610d95565b6100dc565b005b61007a61008a366004610c66565b610428565b61007a6108a1565b600054604080516001600160a01b039092168252519081900360200190f35b61007a6100c4366004610c25565b6108d7565b61007a6100d7366004610cca565b610972565b3360009081526001602052604090205460ff166101145760405162461bcd60e51b815260040161010b906110a3565b60405180910390fd5b80518251146101355760405162461bcd60e51b815260040161010b906110e9565b60005b8251811015610423576000610165848381518110610158576101586111f5565b6020026020010151610a3d565b9050801561020c5783828151811061017f5761017f6111f5565b60209081029190910101516040516301ffc9a760e01b81526380ac58cd60e01b60048201526001600160a01b03909116906301ffc9a79060240160206040518083038186803b1580156101d157600080fd5b505afa1580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102099190610f96565b90505b806102685760405162461bcd60e51b815260206004820152602660248201527f47616d654e46544275726e696e673a20696e76616c696420455243373231206160448201526564647265737360d01b606482015260840161010b565b600084838151811061027c5761027c6111f5565b602002602001015190506000816001600160a01b0316636352211e8686815181106102a9576102a96111f5565b60200260200101516040518263ffffffff1660e01b81526004016102cf91815260200190565b60206040518083038186803b1580156102e757600080fd5b505afa1580156102fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031f9190610c49565b9050816001600160a01b03166342966c68868681518110610342576103426111f5565b60200260200101516040518263ffffffff1660e01b815260040161036891815260200190565b600060405180830381600087803b15801561038257600080fd5b505af1158015610396573d6000803e3d6000fd5b505050508484815181106103ac576103ac6111f5565b60200260200101518685815181106103c6576103c66111f5565b60200260200101516001600160a01b0316826001600160a01b03167fbb9f821ebd56f09e11a6397d5f4ae52dde8ddcf710e835e85f339ea7d234b01160405160405180910390a4505050808061041b906111cc565b915050610138565b505050565b3360009081526001602052604090205460ff166104575760405162461bcd60e51b815260040161010b906110a3565b80518251146104785760405162461bcd60e51b815260040161010b906110e9565b60005b825181101561042357600061049b838381518110610158576101586111f5565b90508015610542578282815181106104b5576104b56111f5565b60209081029190910101516040516301ffc9a760e01b8152636cdb3d1360e11b60048201526001600160a01b03909116906301ffc9a79060240160206040518083038186803b15801561050757600080fd5b505afa15801561051b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053f9190610f96565b90505b8061059f5760405162461bcd60e51b815260206004820152602760248201527f47616d654e46544275726e696e673a20696e76616c69642045524331313535206044820152666164647265737360c81b606482015260840161010b565b60008383815181106105b3576105b36111f5565b602002602001015190506000816001600160a01b0316630f915c648786815181106105e0576105e06111f5565b60200260200101516040518263ffffffff1660e01b815260040161061391906001600160a01b0391909116815260200190565b60006040518083038186803b15801561062b57600080fd5b505afa15801561063f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106679190810190610e47565b90506000815167ffffffffffffffff8111156106855761068561120b565b6040519080825280602002602001820160405280156106ae578160200160208202803683370190505b5090506000825167ffffffffffffffff8111156106cd576106cd61120b565b6040519080825280602002602001820160405280156106f6578160200160208202803683370190505b50905060005b835181101561078e57838181518110610717576107176111f5565b602002602001015160000151838281518110610735576107356111f5565b602002602001018181525050838181518110610753576107536111f5565b602002602001015160200151828281518110610771576107716111f5565b602090810291909101015280610786816111cc565b9150506106fc565b50836001600160a01b0316636b20c4548988815181106107b0576107b06111f5565b602002602001015184846040518463ffffffff1660e01b81526004016107d89392919061100a565b600060405180830381600087803b1580156107f257600080fd5b505af1158015610806573d6000803e3d6000fd5b5050505086868151811061081c5761081c6111f5565b60200260200101516001600160a01b031688878151811061083f5761083f6111f5565b60200260200101516001600160a01b03167f4608fa48020b6d192febfc47c7b96afb63993ce6f9ded5572cd09068eb6d11788484604051610881929190611040565b60405180910390a350505050508080610899906111cc565b91505061047b565b6000546001600160a01b031633146108cb5760405162461bcd60e51b815260040161010b9061106e565b6108d56000610a71565b565b6000546001600160a01b031633146109015760405162461bcd60e51b815260040161010b9061106e565b6001600160a01b0381166109665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161010b565b61096f81610a71565b50565b6000546001600160a01b0316331461099c5760405162461bcd60e51b815260040161010b9061106e565b80518251146109bd5760405162461bcd60e51b815260040161010b906110e9565b60005b8251811015610423578181815181106109db576109db6111f5565b6020026020010151600160008584815181106109f9576109f96111f5565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580610a35816111cc565b9150506109c0565b6000610a50826301ffc9a760e01b610ac1565b8015610a6b5750610a69826001600160e01b0319610ac1565b155b92915050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b0387169061753090610b28908690610fee565b6000604051808303818686fa925050503d8060008114610b64576040519150601f19603f3d011682016040523d82523d6000602084013e610b69565b606091505b5091509150602081511015610b845760009350505050610a6b565b818015610ba0575080806020019051810190610ba09190610f96565b9695505050505050565b600082601f830112610bbb57600080fd5b81356020610bd0610bcb83611178565b611147565b80838252828201915082860187848660051b8901011115610bf057600080fd5b60005b85811015610c18578135610c0681611221565b84529284019290840190600101610bf3565b5090979650505050505050565b600060208284031215610c3757600080fd5b8135610c4281611221565b9392505050565b600060208284031215610c5b57600080fd5b8151610c4281611221565b60008060408385031215610c7957600080fd5b823567ffffffffffffffff80821115610c9157600080fd5b610c9d86838701610baa565b93506020850135915080821115610cb357600080fd5b50610cc085828601610baa565b9150509250929050565b60008060408385031215610cdd57600080fd5b823567ffffffffffffffff80821115610cf557600080fd5b610d0186838701610baa565b9350602091508185013581811115610d1857600080fd5b85019050601f81018613610d2b57600080fd5b8035610d39610bcb82611178565b80828252848201915084840189868560051b8701011115610d5957600080fd5b600094505b83851015610d85578035610d7181611236565b835260019490940193918501918501610d5e565b5080955050505050509250929050565b60008060408385031215610da857600080fd5b823567ffffffffffffffff80821115610dc057600080fd5b610dcc86838701610baa565b9350602091508185013581811115610de357600080fd5b85019050601f81018613610df657600080fd5b8035610e04610bcb82611178565b80828252848201915084840189868560051b8701011115610e2457600080fd5b600094505b83851015610d85578035835260019490940193918501918501610e29565b60006020808385031215610e5a57600080fd5b825167ffffffffffffffff80821115610e7257600080fd5b818501915085601f830112610e8657600080fd5b8151610e94610bcb82611178565b80828252858201915085850189878560051b8801011115610eb457600080fd5b60005b84811015610f8757815186811115610ece57600080fd5b8701601f196060828e0382011215610ee557600080fd5b610eed61111e565b8a83015181526040808401518c83015260608401518a811115610f0f57600080fd5b8085019450508e603f850112610f2457600080fd5b8b8401518a811115610f3857610f3861120b565b610f488d85601f84011601611147565b93508084528f82828701011115610f5e57600080fd5b610f6d818e860184880161119c565b508101919091528552509287019290870190600101610eb7565b50909998505050505050505050565b600060208284031215610fa857600080fd5b8151610c4281611236565b600081518084526020808501945080840160005b83811015610fe357815187529582019590820190600101610fc7565b509495945050505050565b6000825161100081846020870161119c565b9190910192915050565b6001600160a01b038416815260606020820181905260009061102e90830185610fb3565b8281036040840152610ba08185610fb3565b6040815260006110536040830185610fb3565b82810360208401526110658185610fb3565b95945050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526026908201527f47616d654e46544275726e696e673a2063616c6c6572206973206e6f74206f7060408201526532b930ba37b960d11b606082015260800190565b6020808252818101527f47616d654e46544275726e696e673a206c656e67746873206d69736d61746368604082015260600190565b6040516060810167ffffffffffffffff811182821017156111415761114161120b565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156111705761117061120b565b604052919050565b600067ffffffffffffffff8211156111925761119261120b565b5060051b60200190565b60005b838110156111b757818101518382015260200161119f565b838111156111c6576000848401525b50505050565b60006000198214156111ee57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461096f57600080fd5b801515811461096f57600080fdfea2646970667358221220b17cd32f8ba77165c9e2f8bbd0ea8fcfd8e8910975a15e16783f65a713efa43264736f6c63430008060033",
}

// GameNFTBurningABI is the input ABI used to generate the binding from.
// Deprecated: Use GameNFTBurningMetaData.ABI instead.
var GameNFTBurningABI = GameNFTBurningMetaData.ABI

// GameNFTBurningBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameNFTBurningMetaData.Bin instead.
var GameNFTBurningBin = GameNFTBurningMetaData.Bin

// DeployGameNFTBurning deploys a new Ethereum contract, binding an instance of GameNFTBurning to it.
func DeployGameNFTBurning(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameNFTBurning, error) {
	parsed, err := GameNFTBurningMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameNFTBurningBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameNFTBurning{GameNFTBurningCaller: GameNFTBurningCaller{contract: contract}, GameNFTBurningTransactor: GameNFTBurningTransactor{contract: contract}, GameNFTBurningFilterer: GameNFTBurningFilterer{contract: contract}}, nil
}

// GameNFTBurning is an auto generated Go binding around an Ethereum contract.
type GameNFTBurning struct {
	GameNFTBurningCaller     // Read-only binding to the contract
	GameNFTBurningTransactor // Write-only binding to the contract
	GameNFTBurningFilterer   // Log filterer for contract events
}

// GameNFTBurningCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameNFTBurningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTBurningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameNFTBurningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTBurningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameNFTBurningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNFTBurningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameNFTBurningSession struct {
	Contract     *GameNFTBurning   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameNFTBurningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameNFTBurningCallerSession struct {
	Contract *GameNFTBurningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GameNFTBurningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameNFTBurningTransactorSession struct {
	Contract     *GameNFTBurningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GameNFTBurningRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameNFTBurningRaw struct {
	Contract *GameNFTBurning // Generic contract binding to access the raw methods on
}

// GameNFTBurningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameNFTBurningCallerRaw struct {
	Contract *GameNFTBurningCaller // Generic read-only contract binding to access the raw methods on
}

// GameNFTBurningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameNFTBurningTransactorRaw struct {
	Contract *GameNFTBurningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameNFTBurning creates a new instance of GameNFTBurning, bound to a specific deployed contract.
func NewGameNFTBurning(address common.Address, backend bind.ContractBackend) (*GameNFTBurning, error) {
	contract, err := bindGameNFTBurning(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurning{GameNFTBurningCaller: GameNFTBurningCaller{contract: contract}, GameNFTBurningTransactor: GameNFTBurningTransactor{contract: contract}, GameNFTBurningFilterer: GameNFTBurningFilterer{contract: contract}}, nil
}

// NewGameNFTBurningCaller creates a new read-only instance of GameNFTBurning, bound to a specific deployed contract.
func NewGameNFTBurningCaller(address common.Address, caller bind.ContractCaller) (*GameNFTBurningCaller, error) {
	contract, err := bindGameNFTBurning(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningCaller{contract: contract}, nil
}

// NewGameNFTBurningTransactor creates a new write-only instance of GameNFTBurning, bound to a specific deployed contract.
func NewGameNFTBurningTransactor(address common.Address, transactor bind.ContractTransactor) (*GameNFTBurningTransactor, error) {
	contract, err := bindGameNFTBurning(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningTransactor{contract: contract}, nil
}

// NewGameNFTBurningFilterer creates a new log filterer instance of GameNFTBurning, bound to a specific deployed contract.
func NewGameNFTBurningFilterer(address common.Address, filterer bind.ContractFilterer) (*GameNFTBurningFilterer, error) {
	contract, err := bindGameNFTBurning(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningFilterer{contract: contract}, nil
}

// bindGameNFTBurning binds a generic wrapper to an already deployed contract.
func bindGameNFTBurning(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameNFTBurningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFTBurning *GameNFTBurningRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFTBurning.Contract.GameNFTBurningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFTBurning *GameNFTBurningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.GameNFTBurningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFTBurning *GameNFTBurningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.GameNFTBurningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNFTBurning *GameNFTBurningCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNFTBurning.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNFTBurning *GameNFTBurningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNFTBurning *GameNFTBurningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNFTBurning *GameNFTBurningCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameNFTBurning.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNFTBurning *GameNFTBurningSession) Owner() (common.Address, error) {
	return _GameNFTBurning.Contract.Owner(&_GameNFTBurning.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNFTBurning *GameNFTBurningCallerSession) Owner() (common.Address, error) {
	return _GameNFTBurning.Contract.Owner(&_GameNFTBurning.CallOpts)
}

// BurnErc1155IntoGames is a paid mutator transaction binding the contract method 0x552faf9f.
//
// Solidity: function burnErc1155IntoGames(address[] owners, address[] nftAddresses) returns()
func (_GameNFTBurning *GameNFTBurningTransactor) BurnErc1155IntoGames(opts *bind.TransactOpts, owners []common.Address, nftAddresses []common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "burnErc1155IntoGames", owners, nftAddresses)
}

// BurnErc1155IntoGames is a paid mutator transaction binding the contract method 0x552faf9f.
//
// Solidity: function burnErc1155IntoGames(address[] owners, address[] nftAddresses) returns()
func (_GameNFTBurning *GameNFTBurningSession) BurnErc1155IntoGames(owners []common.Address, nftAddresses []common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnErc1155IntoGames(&_GameNFTBurning.TransactOpts, owners, nftAddresses)
}

// BurnErc1155IntoGames is a paid mutator transaction binding the contract method 0x552faf9f.
//
// Solidity: function burnErc1155IntoGames(address[] owners, address[] nftAddresses) returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) BurnErc1155IntoGames(owners []common.Address, nftAddresses []common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnErc1155IntoGames(&_GameNFTBurning.TransactOpts, owners, nftAddresses)
}

// BurnErc721IntoGames is a paid mutator transaction binding the contract method 0x1a2b08d7.
//
// Solidity: function burnErc721IntoGames(address[] nftAddresses, uint256[] tokenIds) returns()
func (_GameNFTBurning *GameNFTBurningTransactor) BurnErc721IntoGames(opts *bind.TransactOpts, nftAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "burnErc721IntoGames", nftAddresses, tokenIds)
}

// BurnErc721IntoGames is a paid mutator transaction binding the contract method 0x1a2b08d7.
//
// Solidity: function burnErc721IntoGames(address[] nftAddresses, uint256[] tokenIds) returns()
func (_GameNFTBurning *GameNFTBurningSession) BurnErc721IntoGames(nftAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnErc721IntoGames(&_GameNFTBurning.TransactOpts, nftAddresses, tokenIds)
}

// BurnErc721IntoGames is a paid mutator transaction binding the contract method 0x1a2b08d7.
//
// Solidity: function burnErc721IntoGames(address[] nftAddresses, uint256[] tokenIds) returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) BurnErc721IntoGames(nftAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnErc721IntoGames(&_GameNFTBurning.TransactOpts, nftAddresses, tokenIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNFTBurning *GameNFTBurningTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNFTBurning *GameNFTBurningSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameNFTBurning.Contract.RenounceOwnership(&_GameNFTBurning.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameNFTBurning.Contract.RenounceOwnership(&_GameNFTBurning.TransactOpts)
}

// SetOperators is a paid mutator transaction binding the contract method 0xf4305035.
//
// Solidity: function setOperators(address[] operators, bool[] isOperators) returns()
func (_GameNFTBurning *GameNFTBurningTransactor) SetOperators(opts *bind.TransactOpts, operators []common.Address, isOperators []bool) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "setOperators", operators, isOperators)
}

// SetOperators is a paid mutator transaction binding the contract method 0xf4305035.
//
// Solidity: function setOperators(address[] operators, bool[] isOperators) returns()
func (_GameNFTBurning *GameNFTBurningSession) SetOperators(operators []common.Address, isOperators []bool) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.SetOperators(&_GameNFTBurning.TransactOpts, operators, isOperators)
}

// SetOperators is a paid mutator transaction binding the contract method 0xf4305035.
//
// Solidity: function setOperators(address[] operators, bool[] isOperators) returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) SetOperators(operators []common.Address, isOperators []bool) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.SetOperators(&_GameNFTBurning.TransactOpts, operators, isOperators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameNFTBurning *GameNFTBurningTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameNFTBurning *GameNFTBurningSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.TransferOwnership(&_GameNFTBurning.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.TransferOwnership(&_GameNFTBurning.TransactOpts, newOwner)
}

// GameNFTBurningErc1155BurnedIntoGamesIterator is returned from FilterErc1155BurnedIntoGames and is used to iterate over the raw logs and unpacked data for Erc1155BurnedIntoGames events raised by the GameNFTBurning contract.
type GameNFTBurningErc1155BurnedIntoGamesIterator struct {
	Event *GameNFTBurningErc1155BurnedIntoGames // Event containing the contract specifics and raw log

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
func (it *GameNFTBurningErc1155BurnedIntoGamesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTBurningErc1155BurnedIntoGames)
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
		it.Event = new(GameNFTBurningErc1155BurnedIntoGames)
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
func (it *GameNFTBurningErc1155BurnedIntoGamesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTBurningErc1155BurnedIntoGamesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTBurningErc1155BurnedIntoGames represents a Erc1155BurnedIntoGames event raised by the GameNFTBurning contract.
type GameNFTBurningErc1155BurnedIntoGames struct {
	Owner      common.Address
	NftAddress common.Address
	TokenIds   []*big.Int
	Quantities []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterErc1155BurnedIntoGames is a free log retrieval operation binding the contract event 0x4608fa48020b6d192febfc47c7b96afb63993ce6f9ded5572cd09068eb6d1178.
//
// Solidity: event Erc1155BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256[] tokenIds, uint256[] quantities)
func (_GameNFTBurning *GameNFTBurningFilterer) FilterErc1155BurnedIntoGames(opts *bind.FilterOpts, owner []common.Address, nftAddress []common.Address) (*GameNFTBurningErc1155BurnedIntoGamesIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _GameNFTBurning.contract.FilterLogs(opts, "Erc1155BurnedIntoGames", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningErc1155BurnedIntoGamesIterator{contract: _GameNFTBurning.contract, event: "Erc1155BurnedIntoGames", logs: logs, sub: sub}, nil
}

// WatchErc1155BurnedIntoGames is a free log subscription operation binding the contract event 0x4608fa48020b6d192febfc47c7b96afb63993ce6f9ded5572cd09068eb6d1178.
//
// Solidity: event Erc1155BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256[] tokenIds, uint256[] quantities)
func (_GameNFTBurning *GameNFTBurningFilterer) WatchErc1155BurnedIntoGames(opts *bind.WatchOpts, sink chan<- *GameNFTBurningErc1155BurnedIntoGames, owner []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _GameNFTBurning.contract.WatchLogs(opts, "Erc1155BurnedIntoGames", ownerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTBurningErc1155BurnedIntoGames)
				if err := _GameNFTBurning.contract.UnpackLog(event, "Erc1155BurnedIntoGames", log); err != nil {
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

// ParseErc1155BurnedIntoGames is a log parse operation binding the contract event 0x4608fa48020b6d192febfc47c7b96afb63993ce6f9ded5572cd09068eb6d1178.
//
// Solidity: event Erc1155BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256[] tokenIds, uint256[] quantities)
func (_GameNFTBurning *GameNFTBurningFilterer) ParseErc1155BurnedIntoGames(log types.Log) (*GameNFTBurningErc1155BurnedIntoGames, error) {
	event := new(GameNFTBurningErc1155BurnedIntoGames)
	if err := _GameNFTBurning.contract.UnpackLog(event, "Erc1155BurnedIntoGames", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTBurningErc721BurnedIntoGamesIterator is returned from FilterErc721BurnedIntoGames and is used to iterate over the raw logs and unpacked data for Erc721BurnedIntoGames events raised by the GameNFTBurning contract.
type GameNFTBurningErc721BurnedIntoGamesIterator struct {
	Event *GameNFTBurningErc721BurnedIntoGames // Event containing the contract specifics and raw log

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
func (it *GameNFTBurningErc721BurnedIntoGamesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTBurningErc721BurnedIntoGames)
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
		it.Event = new(GameNFTBurningErc721BurnedIntoGames)
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
func (it *GameNFTBurningErc721BurnedIntoGamesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTBurningErc721BurnedIntoGamesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTBurningErc721BurnedIntoGames represents a Erc721BurnedIntoGames event raised by the GameNFTBurning contract.
type GameNFTBurningErc721BurnedIntoGames struct {
	Owner      common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterErc721BurnedIntoGames is a free log retrieval operation binding the contract event 0xbb9f821ebd56f09e11a6397d5f4ae52dde8ddcf710e835e85f339ea7d234b011.
//
// Solidity: event Erc721BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) FilterErc721BurnedIntoGames(opts *bind.FilterOpts, owner []common.Address, nftAddress []common.Address, tokenId []*big.Int) (*GameNFTBurningErc721BurnedIntoGamesIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFTBurning.contract.FilterLogs(opts, "Erc721BurnedIntoGames", ownerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningErc721BurnedIntoGamesIterator{contract: _GameNFTBurning.contract, event: "Erc721BurnedIntoGames", logs: logs, sub: sub}, nil
}

// WatchErc721BurnedIntoGames is a free log subscription operation binding the contract event 0xbb9f821ebd56f09e11a6397d5f4ae52dde8ddcf710e835e85f339ea7d234b011.
//
// Solidity: event Erc721BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) WatchErc721BurnedIntoGames(opts *bind.WatchOpts, sink chan<- *GameNFTBurningErc721BurnedIntoGames, owner []common.Address, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFTBurning.contract.WatchLogs(opts, "Erc721BurnedIntoGames", ownerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTBurningErc721BurnedIntoGames)
				if err := _GameNFTBurning.contract.UnpackLog(event, "Erc721BurnedIntoGames", log); err != nil {
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

// ParseErc721BurnedIntoGames is a log parse operation binding the contract event 0xbb9f821ebd56f09e11a6397d5f4ae52dde8ddcf710e835e85f339ea7d234b011.
//
// Solidity: event Erc721BurnedIntoGames(address indexed owner, address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) ParseErc721BurnedIntoGames(log types.Log) (*GameNFTBurningErc721BurnedIntoGames, error) {
	event := new(GameNFTBurningErc721BurnedIntoGames)
	if err := _GameNFTBurning.contract.UnpackLog(event, "Erc721BurnedIntoGames", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNFTBurningOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameNFTBurning contract.
type GameNFTBurningOwnershipTransferredIterator struct {
	Event *GameNFTBurningOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameNFTBurningOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTBurningOwnershipTransferred)
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
		it.Event = new(GameNFTBurningOwnershipTransferred)
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
func (it *GameNFTBurningOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTBurningOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTBurningOwnershipTransferred represents a OwnershipTransferred event raised by the GameNFTBurning contract.
type GameNFTBurningOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameNFTBurning *GameNFTBurningFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameNFTBurningOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameNFTBurning.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningOwnershipTransferredIterator{contract: _GameNFTBurning.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameNFTBurning *GameNFTBurningFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameNFTBurningOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameNFTBurning.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTBurningOwnershipTransferred)
				if err := _GameNFTBurning.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameNFTBurning *GameNFTBurningFilterer) ParseOwnershipTransferred(log types.Log) (*GameNFTBurningOwnershipTransferred, error) {
	event := new(GameNFTBurningOwnershipTransferred)
	if err := _GameNFTBurning.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

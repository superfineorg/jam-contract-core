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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnedIntoGames\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnIntoGames\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6105c38061007e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063320e67c914610051578063715018a6146100665780638da5cb5b1461006e578063f2fde38b1461008d575b600080fd5b61006461005f3660046104d1565b6100a0565b005b610064610255565b600054604080516001600160a01b039092168252519081900360200190f35b61006461009b3660046104af565b61028b565b6000546001600160a01b031633146100d35760405162461bcd60e51b81526004016100ca90610558565b60405180910390fd5b60006100de83610326565b90508015610167576040516301ffc9a760e01b81526380ac58cd60e01b60048201526001600160a01b038416906301ffc9a79060240160206040518083038186803b15801561012c57600080fd5b505afa158015610140573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016491906104fb565b90505b806101c05760405162461bcd60e51b815260206004820152602360248201527f47616d654e46544275726e696e673a20696e76616c6964204e4654206164647260448201526265737360e81b60648201526084016100ca565b604051630852cd8d60e31b8152600481018390526001600160a01b038416906342966c6890602401600060405180830381600087803b15801561020257600080fd5b505af1158015610216573d6000803e3d6000fd5b50506040518492506001600160a01b03861691507f6612d30fc5da458adb3793f8d3f71cc55ec02e9de6b570040f6bf364483677c890600090a3505050565b6000546001600160a01b0316331461027f5760405162461bcd60e51b81526004016100ca90610558565b610289600061035a565b565b6000546001600160a01b031633146102b55760405162461bcd60e51b81526004016100ca90610558565b6001600160a01b03811661031a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016100ca565b6103238161035a565b50565b6000610339826301ffc9a760e01b6103aa565b80156103545750610352826001600160e01b03196103aa565b155b92915050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b179052905160009190829081906001600160a01b038716906175309061041190869061051d565b6000604051808303818686fa925050503d806000811461044d576040519150601f19603f3d011682016040523d82523d6000602084013e610452565b606091505b509150915060208151101561046d5760009350505050610354565b81801561048957508080602001905181019061048991906104fb565b9695505050505050565b80356001600160a01b03811681146104aa57600080fd5b919050565b6000602082840312156104c157600080fd5b6104ca82610493565b9392505050565b600080604083850312156104e457600080fd5b6104ed83610493565b946020939093013593505050565b60006020828403121561050d57600080fd5b815180151581146104ca57600080fd5b6000825160005b8181101561053e5760208186018101518583015201610524565b8181111561054d576000828501525b509190910192915050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea26469706673582212207cef1ba3e83593d08f9499157d837f984bf394abf8f6406b1aeb39b4903ba82e64736f6c63430008060033",
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

// BurnIntoGames is a paid mutator transaction binding the contract method 0x320e67c9.
//
// Solidity: function burnIntoGames(address nftAddress, uint256 tokenId) returns()
func (_GameNFTBurning *GameNFTBurningTransactor) BurnIntoGames(opts *bind.TransactOpts, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.contract.Transact(opts, "burnIntoGames", nftAddress, tokenId)
}

// BurnIntoGames is a paid mutator transaction binding the contract method 0x320e67c9.
//
// Solidity: function burnIntoGames(address nftAddress, uint256 tokenId) returns()
func (_GameNFTBurning *GameNFTBurningSession) BurnIntoGames(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnIntoGames(&_GameNFTBurning.TransactOpts, nftAddress, tokenId)
}

// BurnIntoGames is a paid mutator transaction binding the contract method 0x320e67c9.
//
// Solidity: function burnIntoGames(address nftAddress, uint256 tokenId) returns()
func (_GameNFTBurning *GameNFTBurningTransactorSession) BurnIntoGames(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNFTBurning.Contract.BurnIntoGames(&_GameNFTBurning.TransactOpts, nftAddress, tokenId)
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

// GameNFTBurningBurnedIntoGamesIterator is returned from FilterBurnedIntoGames and is used to iterate over the raw logs and unpacked data for BurnedIntoGames events raised by the GameNFTBurning contract.
type GameNFTBurningBurnedIntoGamesIterator struct {
	Event *GameNFTBurningBurnedIntoGames // Event containing the contract specifics and raw log

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
func (it *GameNFTBurningBurnedIntoGamesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNFTBurningBurnedIntoGames)
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
		it.Event = new(GameNFTBurningBurnedIntoGames)
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
func (it *GameNFTBurningBurnedIntoGamesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNFTBurningBurnedIntoGamesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNFTBurningBurnedIntoGames represents a BurnedIntoGames event raised by the GameNFTBurning contract.
type GameNFTBurningBurnedIntoGames struct {
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurnedIntoGames is a free log retrieval operation binding the contract event 0x6612d30fc5da458adb3793f8d3f71cc55ec02e9de6b570040f6bf364483677c8.
//
// Solidity: event BurnedIntoGames(address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) FilterBurnedIntoGames(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*GameNFTBurningBurnedIntoGamesIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFTBurning.contract.FilterLogs(opts, "BurnedIntoGames", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNFTBurningBurnedIntoGamesIterator{contract: _GameNFTBurning.contract, event: "BurnedIntoGames", logs: logs, sub: sub}, nil
}

// WatchBurnedIntoGames is a free log subscription operation binding the contract event 0x6612d30fc5da458adb3793f8d3f71cc55ec02e9de6b570040f6bf364483677c8.
//
// Solidity: event BurnedIntoGames(address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) WatchBurnedIntoGames(opts *bind.WatchOpts, sink chan<- *GameNFTBurningBurnedIntoGames, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GameNFTBurning.contract.WatchLogs(opts, "BurnedIntoGames", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNFTBurningBurnedIntoGames)
				if err := _GameNFTBurning.contract.UnpackLog(event, "BurnedIntoGames", log); err != nil {
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

// ParseBurnedIntoGames is a log parse operation binding the contract event 0x6612d30fc5da458adb3793f8d3f71cc55ec02e9de6b570040f6bf364483677c8.
//
// Solidity: event BurnedIntoGames(address indexed nftAddress, uint256 indexed tokenId)
func (_GameNFTBurning *GameNFTBurningFilterer) ParseBurnedIntoGames(log types.Log) (*GameNFTBurningBurnedIntoGames, error) {
	event := new(GameNFTBurningBurnedIntoGames)
	if err := _GameNFTBurning.contract.UnpackLog(event, "BurnedIntoGames", log); err != nil {
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

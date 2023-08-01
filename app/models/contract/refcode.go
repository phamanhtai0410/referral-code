// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// RefCodeUser is an auto generated low-level Go binding around an user-defined struct.
type RefCodeUser struct {
	UserAddress common.Address
	Domain      string
	TotalEarn   *big.Int
}

// RefCodeMetaData contains all meta data concerning the RefCode contract.
var RefCodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalEarn\",\"type\":\"uint256\"}],\"internalType\":\"structRefCode.User[]\",\"name\":\"users\",\"type\":\"tuple[]\"}],\"name\":\"uploadData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalEarn\",\"type\":\"uint256\"}],\"name\":\"uploadSingleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RefCodeABI is the input ABI used to generate the binding from.
// Deprecated: Use RefCodeMetaData.ABI instead.
var RefCodeABI = RefCodeMetaData.ABI

// RefCode is an auto generated Go binding around an Ethereum contract.
type RefCode struct {
	RefCodeCaller     // Read-only binding to the contract
	RefCodeTransactor // Write-only binding to the contract
	RefCodeFilterer   // Log filterer for contract events
}

// RefCodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefCodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefCodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefCodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefCodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefCodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefCodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefCodeSession struct {
	Contract     *RefCode          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefCodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefCodeCallerSession struct {
	Contract *RefCodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RefCodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefCodeTransactorSession struct {
	Contract     *RefCodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RefCodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefCodeRaw struct {
	Contract *RefCode // Generic contract binding to access the raw methods on
}

// RefCodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefCodeCallerRaw struct {
	Contract *RefCodeCaller // Generic read-only contract binding to access the raw methods on
}

// RefCodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefCodeTransactorRaw struct {
	Contract *RefCodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefCode creates a new instance of RefCode, bound to a specific deployed contract.
func NewRefCode(address common.Address, backend bind.ContractBackend) (*RefCode, error) {
	contract, err := bindRefCode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefCode{RefCodeCaller: RefCodeCaller{contract: contract}, RefCodeTransactor: RefCodeTransactor{contract: contract}, RefCodeFilterer: RefCodeFilterer{contract: contract}}, nil
}

// NewRefCodeCaller creates a new read-only instance of RefCode, bound to a specific deployed contract.
func NewRefCodeCaller(address common.Address, caller bind.ContractCaller) (*RefCodeCaller, error) {
	contract, err := bindRefCode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefCodeCaller{contract: contract}, nil
}

// NewRefCodeTransactor creates a new write-only instance of RefCode, bound to a specific deployed contract.
func NewRefCodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RefCodeTransactor, error) {
	contract, err := bindRefCode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefCodeTransactor{contract: contract}, nil
}

// NewRefCodeFilterer creates a new log filterer instance of RefCode, bound to a specific deployed contract.
func NewRefCodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RefCodeFilterer, error) {
	contract, err := bindRefCode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefCodeFilterer{contract: contract}, nil
}

// bindRefCode binds a generic wrapper to an already deployed contract.
func bindRefCode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RefCodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefCode *RefCodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefCode.Contract.RefCodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefCode *RefCodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefCode.Contract.RefCodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefCode *RefCodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefCode.Contract.RefCodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefCode *RefCodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RefCode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefCode *RefCodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefCode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefCode *RefCodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefCode.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefCode *RefCodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RefCode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefCode *RefCodeSession) Owner() (common.Address, error) {
	return _RefCode.Contract.Owner(&_RefCode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RefCode *RefCodeCallerSession) Owner() (common.Address, error) {
	return _RefCode.Contract.Owner(&_RefCode.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefCode *RefCodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefCode *RefCodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefCode.Contract.RenounceOwnership(&_RefCode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RefCode *RefCodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RefCode.Contract.RenounceOwnership(&_RefCode.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefCode *RefCodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefCode *RefCodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefCode.Contract.TransferOwnership(&_RefCode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RefCode *RefCodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefCode.Contract.TransferOwnership(&_RefCode.TransactOpts, newOwner)
}

// UploadData is a paid mutator transaction binding the contract method 0x96dfa50f.
//
// Solidity: function uploadData((address,string,uint256)[] users) returns()
func (_RefCode *RefCodeTransactor) UploadData(opts *bind.TransactOpts, users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "uploadData", users)
}

// UploadData is a paid mutator transaction binding the contract method 0x96dfa50f.
//
// Solidity: function uploadData((address,string,uint256)[] users) returns()
func (_RefCode *RefCodeSession) UploadData(users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.Contract.UploadData(&_RefCode.TransactOpts, users)
}

// UploadData is a paid mutator transaction binding the contract method 0x96dfa50f.
//
// Solidity: function uploadData((address,string,uint256)[] users) returns()
func (_RefCode *RefCodeTransactorSession) UploadData(users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.Contract.UploadData(&_RefCode.TransactOpts, users)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0xb6dda933.
//
// Solidity: function uploadSingleData(address addr, string domain, uint256 totalEarn) returns()
func (_RefCode *RefCodeTransactor) UploadSingleData(opts *bind.TransactOpts, addr common.Address, domain string, totalEarn *big.Int) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "uploadSingleData", addr, domain, totalEarn)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0xb6dda933.
//
// Solidity: function uploadSingleData(address addr, string domain, uint256 totalEarn) returns()
func (_RefCode *RefCodeSession) UploadSingleData(addr common.Address, domain string, totalEarn *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.UploadSingleData(&_RefCode.TransactOpts, addr, domain, totalEarn)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0xb6dda933.
//
// Solidity: function uploadSingleData(address addr, string domain, uint256 totalEarn) returns()
func (_RefCode *RefCodeTransactorSession) UploadSingleData(addr common.Address, domain string, totalEarn *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.UploadSingleData(&_RefCode.TransactOpts, addr, domain, totalEarn)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string domain) returns()
func (_RefCode *RefCodeTransactor) Withdraw(opts *bind.TransactOpts, domain string) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "withdraw", domain)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string domain) returns()
func (_RefCode *RefCodeSession) Withdraw(domain string) (*types.Transaction, error) {
	return _RefCode.Contract.Withdraw(&_RefCode.TransactOpts, domain)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string domain) returns()
func (_RefCode *RefCodeTransactorSession) Withdraw(domain string) (*types.Transaction, error) {
	return _RefCode.Contract.Withdraw(&_RefCode.TransactOpts, domain)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefCode *RefCodeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefCode.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefCode *RefCodeSession) Receive() (*types.Transaction, error) {
	return _RefCode.Contract.Receive(&_RefCode.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RefCode *RefCodeTransactorSession) Receive() (*types.Transaction, error) {
	return _RefCode.Contract.Receive(&_RefCode.TransactOpts)
}

// RefCodeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RefCode contract.
type RefCodeDepositedIterator struct {
	Event *RefCodeDeposited // Event containing the contract specifics and raw log

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
func (it *RefCodeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefCodeDeposited)
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
		it.Event = new(RefCodeDeposited)
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
func (it *RefCodeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefCodeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefCodeDeposited represents a Deposited event raised by the RefCode contract.
type RefCodeDeposited struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address account, uint256 amount)
func (_RefCode *RefCodeFilterer) FilterDeposited(opts *bind.FilterOpts) (*RefCodeDepositedIterator, error) {

	logs, sub, err := _RefCode.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &RefCodeDepositedIterator{contract: _RefCode.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address account, uint256 amount)
func (_RefCode *RefCodeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RefCodeDeposited) (event.Subscription, error) {

	logs, sub, err := _RefCode.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefCodeDeposited)
				if err := _RefCode.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address account, uint256 amount)
func (_RefCode *RefCodeFilterer) ParseDeposited(log types.Log) (*RefCodeDeposited, error) {
	event := new(RefCodeDeposited)
	if err := _RefCode.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefCodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RefCode contract.
type RefCodeOwnershipTransferredIterator struct {
	Event *RefCodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefCodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefCodeOwnershipTransferred)
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
		it.Event = new(RefCodeOwnershipTransferred)
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
func (it *RefCodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefCodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefCodeOwnershipTransferred represents a OwnershipTransferred event raised by the RefCode contract.
type RefCodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefCode *RefCodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefCodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefCode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefCodeOwnershipTransferredIterator{contract: _RefCode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RefCode *RefCodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefCodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefCode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefCodeOwnershipTransferred)
				if err := _RefCode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RefCode *RefCodeFilterer) ParseOwnershipTransferred(log types.Log) (*RefCodeOwnershipTransferred, error) {
	event := new(RefCodeOwnershipTransferred)
	if err := _RefCode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RefCodeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RefCode contract.
type RefCodeWithdrawnIterator struct {
	Event *RefCodeWithdrawn // Event containing the contract specifics and raw log

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
func (it *RefCodeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefCodeWithdrawn)
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
		it.Event = new(RefCodeWithdrawn)
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
func (it *RefCodeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefCodeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefCodeWithdrawn represents a Withdrawn event raised by the RefCode contract.
type RefCodeWithdrawn struct {
	Account common.Address
	Domain  string
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x18af30a54a1951aeee806c16cdcb8d087ed1a095f0d9a87261b92cef2c6cc92d.
//
// Solidity: event Withdrawn(address account, string domain, uint256 amount)
func (_RefCode *RefCodeFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*RefCodeWithdrawnIterator, error) {

	logs, sub, err := _RefCode.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &RefCodeWithdrawnIterator{contract: _RefCode.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x18af30a54a1951aeee806c16cdcb8d087ed1a095f0d9a87261b92cef2c6cc92d.
//
// Solidity: event Withdrawn(address account, string domain, uint256 amount)
func (_RefCode *RefCodeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RefCodeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _RefCode.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefCodeWithdrawn)
				if err := _RefCode.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x18af30a54a1951aeee806c16cdcb8d087ed1a095f0d9a87261b92cef2c6cc92d.
//
// Solidity: event Withdrawn(address account, string domain, uint256 amount)
func (_RefCode *RefCodeFilterer) ParseWithdrawn(log types.Log) (*RefCodeWithdrawn, error) {
	event := new(RefCodeWithdrawn)
	if err := _RefCode.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
	TotalSupply *big.Int
}

// RefCodeMetaData contains all meta data concerning the RefCode contract.
var RefCodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"internalType\":\"structRefCode.User[]\",\"name\":\"users\",\"type\":\"tuple[]\"}],\"name\":\"uploadData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"uploadSingleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// UploadData is a paid mutator transaction binding the contract method 0xaa21ea6d.
//
// Solidity: function uploadData((address,uint256)[] users) returns()
func (_RefCode *RefCodeTransactor) UploadData(opts *bind.TransactOpts, users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "uploadData", users)
}

// UploadData is a paid mutator transaction binding the contract method 0xaa21ea6d.
//
// Solidity: function uploadData((address,uint256)[] users) returns()
func (_RefCode *RefCodeSession) UploadData(users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.Contract.UploadData(&_RefCode.TransactOpts, users)
}

// UploadData is a paid mutator transaction binding the contract method 0xaa21ea6d.
//
// Solidity: function uploadData((address,uint256)[] users) returns()
func (_RefCode *RefCodeTransactorSession) UploadData(users []RefCodeUser) (*types.Transaction, error) {
	return _RefCode.Contract.UploadData(&_RefCode.TransactOpts, users)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0x5a0398e0.
//
// Solidity: function uploadSingleData(address addr, uint256 totalSupply) returns()
func (_RefCode *RefCodeTransactor) UploadSingleData(opts *bind.TransactOpts, addr common.Address, totalSupply *big.Int) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "uploadSingleData", addr, totalSupply)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0x5a0398e0.
//
// Solidity: function uploadSingleData(address addr, uint256 totalSupply) returns()
func (_RefCode *RefCodeSession) UploadSingleData(addr common.Address, totalSupply *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.UploadSingleData(&_RefCode.TransactOpts, addr, totalSupply)
}

// UploadSingleData is a paid mutator transaction binding the contract method 0x5a0398e0.
//
// Solidity: function uploadSingleData(address addr, uint256 totalSupply) returns()
func (_RefCode *RefCodeTransactorSession) UploadSingleData(addr common.Address, totalSupply *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.UploadSingleData(&_RefCode.TransactOpts, addr, totalSupply)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RefCode *RefCodeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _RefCode.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RefCode *RefCodeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.Withdraw(&_RefCode.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RefCode *RefCodeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _RefCode.Contract.Withdraw(&_RefCode.TransactOpts, amount)
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
	Account   common.Address
	Amount    *big.Int
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address account, uint256 amount, uint256 remaining)
func (_RefCode *RefCodeFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*RefCodeWithdrawnIterator, error) {

	logs, sub, err := _RefCode.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &RefCodeWithdrawnIterator{contract: _RefCode.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address account, uint256 amount, uint256 remaining)
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address account, uint256 amount, uint256 remaining)
func (_RefCode *RefCodeFilterer) ParseWithdrawn(log types.Log) (*RefCodeWithdrawn, error) {
	event := new(RefCodeWithdrawn)
	if err := _RefCode.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

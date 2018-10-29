// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BirrContractABI is the input ABI used to generate the binding from.
const BirrContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"storedData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BirrContractBin is the compiled bytecode used for deploying new contracts.
const BirrContractBin = `0x608060405234801561001057600080fd5b506040805180820190915260048082527f696e69740000000000000000000000000000000000000000000000000000000060209092019182526100559160009161005b565b506100f6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009c57805160ff19168380011785556100c9565b828001600101855582156100c9579182015b828111156100c95782518255916020019190600101906100ae565b506100d59291506100d9565b5090565b6100f391905b808211156100d557600081556001016100df565b90565b610294806101056000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632a1afcd981146100505780634ed3885e146100da575b600080fd5b34801561005c57600080fd5b50610065610128565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009f578181015183820152602001610087565b50505050905090810190601f1680156100cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805160206004803580820135601f81018490048402850184019095528484526101269436949293602493928401919081908401838280828437509497506101b69650505050505050565b005b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101ae5780601f10610183576101008083540402835291602001916101ae565b820191906000526020600020905b81548152906001019060200180831161019157829003601f168201915b505050505081565b80516101c99060009060208401906101cd565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061020e57805160ff191683800117855561023b565b8280016001018555821561023b579182015b8281111561023b578251825591602001919060010190610220565b5061024792915061024b565b5090565b61026591905b808211156102475760008155600101610251565b905600a165627a7a723058201b03c430a53848d38867a4c4756adf6b34e08dac5a08dd177bf93815b2e1b75b0029`

// DeployBirrContract deploys a new Ethereum contract, binding an instance of BirrContract to it.
func DeployBirrContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BirrContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BirrContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BirrContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BirrContract{BirrContractCaller: BirrContractCaller{contract: contract}, BirrContractTransactor: BirrContractTransactor{contract: contract}, BirrContractFilterer: BirrContractFilterer{contract: contract}}, nil
}

// BirrContract is an auto generated Go binding around an Ethereum contract.
type BirrContract struct {
	BirrContractCaller     // Read-only binding to the contract
	BirrContractTransactor // Write-only binding to the contract
	BirrContractFilterer   // Log filterer for contract events
}

// BirrContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BirrContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirrContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BirrContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirrContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BirrContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BirrContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BirrContractSession struct {
	Contract     *BirrContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BirrContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BirrContractCallerSession struct {
	Contract *BirrContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BirrContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BirrContractTransactorSession struct {
	Contract     *BirrContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BirrContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BirrContractRaw struct {
	Contract *BirrContract // Generic contract binding to access the raw methods on
}

// BirrContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BirrContractCallerRaw struct {
	Contract *BirrContractCaller // Generic read-only contract binding to access the raw methods on
}

// BirrContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BirrContractTransactorRaw struct {
	Contract *BirrContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBirrContract creates a new instance of BirrContract, bound to a specific deployed contract.
func NewBirrContract(address common.Address, backend bind.ContractBackend) (*BirrContract, error) {
	contract, err := bindBirrContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BirrContract{BirrContractCaller: BirrContractCaller{contract: contract}, BirrContractTransactor: BirrContractTransactor{contract: contract}, BirrContractFilterer: BirrContractFilterer{contract: contract}}, nil
}

// NewBirrContractCaller creates a new read-only instance of BirrContract, bound to a specific deployed contract.
func NewBirrContractCaller(address common.Address, caller bind.ContractCaller) (*BirrContractCaller, error) {
	contract, err := bindBirrContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BirrContractCaller{contract: contract}, nil
}

// NewBirrContractTransactor creates a new write-only instance of BirrContract, bound to a specific deployed contract.
func NewBirrContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BirrContractTransactor, error) {
	contract, err := bindBirrContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BirrContractTransactor{contract: contract}, nil
}

// NewBirrContractFilterer creates a new log filterer instance of BirrContract, bound to a specific deployed contract.
func NewBirrContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BirrContractFilterer, error) {
	contract, err := bindBirrContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BirrContractFilterer{contract: contract}, nil
}

// bindBirrContract binds a generic wrapper to an already deployed contract.
func bindBirrContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BirrContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BirrContract *BirrContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BirrContract.Contract.BirrContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BirrContract *BirrContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BirrContract.Contract.BirrContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BirrContract *BirrContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BirrContract.Contract.BirrContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BirrContract *BirrContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BirrContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BirrContract *BirrContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BirrContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BirrContract *BirrContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BirrContract.Contract.contract.Transact(opts, method, params...)
}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() constant returns(string)
func (_BirrContract *BirrContractCaller) StoredData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "storedData")
	return *ret0, err
}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() constant returns(string)
func (_BirrContract *BirrContractSession) StoredData() (string, error) {
	return _BirrContract.Contract.StoredData(&_BirrContract.CallOpts)
}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() constant returns(string)
func (_BirrContract *BirrContractCallerSession) StoredData() (string, error) {
	return _BirrContract.Contract.StoredData(&_BirrContract.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(x string) returns()
func (_BirrContract *BirrContractTransactor) Set(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(x string) returns()
func (_BirrContract *BirrContractSession) Set(x string) (*types.Transaction, error) {
	return _BirrContract.Contract.Set(&_BirrContract.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(x string) returns()
func (_BirrContract *BirrContractTransactorSession) Set(x string) (*types.Transaction, error) {
	return _BirrContract.Contract.Set(&_BirrContract.TransactOpts, x)
}

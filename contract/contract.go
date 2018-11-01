// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BirrContractABI is the input ABI used to generate the binding from.
const BirrContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getObjects\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"objects\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"asNumber\",\"type\":\"uint32\"},{\"name\":\"hash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHowManyOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"asNumber\",\"type\":\"uint32\"},{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BirrContractBin is the compiled bytecode used for deploying new contracts.
const BirrContractBin = `0x608060405234801561001057600080fd5b506040516105e03803806105e08339810160409081528151602080840151336000818152808452949094208054600160a060020a03191690941760a060020a60c060020a0319167401000000000000000000000000000000000000000063ffffffff851602178455909301805191939092610093926001909101918401906100dc565b50506001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018054600160a060020a0319163317905550610177565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011d57805160ff191683800117855561014a565b8280016001018555821561014a579182015b8281111561014a57825182559160200191906001019061012f565b5061015692915061015a565b5090565b61017491905b808211156101565760008155600101610160565b90565b61045a806101866000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630f310b1781146100715780631665ff7c146100a5578063621b23e214610150578063770285ae1461016e578063a2b71511146101f5575b600080fd5b34801561007d57600080fd5b5061008960043561021c565b60408051600160a060020a039092168252519081900360200190f35b3480156100b157600080fd5b506100ba610244565b6040805163ffffffff85168152600160a060020a03831691810191909152606060208083018281528551928401929092528451608084019186019080838360005b838110156101135781810151838201526020016100fb565b50505050905090810190601f1680156101405780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561015c57600080fd5b5061008963ffffffff60043516610325565b34801561017a57600080fd5b5061018f600160a060020a0360043516610357565b6040518084600160a060020a0316600160a060020a031681526020018363ffffffff1663ffffffff1681526020018060200182810382528381815181526020019150805190602001908083836000838110156101135781810151838201526020016100fb565b34801561020157600080fd5b5061020a610428565b60408051918252519081900360200190f35b600180548290811061022a57fe5b600091825260209091200154600160a060020a0316905081565b336000818152602081815260408083208054600191820180548451600261010095831615959095026000190190911693909304601f810186900486028401860190945283835294956060958795919474010000000000000000000000000000000000000000840463ffffffff16949193600160a060020a0316928491908301828280156103125780601f106102e757610100808354040283529160200191610312565b820191906000526020600020905b8154815290600101906020018083116102f557829003601f168201915b5050505050915093509350935050909192565b600060018263ffffffff1681548110151561033c57fe5b600091825260209091200154600160a060020a031692915050565b600060208181529181526040908190208054600180830180548551600261010094831615949094026000190190911692909204601f8101879004870283018701909552848252600160a060020a038316957401000000000000000000000000000000000000000090930463ffffffff1694919290919083018282801561041e5780601f106103f35761010080835404028352916020019161041e565b820191906000526020600020905b81548152906001019060200180831161040157829003601f168201915b5050505050905083565b600154905600a165627a7a7230582076ca6ac850cb713dcd34192e8e5b1cf1fb2317052e8b134ea46886cac3112df40029`

// DeployBirrContract deploys a new Ethereum contract, binding an instance of BirrContract to it.
func DeployBirrContract(auth *bind.TransactOpts, backend bind.ContractBackend, asNumber uint32, ipfsHash string) (common.Address, *types.Transaction, *BirrContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BirrContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BirrContractBin), backend, asNumber, ipfsHash)
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

// GetHowManyOwners is a free data retrieval call binding the contract method 0xa2b71511.
//
// Solidity: function getHowManyOwners() constant returns(uint256)
func (_BirrContract *BirrContractCaller) GetHowManyOwners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "getHowManyOwners")
	return *ret0, err
}

// GetHowManyOwners is a free data retrieval call binding the contract method 0xa2b71511.
//
// Solidity: function getHowManyOwners() constant returns(uint256)
func (_BirrContract *BirrContractSession) GetHowManyOwners() (*big.Int, error) {
	return _BirrContract.Contract.GetHowManyOwners(&_BirrContract.CallOpts)
}

// GetHowManyOwners is a free data retrieval call binding the contract method 0xa2b71511.
//
// Solidity: function getHowManyOwners() constant returns(uint256)
func (_BirrContract *BirrContractCallerSession) GetHowManyOwners() (*big.Int, error) {
	return _BirrContract.Contract.GetHowManyOwners(&_BirrContract.CallOpts)
}

// GetObjects is a free data retrieval call binding the contract method 0x1665ff7c.
//
// Solidity: function getObjects() constant returns(uint32, string, address)
func (_BirrContract *BirrContractCaller) GetObjects(opts *bind.CallOpts) (uint32, string, common.Address, error) {
	var (
		ret0 = new(uint32)
		ret1 = new(string)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _BirrContract.contract.Call(opts, out, "getObjects")
	return *ret0, *ret1, *ret2, err
}

// GetObjects is a free data retrieval call binding the contract method 0x1665ff7c.
//
// Solidity: function getObjects() constant returns(uint32, string, address)
func (_BirrContract *BirrContractSession) GetObjects() (uint32, string, common.Address, error) {
	return _BirrContract.Contract.GetObjects(&_BirrContract.CallOpts)
}

// GetObjects is a free data retrieval call binding the contract method 0x1665ff7c.
//
// Solidity: function getObjects() constant returns(uint32, string, address)
func (_BirrContract *BirrContractCallerSession) GetObjects() (uint32, string, common.Address, error) {
	return _BirrContract.Contract.GetObjects(&_BirrContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(i uint32) constant returns(address)
func (_BirrContract *BirrContractCaller) GetOwner(opts *bind.CallOpts, i uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "getOwner", i)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(i uint32) constant returns(address)
func (_BirrContract *BirrContractSession) GetOwner(i uint32) (common.Address, error) {
	return _BirrContract.Contract.GetOwner(&_BirrContract.CallOpts, i)
}

// GetOwner is a free data retrieval call binding the contract method 0x621b23e2.
//
// Solidity: function getOwner(i uint32) constant returns(address)
func (_BirrContract *BirrContractCallerSession) GetOwner(i uint32) (common.Address, error) {
	return _BirrContract.Contract.GetOwner(&_BirrContract.CallOpts, i)
}

// Objects is a free data retrieval call binding the contract method 0x770285ae.
//
// Solidity: function objects( address) constant returns(owner address, asNumber uint32, hash string)
func (_BirrContract *BirrContractCaller) Objects(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner    common.Address
	AsNumber uint32
	Hash     string
}, error) {
	ret := new(struct {
		Owner    common.Address
		AsNumber uint32
		Hash     string
	})
	out := ret
	err := _BirrContract.contract.Call(opts, out, "objects", arg0)
	return *ret, err
}

// Objects is a free data retrieval call binding the contract method 0x770285ae.
//
// Solidity: function objects( address) constant returns(owner address, asNumber uint32, hash string)
func (_BirrContract *BirrContractSession) Objects(arg0 common.Address) (struct {
	Owner    common.Address
	AsNumber uint32
	Hash     string
}, error) {
	return _BirrContract.Contract.Objects(&_BirrContract.CallOpts, arg0)
}

// Objects is a free data retrieval call binding the contract method 0x770285ae.
//
// Solidity: function objects( address) constant returns(owner address, asNumber uint32, hash string)
func (_BirrContract *BirrContractCallerSession) Objects(arg0 common.Address) (struct {
	Owner    common.Address
	AsNumber uint32
	Hash     string
}, error) {
	return _BirrContract.Contract.Objects(&_BirrContract.CallOpts, arg0)
}

// OwnerAddresses is a free data retrieval call binding the contract method 0x0f310b17.
//
// Solidity: function ownerAddresses( uint256) constant returns(address)
func (_BirrContract *BirrContractCaller) OwnerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "ownerAddresses", arg0)
	return *ret0, err
}

// OwnerAddresses is a free data retrieval call binding the contract method 0x0f310b17.
//
// Solidity: function ownerAddresses( uint256) constant returns(address)
func (_BirrContract *BirrContractSession) OwnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _BirrContract.Contract.OwnerAddresses(&_BirrContract.CallOpts, arg0)
}

// OwnerAddresses is a free data retrieval call binding the contract method 0x0f310b17.
//
// Solidity: function ownerAddresses( uint256) constant returns(address)
func (_BirrContract *BirrContractCallerSession) OwnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _BirrContract.Contract.OwnerAddresses(&_BirrContract.CallOpts, arg0)
}

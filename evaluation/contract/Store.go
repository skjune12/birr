// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAutNum\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getRoute\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAutNum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute6\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeRoute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsRouteKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAsSet\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouteKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAsSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRoute6\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint8\"}],\"name\":\"ItemSet\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `0x608060405234801561001057600080fd5b50604051610d77380380610d778339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815164010000000081118282018710171561007857600080fd5b5050805190935061009292506001915060208401906100ab565b505060008054600160a060020a03191633179055610146565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ec57805160ff1916838001178555610119565b82800160010185558215610119579182015b828111156101195782518255916020019190600101906100fe565b50610125929150610129565b5090565b61014391905b80821115610125576000815560010161012f565b90565b610c22806101556000396000f3fe6080604052600436106100c9577c01000000000000000000000000000000000000000000000000000000006000350463084fa70281146100ce5780630b014d8414610123578063107997551461015c57806341c0e1b51461019957806347fb6379146101ae5780634d0fdcae146101e95780635129efce1461021357806354fd4d50146102515780636c28e70c146102db5780638819bb561461030e5780638da5cb5b1461034f578063a15b186714610380578063bbc9a14f146103e5578063c8b8089514610420575b600080fd5b3480156100da57600080fd5b50610101600480360360208110156100f157600080fd5b5035600160a060020a0316610453565b6040805193845260ff9283166020850152911682820152519081900360600190f35b34801561012f57600080fd5b506101016004803603604081101561014657600080fd5b50600160a060020a0381351690602001356104b2565b34801561016857600080fd5b506101976004803603606081101561017f57600080fd5b5080359060ff6020820135811691604001351661051d565b005b3480156101a557600080fd5b506101976105c3565b3480156101ba57600080fd5b50610197600480360360608110156101d157600080fd5b5080359060ff602082013581169160400135166105e8565b3480156101f557600080fd5b506101976004803603602081101561020c57600080fd5b503561068e565b34801561021f57600080fd5b5061023d6004803603602081101561023657600080fd5b503561069b565b604080519115158252519081900360200190f35b34801561025d57600080fd5b506102666106bf565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102a0578181015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102e757600080fd5b50610101600480360360208110156102fe57600080fd5b5035600160a060020a031661074c565b34801561031a57600080fd5b506101976004803603608081101561033157600080fd5b5080359060208101359060ff604082013581169160600135166107ab565b34801561035b57600080fd5b50610364610887565b60408051600160a060020a039092168252519081900360200190f35b34801561038c57600080fd5b50610395610896565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103d15781810151838201526020016103b9565b505050509050019250505060405180910390f35b3480156103f157600080fd5b506101976004803603606081101561040857600080fd5b5080359060ff602082013581169160400135166108fb565b34801561042c57600080fd5b506101016004803603602081101561044357600080fd5b5035600160a060020a03166109a1565b6000806000610460610b79565b50505050600160a060020a03166000908152600260209081526040918290208251606081018452600582015480825260069092015460ff80821694830185905261010090910416930183905292909190565b60008060006104bf610b79565b505050600160a060020a03831660009081526002602090815260408083208584526009018252918290208251606081018452815480825260019092015460ff8082169483018590526101009091041693018390529250909250925092565b610525610b79565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a902089516005820155935160069094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a815291820192909252808601929092529351929392600080516020610bd7833981519152929181900390910190a250505050565b600054600160a060020a031633146105da57600080fd5b600054600160a060020a0316ff5b6105f0610b79565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a902089516003820155935160049094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a815291820192909252808601929092529351929392600080516020610bd7833981519152929181900390910190a250505050565b6106983382610a00565b50565b336000908152600260209081526040808320938352600a9093019052205460ff1690565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107445780601f1061071957610100808354040283529160200191610744565b820191906000526020600020905b81548152906001019060200180831161072757829003601f168201915b505050505081565b6000806000610759610b79565b50505050600160a060020a03166000908152600260209081526040918290208251606081018452600782015480825260089092015460ff80821694830185905261010090910416930183905292909190565b6107b3610b79565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a81208f82526009810187528b82208b51815595516001968701805495518b166101000261ff001992909b1660ff19968716179190911699909917909855600b880180548087018255908252868220018f90558e8152600a909701855295899020805490911690921790915586518a81529182019290925280860191909152935192939092600080516020610bd7833981519152929181900390910190a25050505050565b600054600160a060020a031681565b33600090815260026020908152604091829020600b018054835181840281018401909452808452606093928301828280156108f057602002820191906000526020600020905b8154815260200190600101908083116108dc575b505050505090505b90565b610903610b79565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a902089516007820155935160089094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a815291820192909252808601929092529351929392600080516020610bd7833981519152929181900390910190a250505050565b60008060006109ae610b79565b50505050600160a060020a03166000908152600260209081526040918290208251606081018452600382015480825260049092015460ff80821694830185905261010090910416930183905292909190565b6000610a0c8383610a63565b9050610a188382610ab0565b50600160a060020a03919091166000908152600260209081526040808320938352600984018252808320838155600101805461ffff19169055600a909301905220805460ff19169055565b60005b600160a060020a0383166000908152600260205260409020600b01805483919083908110610a9057fe5b9060005260206000200154141515610aaa57600101610a66565b92915050565b600160a060020a0382166000908152600260205260409020600b015460001901811015610b4857600160a060020a0382166000908152600260205260409020600b01805460018301908110610b0157fe5b6000918252602080832090910154600160a060020a038516835260029091526040909120600b01805483908110610b3457fe5b600091825260209091200155600101610ab0565b600160a060020a0382166000908152600260205260409020600b01805490610b74906000198301610b99565b505050565b604080516060810182526000808252602082018190529181019190915290565b815481835581811115610b7457600083815260209020610b749181019083016108f891905b80821115610bd25760008155600101610bbe565b509056fe7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2da165627a7a72305820ac498eeccffb8c577bf31f76994bb47d2be603f2970558a3def4b0894a7714870029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_Store *StoreCaller) ExistsRouteKey(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "existsRouteKey", key)
	return *ret0, err
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_Store *StoreSession) ExistsRouteKey(key [32]byte) (bool, error) {
	return _Store.Contract.ExistsRouteKey(&_Store.CallOpts, key)
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_Store *StoreCallerSession) ExistsRouteKey(key [32]byte) (bool, error) {
	return _Store.Contract.ExistsRouteKey(&_Store.CallOpts, key)
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) GetAsSet(opts *bind.CallOpts, addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getAsSet", addr)
	return *ret, err
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) GetAsSet(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetAsSet(&_Store.CallOpts, addr)
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) GetAsSet(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetAsSet(&_Store.CallOpts, addr)
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) GetAutNum(opts *bind.CallOpts, addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getAutNum", addr)
	return *ret, err
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) GetAutNum(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetAutNum(&_Store.CallOpts, addr)
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) GetAutNum(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetAutNum(&_Store.CallOpts, addr)
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) GetRoute(opts *bind.CallOpts, addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getRoute", addr, key)
	return *ret, err
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) GetRoute(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute(&_Store.CallOpts, addr, key)
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) GetRoute(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute(&_Store.CallOpts, addr, key)
}

// GetRoute6 is a free data retrieval call binding the contract method 0xc8b80895.
//
// Solidity: function getRoute6(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) GetRoute6(opts *bind.CallOpts, addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	ret := new(struct {
		Digest       [32]byte
		HashFunction uint8
		Size         uint8
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getRoute6", addr)
	return *ret, err
}

// GetRoute6 is a free data retrieval call binding the contract method 0xc8b80895.
//
// Solidity: function getRoute6(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) GetRoute6(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute6(&_Store.CallOpts, addr)
}

// GetRoute6 is a free data retrieval call binding the contract method 0xc8b80895.
//
// Solidity: function getRoute6(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) GetRoute6(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute6(&_Store.CallOpts, addr)
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_Store *StoreCaller) GetRouteKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getRouteKeys")
	return *ret0, err
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_Store *StoreSession) GetRouteKeys() ([][32]byte, error) {
	return _Store.Contract.GetRouteKeys(&_Store.CallOpts)
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_Store *StoreCallerSession) GetRouteKeys() ([][32]byte, error) {
	return _Store.Contract.GetRouteKeys(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Store *StoreTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Store *StoreSession) Kill() (*types.Transaction, error) {
	return _Store.Contract.Kill(&_Store.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Store *StoreTransactorSession) Kill() (*types.Transaction, error) {
	return _Store.Contract.Kill(&_Store.TransactOpts)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_Store *StoreTransactor) RemoveRoute(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "removeRoute", key)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_Store *StoreSession) RemoveRoute(key [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RemoveRoute(&_Store.TransactOpts, key)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_Store *StoreTransactorSession) RemoveRoute(key [32]byte) (*types.Transaction, error) {
	return _Store.Contract.RemoveRoute(&_Store.TransactOpts, key)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactor) SetAsSet(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setAsSet", _digest, _hashFunction, _size)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreSession) SetAsSet(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetAsSet(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactorSession) SetAsSet(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetAsSet(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactor) SetAutNum(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setAutNum", _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreSession) SetAutNum(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetAutNum(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactorSession) SetAutNum(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetAutNum(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactor) SetRoute(opts *bind.TransactOpts, key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setRoute", key, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreSession) SetRoute(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute(&_Store.TransactOpts, key, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactorSession) SetRoute(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute(&_Store.TransactOpts, key, _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x47fb6379.
//
// Solidity: function setRoute6(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactor) SetRoute6(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setRoute6", _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x47fb6379.
//
// Solidity: function setRoute6(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreSession) SetRoute6(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute6(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x47fb6379.
//
// Solidity: function setRoute6(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactorSession) SetRoute6(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute6(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// StoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store contract.
type StoreItemSetIterator struct {
	Event *StoreItemSet // Event containing the contract specifics and raw log

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
func (it *StoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemSet)
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
		it.Event = new(StoreItemSet)
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
func (it *StoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemSet represents a ItemSet event raised by the Store contract.
type StoreItemSet struct {
	Key          common.Address
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2d.
//
// Solidity: e ItemSet(key indexed address, digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreFilterer) FilterItemSet(opts *bind.FilterOpts, key []common.Address) (*StoreItemSetIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemSet", keyRule)
	if err != nil {
		return nil, err
	}
	return &StoreItemSetIterator{contract: _Store.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2d.
//
// Solidity: e ItemSet(key indexed address, digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StoreItemSet, key []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemSet", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemSet)
				if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

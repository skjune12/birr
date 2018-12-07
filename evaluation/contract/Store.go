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
const StoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAutNum\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAutNum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute6\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRoute\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAsSet\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAsSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRoute6\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"items\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint8\"}],\"name\":\"ItemSet\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `0x608060405234801561001057600080fd5b50604051610a59380380610a598339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815164010000000081118282018710171561007857600080fd5b5050805190935061009292506001915060208401906100ab565b505060008054600160a060020a03191633179055610146565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ec57805160ff1916838001178555610119565b82800160010185558215610119579182015b828111156101195782518255916020019190600101906100fe565b50610125929150610129565b5090565b61014391905b80821115610125576000815560010161012f565b90565b610904806101556000396000f3fe6080604052600436106100b95763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663084fa70281146100be578063107997551461011357806341c0e1b51461015057806347fb637914610165578063481fd65c146101a057806354fd4d50146101d35780636c28e70c1461025d5780638da5cb5b14610290578063bbc9a14f146102c1578063c8b80895146102fc578063d321411c1461032f578063ff4face714610362575b600080fd5b3480156100ca57600080fd5b506100f1600480360360208110156100e157600080fd5b5035600160a060020a031661039d565b6040805193845260ff9283166020850152911682820152519081900360600190f35b34801561011f57600080fd5b5061014e6004803603606081101561013657600080fd5b5080359060ff602082013581169160400135166103fc565b005b34801561015c57600080fd5b5061014e6104a2565b34801561017157600080fd5b5061014e6004803603606081101561018857600080fd5b5080359060ff602082013581169160400135166104c5565b3480156101ac57600080fd5b506100f1600480360360208110156101c357600080fd5b5035600160a060020a031661056f565b3480156101df57600080fd5b506101e86105ce565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022257818101518382015260200161020a565b50505050905090810190601f16801561024f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026957600080fd5b506100f16004803603602081101561028057600080fd5b5035600160a060020a031661065b565b34801561029c57600080fd5b506102a56106ba565b60408051600160a060020a039092168252519081900360200190f35b3480156102cd57600080fd5b5061014e600480360360608110156102e457600080fd5b5080359060ff602082013581169160400135166106c9565b34801561030857600080fd5b506100f16004803603602081101561031f57600080fd5b5035600160a060020a031661076f565b34801561033b57600080fd5b506100f16004803603602081101561035257600080fd5b5035600160a060020a03166107cd565b34801561036e57600080fd5b5061014e6004803603606081101561038557600080fd5b5080359060ff602082013581169160400135166107f2565b60008060006103aa610898565b50505050600160a060020a03166000908152600360209081526040918290208251606081018452600582015480825260069092015460ff80821694830185905261010090910416930183905292909190565b610404610898565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600386528a902089516005820155935160069094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a8152918201929092528086019290925293519293926000805160206108b9833981519152929181900390910190a250505050565b600054600160a060020a03163314156104c357600054600160a060020a0316ff5b565b6104cd610898565b50604080516060818101835285825260ff80861660208085018281528784168688018181523360008181526003808752908b90208a5191810191909155935160049094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a8152918201929092528086019290925293519293926000805160206108b9833981519152929181900390910190a250505050565b600080600061057c610898565b50505050600160a060020a03166000908152600360209081526040918290208251606081018452600182015480825260029092015460ff80821694830185905261010090910416930183905292909190565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106535780601f1061062857610100808354040283529160200191610653565b820191906000526020600020905b81548152906001019060200180831161063657829003601f168201915b505050505081565b6000806000610668610898565b50505050600160a060020a03166000908152600360209081526040918290208251606081018452600782015480825260089092015460ff80821694830185905261010090910416930183905292909190565b600054600160a060020a031681565b6106d1610898565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600386528a902089516007820155935160089094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a8152918201929092528086019290925293519293926000805160206108b9833981519152929181900390910190a250505050565b600080600061077c610898565b50505050600160a060020a031660009081526003602081815260409283902083516060810185529281015480845260049091015460ff80821693850184905261010090910416929093018290529192565b6002602052600090815260409020805460019091015460ff8082169161010090041683565b6107fa610898565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600386528a902089516001820155935160029094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a8152918201929092528086019290925293519293926000805160206108b9833981519152929181900390910190a250505050565b60408051606081018252600080825260208201819052918101919091529056fe7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2da165627a7a723058207c5d98e36dec63d2f6e09617c2f874c350fbae3eb840ed8b135a550b123db16b0029`

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

// GetRoute is a free data retrieval call binding the contract method 0x481fd65c.
//
// Solidity: function getRoute(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) GetRoute(opts *bind.CallOpts, addr common.Address) (struct {
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
	err := _Store.contract.Call(opts, out, "getRoute", addr)
	return *ret, err
}

// GetRoute is a free data retrieval call binding the contract method 0x481fd65c.
//
// Solidity: function getRoute(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) GetRoute(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute(&_Store.CallOpts, addr)
}

// GetRoute is a free data retrieval call binding the contract method 0x481fd65c.
//
// Solidity: function getRoute(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) GetRoute(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.GetRoute(&_Store.CallOpts, addr)
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

// Items is a free data retrieval call binding the contract method 0xd321411c.
//
// Solidity: function items( address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Store.contract.Call(opts, out, "items", arg0)
	return *ret, err
}

// Items is a free data retrieval call binding the contract method 0xd321411c.
//
// Solidity: function items( address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreSession) Items(arg0 common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xd321411c.
//
// Solidity: function items( address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_Store *StoreCallerSession) Items(arg0 common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
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

// SetRoute is a paid mutator transaction binding the contract method 0xff4face7.
//
// Solidity: function setRoute(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactor) SetRoute(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setRoute", _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0xff4face7.
//
// Solidity: function setRoute(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreSession) SetRoute(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute(&_Store.TransactOpts, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0xff4face7.
//
// Solidity: function setRoute(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_Store *StoreTransactorSession) SetRoute(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _Store.Contract.SetRoute(&_Store.TransactOpts, _digest, _hashFunction, _size)
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

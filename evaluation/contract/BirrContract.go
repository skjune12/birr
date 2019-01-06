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

// BirrContractABI is the input ABI used to generate the binding from.
const BirrContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAutNum\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getRoute\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAutNum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeRoute6\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeRoute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsRouteKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute6\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAsSet\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsRoute6Key\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setRoute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouteKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRoute6Keys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\"},{\"name\":\"_hashFunction\",\"type\":\"uint8\"},{\"name\":\"_size\",\"type\":\"uint8\"}],\"name\":\"setAsSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getRoute6\",\"outputs\":[{\"name\":\"digest\",\"type\":\"bytes32\"},{\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"name\":\"size\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashFunction\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"size\",\"type\":\"uint8\"}],\"name\":\"ItemSet\",\"type\":\"event\"}]"

// BirrContractBin is the compiled bytecode used for deploying new contracts.
const BirrContractBin = `0x60806040523480156200001157600080fd5b5060405162000fc138038062000fc1833981018060405260208110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b820160208101848111156200006457600080fd5b81516401000000008111828201871017156200007f57600080fd5b505080519093506200009b9250600191506020840190620000b5565b505060008054600160a060020a031916331790556200015a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f857805160ff191683800117855562000128565b8280016001018555821562000128579182015b82811115620001285782518255916020019190600101906200010b565b50620001369291506200013a565b5090565b6200015791905b8082111562000136576000815560010162000141565b90565b610e57806200016a6000396000f3fe608060405234801561001057600080fd5b5060043610610128576000357c0100000000000000000000000000000000000000000000000000000000900480636b98c1a0116100bf5780638da5cb5b1161008e5780638da5cb5b1461036c578063a15b186714610390578063b61c812c146103e8578063bbc9a14f146103f0578063d5d496e91461041e57610128565b80636b98c1a0146102c15780636c28e70c146102f5578063783e6b7a1461031b5780638819bb561461033857610128565b806341c0e1b5116100fb57806341c0e1b5146101ee5780634d0fdcae146101f65780635129efce1461021357806354fd4d501461024457610128565b8063084fa7021461012d5780630b014d841461017557806310799755146101a15780631d8700a1146101d1575b600080fd5b6101536004803603602081101561014357600080fd5b5035600160a060020a031661044a565b6040805193845260ff9283166020850152911682820152519081900360600190f35b6101536004803603604081101561018b57600080fd5b50600160a060020a0381351690602001356104a8565b6101cf600480360360608110156101b757600080fd5b5080359060ff60208201358116916040013516610513565b005b6101cf600480360360208110156101e757600080fd5b50356105b8565b6101cf6105c5565b6101cf6004803603602081101561020c57600080fd5b50356105ea565b6102306004803603602081101561022957600080fd5b50356105f4565b604080519115158252519081900360200190f35b61024c610618565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561028657818101518382015260200161026e565b50505050905090810190601f1680156102b35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101cf600480360360808110156102d757600080fd5b5080359060208101359060ff604082013581169160600135166106a5565b6101536004803603602081101561030b57600080fd5b5035600160a060020a0316610781565b6102306004803603602081101561033157600080fd5b50356107e0565b6101cf6004803603608081101561034e57600080fd5b5080359060208101359060ff60408201358116916060013516610804565b6103746108e0565b60408051600160a060020a039092168252519081900360200190f35b6103986108ef565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103d45781810151838201526020016103bc565b505050509050019250505060405180910390f35b610398610954565b6101cf6004803603606081101561040657600080fd5b5080359060ff602082013581169160400135166109b6565b6101536004803603604081101561043457600080fd5b50600160a060020a038135169060200135610a5c565b6000806000610457610dae565b50505050600160a060020a031660009081526002602081815260409283902083516060810185526001820154808252919093015460ff80821693850184905261010090910416929093018290529192565b60008060006104b5610dae565b505050600160a060020a03831660009081526002602090815260408083208584526005018252918290208251606081018452815480825260019092015460ff8082169483018590526101009091041693018390529250909250925092565b61051b610dae565b50604080516060818101835285825260ff80861660208085018281528784168688018181523360008181526002808752908b90208a516001820155945194018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a815291820192909252808601929092529351929392600080516020610e0c833981519152929181900390910190a250505050565b6105c23382610ac7565b50565b600054600160a060020a031633146105dc57600080fd5b600054600160a060020a0316ff5b6105c23382610b2a565b33600090815260026020908152604080832093835260069093019052205460ff1690565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561069d5780601f106106725761010080835404028352916020019161069d565b820191906000526020600020905b81548152906001019060200180831161068057829003601f168201915b505050505081565b6106ad610dae565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a81208f82526008810187528b82208b51815595516001968701805495518b166101000261ff001992909b1660ff19968716179190911699909917909855600a880180548087018255908252868220018f90558e81526009909701855295899020805490911690921790915586518a81529182019290925280860191909152935192939092600080516020610e0c833981519152929181900390910190a25050505050565b600080600061078e610dae565b50505050600160a060020a03166000908152600260209081526040918290208251606081018452600382015480825260049092015460ff80821694830185905261010090910416930183905292909190565b33600090815260026020908152604080832093835260099093019052205460ff1690565b61080c610dae565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a81208f82526005810187528b82208b51815595516001968701805495518b166101000261ff001992909b1660ff199687161791909116999099179098556007880180548087018255908252868220018f90558e81526006909701855295899020805490911690921790915586518a81529182019290925280860191909152935192939092600080516020610e0c833981519152929181900390910190a25050505050565b600054600160a060020a031681565b3360009081526002602090815260409182902060070180548351818402810184019094528084526060939283018282801561094957602002820191906000526020600020905b815481526020019060010190808311610935575b505050505090505b90565b33600090815260026020908152604091829020600a018054835181840281018401909452808452606093928301828280156109495760200282019190600052602060002090815481526020019060010190808311610935575050505050905090565b6109be610dae565b50604080516060818101835285825260ff8086166020808501828152878416868801818152336000818152600286528a902089516003820155935160049094018054925188166101000261ff00199590981660ff1990931692909217939093169590951790945586518a815291820192909252808601929092529351929392600080516020610e0c833981519152929181900390910190a250505050565b6000806000610a69610dae565b505050600160a060020a03831660009081526002602090815260408083208584526008018252918290208251606081018452815480825260019092015460ff8082169483018590526101009091041693018390529250909250925092565b6000610ad38383610b8d565b9050610adf8382610bda565b50600160a060020a03919091166000908152600260209081526040808320938352600884018252808320838155600101805461ffff191690556009909301905220805460ff19169055565b6000610b368383610ca3565b9050610b428382610cea565b50600160a060020a03919091166000908152600260209081526040808320938352600584018252808320838155600101805461ffff191690556006909301905220805460ff19169055565b60005b600160a060020a0383166000908152600260205260409020600a01805483919083908110610bba57fe5b9060005260206000200154141515610bd457600101610b90565b92915050565b600160a060020a0382166000908152600260205260409020600a015460001901811015610c7257600160a060020a0382166000908152600260205260409020600a01805460018301908110610c2b57fe5b6000918252602080832090910154600160a060020a038516835260029091526040909120600a01805483908110610c5e57fe5b600091825260209091200155600101610bda565b600160a060020a0382166000908152600260205260409020600a01805490610c9e906000198301610dce565b505050565b60005b600160a060020a0383166000908152600260205260409020600701805483919083908110610cd057fe5b9060005260206000200154141515610bd457600101610ca6565b600160a060020a03821660009081526002602052604090206007015460001901811015610d8257600160a060020a0382166000908152600260205260409020600701805460018301908110610d3b57fe5b6000918252602080832090910154600160a060020a038516835260029091526040909120600701805483908110610d6e57fe5b600091825260209091200155600101610cea565b600160a060020a0382166000908152600260205260409020600701805490610c9e906000198301610dce565b604080516060810182526000808252602082018190529181019190915290565b815481835581811115610c9e57600083815260209020610c9e91810190830161095191905b80821115610e075760008155600101610df3565b509056fe7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2da165627a7a7230582087f67f81118f98d3e41232fdaaf9a10d5070d56ea004a758499b6193d1653fb30029`

// DeployBirrContract deploys a new Ethereum contract, binding an instance of BirrContract to it.
func DeployBirrContract(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *BirrContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BirrContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BirrContractBin), backend, _version)
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

// ExistsRoute6Key is a free data retrieval call binding the contract method 0x783e6b7a.
//
// Solidity: function existsRoute6Key(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractCaller) ExistsRoute6Key(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "existsRoute6Key", key)
	return *ret0, err
}

// ExistsRoute6Key is a free data retrieval call binding the contract method 0x783e6b7a.
//
// Solidity: function existsRoute6Key(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractSession) ExistsRoute6Key(key [32]byte) (bool, error) {
	return _BirrContract.Contract.ExistsRoute6Key(&_BirrContract.CallOpts, key)
}

// ExistsRoute6Key is a free data retrieval call binding the contract method 0x783e6b7a.
//
// Solidity: function existsRoute6Key(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractCallerSession) ExistsRoute6Key(key [32]byte) (bool, error) {
	return _BirrContract.Contract.ExistsRoute6Key(&_BirrContract.CallOpts, key)
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractCaller) ExistsRouteKey(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "existsRouteKey", key)
	return *ret0, err
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractSession) ExistsRouteKey(key [32]byte) (bool, error) {
	return _BirrContract.Contract.ExistsRouteKey(&_BirrContract.CallOpts, key)
}

// ExistsRouteKey is a free data retrieval call binding the contract method 0x5129efce.
//
// Solidity: function existsRouteKey(key bytes32) constant returns(bool)
func (_BirrContract *BirrContractCallerSession) ExistsRouteKey(key [32]byte) (bool, error) {
	return _BirrContract.Contract.ExistsRouteKey(&_BirrContract.CallOpts, key)
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCaller) GetAsSet(opts *bind.CallOpts, addr common.Address) (struct {
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
	err := _BirrContract.contract.Call(opts, out, "getAsSet", addr)
	return *ret, err
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractSession) GetAsSet(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetAsSet(&_BirrContract.CallOpts, addr)
}

// GetAsSet is a free data retrieval call binding the contract method 0x6c28e70c.
//
// Solidity: function getAsSet(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCallerSession) GetAsSet(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetAsSet(&_BirrContract.CallOpts, addr)
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCaller) GetAutNum(opts *bind.CallOpts, addr common.Address) (struct {
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
	err := _BirrContract.contract.Call(opts, out, "getAutNum", addr)
	return *ret, err
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractSession) GetAutNum(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetAutNum(&_BirrContract.CallOpts, addr)
}

// GetAutNum is a free data retrieval call binding the contract method 0x084fa702.
//
// Solidity: function getAutNum(addr address) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCallerSession) GetAutNum(addr common.Address) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetAutNum(&_BirrContract.CallOpts, addr)
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCaller) GetRoute(opts *bind.CallOpts, addr common.Address, key [32]byte) (struct {
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
	err := _BirrContract.contract.Call(opts, out, "getRoute", addr, key)
	return *ret, err
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractSession) GetRoute(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetRoute(&_BirrContract.CallOpts, addr, key)
}

// GetRoute is a free data retrieval call binding the contract method 0x0b014d84.
//
// Solidity: function getRoute(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCallerSession) GetRoute(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetRoute(&_BirrContract.CallOpts, addr, key)
}

// GetRoute6 is a free data retrieval call binding the contract method 0xd5d496e9.
//
// Solidity: function getRoute6(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCaller) GetRoute6(opts *bind.CallOpts, addr common.Address, key [32]byte) (struct {
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
	err := _BirrContract.contract.Call(opts, out, "getRoute6", addr, key)
	return *ret, err
}

// GetRoute6 is a free data retrieval call binding the contract method 0xd5d496e9.
//
// Solidity: function getRoute6(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractSession) GetRoute6(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetRoute6(&_BirrContract.CallOpts, addr, key)
}

// GetRoute6 is a free data retrieval call binding the contract method 0xd5d496e9.
//
// Solidity: function getRoute6(addr address, key bytes32) constant returns(digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractCallerSession) GetRoute6(addr common.Address, key [32]byte) (struct {
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
}, error) {
	return _BirrContract.Contract.GetRoute6(&_BirrContract.CallOpts, addr, key)
}

// GetRoute6Keys is a free data retrieval call binding the contract method 0xb61c812c.
//
// Solidity: function getRoute6Keys() constant returns(bytes32[])
func (_BirrContract *BirrContractCaller) GetRoute6Keys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "getRoute6Keys")
	return *ret0, err
}

// GetRoute6Keys is a free data retrieval call binding the contract method 0xb61c812c.
//
// Solidity: function getRoute6Keys() constant returns(bytes32[])
func (_BirrContract *BirrContractSession) GetRoute6Keys() ([][32]byte, error) {
	return _BirrContract.Contract.GetRoute6Keys(&_BirrContract.CallOpts)
}

// GetRoute6Keys is a free data retrieval call binding the contract method 0xb61c812c.
//
// Solidity: function getRoute6Keys() constant returns(bytes32[])
func (_BirrContract *BirrContractCallerSession) GetRoute6Keys() ([][32]byte, error) {
	return _BirrContract.Contract.GetRoute6Keys(&_BirrContract.CallOpts)
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_BirrContract *BirrContractCaller) GetRouteKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "getRouteKeys")
	return *ret0, err
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_BirrContract *BirrContractSession) GetRouteKeys() ([][32]byte, error) {
	return _BirrContract.Contract.GetRouteKeys(&_BirrContract.CallOpts)
}

// GetRouteKeys is a free data retrieval call binding the contract method 0xa15b1867.
//
// Solidity: function getRouteKeys() constant returns(bytes32[])
func (_BirrContract *BirrContractCallerSession) GetRouteKeys() ([][32]byte, error) {
	return _BirrContract.Contract.GetRouteKeys(&_BirrContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BirrContract *BirrContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BirrContract *BirrContractSession) Owner() (common.Address, error) {
	return _BirrContract.Contract.Owner(&_BirrContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BirrContract *BirrContractCallerSession) Owner() (common.Address, error) {
	return _BirrContract.Contract.Owner(&_BirrContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BirrContract *BirrContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BirrContract *BirrContractSession) Version() (string, error) {
	return _BirrContract.Contract.Version(&_BirrContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BirrContract *BirrContractCallerSession) Version() (string, error) {
	return _BirrContract.Contract.Version(&_BirrContract.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BirrContract *BirrContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BirrContract *BirrContractSession) Kill() (*types.Transaction, error) {
	return _BirrContract.Contract.Kill(&_BirrContract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BirrContract *BirrContractTransactorSession) Kill() (*types.Transaction, error) {
	return _BirrContract.Contract.Kill(&_BirrContract.TransactOpts)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_BirrContract *BirrContractTransactor) RemoveRoute(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "removeRoute", key)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_BirrContract *BirrContractSession) RemoveRoute(key [32]byte) (*types.Transaction, error) {
	return _BirrContract.Contract.RemoveRoute(&_BirrContract.TransactOpts, key)
}

// RemoveRoute is a paid mutator transaction binding the contract method 0x4d0fdcae.
//
// Solidity: function removeRoute(key bytes32) returns()
func (_BirrContract *BirrContractTransactorSession) RemoveRoute(key [32]byte) (*types.Transaction, error) {
	return _BirrContract.Contract.RemoveRoute(&_BirrContract.TransactOpts, key)
}

// RemoveRoute6 is a paid mutator transaction binding the contract method 0x1d8700a1.
//
// Solidity: function removeRoute6(key bytes32) returns()
func (_BirrContract *BirrContractTransactor) RemoveRoute6(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "removeRoute6", key)
}

// RemoveRoute6 is a paid mutator transaction binding the contract method 0x1d8700a1.
//
// Solidity: function removeRoute6(key bytes32) returns()
func (_BirrContract *BirrContractSession) RemoveRoute6(key [32]byte) (*types.Transaction, error) {
	return _BirrContract.Contract.RemoveRoute6(&_BirrContract.TransactOpts, key)
}

// RemoveRoute6 is a paid mutator transaction binding the contract method 0x1d8700a1.
//
// Solidity: function removeRoute6(key bytes32) returns()
func (_BirrContract *BirrContractTransactorSession) RemoveRoute6(key [32]byte) (*types.Transaction, error) {
	return _BirrContract.Contract.RemoveRoute6(&_BirrContract.TransactOpts, key)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactor) SetAsSet(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setAsSet", _digest, _hashFunction, _size)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractSession) SetAsSet(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAsSet(&_BirrContract.TransactOpts, _digest, _hashFunction, _size)
}

// SetAsSet is a paid mutator transaction binding the contract method 0xbbc9a14f.
//
// Solidity: function setAsSet(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactorSession) SetAsSet(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAsSet(&_BirrContract.TransactOpts, _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactor) SetAutNum(opts *bind.TransactOpts, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setAutNum", _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractSession) SetAutNum(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAutNum(&_BirrContract.TransactOpts, _digest, _hashFunction, _size)
}

// SetAutNum is a paid mutator transaction binding the contract method 0x10799755.
//
// Solidity: function setAutNum(_digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactorSession) SetAutNum(_digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAutNum(&_BirrContract.TransactOpts, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactor) SetRoute(opts *bind.TransactOpts, key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setRoute", key, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractSession) SetRoute(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute(&_BirrContract.TransactOpts, key, _digest, _hashFunction, _size)
}

// SetRoute is a paid mutator transaction binding the contract method 0x8819bb56.
//
// Solidity: function setRoute(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactorSession) SetRoute(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute(&_BirrContract.TransactOpts, key, _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x6b98c1a0.
//
// Solidity: function setRoute6(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactor) SetRoute6(opts *bind.TransactOpts, key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setRoute6", key, _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x6b98c1a0.
//
// Solidity: function setRoute6(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractSession) SetRoute6(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute6(&_BirrContract.TransactOpts, key, _digest, _hashFunction, _size)
}

// SetRoute6 is a paid mutator transaction binding the contract method 0x6b98c1a0.
//
// Solidity: function setRoute6(key bytes32, _digest bytes32, _hashFunction uint8, _size uint8) returns()
func (_BirrContract *BirrContractTransactorSession) SetRoute6(key [32]byte, _digest [32]byte, _hashFunction uint8, _size uint8) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute6(&_BirrContract.TransactOpts, key, _digest, _hashFunction, _size)
}

// BirrContractItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the BirrContract contract.
type BirrContractItemSetIterator struct {
	Event *BirrContractItemSet // Event containing the contract specifics and raw log

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
func (it *BirrContractItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BirrContractItemSet)
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
		it.Event = new(BirrContractItemSet)
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
func (it *BirrContractItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BirrContractItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BirrContractItemSet represents a ItemSet event raised by the BirrContract contract.
type BirrContractItemSet struct {
	Key          common.Address
	Digest       [32]byte
	HashFunction uint8
	Size         uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2d.
//
// Solidity: e ItemSet(key indexed address, digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractFilterer) FilterItemSet(opts *bind.FilterOpts, key []common.Address) (*BirrContractItemSetIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _BirrContract.contract.FilterLogs(opts, "ItemSet", keyRule)
	if err != nil {
		return nil, err
	}
	return &BirrContractItemSetIterator{contract: _BirrContract.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x7abc65a09f41055eede1f559e9c49c9b294484af9110a5fde063cc2ff243be2d.
//
// Solidity: e ItemSet(key indexed address, digest bytes32, hashFunction uint8, size uint8)
func (_BirrContract *BirrContractFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *BirrContractItemSet, key []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _BirrContract.contract.WatchLogs(opts, "ItemSet", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BirrContractItemSet)
				if err := _BirrContract.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

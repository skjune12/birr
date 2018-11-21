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
const BirrContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setAutNumObj\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getObject\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setRoute6Obj\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setRouteObj\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfAddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"setAsSetObj\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asNumber\",\"type\":\"uint32\"}],\"name\":\"newObject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BirrContractBin is the compiled bytecode used for deploying new contracts.
const BirrContractBin = `0x608060405234801561001057600080fd5b50610a62806100206000396000f3fe60806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630e70965581146100925780631edc801d14610147578063428761ea1461034457806379353684146103f7578063b810fb43146104aa578063c0f9bde0146104fd578063c115aa0414610524578063c4be1f9f146105d7575b600080fd5b34801561009e57600080fd5b50610145600480360360208110156100b557600080fd5b8101906020810181356401000000008111156100d057600080fd5b8201836020820111156100e257600080fd5b8035906020019184600183028401116401000000008311171561010457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610607945050505050565b005b34801561015357600080fd5b506101876004803603602081101561016a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610626565b604051808663ffffffff1663ffffffff16815260200180602001806020018060200180602001858103855289818151815260200191508051906020019080838360005b838110156101e25781810151838201526020016101ca565b50505050905090810190601f16801561020f5780820380516001836020036101000a031916815260200191505b5085810384528851815288516020918201918a019080838360005b8381101561024257818101518382015260200161022a565b50505050905090810190601f16801561026f5780820380516001836020036101000a031916815260200191505b50858103835287518152875160209182019189019080838360005b838110156102a257818101518382015260200161028a565b50505050905090810190601f1680156102cf5780820380516001836020036101000a031916815260200191505b50858103825286518152865160209182019188019080838360005b838110156103025781810151838201526020016102ea565b50505050905090810190601f16801561032f5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34801561035057600080fd5b506101456004803603602081101561036757600080fd5b81019060208101813564010000000081111561038257600080fd5b82018360208201111561039457600080fd5b803590602001918460018302840111640100000000831117156103b657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106f1945050505050565b34801561040357600080fd5b506101456004803603602081101561041a57600080fd5b81019060208101813564010000000081111561043557600080fd5b82018360208201111561044757600080fd5b8035906020019184600183028401116401000000008311171561046957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610710945050505050565b3480156104b657600080fd5b506104d4600480360360208110156104cd57600080fd5b503561072f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561050957600080fd5b50610512610764565b60408051918252519081900360200190f35b34801561053057600080fd5b506101456004803603602081101561054757600080fd5b81019060208101813564010000000081111561056257600080fd5b82018360208201111561057457600080fd5b8035906020019184600183028401116401000000008311171561059657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061076a945050505050565b3480156105e357600080fd5b50610145600480360360208110156105fa57600080fd5b503563ffffffff16610789565b6106108161084b565b3360009081526020819052604090206003015550565b6000606080606080610636610a08565b5073ffffffffffffffffffffffffffffffffffffffff861660009081526020818152604091829020825160a081018452815463ffffffff168152600182015492810183905260028201549381019390935260038101546060808501919091526004909101546080840152906106aa9061086f565b905060606106bb836040015161086f565b905060606106cc846060015161086f565b905060606106dd856080015161086f565b94519b939a50919850965091945092505050565b6106fa8161084b565b3360009081526020819052604090206002015550565b6107198161084b565b3360009081526020819052604090206001015550565b600180548290811061073d57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60015490565b6107738161084b565b3360009081526020819052604090206004015550565b6040805160a08101825263ffffffff928316815260006020808301828152838501838152606085018481526080860185815233808752948690529685209551865463ffffffff19169816979097178555905160018086019190915590516002850155945160038401559251600490920191909155825480840184559290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6909101805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b80516000908290151561086257506000905061086a565b505060208101515b919050565b6040805160208082528183019092526060918291906020820181803883390190505090506000805b6020811015610920576008810260020a85027fff00000000000000000000000000000000000000000000000000000000000000811615610917578084848151811015156108e057fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506001909201915b50600101610897565b6060826040519080825280601f01601f19166020018201604052801561094d576020820181803883390190505b509050600091505b828210156109ff57838281518110151561096b57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000281838151811015156109c457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190910190610955565b95945050505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea165627a7a72305820ed8ea866cfb66d72e0a416774011b858300a144e51214128d006bc1ba49d25de0029`

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

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList( uint256) constant returns(address)
func (_BirrContract *BirrContractCaller) AddressList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "addressList", arg0)
	return *ret0, err
}

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList( uint256) constant returns(address)
func (_BirrContract *BirrContractSession) AddressList(arg0 *big.Int) (common.Address, error) {
	return _BirrContract.Contract.AddressList(&_BirrContract.CallOpts, arg0)
}

// AddressList is a free data retrieval call binding the contract method 0xb810fb43.
//
// Solidity: function addressList( uint256) constant returns(address)
func (_BirrContract *BirrContractCallerSession) AddressList(arg0 *big.Int) (common.Address, error) {
	return _BirrContract.Contract.AddressList(&_BirrContract.CallOpts, arg0)
}

// GetObject is a free data retrieval call binding the contract method 0x1edc801d.
//
// Solidity: function getObject(addr address) constant returns(uint32, string, string, string, string)
func (_BirrContract *BirrContractCaller) GetObject(opts *bind.CallOpts, addr common.Address) (uint32, string, string, string, string, error) {
	var (
		ret0 = new(uint32)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BirrContract.contract.Call(opts, out, "getObject", addr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetObject is a free data retrieval call binding the contract method 0x1edc801d.
//
// Solidity: function getObject(addr address) constant returns(uint32, string, string, string, string)
func (_BirrContract *BirrContractSession) GetObject(addr common.Address) (uint32, string, string, string, string, error) {
	return _BirrContract.Contract.GetObject(&_BirrContract.CallOpts, addr)
}

// GetObject is a free data retrieval call binding the contract method 0x1edc801d.
//
// Solidity: function getObject(addr address) constant returns(uint32, string, string, string, string)
func (_BirrContract *BirrContractCallerSession) GetObject(addr common.Address) (uint32, string, string, string, string, error) {
	return _BirrContract.Contract.GetObject(&_BirrContract.CallOpts, addr)
}

// NumOfAddrs is a free data retrieval call binding the contract method 0xc0f9bde0.
//
// Solidity: function numOfAddrs() constant returns(uint256)
func (_BirrContract *BirrContractCaller) NumOfAddrs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BirrContract.contract.Call(opts, out, "numOfAddrs")
	return *ret0, err
}

// NumOfAddrs is a free data retrieval call binding the contract method 0xc0f9bde0.
//
// Solidity: function numOfAddrs() constant returns(uint256)
func (_BirrContract *BirrContractSession) NumOfAddrs() (*big.Int, error) {
	return _BirrContract.Contract.NumOfAddrs(&_BirrContract.CallOpts)
}

// NumOfAddrs is a free data retrieval call binding the contract method 0xc0f9bde0.
//
// Solidity: function numOfAddrs() constant returns(uint256)
func (_BirrContract *BirrContractCallerSession) NumOfAddrs() (*big.Int, error) {
	return _BirrContract.Contract.NumOfAddrs(&_BirrContract.CallOpts)
}

// NewObject is a paid mutator transaction binding the contract method 0xc4be1f9f.
//
// Solidity: function newObject(asNumber uint32) returns()
func (_BirrContract *BirrContractTransactor) NewObject(opts *bind.TransactOpts, asNumber uint32) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "newObject", asNumber)
}

// NewObject is a paid mutator transaction binding the contract method 0xc4be1f9f.
//
// Solidity: function newObject(asNumber uint32) returns()
func (_BirrContract *BirrContractSession) NewObject(asNumber uint32) (*types.Transaction, error) {
	return _BirrContract.Contract.NewObject(&_BirrContract.TransactOpts, asNumber)
}

// NewObject is a paid mutator transaction binding the contract method 0xc4be1f9f.
//
// Solidity: function newObject(asNumber uint32) returns()
func (_BirrContract *BirrContractTransactorSession) NewObject(asNumber uint32) (*types.Transaction, error) {
	return _BirrContract.Contract.NewObject(&_BirrContract.TransactOpts, asNumber)
}

// SetAsSetObj is a paid mutator transaction binding the contract method 0xc115aa04.
//
// Solidity: function setAsSetObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactor) SetAsSetObj(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setAsSetObj", ipfsHash)
}

// SetAsSetObj is a paid mutator transaction binding the contract method 0xc115aa04.
//
// Solidity: function setAsSetObj(ipfsHash string) returns()
func (_BirrContract *BirrContractSession) SetAsSetObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAsSetObj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetAsSetObj is a paid mutator transaction binding the contract method 0xc115aa04.
//
// Solidity: function setAsSetObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactorSession) SetAsSetObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAsSetObj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetAutNumObj is a paid mutator transaction binding the contract method 0x0e709655.
//
// Solidity: function setAutNumObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactor) SetAutNumObj(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setAutNumObj", ipfsHash)
}

// SetAutNumObj is a paid mutator transaction binding the contract method 0x0e709655.
//
// Solidity: function setAutNumObj(ipfsHash string) returns()
func (_BirrContract *BirrContractSession) SetAutNumObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAutNumObj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetAutNumObj is a paid mutator transaction binding the contract method 0x0e709655.
//
// Solidity: function setAutNumObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactorSession) SetAutNumObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetAutNumObj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetRoute6Obj is a paid mutator transaction binding the contract method 0x428761ea.
//
// Solidity: function setRoute6Obj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactor) SetRoute6Obj(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setRoute6Obj", ipfsHash)
}

// SetRoute6Obj is a paid mutator transaction binding the contract method 0x428761ea.
//
// Solidity: function setRoute6Obj(ipfsHash string) returns()
func (_BirrContract *BirrContractSession) SetRoute6Obj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute6Obj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetRoute6Obj is a paid mutator transaction binding the contract method 0x428761ea.
//
// Solidity: function setRoute6Obj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactorSession) SetRoute6Obj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRoute6Obj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetRouteObj is a paid mutator transaction binding the contract method 0x79353684.
//
// Solidity: function setRouteObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactor) SetRouteObj(opts *bind.TransactOpts, ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.contract.Transact(opts, "setRouteObj", ipfsHash)
}

// SetRouteObj is a paid mutator transaction binding the contract method 0x79353684.
//
// Solidity: function setRouteObj(ipfsHash string) returns()
func (_BirrContract *BirrContractSession) SetRouteObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRouteObj(&_BirrContract.TransactOpts, ipfsHash)
}

// SetRouteObj is a paid mutator transaction binding the contract method 0x79353684.
//
// Solidity: function setRouteObj(ipfsHash string) returns()
func (_BirrContract *BirrContractTransactorSession) SetRouteObj(ipfsHash string) (*types.Transaction, error) {
	return _BirrContract.Contract.SetRouteObj(&_BirrContract.TransactOpts, ipfsHash)
}

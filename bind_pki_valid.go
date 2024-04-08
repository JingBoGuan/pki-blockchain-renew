// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// LuxUniPKIValidMetaData contains all meta data concerning the LuxUniPKIValid contract.
var LuxUniPKIValidMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_res\",\"type\":\"uint256\"}],\"name\":\"DecodeReturnErr\",\"outputs\":[{\"name\":\"_err\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_res\",\"type\":\"uint256\"}],\"name\":\"DecodeReturnIter\",\"outputs\":[{\"name\":\"_iter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_newHash\",\"type\":\"bytes32\"},{\"name\":\"_addrCA\",\"type\":\"address\"},{\"name\":\"_addrRoot\",\"type\":\"address\"}],\"name\":\"CheckCert\",\"outputs\":[{\"name\":\"_result\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_iter\",\"type\":\"uint256\"},{\"name\":\"_err\",\"type\":\"uint256\"}],\"name\":\"EncodeReturn\",\"outputs\":[{\"name\":\"_res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LuxUniPKIValidABI is the input ABI used to generate the binding from.
// Deprecated: Use LuxUniPKIValidMetaData.ABI instead.
var LuxUniPKIValidABI = LuxUniPKIValidMetaData.ABI

// LuxUniPKIValid is an auto generated Go binding around an Ethereum contract.
type LuxUniPKIValid struct {
	LuxUniPKIValidCaller     // Read-only binding to the contract
	LuxUniPKIValidTransactor // Write-only binding to the contract
	LuxUniPKIValidFilterer   // Log filterer for contract events
}

// LuxUniPKIValidCaller is an auto generated read-only Go binding around an Ethereum contract.
type LuxUniPKIValidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIValidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LuxUniPKIValidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIValidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LuxUniPKIValidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIValidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LuxUniPKIValidSession struct {
	Contract     *LuxUniPKIValid   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LuxUniPKIValidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LuxUniPKIValidCallerSession struct {
	Contract *LuxUniPKIValidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LuxUniPKIValidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LuxUniPKIValidTransactorSession struct {
	Contract     *LuxUniPKIValidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LuxUniPKIValidRaw is an auto generated low-level Go binding around an Ethereum contract.
type LuxUniPKIValidRaw struct {
	Contract *LuxUniPKIValid // Generic contract binding to access the raw methods on
}

// LuxUniPKIValidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LuxUniPKIValidCallerRaw struct {
	Contract *LuxUniPKIValidCaller // Generic read-only contract binding to access the raw methods on
}

// LuxUniPKIValidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LuxUniPKIValidTransactorRaw struct {
	Contract *LuxUniPKIValidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLuxUniPKIValid creates a new instance of LuxUniPKIValid, bound to a specific deployed contract.
func NewLuxUniPKIValid(address common.Address, backend bind.ContractBackend) (*LuxUniPKIValid, error) {
	contract, err := bindLuxUniPKIValid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIValid{LuxUniPKIValidCaller: LuxUniPKIValidCaller{contract: contract}, LuxUniPKIValidTransactor: LuxUniPKIValidTransactor{contract: contract}, LuxUniPKIValidFilterer: LuxUniPKIValidFilterer{contract: contract}}, nil
}

// NewLuxUniPKIValidCaller creates a new read-only instance of LuxUniPKIValid, bound to a specific deployed contract.
func NewLuxUniPKIValidCaller(address common.Address, caller bind.ContractCaller) (*LuxUniPKIValidCaller, error) {
	contract, err := bindLuxUniPKIValid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIValidCaller{contract: contract}, nil
}

// NewLuxUniPKIValidTransactor creates a new write-only instance of LuxUniPKIValid, bound to a specific deployed contract.
func NewLuxUniPKIValidTransactor(address common.Address, transactor bind.ContractTransactor) (*LuxUniPKIValidTransactor, error) {
	contract, err := bindLuxUniPKIValid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIValidTransactor{contract: contract}, nil
}

// NewLuxUniPKIValidFilterer creates a new log filterer instance of LuxUniPKIValid, bound to a specific deployed contract.
func NewLuxUniPKIValidFilterer(address common.Address, filterer bind.ContractFilterer) (*LuxUniPKIValidFilterer, error) {
	contract, err := bindLuxUniPKIValid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIValidFilterer{contract: contract}, nil
}

// bindLuxUniPKIValid binds a generic wrapper to an already deployed contract.
func bindLuxUniPKIValid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LuxUniPKIValidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuxUniPKIValid *LuxUniPKIValidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuxUniPKIValid.Contract.LuxUniPKIValidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuxUniPKIValid *LuxUniPKIValidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuxUniPKIValid.Contract.LuxUniPKIValidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuxUniPKIValid *LuxUniPKIValidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuxUniPKIValid.Contract.LuxUniPKIValidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuxUniPKIValid *LuxUniPKIValidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuxUniPKIValid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuxUniPKIValid *LuxUniPKIValidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuxUniPKIValid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuxUniPKIValid *LuxUniPKIValidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuxUniPKIValid.Contract.contract.Transact(opts, method, params...)
}

// CheckCert is a free data retrieval call binding the contract method 0xb31c6071.
//
// Solidity: function CheckCert(bytes32 _newHash, address _addrCA, address _addrRoot) view returns(uint256 _result)
func (_LuxUniPKIValid *LuxUniPKIValidCaller) CheckCert(opts *bind.CallOpts, _newHash [32]byte, _addrCA common.Address, _addrRoot common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIValid.contract.Call(opts, &out, "CheckCert", _newHash, _addrCA, _addrRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckCert is a free data retrieval call binding the contract method 0xb31c6071.
//
// Solidity: function CheckCert(bytes32 _newHash, address _addrCA, address _addrRoot) view returns(uint256 _result)
func (_LuxUniPKIValid *LuxUniPKIValidSession) CheckCert(_newHash [32]byte, _addrCA common.Address, _addrRoot common.Address) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.CheckCert(&_LuxUniPKIValid.CallOpts, _newHash, _addrCA, _addrRoot)
}

// CheckCert is a free data retrieval call binding the contract method 0xb31c6071.
//
// Solidity: function CheckCert(bytes32 _newHash, address _addrCA, address _addrRoot) view returns(uint256 _result)
func (_LuxUniPKIValid *LuxUniPKIValidCallerSession) CheckCert(_newHash [32]byte, _addrCA common.Address, _addrRoot common.Address) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.CheckCert(&_LuxUniPKIValid.CallOpts, _newHash, _addrCA, _addrRoot)
}

// DecodeReturnErr is a free data retrieval call binding the contract method 0x7ce13531.
//
// Solidity: function DecodeReturnErr(uint256 _res) view returns(uint256 _err)
func (_LuxUniPKIValid *LuxUniPKIValidCaller) DecodeReturnErr(opts *bind.CallOpts, _res *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIValid.contract.Call(opts, &out, "DecodeReturnErr", _res)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecodeReturnErr is a free data retrieval call binding the contract method 0x7ce13531.
//
// Solidity: function DecodeReturnErr(uint256 _res) view returns(uint256 _err)
func (_LuxUniPKIValid *LuxUniPKIValidSession) DecodeReturnErr(_res *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.DecodeReturnErr(&_LuxUniPKIValid.CallOpts, _res)
}

// DecodeReturnErr is a free data retrieval call binding the contract method 0x7ce13531.
//
// Solidity: function DecodeReturnErr(uint256 _res) view returns(uint256 _err)
func (_LuxUniPKIValid *LuxUniPKIValidCallerSession) DecodeReturnErr(_res *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.DecodeReturnErr(&_LuxUniPKIValid.CallOpts, _res)
}

// DecodeReturnIter is a free data retrieval call binding the contract method 0x82ff0002.
//
// Solidity: function DecodeReturnIter(uint256 _res) view returns(uint256 _iter)
func (_LuxUniPKIValid *LuxUniPKIValidCaller) DecodeReturnIter(opts *bind.CallOpts, _res *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIValid.contract.Call(opts, &out, "DecodeReturnIter", _res)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecodeReturnIter is a free data retrieval call binding the contract method 0x82ff0002.
//
// Solidity: function DecodeReturnIter(uint256 _res) view returns(uint256 _iter)
func (_LuxUniPKIValid *LuxUniPKIValidSession) DecodeReturnIter(_res *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.DecodeReturnIter(&_LuxUniPKIValid.CallOpts, _res)
}

// DecodeReturnIter is a free data retrieval call binding the contract method 0x82ff0002.
//
// Solidity: function DecodeReturnIter(uint256 _res) view returns(uint256 _iter)
func (_LuxUniPKIValid *LuxUniPKIValidCallerSession) DecodeReturnIter(_res *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.DecodeReturnIter(&_LuxUniPKIValid.CallOpts, _res)
}

// EncodeReturn is a free data retrieval call binding the contract method 0xeb31f8ab.
//
// Solidity: function EncodeReturn(uint256 _iter, uint256 _err) view returns(uint256 _res)
func (_LuxUniPKIValid *LuxUniPKIValidCaller) EncodeReturn(opts *bind.CallOpts, _iter *big.Int, _err *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIValid.contract.Call(opts, &out, "EncodeReturn", _iter, _err)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeReturn is a free data retrieval call binding the contract method 0xeb31f8ab.
//
// Solidity: function EncodeReturn(uint256 _iter, uint256 _err) view returns(uint256 _res)
func (_LuxUniPKIValid *LuxUniPKIValidSession) EncodeReturn(_iter *big.Int, _err *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.EncodeReturn(&_LuxUniPKIValid.CallOpts, _iter, _err)
}

// EncodeReturn is a free data retrieval call binding the contract method 0xeb31f8ab.
//
// Solidity: function EncodeReturn(uint256 _iter, uint256 _err) view returns(uint256 _res)
func (_LuxUniPKIValid *LuxUniPKIValidCallerSession) EncodeReturn(_iter *big.Int, _err *big.Int) (*big.Int, error) {
	return _LuxUniPKIValid.Contract.EncodeReturn(&_LuxUniPKIValid.CallOpts, _iter, _err)
}

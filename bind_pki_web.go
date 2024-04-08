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

// LuxUniPKIWebMetaData contains all meta data concerning the LuxUniPKIWeb contract.
var LuxUniPKIWebMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"getRegCreationDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parentAddr\",\"type\":\"address\"},{\"name\":\"_arrInd\",\"type\":\"uint256\"},{\"name\":\"_ethAccCA\",\"type\":\"address\"},{\"name\":\"_contrAddr\",\"type\":\"address\"},{\"name\":\"_fileName\",\"type\":\"string\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"newRegDatum\",\"outputs\":[{\"name\":\"err\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"getRegEthAccCA\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"getRegContrAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"getRegDescription\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"getRegFileName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addrContr\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint96\"}],\"name\":\"EncodeMapID\",\"outputs\":[{\"name\":\"_res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LuxUniPKIWebABI is the input ABI used to generate the binding from.
// Deprecated: Use LuxUniPKIWebMetaData.ABI instead.
var LuxUniPKIWebABI = LuxUniPKIWebMetaData.ABI

// LuxUniPKIWeb is an auto generated Go binding around an Ethereum contract.
type LuxUniPKIWeb struct {
	LuxUniPKIWebCaller     // Read-only binding to the contract
	LuxUniPKIWebTransactor // Write-only binding to the contract
	LuxUniPKIWebFilterer   // Log filterer for contract events
}

// LuxUniPKIWebCaller is an auto generated read-only Go binding around an Ethereum contract.
type LuxUniPKIWebCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIWebTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LuxUniPKIWebTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIWebFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LuxUniPKIWebFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuxUniPKIWebSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LuxUniPKIWebSession struct {
	Contract     *LuxUniPKIWeb     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LuxUniPKIWebCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LuxUniPKIWebCallerSession struct {
	Contract *LuxUniPKIWebCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LuxUniPKIWebTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LuxUniPKIWebTransactorSession struct {
	Contract     *LuxUniPKIWebTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LuxUniPKIWebRaw is an auto generated low-level Go binding around an Ethereum contract.
type LuxUniPKIWebRaw struct {
	Contract *LuxUniPKIWeb // Generic contract binding to access the raw methods on
}

// LuxUniPKIWebCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LuxUniPKIWebCallerRaw struct {
	Contract *LuxUniPKIWebCaller // Generic read-only contract binding to access the raw methods on
}

// LuxUniPKIWebTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LuxUniPKIWebTransactorRaw struct {
	Contract *LuxUniPKIWebTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLuxUniPKIWeb creates a new instance of LuxUniPKIWeb, bound to a specific deployed contract.
func NewLuxUniPKIWeb(address common.Address, backend bind.ContractBackend) (*LuxUniPKIWeb, error) {
	contract, err := bindLuxUniPKIWeb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIWeb{LuxUniPKIWebCaller: LuxUniPKIWebCaller{contract: contract}, LuxUniPKIWebTransactor: LuxUniPKIWebTransactor{contract: contract}, LuxUniPKIWebFilterer: LuxUniPKIWebFilterer{contract: contract}}, nil
}

// NewLuxUniPKIWebCaller creates a new read-only instance of LuxUniPKIWeb, bound to a specific deployed contract.
func NewLuxUniPKIWebCaller(address common.Address, caller bind.ContractCaller) (*LuxUniPKIWebCaller, error) {
	contract, err := bindLuxUniPKIWeb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIWebCaller{contract: contract}, nil
}

// NewLuxUniPKIWebTransactor creates a new write-only instance of LuxUniPKIWeb, bound to a specific deployed contract.
func NewLuxUniPKIWebTransactor(address common.Address, transactor bind.ContractTransactor) (*LuxUniPKIWebTransactor, error) {
	contract, err := bindLuxUniPKIWeb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIWebTransactor{contract: contract}, nil
}

// NewLuxUniPKIWebFilterer creates a new log filterer instance of LuxUniPKIWeb, bound to a specific deployed contract.
func NewLuxUniPKIWebFilterer(address common.Address, filterer bind.ContractFilterer) (*LuxUniPKIWebFilterer, error) {
	contract, err := bindLuxUniPKIWeb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LuxUniPKIWebFilterer{contract: contract}, nil
}

// bindLuxUniPKIWeb binds a generic wrapper to an already deployed contract.
func bindLuxUniPKIWeb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LuxUniPKIWebMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuxUniPKIWeb *LuxUniPKIWebRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuxUniPKIWeb.Contract.LuxUniPKIWebCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuxUniPKIWeb *LuxUniPKIWebRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.LuxUniPKIWebTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuxUniPKIWeb *LuxUniPKIWebRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.LuxUniPKIWebTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuxUniPKIWeb *LuxUniPKIWebCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuxUniPKIWeb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuxUniPKIWeb *LuxUniPKIWebTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuxUniPKIWeb *LuxUniPKIWebTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.contract.Transact(opts, method, params...)
}

// EncodeMapID is a free data retrieval call binding the contract method 0xe2ecd637.
//
// Solidity: function EncodeMapID(address _addrContr, uint96 _index) view returns(uint256 _res)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) EncodeMapID(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "EncodeMapID", _addrContr, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeMapID is a free data retrieval call binding the contract method 0xe2ecd637.
//
// Solidity: function EncodeMapID(address _addrContr, uint96 _index) view returns(uint256 _res)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) EncodeMapID(_addrContr common.Address, _index *big.Int) (*big.Int, error) {
	return _LuxUniPKIWeb.Contract.EncodeMapID(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// EncodeMapID is a free data retrieval call binding the contract method 0xe2ecd637.
//
// Solidity: function EncodeMapID(address _addrContr, uint96 _index) view returns(uint256 _res)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) EncodeMapID(_addrContr common.Address, _index *big.Int) (*big.Int, error) {
	return _LuxUniPKIWeb.Contract.EncodeMapID(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegContrAddr is a free data retrieval call binding the contract method 0xa46eb2c5.
//
// Solidity: function getRegContrAddr(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) GetRegContrAddr(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "getRegContrAddr", _addrContr, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegContrAddr is a free data retrieval call binding the contract method 0xa46eb2c5.
//
// Solidity: function getRegContrAddr(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) GetRegContrAddr(_addrContr common.Address, _index *big.Int) (common.Address, error) {
	return _LuxUniPKIWeb.Contract.GetRegContrAddr(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegContrAddr is a free data retrieval call binding the contract method 0xa46eb2c5.
//
// Solidity: function getRegContrAddr(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) GetRegContrAddr(_addrContr common.Address, _index *big.Int) (common.Address, error) {
	return _LuxUniPKIWeb.Contract.GetRegContrAddr(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegCreationDate is a free data retrieval call binding the contract method 0x06862d3e.
//
// Solidity: function getRegCreationDate(address _addrContr, uint96 _index) view returns(uint256)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) GetRegCreationDate(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "getRegCreationDate", _addrContr, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegCreationDate is a free data retrieval call binding the contract method 0x06862d3e.
//
// Solidity: function getRegCreationDate(address _addrContr, uint96 _index) view returns(uint256)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) GetRegCreationDate(_addrContr common.Address, _index *big.Int) (*big.Int, error) {
	return _LuxUniPKIWeb.Contract.GetRegCreationDate(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegCreationDate is a free data retrieval call binding the contract method 0x06862d3e.
//
// Solidity: function getRegCreationDate(address _addrContr, uint96 _index) view returns(uint256)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) GetRegCreationDate(_addrContr common.Address, _index *big.Int) (*big.Int, error) {
	return _LuxUniPKIWeb.Contract.GetRegCreationDate(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegDescription is a free data retrieval call binding the contract method 0xd2b9596b.
//
// Solidity: function getRegDescription(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) GetRegDescription(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (string, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "getRegDescription", _addrContr, _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRegDescription is a free data retrieval call binding the contract method 0xd2b9596b.
//
// Solidity: function getRegDescription(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) GetRegDescription(_addrContr common.Address, _index *big.Int) (string, error) {
	return _LuxUniPKIWeb.Contract.GetRegDescription(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegDescription is a free data retrieval call binding the contract method 0xd2b9596b.
//
// Solidity: function getRegDescription(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) GetRegDescription(_addrContr common.Address, _index *big.Int) (string, error) {
	return _LuxUniPKIWeb.Contract.GetRegDescription(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegEthAccCA is a free data retrieval call binding the contract method 0x72994190.
//
// Solidity: function getRegEthAccCA(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) GetRegEthAccCA(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "getRegEthAccCA", _addrContr, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegEthAccCA is a free data retrieval call binding the contract method 0x72994190.
//
// Solidity: function getRegEthAccCA(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) GetRegEthAccCA(_addrContr common.Address, _index *big.Int) (common.Address, error) {
	return _LuxUniPKIWeb.Contract.GetRegEthAccCA(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegEthAccCA is a free data retrieval call binding the contract method 0x72994190.
//
// Solidity: function getRegEthAccCA(address _addrContr, uint96 _index) view returns(address)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) GetRegEthAccCA(_addrContr common.Address, _index *big.Int) (common.Address, error) {
	return _LuxUniPKIWeb.Contract.GetRegEthAccCA(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegFileName is a free data retrieval call binding the contract method 0xd5735f5c.
//
// Solidity: function getRegFileName(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebCaller) GetRegFileName(opts *bind.CallOpts, _addrContr common.Address, _index *big.Int) (string, error) {
	var out []interface{}
	err := _LuxUniPKIWeb.contract.Call(opts, &out, "getRegFileName", _addrContr, _index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRegFileName is a free data retrieval call binding the contract method 0xd5735f5c.
//
// Solidity: function getRegFileName(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) GetRegFileName(_addrContr common.Address, _index *big.Int) (string, error) {
	return _LuxUniPKIWeb.Contract.GetRegFileName(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// GetRegFileName is a free data retrieval call binding the contract method 0xd5735f5c.
//
// Solidity: function getRegFileName(address _addrContr, uint96 _index) view returns(string)
func (_LuxUniPKIWeb *LuxUniPKIWebCallerSession) GetRegFileName(_addrContr common.Address, _index *big.Int) (string, error) {
	return _LuxUniPKIWeb.Contract.GetRegFileName(&_LuxUniPKIWeb.CallOpts, _addrContr, _index)
}

// NewRegDatum is a paid mutator transaction binding the contract method 0x6fd4c3b9.
//
// Solidity: function newRegDatum(address _parentAddr, uint256 _arrInd, address _ethAccCA, address _contrAddr, string _fileName, string _description) returns(uint256 err)
func (_LuxUniPKIWeb *LuxUniPKIWebTransactor) NewRegDatum(opts *bind.TransactOpts, _parentAddr common.Address, _arrInd *big.Int, _ethAccCA common.Address, _contrAddr common.Address, _fileName string, _description string) (*types.Transaction, error) {
	return _LuxUniPKIWeb.contract.Transact(opts, "newRegDatum", _parentAddr, _arrInd, _ethAccCA, _contrAddr, _fileName, _description)
}

// NewRegDatum is a paid mutator transaction binding the contract method 0x6fd4c3b9.
//
// Solidity: function newRegDatum(address _parentAddr, uint256 _arrInd, address _ethAccCA, address _contrAddr, string _fileName, string _description) returns(uint256 err)
func (_LuxUniPKIWeb *LuxUniPKIWebSession) NewRegDatum(_parentAddr common.Address, _arrInd *big.Int, _ethAccCA common.Address, _contrAddr common.Address, _fileName string, _description string) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.NewRegDatum(&_LuxUniPKIWeb.TransactOpts, _parentAddr, _arrInd, _ethAccCA, _contrAddr, _fileName, _description)
}

// NewRegDatum is a paid mutator transaction binding the contract method 0x6fd4c3b9.
//
// Solidity: function newRegDatum(address _parentAddr, uint256 _arrInd, address _ethAccCA, address _contrAddr, string _fileName, string _description) returns(uint256 err)
func (_LuxUniPKIWeb *LuxUniPKIWebTransactorSession) NewRegDatum(_parentAddr common.Address, _arrInd *big.Int, _ethAccCA common.Address, _contrAddr common.Address, _fileName string, _description string) (*types.Transaction, error) {
	return _LuxUniPKIWeb.Contract.NewRegDatum(&_LuxUniPKIWeb.TransactOpts, _parentAddr, _arrInd, _ethAccCA, _contrAddr, _fileName, _description)
}

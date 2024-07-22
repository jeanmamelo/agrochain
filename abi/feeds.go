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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LatestHumiditySensorData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LatestPressureSensorData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LatestTemperatureSensorData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50731b44f3514812d835eb1bdb0acb33d3fa3351ee435f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073694aa1769357215de4fac081bf1f309adc32530660015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507342585ed362b3f1bca95c640fdff35ef89921273460025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610444806101185f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80638a331ca914610043578063e9f54d5614610062578063ec1b5b8414610081575b5f80fd5b61004b6100a0565b6040516100599291906102e5565b60405180910390f35b61006a61014e565b6040516100789291906102e5565b60405180910390f35b6100896101fb565b6040516100979291906102e5565b60405180910390f35b5f805f805f805f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610111573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101359190610397565b9450945094509450945084849650965050505050509091565b5f805f805f805f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa1580156101be573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e29190610397565b9450945094509450945084849650965050505050509091565b5f805f805f805f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801561026c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102909190610397565b9450945094509450945084849650965050505050509091565b5f69ffffffffffffffffffff82169050919050565b6102c7816102a9565b82525050565b5f819050919050565b6102df816102cd565b82525050565b5f6040820190506102f85f8301856102be565b61030560208301846102d6565b9392505050565b5f80fd5b610319816102a9565b8114610323575f80fd5b50565b5f8151905061033481610310565b92915050565b610343816102cd565b811461034d575f80fd5b50565b5f8151905061035e8161033a565b92915050565b5f819050919050565b61037681610364565b8114610380575f80fd5b50565b5f815190506103918161036d565b92915050565b5f805f805f60a086880312156103b0576103af61030c565b5b5f6103bd88828901610326565b95505060206103ce88828901610350565b94505060406103df88828901610383565b93505060606103f088828901610383565b925050608061040188828901610326565b915050929550929590935056fea2646970667358221220940a5589c58220f1ae367e9e83e10fa377c6cf56942bcc2d9bddb29daf32fc6f64736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// LatestHumiditySensorData is a free data retrieval call binding the contract method 0x8a331ca9.
//
// Solidity: function LatestHumiditySensorData() view returns(uint80, int256)
func (_Contract *ContractCaller) LatestHumiditySensorData(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "LatestHumiditySensorData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LatestHumiditySensorData is a free data retrieval call binding the contract method 0x8a331ca9.
//
// Solidity: function LatestHumiditySensorData() view returns(uint80, int256)
func (_Contract *ContractSession) LatestHumiditySensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestHumiditySensorData(&_Contract.CallOpts)
}

// LatestHumiditySensorData is a free data retrieval call binding the contract method 0x8a331ca9.
//
// Solidity: function LatestHumiditySensorData() view returns(uint80, int256)
func (_Contract *ContractCallerSession) LatestHumiditySensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestHumiditySensorData(&_Contract.CallOpts)
}

// LatestPressureSensorData is a free data retrieval call binding the contract method 0xec1b5b84.
//
// Solidity: function LatestPressureSensorData() view returns(uint80, int256)
func (_Contract *ContractCaller) LatestPressureSensorData(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "LatestPressureSensorData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LatestPressureSensorData is a free data retrieval call binding the contract method 0xec1b5b84.
//
// Solidity: function LatestPressureSensorData() view returns(uint80, int256)
func (_Contract *ContractSession) LatestPressureSensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestPressureSensorData(&_Contract.CallOpts)
}

// LatestPressureSensorData is a free data retrieval call binding the contract method 0xec1b5b84.
//
// Solidity: function LatestPressureSensorData() view returns(uint80, int256)
func (_Contract *ContractCallerSession) LatestPressureSensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestPressureSensorData(&_Contract.CallOpts)
}

// LatestTemperatureSensorData is a free data retrieval call binding the contract method 0xe9f54d56.
//
// Solidity: function LatestTemperatureSensorData() view returns(uint80, int256)
func (_Contract *ContractCaller) LatestTemperatureSensorData(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "LatestTemperatureSensorData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LatestTemperatureSensorData is a free data retrieval call binding the contract method 0xe9f54d56.
//
// Solidity: function LatestTemperatureSensorData() view returns(uint80, int256)
func (_Contract *ContractSession) LatestTemperatureSensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestTemperatureSensorData(&_Contract.CallOpts)
}

// LatestTemperatureSensorData is a free data retrieval call binding the contract method 0xe9f54d56.
//
// Solidity: function LatestTemperatureSensorData() view returns(uint80, int256)
func (_Contract *ContractCallerSession) LatestTemperatureSensorData() (*big.Int, *big.Int, error) {
	return _Contract.Contract.LatestTemperatureSensorData(&_Contract.CallOpts)
}

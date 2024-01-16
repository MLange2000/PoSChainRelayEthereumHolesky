// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relays

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
)

// ChainRelayUpdate is an auto generated low-level Go binding around an user-defined struct.
type ChainRelayUpdate struct {
	Signature                []byte
	Participants             [512]bool
	LatestBlockRoot          [32]byte
	SigningDomain            [32]byte
	LatestSlot               uint64
	LatestSlotBranch         [][32]byte
	FinalizedBlockRoot       [32]byte
	FinalizingBranch         [][32]byte
	FinalizedSlot            uint64
	FinalizedSlotBranch      [][32]byte
	FinalizedStateRoot       [32]byte
	FinalizedStateRootBranch [][32]byte
	StateRoot                [32]byte
	StateRootBranch          [][32]byte
}

// SyncCommitteeUpdate is an auto generated low-level Go binding around an user-defined struct.
type SyncCommitteeUpdate struct {
	NextNextValidatorSet          [512][]byte
	NextNextValidatorSetAggregate []byte
	NextNextValidatorSetBranch    [][32]byte
}

// RelaysMetaData contains all meta data concerning the Relays contract.
var RelaysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_signatureThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_finalizedBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_finalizedStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_finalizedSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_latestSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_latestSlotWithValidatorSetChange\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"bitLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"_input\",\"type\":\"bytes2\"}],\"name\":\"bytes2ToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_input\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes8\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"result\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingDomain\",\"type\":\"bytes32\"}],\"name\":\"computeSigningRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[512]\",\"name\":\"_bools\",\"type\":\"bool[512]\"}],\"name\":\"countTrueBools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"fastAggregateVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"o\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"floorLog2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_pubkeys\",\"type\":\"bytes[512]\"},{\"internalType\":\"bool[512]\",\"name\":\"_isActive\",\"type\":\"bool[512]\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfActive\",\"type\":\"uint256\"}],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSlotWithValidatorSetChange\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generalizedIndex\",\"type\":\"uint256\"}],\"name\":\"getSubtreeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blspubkey\",\"type\":\"bytes\"}],\"name\":\"hashTreeRootBlspubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashTreeRootPair\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_syncCommittee\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_syncAggregate\",\"type\":\"bytes\"}],\"name\":\"hashTreeRootSyncCommittee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_currentValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_currentValidatorSetAggregate\",\"type\":\"bytes\"}],\"name\":\"initializeCurrent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_nextValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_nextValidatorSetAggregate\",\"type\":\"bytes\"}],\"name\":\"initializeNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[512]\",\"name\":\"_chunks\",\"type\":\"bytes32[512]\"}],\"name\":\"merkleize\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_input\",\"type\":\"uint64\"}],\"name\":\"merklelizeSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"nextPowOfTwo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"_input\",\"type\":\"bytes8\"}],\"name\":\"revertBytes8\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"serializeAggregateSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"slotToUnixTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool[512]\",\"name\":\"participants\",\"type\":\"bool[512]\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingDomain\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"latestSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestSlotBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizingBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizedSlotBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizedStateRootBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"stateRootBranch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structChainRelayUpdate\",\"name\":\"_chainRelayUpdate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[512]\",\"name\":\"nextNextValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"nextNextValidatorSetAggregate\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nextNextValidatorSetBranch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structSyncCommitteeUpdate\",\"name\":\"_syncCommitteeUpdate\",\"type\":\"tuple\"}],\"name\":\"submitUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_generalizedIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_branch\",\"type\":\"bytes32[]\"}],\"name\":\"validateMerkleBranch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// RelaysABI is the input ABI used to generate the binding from.
// Deprecated: Use RelaysMetaData.ABI instead.
var RelaysABI = RelaysMetaData.ABI

// Relays is an auto generated Go binding around an Ethereum contract.
type Relays struct {
	RelaysCaller     // Read-only binding to the contract
	RelaysTransactor // Write-only binding to the contract
	RelaysFilterer   // Log filterer for contract events
}

// RelaysCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelaysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelaysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelaysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaysSession struct {
	Contract     *Relays           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelaysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelaysCallerSession struct {
	Contract *RelaysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RelaysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelaysTransactorSession struct {
	Contract     *RelaysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelaysRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelaysRaw struct {
	Contract *Relays // Generic contract binding to access the raw methods on
}

// RelaysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelaysCallerRaw struct {
	Contract *RelaysCaller // Generic read-only contract binding to access the raw methods on
}

// RelaysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelaysTransactorRaw struct {
	Contract *RelaysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelays creates a new instance of Relays, bound to a specific deployed contract.
func NewRelays(address common.Address, backend bind.ContractBackend) (*Relays, error) {
	contract, err := bindRelays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relays{RelaysCaller: RelaysCaller{contract: contract}, RelaysTransactor: RelaysTransactor{contract: contract}, RelaysFilterer: RelaysFilterer{contract: contract}}, nil
}

// NewRelaysCaller creates a new read-only instance of Relays, bound to a specific deployed contract.
func NewRelaysCaller(address common.Address, caller bind.ContractCaller) (*RelaysCaller, error) {
	contract, err := bindRelays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelaysCaller{contract: contract}, nil
}

// NewRelaysTransactor creates a new write-only instance of Relays, bound to a specific deployed contract.
func NewRelaysTransactor(address common.Address, transactor bind.ContractTransactor) (*RelaysTransactor, error) {
	contract, err := bindRelays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelaysTransactor{contract: contract}, nil
}

// NewRelaysFilterer creates a new log filterer instance of Relays, bound to a specific deployed contract.
func NewRelaysFilterer(address common.Address, filterer bind.ContractFilterer) (*RelaysFilterer, error) {
	contract, err := bindRelays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelaysFilterer{contract: contract}, nil
}

// bindRelays binds a generic wrapper to an already deployed contract.
func bindRelays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelaysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relays *RelaysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relays.Contract.RelaysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relays *RelaysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relays.Contract.RelaysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relays *RelaysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relays.Contract.RelaysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relays *RelaysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relays *RelaysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relays *RelaysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relays.Contract.contract.Transact(opts, method, params...)
}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCaller) BitLength(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "bitLength", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Relays *RelaysSession) BitLength(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.BitLength(&_Relays.CallOpts, _input)
}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCallerSession) BitLength(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.BitLength(&_Relays.CallOpts, _input)
}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Relays *RelaysCaller) Bytes2ToBytes(opts *bind.CallOpts, _input [2]byte) ([]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "bytes2ToBytes", _input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Relays *RelaysSession) Bytes2ToBytes(_input [2]byte) ([]byte, error) {
	return _Relays.Contract.Bytes2ToBytes(&_Relays.CallOpts, _input)
}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Relays *RelaysCallerSession) Bytes2ToBytes(_input [2]byte) ([]byte, error) {
	return _Relays.Contract.Bytes2ToBytes(&_Relays.CallOpts, _input)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Relays *RelaysCaller) Bytes32ToBytes(opts *bind.CallOpts, _input [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "bytes32ToBytes", _input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Relays *RelaysSession) Bytes32ToBytes(_input [32]byte) ([]byte, error) {
	return _Relays.Contract.Bytes32ToBytes(&_Relays.CallOpts, _input)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Relays *RelaysCallerSession) Bytes32ToBytes(_input [32]byte) ([]byte, error) {
	return _Relays.Contract.Bytes32ToBytes(&_Relays.CallOpts, _input)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Relays *RelaysCaller) BytesToBytes32(opts *bind.CallOpts, _input []byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "bytesToBytes32", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Relays *RelaysSession) BytesToBytes32(_input []byte) ([32]byte, error) {
	return _Relays.Contract.BytesToBytes32(&_Relays.CallOpts, _input)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Relays *RelaysCallerSession) BytesToBytes32(_input []byte) ([32]byte, error) {
	return _Relays.Contract.BytesToBytes32(&_Relays.CallOpts, _input)
}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Relays *RelaysCaller) BytesToBytes8(opts *bind.CallOpts, _input []byte) ([8]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "bytesToBytes8", _input)

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Relays *RelaysSession) BytesToBytes8(_input []byte) ([8]byte, error) {
	return _Relays.Contract.BytesToBytes8(&_Relays.CallOpts, _input)
}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Relays *RelaysCallerSession) BytesToBytes8(_input []byte) ([8]byte, error) {
	return _Relays.Contract.BytesToBytes8(&_Relays.CallOpts, _input)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Relays *RelaysCaller) ComputeSigningRoot(opts *bind.CallOpts, _blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "computeSigningRoot", _blockRoot, _signingDomain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Relays *RelaysSession) ComputeSigningRoot(_blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	return _Relays.Contract.ComputeSigningRoot(&_Relays.CallOpts, _blockRoot, _signingDomain)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Relays *RelaysCallerSession) ComputeSigningRoot(_blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	return _Relays.Contract.ComputeSigningRoot(&_Relays.CallOpts, _blockRoot, _signingDomain)
}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Relays *RelaysCaller) Concat(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "concat", _left, _right)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Relays *RelaysSession) Concat(_left [32]byte, _right [32]byte) ([]byte, error) {
	return _Relays.Contract.Concat(&_Relays.CallOpts, _left, _right)
}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Relays *RelaysCallerSession) Concat(_left [32]byte, _right [32]byte) ([]byte, error) {
	return _Relays.Contract.Concat(&_Relays.CallOpts, _left, _right)
}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Relays *RelaysCaller) CountTrueBools(opts *bind.CallOpts, _bools [512]bool) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "countTrueBools", _bools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Relays *RelaysSession) CountTrueBools(_bools [512]bool) (*big.Int, error) {
	return _Relays.Contract.CountTrueBools(&_Relays.CallOpts, _bools)
}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Relays *RelaysCallerSession) CountTrueBools(_bools [512]bool) (*big.Int, error) {
	return _Relays.Contract.CountTrueBools(&_Relays.CallOpts, _bools)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Relays *RelaysCaller) FastAggregateVerify(opts *bind.CallOpts, _input []byte) (bool, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "fastAggregateVerify", _input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Relays *RelaysSession) FastAggregateVerify(_input []byte) (bool, error) {
	return _Relays.Contract.FastAggregateVerify(&_Relays.CallOpts, _input)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Relays *RelaysCallerSession) FastAggregateVerify(_input []byte) (bool, error) {
	return _Relays.Contract.FastAggregateVerify(&_Relays.CallOpts, _input)
}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCaller) FloorLog2(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "floorLog2", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Relays *RelaysSession) FloorLog2(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.FloorLog2(&_Relays.CallOpts, _input)
}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCallerSession) FloorLog2(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.FloorLog2(&_Relays.CallOpts, _input)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Relays *RelaysCaller) GetActiveValidators(opts *bind.CallOpts, _pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getActiveValidators", _pubkeys, _isActive, _numberOfActive)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Relays *RelaysSession) GetActiveValidators(_pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	return _Relays.Contract.GetActiveValidators(&_Relays.CallOpts, _pubkeys, _isActive, _numberOfActive)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Relays *RelaysCallerSession) GetActiveValidators(_pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	return _Relays.Contract.GetActiveValidators(&_Relays.CallOpts, _pubkeys, _isActive, _numberOfActive)
}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Relays *RelaysCaller) GetFinalizedBlockRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getFinalizedBlockRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Relays *RelaysSession) GetFinalizedBlockRoot() ([32]byte, error) {
	return _Relays.Contract.GetFinalizedBlockRoot(&_Relays.CallOpts)
}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Relays *RelaysCallerSession) GetFinalizedBlockRoot() ([32]byte, error) {
	return _Relays.Contract.GetFinalizedBlockRoot(&_Relays.CallOpts)
}

// GetFinalizedSlot is a free data retrieval call binding the contract method 0x8a83458d.
//
// Solidity: function getFinalizedSlot() view returns(uint64)
func (_Relays *RelaysCaller) GetFinalizedSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getFinalizedSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFinalizedSlot is a free data retrieval call binding the contract method 0x8a83458d.
//
// Solidity: function getFinalizedSlot() view returns(uint64)
func (_Relays *RelaysSession) GetFinalizedSlot() (uint64, error) {
	return _Relays.Contract.GetFinalizedSlot(&_Relays.CallOpts)
}

// GetFinalizedSlot is a free data retrieval call binding the contract method 0x8a83458d.
//
// Solidity: function getFinalizedSlot() view returns(uint64)
func (_Relays *RelaysCallerSession) GetFinalizedSlot() (uint64, error) {
	return _Relays.Contract.GetFinalizedSlot(&_Relays.CallOpts)
}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Relays *RelaysCaller) GetFinalizedStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getFinalizedStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Relays *RelaysSession) GetFinalizedStateRoot() ([32]byte, error) {
	return _Relays.Contract.GetFinalizedStateRoot(&_Relays.CallOpts)
}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Relays *RelaysCallerSession) GetFinalizedStateRoot() ([32]byte, error) {
	return _Relays.Contract.GetFinalizedStateRoot(&_Relays.CallOpts)
}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Relays *RelaysCaller) GetLatestSlotWithValidatorSetChange(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getLatestSlotWithValidatorSetChange")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Relays *RelaysSession) GetLatestSlotWithValidatorSetChange() (uint64, error) {
	return _Relays.Contract.GetLatestSlotWithValidatorSetChange(&_Relays.CallOpts)
}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Relays *RelaysCallerSession) GetLatestSlotWithValidatorSetChange() (uint64, error) {
	return _Relays.Contract.GetLatestSlotWithValidatorSetChange(&_Relays.CallOpts)
}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Relays *RelaysCaller) GetSubtreeIndex(opts *bind.CallOpts, _generalizedIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "getSubtreeIndex", _generalizedIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Relays *RelaysSession) GetSubtreeIndex(_generalizedIndex *big.Int) (*big.Int, error) {
	return _Relays.Contract.GetSubtreeIndex(&_Relays.CallOpts, _generalizedIndex)
}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Relays *RelaysCallerSession) GetSubtreeIndex(_generalizedIndex *big.Int) (*big.Int, error) {
	return _Relays.Contract.GetSubtreeIndex(&_Relays.CallOpts, _generalizedIndex)
}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Relays *RelaysCaller) HashTreeRootBlspubkey(opts *bind.CallOpts, _blspubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "hashTreeRootBlspubkey", _blspubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Relays *RelaysSession) HashTreeRootBlspubkey(_blspubkey []byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootBlspubkey(&_Relays.CallOpts, _blspubkey)
}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Relays *RelaysCallerSession) HashTreeRootBlspubkey(_blspubkey []byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootBlspubkey(&_Relays.CallOpts, _blspubkey)
}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Relays *RelaysCaller) HashTreeRootPair(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "hashTreeRootPair", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Relays *RelaysSession) HashTreeRootPair(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootPair(&_Relays.CallOpts, _left, _right)
}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Relays *RelaysCallerSession) HashTreeRootPair(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootPair(&_Relays.CallOpts, _left, _right)
}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Relays *RelaysCaller) HashTreeRootSyncCommittee(opts *bind.CallOpts, _syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "hashTreeRootSyncCommittee", _syncCommittee, _syncAggregate)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Relays *RelaysSession) HashTreeRootSyncCommittee(_syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootSyncCommittee(&_Relays.CallOpts, _syncCommittee, _syncAggregate)
}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Relays *RelaysCallerSession) HashTreeRootSyncCommittee(_syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	return _Relays.Contract.HashTreeRootSyncCommittee(&_Relays.CallOpts, _syncCommittee, _syncAggregate)
}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Relays *RelaysCaller) Merkleize(opts *bind.CallOpts, _chunks [512][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "merkleize", _chunks)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Relays *RelaysSession) Merkleize(_chunks [512][32]byte) ([32]byte, error) {
	return _Relays.Contract.Merkleize(&_Relays.CallOpts, _chunks)
}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Relays *RelaysCallerSession) Merkleize(_chunks [512][32]byte) ([32]byte, error) {
	return _Relays.Contract.Merkleize(&_Relays.CallOpts, _chunks)
}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Relays *RelaysCaller) MerklelizeSlot(opts *bind.CallOpts, _input uint64) ([32]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "merklelizeSlot", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Relays *RelaysSession) MerklelizeSlot(_input uint64) ([32]byte, error) {
	return _Relays.Contract.MerklelizeSlot(&_Relays.CallOpts, _input)
}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Relays *RelaysCallerSession) MerklelizeSlot(_input uint64) ([32]byte, error) {
	return _Relays.Contract.MerklelizeSlot(&_Relays.CallOpts, _input)
}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCaller) NextPowOfTwo(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "nextPowOfTwo", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Relays *RelaysSession) NextPowOfTwo(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.NextPowOfTwo(&_Relays.CallOpts, _input)
}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Relays *RelaysCallerSession) NextPowOfTwo(_input *big.Int) (*big.Int, error) {
	return _Relays.Contract.NextPowOfTwo(&_Relays.CallOpts, _input)
}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Relays *RelaysCaller) RevertBytes8(opts *bind.CallOpts, _input [8]byte) ([8]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "revertBytes8", _input)

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Relays *RelaysSession) RevertBytes8(_input [8]byte) ([8]byte, error) {
	return _Relays.Contract.RevertBytes8(&_Relays.CallOpts, _input)
}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Relays *RelaysCallerSession) RevertBytes8(_input [8]byte) ([8]byte, error) {
	return _Relays.Contract.RevertBytes8(&_Relays.CallOpts, _input)
}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Relays *RelaysCaller) SerializeAggregateSignature(opts *bind.CallOpts, _message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "serializeAggregateSignature", _message, _signature, _pubkeys)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Relays *RelaysSession) SerializeAggregateSignature(_message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	return _Relays.Contract.SerializeAggregateSignature(&_Relays.CallOpts, _message, _signature, _pubkeys)
}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Relays *RelaysCallerSession) SerializeAggregateSignature(_message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	return _Relays.Contract.SerializeAggregateSignature(&_Relays.CallOpts, _message, _signature, _pubkeys)
}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Relays *RelaysCaller) SlotToUnixTimestamp(opts *bind.CallOpts, _slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "slotToUnixTimestamp", _slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Relays *RelaysSession) SlotToUnixTimestamp(_slot *big.Int) (*big.Int, error) {
	return _Relays.Contract.SlotToUnixTimestamp(&_Relays.CallOpts, _slot)
}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Relays *RelaysCallerSession) SlotToUnixTimestamp(_slot *big.Int) (*big.Int, error) {
	return _Relays.Contract.SlotToUnixTimestamp(&_Relays.CallOpts, _slot)
}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Relays *RelaysCaller) ValidateMerkleBranch(opts *bind.CallOpts, _root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	var out []interface{}
	err := _Relays.contract.Call(opts, &out, "validateMerkleBranch", _root, _leaf, _generalizedIndex, _branch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Relays *RelaysSession) ValidateMerkleBranch(_root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	return _Relays.Contract.ValidateMerkleBranch(&_Relays.CallOpts, _root, _leaf, _generalizedIndex, _branch)
}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Relays *RelaysCallerSession) ValidateMerkleBranch(_root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	return _Relays.Contract.ValidateMerkleBranch(&_Relays.CallOpts, _root, _leaf, _generalizedIndex, _branch)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Relays *RelaysTransactor) InitializeCurrent(opts *bind.TransactOpts, _currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.contract.Transact(opts, "initializeCurrent", _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Relays *RelaysSession) InitializeCurrent(_currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.Contract.InitializeCurrent(&_Relays.TransactOpts, _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Relays *RelaysTransactorSession) InitializeCurrent(_currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.Contract.InitializeCurrent(&_Relays.TransactOpts, _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Relays *RelaysTransactor) InitializeNext(opts *bind.TransactOpts, _nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.contract.Transact(opts, "initializeNext", _nextValidatorSet, _nextValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Relays *RelaysSession) InitializeNext(_nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.Contract.InitializeNext(&_Relays.TransactOpts, _nextValidatorSet, _nextValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Relays *RelaysTransactorSession) InitializeNext(_nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Relays.Contract.InitializeNext(&_Relays.TransactOpts, _nextValidatorSet, _nextValidatorSetAggregate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Relays *RelaysTransactor) SubmitUpdate(opts *bind.TransactOpts, _chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Relays.contract.Transact(opts, "submitUpdate", _chainRelayUpdate, _syncCommitteeUpdate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Relays *RelaysSession) SubmitUpdate(_chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Relays.Contract.SubmitUpdate(&_Relays.TransactOpts, _chainRelayUpdate, _syncCommitteeUpdate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Relays *RelaysTransactorSession) SubmitUpdate(_chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Relays.Contract.SubmitUpdate(&_Relays.TransactOpts, _chainRelayUpdate, _syncCommitteeUpdate)
}

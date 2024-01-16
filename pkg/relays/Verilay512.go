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

// VerilayMetaData contains all meta data concerning the Verilay contract.
var VerilayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_signatureThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_finalizedBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_finalizedStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_finalizedSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_latestSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_latestSlotWithValidatorSetChange\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"bitLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"_input\",\"type\":\"bytes2\"}],\"name\":\"bytes2ToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_input\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes8\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"result\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingDomain\",\"type\":\"bytes32\"}],\"name\":\"computeSigningRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[512]\",\"name\":\"_bools\",\"type\":\"bool[512]\"}],\"name\":\"countTrueBools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"fastAggregateVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"o\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"floorLog2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_pubkeys\",\"type\":\"bytes[512]\"},{\"internalType\":\"bool[512]\",\"name\":\"_isActive\",\"type\":\"bool[512]\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfActive\",\"type\":\"uint256\"}],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSlotWithValidatorSetChange\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generalizedIndex\",\"type\":\"uint256\"}],\"name\":\"getSubtreeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blspubkey\",\"type\":\"bytes\"}],\"name\":\"hashTreeRootBlspubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashTreeRootPair\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_syncCommittee\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_syncAggregate\",\"type\":\"bytes\"}],\"name\":\"hashTreeRootSyncCommittee\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_currentValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_currentValidatorSetAggregate\",\"type\":\"bytes\"}],\"name\":\"initializeCurrent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[512]\",\"name\":\"_nextValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"_nextValidatorSetAggregate\",\"type\":\"bytes\"}],\"name\":\"initializeNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[512]\",\"name\":\"_chunks\",\"type\":\"bytes32[512]\"}],\"name\":\"merkleize\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_input\",\"type\":\"uint64\"}],\"name\":\"merklelizeSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_input\",\"type\":\"uint256\"}],\"name\":\"nextPowOfTwo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"_input\",\"type\":\"bytes8\"}],\"name\":\"revertBytes8\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"serializeAggregateSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"slotToUnixTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool[512]\",\"name\":\"participants\",\"type\":\"bool[512]\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingDomain\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"latestSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestSlotBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedBlockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizingBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizedSlotBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"finalizedStateRootBranch\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"stateRootBranch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structChainRelayUpdate\",\"name\":\"_chainRelayUpdate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[512]\",\"name\":\"nextNextValidatorSet\",\"type\":\"bytes[512]\"},{\"internalType\":\"bytes\",\"name\":\"nextNextValidatorSetAggregate\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nextNextValidatorSetBranch\",\"type\":\"bytes32[]\"}],\"internalType\":\"structSyncCommitteeUpdate\",\"name\":\"_syncCommitteeUpdate\",\"type\":\"tuple\"}],\"name\":\"submitUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_generalizedIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_branch\",\"type\":\"bytes32[]\"}],\"name\":\"validateMerkleBranch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VerilayABI is the input ABI used to generate the binding from.
// Deprecated: Use VerilayMetaData.ABI instead.
var VerilayABI = VerilayMetaData.ABI

// Verilay is an auto generated Go binding around an Ethereum contract.
type Verilay struct {
	VerilayCaller     // Read-only binding to the contract
	VerilayTransactor // Write-only binding to the contract
	VerilayFilterer   // Log filterer for contract events
}

// VerilayCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerilayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerilayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerilayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerilayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerilayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerilaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerilaySession struct {
	Contract     *Verilay          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerilayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerilayCallerSession struct {
	Contract *VerilayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VerilayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerilayTransactorSession struct {
	Contract     *VerilayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VerilayRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerilayRaw struct {
	Contract *Verilay // Generic contract binding to access the raw methods on
}

// VerilayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerilayCallerRaw struct {
	Contract *VerilayCaller // Generic read-only contract binding to access the raw methods on
}

// VerilayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerilayTransactorRaw struct {
	Contract *VerilayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerilay creates a new instance of Verilay, bound to a specific deployed contract.
func NewVerilay(address common.Address, backend bind.ContractBackend) (*Verilay, error) {
	contract, err := bindVerilay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verilay{VerilayCaller: VerilayCaller{contract: contract}, VerilayTransactor: VerilayTransactor{contract: contract}, VerilayFilterer: VerilayFilterer{contract: contract}}, nil
}

// NewVerilayCaller creates a new read-only instance of Verilay, bound to a specific deployed contract.
func NewVerilayCaller(address common.Address, caller bind.ContractCaller) (*VerilayCaller, error) {
	contract, err := bindVerilay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerilayCaller{contract: contract}, nil
}

// NewVerilayTransactor creates a new write-only instance of Verilay, bound to a specific deployed contract.
func NewVerilayTransactor(address common.Address, transactor bind.ContractTransactor) (*VerilayTransactor, error) {
	contract, err := bindVerilay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerilayTransactor{contract: contract}, nil
}

// NewVerilayFilterer creates a new log filterer instance of Verilay, bound to a specific deployed contract.
func NewVerilayFilterer(address common.Address, filterer bind.ContractFilterer) (*VerilayFilterer, error) {
	contract, err := bindVerilay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerilayFilterer{contract: contract}, nil
}

// bindVerilay binds a generic wrapper to an already deployed contract.
func bindVerilay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerilayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verilay *VerilayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verilay.Contract.VerilayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verilay *VerilayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verilay.Contract.VerilayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verilay *VerilayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verilay.Contract.VerilayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verilay *VerilayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verilay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verilay *VerilayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verilay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verilay *VerilayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verilay.Contract.contract.Transact(opts, method, params...)
}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCaller) BitLength(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "bitLength", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Verilay *VerilaySession) BitLength(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.BitLength(&_Verilay.CallOpts, _input)
}

// BitLength is a free data retrieval call binding the contract method 0xe41bbcf8.
//
// Solidity: function bitLength(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCallerSession) BitLength(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.BitLength(&_Verilay.CallOpts, _input)
}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Verilay *VerilayCaller) Bytes2ToBytes(opts *bind.CallOpts, _input [2]byte) ([]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "bytes2ToBytes", _input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Verilay *VerilaySession) Bytes2ToBytes(_input [2]byte) ([]byte, error) {
	return _Verilay.Contract.Bytes2ToBytes(&_Verilay.CallOpts, _input)
}

// Bytes2ToBytes is a free data retrieval call binding the contract method 0xa9755eb1.
//
// Solidity: function bytes2ToBytes(bytes2 _input) pure returns(bytes output)
func (_Verilay *VerilayCallerSession) Bytes2ToBytes(_input [2]byte) ([]byte, error) {
	return _Verilay.Contract.Bytes2ToBytes(&_Verilay.CallOpts, _input)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Verilay *VerilayCaller) Bytes32ToBytes(opts *bind.CallOpts, _input [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "bytes32ToBytes", _input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Verilay *VerilaySession) Bytes32ToBytes(_input [32]byte) ([]byte, error) {
	return _Verilay.Contract.Bytes32ToBytes(&_Verilay.CallOpts, _input)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 _input) pure returns(bytes output)
func (_Verilay *VerilayCallerSession) Bytes32ToBytes(_input [32]byte) ([]byte, error) {
	return _Verilay.Contract.Bytes32ToBytes(&_Verilay.CallOpts, _input)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Verilay *VerilayCaller) BytesToBytes32(opts *bind.CallOpts, _input []byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "bytesToBytes32", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Verilay *VerilaySession) BytesToBytes32(_input []byte) ([32]byte, error) {
	return _Verilay.Contract.BytesToBytes32(&_Verilay.CallOpts, _input)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes _input) pure returns(bytes32 result)
func (_Verilay *VerilayCallerSession) BytesToBytes32(_input []byte) ([32]byte, error) {
	return _Verilay.Contract.BytesToBytes32(&_Verilay.CallOpts, _input)
}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Verilay *VerilayCaller) BytesToBytes8(opts *bind.CallOpts, _input []byte) ([8]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "bytesToBytes8", _input)

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Verilay *VerilaySession) BytesToBytes8(_input []byte) ([8]byte, error) {
	return _Verilay.Contract.BytesToBytes8(&_Verilay.CallOpts, _input)
}

// BytesToBytes8 is a free data retrieval call binding the contract method 0x5c552879.
//
// Solidity: function bytesToBytes8(bytes _input) pure returns(bytes8 result)
func (_Verilay *VerilayCallerSession) BytesToBytes8(_input []byte) ([8]byte, error) {
	return _Verilay.Contract.BytesToBytes8(&_Verilay.CallOpts, _input)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Verilay *VerilayCaller) ComputeSigningRoot(opts *bind.CallOpts, _blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "computeSigningRoot", _blockRoot, _signingDomain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Verilay *VerilaySession) ComputeSigningRoot(_blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	return _Verilay.Contract.ComputeSigningRoot(&_Verilay.CallOpts, _blockRoot, _signingDomain)
}

// ComputeSigningRoot is a free data retrieval call binding the contract method 0xe92938ab.
//
// Solidity: function computeSigningRoot(bytes32 _blockRoot, bytes32 _signingDomain) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) ComputeSigningRoot(_blockRoot [32]byte, _signingDomain [32]byte) ([32]byte, error) {
	return _Verilay.Contract.ComputeSigningRoot(&_Verilay.CallOpts, _blockRoot, _signingDomain)
}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Verilay *VerilayCaller) Concat(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "concat", _left, _right)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Verilay *VerilaySession) Concat(_left [32]byte, _right [32]byte) ([]byte, error) {
	return _Verilay.Contract.Concat(&_Verilay.CallOpts, _left, _right)
}

// Concat is a free data retrieval call binding the contract method 0x7b9e6200.
//
// Solidity: function concat(bytes32 _left, bytes32 _right) pure returns(bytes)
func (_Verilay *VerilayCallerSession) Concat(_left [32]byte, _right [32]byte) ([]byte, error) {
	return _Verilay.Contract.Concat(&_Verilay.CallOpts, _left, _right)
}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Verilay *VerilayCaller) CountTrueBools(opts *bind.CallOpts, _bools [512]bool) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "countTrueBools", _bools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Verilay *VerilaySession) CountTrueBools(_bools [512]bool) (*big.Int, error) {
	return _Verilay.Contract.CountTrueBools(&_Verilay.CallOpts, _bools)
}

// CountTrueBools is a free data retrieval call binding the contract method 0x809d74ba.
//
// Solidity: function countTrueBools(bool[512] _bools) pure returns(uint256)
func (_Verilay *VerilayCallerSession) CountTrueBools(_bools [512]bool) (*big.Int, error) {
	return _Verilay.Contract.CountTrueBools(&_Verilay.CallOpts, _bools)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Verilay *VerilayCaller) FastAggregateVerify(opts *bind.CallOpts, _input []byte) (bool, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "fastAggregateVerify", _input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Verilay *VerilaySession) FastAggregateVerify(_input []byte) (bool, error) {
	return _Verilay.Contract.FastAggregateVerify(&_Verilay.CallOpts, _input)
}

// FastAggregateVerify is a free data retrieval call binding the contract method 0x398fd9d4.
//
// Solidity: function fastAggregateVerify(bytes _input) view returns(bool o)
func (_Verilay *VerilayCallerSession) FastAggregateVerify(_input []byte) (bool, error) {
	return _Verilay.Contract.FastAggregateVerify(&_Verilay.CallOpts, _input)
}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCaller) FloorLog2(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "floorLog2", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Verilay *VerilaySession) FloorLog2(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.FloorLog2(&_Verilay.CallOpts, _input)
}

// FloorLog2 is a free data retrieval call binding the contract method 0x45b8bafc.
//
// Solidity: function floorLog2(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCallerSession) FloorLog2(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.FloorLog2(&_Verilay.CallOpts, _input)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Verilay *VerilayCaller) GetActiveValidators(opts *bind.CallOpts, _pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "getActiveValidators", _pubkeys, _isActive, _numberOfActive)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Verilay *VerilaySession) GetActiveValidators(_pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	return _Verilay.Contract.GetActiveValidators(&_Verilay.CallOpts, _pubkeys, _isActive, _numberOfActive)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0xfe24511e.
//
// Solidity: function getActiveValidators(bytes[512] _pubkeys, bool[512] _isActive, uint256 _numberOfActive) pure returns(bytes[])
func (_Verilay *VerilayCallerSession) GetActiveValidators(_pubkeys [512][]byte, _isActive [512]bool, _numberOfActive *big.Int) ([][]byte, error) {
	return _Verilay.Contract.GetActiveValidators(&_Verilay.CallOpts, _pubkeys, _isActive, _numberOfActive)
}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Verilay *VerilayCaller) GetFinalizedBlockRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "getFinalizedBlockRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Verilay *VerilaySession) GetFinalizedBlockRoot() ([32]byte, error) {
	return _Verilay.Contract.GetFinalizedBlockRoot(&_Verilay.CallOpts)
}

// GetFinalizedBlockRoot is a free data retrieval call binding the contract method 0x2c26bfd4.
//
// Solidity: function getFinalizedBlockRoot() view returns(bytes32)
func (_Verilay *VerilayCallerSession) GetFinalizedBlockRoot() ([32]byte, error) {
	return _Verilay.Contract.GetFinalizedBlockRoot(&_Verilay.CallOpts)
}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Verilay *VerilayCaller) GetFinalizedStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "getFinalizedStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Verilay *VerilaySession) GetFinalizedStateRoot() ([32]byte, error) {
	return _Verilay.Contract.GetFinalizedStateRoot(&_Verilay.CallOpts)
}

// GetFinalizedStateRoot is a free data retrieval call binding the contract method 0x56f07c06.
//
// Solidity: function getFinalizedStateRoot() view returns(bytes32)
func (_Verilay *VerilayCallerSession) GetFinalizedStateRoot() ([32]byte, error) {
	return _Verilay.Contract.GetFinalizedStateRoot(&_Verilay.CallOpts)
}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Verilay *VerilayCaller) GetLatestSlotWithValidatorSetChange(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "getLatestSlotWithValidatorSetChange")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Verilay *VerilaySession) GetLatestSlotWithValidatorSetChange() (uint64, error) {
	return _Verilay.Contract.GetLatestSlotWithValidatorSetChange(&_Verilay.CallOpts)
}

// GetLatestSlotWithValidatorSetChange is a free data retrieval call binding the contract method 0x639f30c3.
//
// Solidity: function getLatestSlotWithValidatorSetChange() view returns(uint64)
func (_Verilay *VerilayCallerSession) GetLatestSlotWithValidatorSetChange() (uint64, error) {
	return _Verilay.Contract.GetLatestSlotWithValidatorSetChange(&_Verilay.CallOpts)
}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Verilay *VerilayCaller) GetSubtreeIndex(opts *bind.CallOpts, _generalizedIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "getSubtreeIndex", _generalizedIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Verilay *VerilaySession) GetSubtreeIndex(_generalizedIndex *big.Int) (*big.Int, error) {
	return _Verilay.Contract.GetSubtreeIndex(&_Verilay.CallOpts, _generalizedIndex)
}

// GetSubtreeIndex is a free data retrieval call binding the contract method 0xb1e9f2a0.
//
// Solidity: function getSubtreeIndex(uint256 _generalizedIndex) pure returns(uint256)
func (_Verilay *VerilayCallerSession) GetSubtreeIndex(_generalizedIndex *big.Int) (*big.Int, error) {
	return _Verilay.Contract.GetSubtreeIndex(&_Verilay.CallOpts, _generalizedIndex)
}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Verilay *VerilayCaller) HashTreeRootBlspubkey(opts *bind.CallOpts, _blspubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "hashTreeRootBlspubkey", _blspubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Verilay *VerilaySession) HashTreeRootBlspubkey(_blspubkey []byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootBlspubkey(&_Verilay.CallOpts, _blspubkey)
}

// HashTreeRootBlspubkey is a free data retrieval call binding the contract method 0xc36a13da.
//
// Solidity: function hashTreeRootBlspubkey(bytes _blspubkey) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) HashTreeRootBlspubkey(_blspubkey []byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootBlspubkey(&_Verilay.CallOpts, _blspubkey)
}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Verilay *VerilayCaller) HashTreeRootPair(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "hashTreeRootPair", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Verilay *VerilaySession) HashTreeRootPair(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootPair(&_Verilay.CallOpts, _left, _right)
}

// HashTreeRootPair is a free data retrieval call binding the contract method 0x9eb51fce.
//
// Solidity: function hashTreeRootPair(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) HashTreeRootPair(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootPair(&_Verilay.CallOpts, _left, _right)
}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Verilay *VerilayCaller) HashTreeRootSyncCommittee(opts *bind.CallOpts, _syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "hashTreeRootSyncCommittee", _syncCommittee, _syncAggregate)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Verilay *VerilaySession) HashTreeRootSyncCommittee(_syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootSyncCommittee(&_Verilay.CallOpts, _syncCommittee, _syncAggregate)
}

// HashTreeRootSyncCommittee is a free data retrieval call binding the contract method 0x255f3791.
//
// Solidity: function hashTreeRootSyncCommittee(bytes[512] _syncCommittee, bytes _syncAggregate) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) HashTreeRootSyncCommittee(_syncCommittee [512][]byte, _syncAggregate []byte) ([32]byte, error) {
	return _Verilay.Contract.HashTreeRootSyncCommittee(&_Verilay.CallOpts, _syncCommittee, _syncAggregate)
}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Verilay *VerilayCaller) Merkleize(opts *bind.CallOpts, _chunks [512][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "merkleize", _chunks)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Verilay *VerilaySession) Merkleize(_chunks [512][32]byte) ([32]byte, error) {
	return _Verilay.Contract.Merkleize(&_Verilay.CallOpts, _chunks)
}

// Merkleize is a free data retrieval call binding the contract method 0x2b224751.
//
// Solidity: function merkleize(bytes32[512] _chunks) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) Merkleize(_chunks [512][32]byte) ([32]byte, error) {
	return _Verilay.Contract.Merkleize(&_Verilay.CallOpts, _chunks)
}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Verilay *VerilayCaller) MerklelizeSlot(opts *bind.CallOpts, _input uint64) ([32]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "merklelizeSlot", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Verilay *VerilaySession) MerklelizeSlot(_input uint64) ([32]byte, error) {
	return _Verilay.Contract.MerklelizeSlot(&_Verilay.CallOpts, _input)
}

// MerklelizeSlot is a free data retrieval call binding the contract method 0x404c2a71.
//
// Solidity: function merklelizeSlot(uint64 _input) pure returns(bytes32)
func (_Verilay *VerilayCallerSession) MerklelizeSlot(_input uint64) ([32]byte, error) {
	return _Verilay.Contract.MerklelizeSlot(&_Verilay.CallOpts, _input)
}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCaller) NextPowOfTwo(opts *bind.CallOpts, _input *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "nextPowOfTwo", _input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Verilay *VerilaySession) NextPowOfTwo(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.NextPowOfTwo(&_Verilay.CallOpts, _input)
}

// NextPowOfTwo is a free data retrieval call binding the contract method 0xb8ffb2e3.
//
// Solidity: function nextPowOfTwo(uint256 _input) pure returns(uint256)
func (_Verilay *VerilayCallerSession) NextPowOfTwo(_input *big.Int) (*big.Int, error) {
	return _Verilay.Contract.NextPowOfTwo(&_Verilay.CallOpts, _input)
}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Verilay *VerilayCaller) RevertBytes8(opts *bind.CallOpts, _input [8]byte) ([8]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "revertBytes8", _input)

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Verilay *VerilaySession) RevertBytes8(_input [8]byte) ([8]byte, error) {
	return _Verilay.Contract.RevertBytes8(&_Verilay.CallOpts, _input)
}

// RevertBytes8 is a free data retrieval call binding the contract method 0x6e1e76b9.
//
// Solidity: function revertBytes8(bytes8 _input) pure returns(bytes8)
func (_Verilay *VerilayCallerSession) RevertBytes8(_input [8]byte) ([8]byte, error) {
	return _Verilay.Contract.RevertBytes8(&_Verilay.CallOpts, _input)
}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Verilay *VerilayCaller) SerializeAggregateSignature(opts *bind.CallOpts, _message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "serializeAggregateSignature", _message, _signature, _pubkeys)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Verilay *VerilaySession) SerializeAggregateSignature(_message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	return _Verilay.Contract.SerializeAggregateSignature(&_Verilay.CallOpts, _message, _signature, _pubkeys)
}

// SerializeAggregateSignature is a free data retrieval call binding the contract method 0x7e7f232a.
//
// Solidity: function serializeAggregateSignature(bytes32 _message, bytes _signature, bytes[] _pubkeys) pure returns(bytes)
func (_Verilay *VerilayCallerSession) SerializeAggregateSignature(_message [32]byte, _signature []byte, _pubkeys [][]byte) ([]byte, error) {
	return _Verilay.Contract.SerializeAggregateSignature(&_Verilay.CallOpts, _message, _signature, _pubkeys)
}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Verilay *VerilayCaller) SlotToUnixTimestamp(opts *bind.CallOpts, _slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "slotToUnixTimestamp", _slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Verilay *VerilaySession) SlotToUnixTimestamp(_slot *big.Int) (*big.Int, error) {
	return _Verilay.Contract.SlotToUnixTimestamp(&_Verilay.CallOpts, _slot)
}

// SlotToUnixTimestamp is a free data retrieval call binding the contract method 0x3fccb97b.
//
// Solidity: function slotToUnixTimestamp(uint256 _slot) pure returns(uint256 timestamp)
func (_Verilay *VerilayCallerSession) SlotToUnixTimestamp(_slot *big.Int) (*big.Int, error) {
	return _Verilay.Contract.SlotToUnixTimestamp(&_Verilay.CallOpts, _slot)
}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Verilay *VerilayCaller) ValidateMerkleBranch(opts *bind.CallOpts, _root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	var out []interface{}
	err := _Verilay.contract.Call(opts, &out, "validateMerkleBranch", _root, _leaf, _generalizedIndex, _branch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Verilay *VerilaySession) ValidateMerkleBranch(_root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	return _Verilay.Contract.ValidateMerkleBranch(&_Verilay.CallOpts, _root, _leaf, _generalizedIndex, _branch)
}

// ValidateMerkleBranch is a free data retrieval call binding the contract method 0xf218850e.
//
// Solidity: function validateMerkleBranch(bytes32 _root, bytes32 _leaf, uint256 _generalizedIndex, bytes32[] _branch) pure returns(bool)
func (_Verilay *VerilayCallerSession) ValidateMerkleBranch(_root [32]byte, _leaf [32]byte, _generalizedIndex *big.Int, _branch [][32]byte) (bool, error) {
	return _Verilay.Contract.ValidateMerkleBranch(&_Verilay.CallOpts, _root, _leaf, _generalizedIndex, _branch)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Verilay *VerilayTransactor) InitializeCurrent(opts *bind.TransactOpts, _currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.contract.Transact(opts, "initializeCurrent", _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Verilay *VerilaySession) InitializeCurrent(_currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.Contract.InitializeCurrent(&_Verilay.TransactOpts, _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeCurrent is a paid mutator transaction binding the contract method 0x201112e6.
//
// Solidity: function initializeCurrent(bytes[512] _currentValidatorSet, bytes _currentValidatorSetAggregate) returns()
func (_Verilay *VerilayTransactorSession) InitializeCurrent(_currentValidatorSet [512][]byte, _currentValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.Contract.InitializeCurrent(&_Verilay.TransactOpts, _currentValidatorSet, _currentValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Verilay *VerilayTransactor) InitializeNext(opts *bind.TransactOpts, _nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.contract.Transact(opts, "initializeNext", _nextValidatorSet, _nextValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Verilay *VerilaySession) InitializeNext(_nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.Contract.InitializeNext(&_Verilay.TransactOpts, _nextValidatorSet, _nextValidatorSetAggregate)
}

// InitializeNext is a paid mutator transaction binding the contract method 0x4466b430.
//
// Solidity: function initializeNext(bytes[512] _nextValidatorSet, bytes _nextValidatorSetAggregate) returns()
func (_Verilay *VerilayTransactorSession) InitializeNext(_nextValidatorSet [512][]byte, _nextValidatorSetAggregate []byte) (*types.Transaction, error) {
	return _Verilay.Contract.InitializeNext(&_Verilay.TransactOpts, _nextValidatorSet, _nextValidatorSetAggregate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Verilay *VerilayTransactor) SubmitUpdate(opts *bind.TransactOpts, _chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Verilay.contract.Transact(opts, "submitUpdate", _chainRelayUpdate, _syncCommitteeUpdate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Verilay *VerilaySession) SubmitUpdate(_chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Verilay.Contract.SubmitUpdate(&_Verilay.TransactOpts, _chainRelayUpdate, _syncCommitteeUpdate)
}

// SubmitUpdate is a paid mutator transaction binding the contract method 0x8e8a29eb.
//
// Solidity: function submitUpdate((bytes,bool[512],bytes32,bytes32,uint64,bytes32[],bytes32,bytes32[],uint64,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) _chainRelayUpdate, (bytes[512],bytes,bytes32[]) _syncCommitteeUpdate) returns(bool)
func (_Verilay *VerilayTransactorSession) SubmitUpdate(_chainRelayUpdate ChainRelayUpdate, _syncCommitteeUpdate SyncCommitteeUpdate) (*types.Transaction, error) {
	return _Verilay.Contract.SubmitUpdate(&_Verilay.TransactOpts, _chainRelayUpdate, _syncCommitteeUpdate)
}

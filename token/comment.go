// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CommentABI is the input ABI used to generate the binding from.
const CommentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"name\":\"len\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commentList\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"contructor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CommentBin is the compiled bytecode used for deploying new contracts.
const CommentBin = `0x608060405234801561001057600080fd5b50610674806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631f7b6d3281146100715780637d4cd65f146100985780639507d39a14610195578063969fbf12146101ad57806397c3eb9b146101c4575b600080fd5b34801561007d57600080fd5b5061008661025d565b60408051918252519081900360200190f35b3480156100a457600080fd5b506100b0600435610263565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156100f75781810151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561015757818101518382015260200161013f565b50505050905090810190601f1680156101845780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156101a157600080fd5b506100b06004356103b8565b3480156101b957600080fd5b506101c261051b565b005b3480156101d057600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101c294369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750509335945061051d9350505050565b60005490565b600080548290811061027157fe5b60009182526020918290206003919091020180546040805160026001841615610100026000190190931692909204601f81018590048502830185019091528082529193509183919083018282801561030a5780601f106102df5761010080835404028352916020019161030a565b820191906000526020600020905b8154815290600101906020018083116102ed57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a85780601f1061037d576101008083540402835291602001916103a8565b820191906000526020600020905b81548152906001019060200180831161038b57829003601f168201915b5050505050908060020154905083565b6060806000806000858154811015156103cd57fe5b60009182526020918290206003919091020180546040805160026001841615610100026000190190931692909204601f8101859004850283018501909152808252919350918391908301828280156104665780601f1061043b57610100808354040283529160200191610466565b820191906000526020600020905b81548152906001019060200180831161044957829003601f168201915b50505050509350806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105055780601f106104da57610100808354040283529160200191610505565b820191906000526020600020905b8154815290600101906020018083116104e857829003601f168201915b5050505050925080600201549150509193909250565b565b604080516060810182528481526020808201859052918101839052600080546001810180835591805282518051929460039092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019261058192849201906105ad565b50602082810151805161059a92600185019201906105ad565b5060408201518160020155505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105ee57805160ff191683800117855561061b565b8280016001018555821561061b579182015b8281111561061b578251825591602001919060010190610600565b5061062792915061062b565b5090565b61064591905b808211156106275760008155600101610631565b905600a165627a7a72305820fd2e179ac281a80a016c176dc80247e1f3c687922827ff44ca4b048783246a420029`

// DeployComment deploys a new Ethereum contract, binding an instance of Comment to it.
func DeployComment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Comment, error) {
	parsed, err := abi.JSON(strings.NewReader(CommentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Comment{CommentCaller: CommentCaller{contract: contract}, CommentTransactor: CommentTransactor{contract: contract}, CommentFilterer: CommentFilterer{contract: contract}}, nil
}

// Comment is an auto generated Go binding around an Ethereum contract.
type Comment struct {
	CommentCaller     // Read-only binding to the contract
	CommentTransactor // Write-only binding to the contract
	CommentFilterer   // Log filterer for contract events
}

// CommentCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommentSession struct {
	Contract     *Comment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommentCallerSession struct {
	Contract *CommentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CommentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommentTransactorSession struct {
	Contract     *CommentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CommentRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommentRaw struct {
	Contract *Comment // Generic contract binding to access the raw methods on
}

// CommentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommentCallerRaw struct {
	Contract *CommentCaller // Generic read-only contract binding to access the raw methods on
}

// CommentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommentTransactorRaw struct {
	Contract *CommentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComment creates a new instance of Comment, bound to a specific deployed contract.
func NewComment(address common.Address, backend bind.ContractBackend) (*Comment, error) {
	contract, err := bindComment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comment{CommentCaller: CommentCaller{contract: contract}, CommentTransactor: CommentTransactor{contract: contract}, CommentFilterer: CommentFilterer{contract: contract}}, nil
}

// NewCommentCaller creates a new read-only instance of Comment, bound to a specific deployed contract.
func NewCommentCaller(address common.Address, caller bind.ContractCaller) (*CommentCaller, error) {
	contract, err := bindComment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommentCaller{contract: contract}, nil
}

// NewCommentTransactor creates a new write-only instance of Comment, bound to a specific deployed contract.
func NewCommentTransactor(address common.Address, transactor bind.ContractTransactor) (*CommentTransactor, error) {
	contract, err := bindComment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommentTransactor{contract: contract}, nil
}

// NewCommentFilterer creates a new log filterer instance of Comment, bound to a specific deployed contract.
func NewCommentFilterer(address common.Address, filterer bind.ContractFilterer) (*CommentFilterer, error) {
	contract, err := bindComment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommentFilterer{contract: contract}, nil
}

// bindComment binds a generic wrapper to an already deployed contract.
func bindComment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comment *CommentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comment.Contract.CommentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comment *CommentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comment.Contract.CommentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comment *CommentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comment.Contract.CommentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comment *CommentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comment *CommentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comment *CommentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comment.Contract.contract.Transact(opts, method, params...)
}

// CommentList is a free data retrieval call binding the contract method 0x7d4cd65f.
//
// Solidity: function commentList( uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentCaller) CommentList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	ret := new(struct {
		Name    []byte
		Message []byte
		Time    *big.Int
	})
	out := ret
	err := _Comment.contract.Call(opts, out, "commentList", arg0)
	return *ret, err
}

// CommentList is a free data retrieval call binding the contract method 0x7d4cd65f.
//
// Solidity: function commentList( uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentSession) CommentList(arg0 *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	return _Comment.Contract.CommentList(&_Comment.CallOpts, arg0)
}

// CommentList is a free data retrieval call binding the contract method 0x7d4cd65f.
//
// Solidity: function commentList( uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentCallerSession) CommentList(arg0 *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	return _Comment.Contract.CommentList(&_Comment.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(id uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentCaller) Get(opts *bind.CallOpts, id *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	ret := new(struct {
		Name    []byte
		Message []byte
		Time    *big.Int
	})
	out := ret
	err := _Comment.contract.Call(opts, out, "get", id)
	return *ret, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(id uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentSession) Get(id *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	return _Comment.Contract.Get(&_Comment.CallOpts, id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(id uint256) constant returns(name bytes, message bytes, time uint256)
func (_Comment *CommentCallerSession) Get(id *big.Int) (struct {
	Name    []byte
	Message []byte
	Time    *big.Int
}, error) {
	return _Comment.Contract.Get(&_Comment.CallOpts, id)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() constant returns(len uint256)
func (_Comment *CommentCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comment.contract.Call(opts, out, "length")
	return *ret0, err
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() constant returns(len uint256)
func (_Comment *CommentSession) Length() (*big.Int, error) {
	return _Comment.Contract.Length(&_Comment.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() constant returns(len uint256)
func (_Comment *CommentCallerSession) Length() (*big.Int, error) {
	return _Comment.Contract.Length(&_Comment.CallOpts)
}

// Contructor is a paid mutator transaction binding the contract method 0x969fbf12.
//
// Solidity: function contructor() returns()
func (_Comment *CommentTransactor) Contructor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comment.contract.Transact(opts, "contructor")
}

// Contructor is a paid mutator transaction binding the contract method 0x969fbf12.
//
// Solidity: function contructor() returns()
func (_Comment *CommentSession) Contructor() (*types.Transaction, error) {
	return _Comment.Contract.Contructor(&_Comment.TransactOpts)
}

// Contructor is a paid mutator transaction binding the contract method 0x969fbf12.
//
// Solidity: function contructor() returns()
func (_Comment *CommentTransactorSession) Contructor() (*types.Transaction, error) {
	return _Comment.Contract.Contructor(&_Comment.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x97c3eb9b.
//
// Solidity: function set(name bytes, message bytes, time uint256) returns()
func (_Comment *CommentTransactor) Set(opts *bind.TransactOpts, name []byte, message []byte, time *big.Int) (*types.Transaction, error) {
	return _Comment.contract.Transact(opts, "set", name, message, time)
}

// Set is a paid mutator transaction binding the contract method 0x97c3eb9b.
//
// Solidity: function set(name bytes, message bytes, time uint256) returns()
func (_Comment *CommentSession) Set(name []byte, message []byte, time *big.Int) (*types.Transaction, error) {
	return _Comment.Contract.Set(&_Comment.TransactOpts, name, message, time)
}

// Set is a paid mutator transaction binding the contract method 0x97c3eb9b.
//
// Solidity: function set(name bytes, message bytes, time uint256) returns()
func (_Comment *CommentTransactorSession) Set(name []byte, message []byte, time *big.Int) (*types.Transaction, error) {
	return _Comment.Contract.Set(&_Comment.TransactOpts, name, message, time)
}

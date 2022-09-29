// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty struct {
	Recipient   common.Address
	FeeFraction *big.Int
	SetBy       common.Address
}

// INFTKEYMarketplaceV2Bid is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceV2Bid struct {
	TokenId         *big.Int
	Value           *big.Int
	Bidder          common.Address
	ExpireTimestamp *big.Int
}

// INFTKEYMarketplaceV2Listing is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceV2Listing struct {
	TokenId         *big.Int
	Value           *big.Int
	Seller          common.Address
	ExpireTimestamp *big.Int
}

// NftkeyMetaData contains all meta data concerning the Nftkey contract.
var NftkeyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"SetRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyFee\",\"type\":\"uint256\"}],\"name\":\"TokenBidAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"TokenBidEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"TokenBidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyFee\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"TokenDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"TokenListed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"acceptBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionTimeOutRangeMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionTimeOutRangeMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"changeMarketplaceStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeInSec\",\"type\":\"uint256\"}],\"name\":\"changeMaxActionTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeInSec\",\"type\":\"uint256\"}],\"name\":\"changeMinActionTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"serviceFeeFraction_\",\"type\":\"uint8\"}],\"name\":\"changeSeriveFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRoyaltyFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"delistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"name\":\"enterBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getBidderBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"bidderBids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getBidderTokenBid\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"validBid\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenHighestBid\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"highestBid\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getTokenHighestBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"highestBids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"validListing\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getTokenListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Listing[]\",\"name\":\"listings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTradingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"numTokenListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"numTokenWithBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"royalty\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"internalType\":\"structINFTKEYMarketplaceRoyalty.ERC721CollectionRoyalty\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyUpperLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyForCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newUpperLimit\",\"type\":\"uint256\"}],\"name\":\"updateRoyaltyUpperLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NftkeyABI is the input ABI used to generate the binding from.
// Deprecated: Use NftkeyMetaData.ABI instead.
var NftkeyABI = NftkeyMetaData.ABI

// Nftkey is an auto generated Go binding around an Ethereum contract.
type Nftkey struct {
	NftkeyCaller     // Read-only binding to the contract
	NftkeyTransactor // Write-only binding to the contract
	NftkeyFilterer   // Log filterer for contract events
}

// NftkeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftkeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftkeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftkeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftkeyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftkeyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftkeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftkeySession struct {
	Contract     *Nftkey           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftkeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftkeyCallerSession struct {
	Contract *NftkeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftkeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftkeyTransactorSession struct {
	Contract     *NftkeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftkeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftkeyRaw struct {
	Contract *Nftkey // Generic contract binding to access the raw methods on
}

// NftkeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftkeyCallerRaw struct {
	Contract *NftkeyCaller // Generic read-only contract binding to access the raw methods on
}

// NftkeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftkeyTransactorRaw struct {
	Contract *NftkeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftkey creates a new instance of Nftkey, bound to a specific deployed contract.
func NewNftkey(address common.Address, backend bind.ContractBackend) (*Nftkey, error) {
	contract, err := bindNftkey(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nftkey{NftkeyCaller: NftkeyCaller{contract: contract}, NftkeyTransactor: NftkeyTransactor{contract: contract}, NftkeyFilterer: NftkeyFilterer{contract: contract}}, nil
}

// NewNftkeyCaller creates a new read-only instance of Nftkey, bound to a specific deployed contract.
func NewNftkeyCaller(address common.Address, caller bind.ContractCaller) (*NftkeyCaller, error) {
	contract, err := bindNftkey(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftkeyCaller{contract: contract}, nil
}

// NewNftkeyTransactor creates a new write-only instance of Nftkey, bound to a specific deployed contract.
func NewNftkeyTransactor(address common.Address, transactor bind.ContractTransactor) (*NftkeyTransactor, error) {
	contract, err := bindNftkey(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftkeyTransactor{contract: contract}, nil
}

// NewNftkeyFilterer creates a new log filterer instance of Nftkey, bound to a specific deployed contract.
func NewNftkeyFilterer(address common.Address, filterer bind.ContractFilterer) (*NftkeyFilterer, error) {
	contract, err := bindNftkey(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftkeyFilterer{contract: contract}, nil
}

// bindNftkey binds a generic wrapper to an already deployed contract.
func bindNftkey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftkeyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftkey *NftkeyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftkey.Contract.NftkeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftkey *NftkeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftkey.Contract.NftkeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftkey *NftkeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftkey.Contract.NftkeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nftkey *NftkeyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nftkey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nftkey *NftkeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftkey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nftkey *NftkeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nftkey.Contract.contract.Transact(opts, method, params...)
}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_Nftkey *NftkeyCaller) ActionTimeOutRangeMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "actionTimeOutRangeMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_Nftkey *NftkeySession) ActionTimeOutRangeMax() (*big.Int, error) {
	return _Nftkey.Contract.ActionTimeOutRangeMax(&_Nftkey.CallOpts)
}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_Nftkey *NftkeyCallerSession) ActionTimeOutRangeMax() (*big.Int, error) {
	return _Nftkey.Contract.ActionTimeOutRangeMax(&_Nftkey.CallOpts)
}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_Nftkey *NftkeyCaller) ActionTimeOutRangeMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "actionTimeOutRangeMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_Nftkey *NftkeySession) ActionTimeOutRangeMin() (*big.Int, error) {
	return _Nftkey.Contract.ActionTimeOutRangeMin(&_Nftkey.CallOpts)
}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_Nftkey *NftkeyCallerSession) ActionTimeOutRangeMin() (*big.Int, error) {
	return _Nftkey.Contract.ActionTimeOutRangeMin(&_Nftkey.CallOpts)
}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_Nftkey *NftkeyCaller) DefaultRoyaltyFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "defaultRoyaltyFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_Nftkey *NftkeySession) DefaultRoyaltyFraction() (*big.Int, error) {
	return _Nftkey.Contract.DefaultRoyaltyFraction(&_Nftkey.CallOpts)
}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_Nftkey *NftkeyCallerSession) DefaultRoyaltyFraction() (*big.Int, error) {
	return _Nftkey.Contract.DefaultRoyaltyFraction(&_Nftkey.CallOpts)
}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_Nftkey *NftkeyCaller) GetBidderBids(opts *bind.CallOpts, erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getBidderBids", erc721Address, bidder, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_Nftkey *NftkeySession) GetBidderBids(erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetBidderBids(&_Nftkey.CallOpts, erc721Address, bidder, from, size)
}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_Nftkey *NftkeyCallerSession) GetBidderBids(erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetBidderBids(&_Nftkey.CallOpts, erc721Address, bidder, from, size)
}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_Nftkey *NftkeyCaller) GetBidderTokenBid(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getBidderTokenBid", erc721Address, tokenId, bidder)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Bid)).(*INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_Nftkey *NftkeySession) GetBidderTokenBid(erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetBidderTokenBid(&_Nftkey.CallOpts, erc721Address, tokenId, bidder)
}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_Nftkey *NftkeyCallerSession) GetBidderTokenBid(erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetBidderTokenBid(&_Nftkey.CallOpts, erc721Address, tokenId, bidder)
}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_Nftkey *NftkeyCaller) GetTokenBids(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getTokenBids", erc721Address, tokenId)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_Nftkey *NftkeySession) GetTokenBids(erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenBids(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_Nftkey *NftkeyCallerSession) GetTokenBids(erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenBids(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_Nftkey *NftkeyCaller) GetTokenHighestBid(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getTokenHighestBid", erc721Address, tokenId)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Bid)).(*INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_Nftkey *NftkeySession) GetTokenHighestBid(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenHighestBid(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_Nftkey *NftkeyCallerSession) GetTokenHighestBid(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenHighestBid(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_Nftkey *NftkeyCaller) GetTokenHighestBids(opts *bind.CallOpts, erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getTokenHighestBids", erc721Address, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_Nftkey *NftkeySession) GetTokenHighestBids(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenHighestBids(&_Nftkey.CallOpts, erc721Address, from, size)
}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_Nftkey *NftkeyCallerSession) GetTokenHighestBids(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _Nftkey.Contract.GetTokenHighestBids(&_Nftkey.CallOpts, erc721Address, from, size)
}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_Nftkey *NftkeyCaller) GetTokenListing(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getTokenListing", erc721Address, tokenId)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Listing)).(*INFTKEYMarketplaceV2Listing)

	return out0, err

}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_Nftkey *NftkeySession) GetTokenListing(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	return _Nftkey.Contract.GetTokenListing(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_Nftkey *NftkeyCallerSession) GetTokenListing(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	return _Nftkey.Contract.GetTokenListing(&_Nftkey.CallOpts, erc721Address, tokenId)
}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_Nftkey *NftkeyCaller) GetTokenListings(opts *bind.CallOpts, erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "getTokenListings", erc721Address, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Listing)).(*[]INFTKEYMarketplaceV2Listing)

	return out0, err

}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_Nftkey *NftkeySession) GetTokenListings(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	return _Nftkey.Contract.GetTokenListings(&_Nftkey.CallOpts, erc721Address, from, size)
}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_Nftkey *NftkeyCallerSession) GetTokenListings(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	return _Nftkey.Contract.GetTokenListings(&_Nftkey.CallOpts, erc721Address, from, size)
}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_Nftkey *NftkeyCaller) IsTradingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "isTradingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_Nftkey *NftkeySession) IsTradingEnabled() (bool, error) {
	return _Nftkey.Contract.IsTradingEnabled(&_Nftkey.CallOpts)
}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_Nftkey *NftkeyCallerSession) IsTradingEnabled() (bool, error) {
	return _Nftkey.Contract.IsTradingEnabled(&_Nftkey.CallOpts)
}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeyCaller) NumTokenListings(opts *bind.CallOpts, erc721Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "numTokenListings", erc721Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeySession) NumTokenListings(erc721Address common.Address) (*big.Int, error) {
	return _Nftkey.Contract.NumTokenListings(&_Nftkey.CallOpts, erc721Address)
}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeyCallerSession) NumTokenListings(erc721Address common.Address) (*big.Int, error) {
	return _Nftkey.Contract.NumTokenListings(&_Nftkey.CallOpts, erc721Address)
}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeyCaller) NumTokenWithBids(opts *bind.CallOpts, erc721Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "numTokenWithBids", erc721Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeySession) NumTokenWithBids(erc721Address common.Address) (*big.Int, error) {
	return _Nftkey.Contract.NumTokenWithBids(&_Nftkey.CallOpts, erc721Address)
}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_Nftkey *NftkeyCallerSession) NumTokenWithBids(erc721Address common.Address) (*big.Int, error) {
	return _Nftkey.Contract.NumTokenWithBids(&_Nftkey.CallOpts, erc721Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nftkey *NftkeyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nftkey *NftkeySession) Owner() (common.Address, error) {
	return _Nftkey.Contract.Owner(&_Nftkey.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nftkey *NftkeyCallerSession) Owner() (common.Address, error) {
	return _Nftkey.Contract.Owner(&_Nftkey.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Nftkey *NftkeyCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Nftkey *NftkeySession) PaymentToken() (common.Address, error) {
	return _Nftkey.Contract.PaymentToken(&_Nftkey.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Nftkey *NftkeyCallerSession) PaymentToken() (common.Address, error) {
	return _Nftkey.Contract.PaymentToken(&_Nftkey.CallOpts)
}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_Nftkey *NftkeyCaller) Royalty(opts *bind.CallOpts, erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "royalty", erc721Address)

	if err != nil {
		return *new(INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty)).(*INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty)

	return out0, err

}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_Nftkey *NftkeySession) Royalty(erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	return _Nftkey.Contract.Royalty(&_Nftkey.CallOpts, erc721Address)
}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_Nftkey *NftkeyCallerSession) Royalty(erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	return _Nftkey.Contract.Royalty(&_Nftkey.CallOpts, erc721Address)
}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_Nftkey *NftkeyCaller) RoyaltyUpperLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "royaltyUpperLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_Nftkey *NftkeySession) RoyaltyUpperLimit() (*big.Int, error) {
	return _Nftkey.Contract.RoyaltyUpperLimit(&_Nftkey.CallOpts)
}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_Nftkey *NftkeyCallerSession) RoyaltyUpperLimit() (*big.Int, error) {
	return _Nftkey.Contract.RoyaltyUpperLimit(&_Nftkey.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_Nftkey *NftkeyCaller) ServiceFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nftkey.contract.Call(opts, &out, "serviceFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_Nftkey *NftkeySession) ServiceFee() (uint8, error) {
	return _Nftkey.Contract.ServiceFee(&_Nftkey.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_Nftkey *NftkeyCallerSession) ServiceFee() (uint8, error) {
	return _Nftkey.Contract.ServiceFee(&_Nftkey.CallOpts)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_Nftkey *NftkeyTransactor) AcceptBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "acceptBidForToken", erc721Address, tokenId, bidder, value)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_Nftkey *NftkeySession) AcceptBidForToken(erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.AcceptBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId, bidder, value)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_Nftkey *NftkeyTransactorSession) AcceptBidForToken(erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.AcceptBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId, bidder, value)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_Nftkey *NftkeyTransactor) BuyToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "buyToken", erc721Address, tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_Nftkey *NftkeySession) BuyToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.BuyToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_Nftkey *NftkeyTransactorSession) BuyToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.BuyToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_Nftkey *NftkeyTransactor) ChangeMarketplaceStatus(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "changeMarketplaceStatus", enabled)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_Nftkey *NftkeySession) ChangeMarketplaceStatus(enabled bool) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMarketplaceStatus(&_Nftkey.TransactOpts, enabled)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_Nftkey *NftkeyTransactorSession) ChangeMarketplaceStatus(enabled bool) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMarketplaceStatus(&_Nftkey.TransactOpts, enabled)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeyTransactor) ChangeMaxActionTimeLimit(opts *bind.TransactOpts, timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "changeMaxActionTimeLimit", timeInSec)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeySession) ChangeMaxActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMaxActionTimeLimit(&_Nftkey.TransactOpts, timeInSec)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeyTransactorSession) ChangeMaxActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMaxActionTimeLimit(&_Nftkey.TransactOpts, timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeyTransactor) ChangeMinActionTimeLimit(opts *bind.TransactOpts, timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "changeMinActionTimeLimit", timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeySession) ChangeMinActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMinActionTimeLimit(&_Nftkey.TransactOpts, timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_Nftkey *NftkeyTransactorSession) ChangeMinActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeMinActionTimeLimit(&_Nftkey.TransactOpts, timeInSec)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_Nftkey *NftkeyTransactor) ChangeSeriveFee(opts *bind.TransactOpts, serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "changeSeriveFee", serviceFeeFraction_)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_Nftkey *NftkeySession) ChangeSeriveFee(serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeSeriveFee(&_Nftkey.TransactOpts, serviceFeeFraction_)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_Nftkey *NftkeyTransactorSession) ChangeSeriveFee(serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _Nftkey.Contract.ChangeSeriveFee(&_Nftkey.TransactOpts, serviceFeeFraction_)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeyTransactor) DelistToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "delistToken", erc721Address, tokenId)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeySession) DelistToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.DelistToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeyTransactorSession) DelistToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.DelistToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeyTransactor) EnterBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "enterBidForToken", erc721Address, tokenId, value, expireTimestamp)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeySession) EnterBidForToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.EnterBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeyTransactorSession) EnterBidForToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.EnterBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeyTransactor) ListToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "listToken", erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeySession) ListToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ListToken(&_Nftkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_Nftkey *NftkeyTransactorSession) ListToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.ListToken(&_Nftkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nftkey *NftkeyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nftkey *NftkeySession) RenounceOwnership() (*types.Transaction, error) {
	return _Nftkey.Contract.RenounceOwnership(&_Nftkey.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nftkey *NftkeyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nftkey.Contract.RenounceOwnership(&_Nftkey.TransactOpts)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeyTransactor) SetRoyalty(opts *bind.TransactOpts, erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "setRoyalty", erc721Address, newRecipient, feeFraction)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeySession) SetRoyalty(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.SetRoyalty(&_Nftkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeyTransactorSession) SetRoyalty(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.SetRoyalty(&_Nftkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeyTransactor) SetRoyaltyForCollection(opts *bind.TransactOpts, erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "setRoyaltyForCollection", erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeySession) SetRoyaltyForCollection(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.SetRoyaltyForCollection(&_Nftkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_Nftkey *NftkeyTransactorSession) SetRoyaltyForCollection(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.SetRoyaltyForCollection(&_Nftkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nftkey *NftkeyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nftkey *NftkeySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nftkey.Contract.TransferOwnership(&_Nftkey.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nftkey *NftkeyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nftkey.Contract.TransferOwnership(&_Nftkey.TransactOpts, newOwner)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_Nftkey *NftkeyTransactor) UpdateRoyaltyUpperLimit(opts *bind.TransactOpts, _newUpperLimit *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "updateRoyaltyUpperLimit", _newUpperLimit)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_Nftkey *NftkeySession) UpdateRoyaltyUpperLimit(_newUpperLimit *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.UpdateRoyaltyUpperLimit(&_Nftkey.TransactOpts, _newUpperLimit)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_Nftkey *NftkeyTransactorSession) UpdateRoyaltyUpperLimit(_newUpperLimit *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.UpdateRoyaltyUpperLimit(&_Nftkey.TransactOpts, _newUpperLimit)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeyTransactor) WithdrawBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.contract.Transact(opts, "withdrawBidForToken", erc721Address, tokenId)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeySession) WithdrawBidForToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.WithdrawBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_Nftkey *NftkeyTransactorSession) WithdrawBidForToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nftkey.Contract.WithdrawBidForToken(&_Nftkey.TransactOpts, erc721Address, tokenId)
}

// NftkeyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nftkey contract.
type NftkeyOwnershipTransferredIterator struct {
	Event *NftkeyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftkeyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyOwnershipTransferred)
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
		it.Event = new(NftkeyOwnershipTransferred)
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
func (it *NftkeyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyOwnershipTransferred represents a OwnershipTransferred event raised by the Nftkey contract.
type NftkeyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nftkey *NftkeyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftkeyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyOwnershipTransferredIterator{contract: _Nftkey.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nftkey *NftkeyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftkeyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyOwnershipTransferred)
				if err := _Nftkey.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nftkey *NftkeyFilterer) ParseOwnershipTransferred(log types.Log) (*NftkeyOwnershipTransferred, error) {
	event := new(NftkeyOwnershipTransferred)
	if err := _Nftkey.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeySetRoyaltyIterator is returned from FilterSetRoyalty and is used to iterate over the raw logs and unpacked data for SetRoyalty events raised by the Nftkey contract.
type NftkeySetRoyaltyIterator struct {
	Event *NftkeySetRoyalty // Event containing the contract specifics and raw log

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
func (it *NftkeySetRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeySetRoyalty)
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
		it.Event = new(NftkeySetRoyalty)
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
func (it *NftkeySetRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeySetRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeySetRoyalty represents a SetRoyalty event raised by the Nftkey contract.
type NftkeySetRoyalty struct {
	Erc721Address common.Address
	Recipient     common.Address
	FeeFraction   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetRoyalty is a free log retrieval operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_Nftkey *NftkeyFilterer) FilterSetRoyalty(opts *bind.FilterOpts, erc721Address []common.Address, recipient []common.Address) (*NftkeySetRoyaltyIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "SetRoyalty", erc721AddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NftkeySetRoyaltyIterator{contract: _Nftkey.contract, event: "SetRoyalty", logs: logs, sub: sub}, nil
}

// WatchSetRoyalty is a free log subscription operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_Nftkey *NftkeyFilterer) WatchSetRoyalty(opts *bind.WatchOpts, sink chan<- *NftkeySetRoyalty, erc721Address []common.Address, recipient []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "SetRoyalty", erc721AddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeySetRoyalty)
				if err := _Nftkey.contract.UnpackLog(event, "SetRoyalty", log); err != nil {
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

// ParseSetRoyalty is a log parse operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_Nftkey *NftkeyFilterer) ParseSetRoyalty(log types.Log) (*NftkeySetRoyalty, error) {
	event := new(NftkeySetRoyalty)
	if err := _Nftkey.contract.UnpackLog(event, "SetRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenBidAcceptedIterator is returned from FilterTokenBidAccepted and is used to iterate over the raw logs and unpacked data for TokenBidAccepted events raised by the Nftkey contract.
type NftkeyTokenBidAcceptedIterator struct {
	Event *NftkeyTokenBidAccepted // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenBidAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenBidAccepted)
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
		it.Event = new(NftkeyTokenBidAccepted)
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
func (it *NftkeyTokenBidAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenBidAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenBidAccepted represents a TokenBidAccepted event raised by the Nftkey contract.
type NftkeyTokenBidAccepted struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Seller        common.Address
	Bid           INFTKEYMarketplaceV2Bid
	ServiceFee    *big.Int
	RoyaltyFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidAccepted is a free log retrieval operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) FilterTokenBidAccepted(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int, seller []common.Address) (*NftkeyTokenBidAcceptedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenBidAccepted", erc721AddressRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenBidAcceptedIterator{contract: _Nftkey.contract, event: "TokenBidAccepted", logs: logs, sub: sub}, nil
}

// WatchTokenBidAccepted is a free log subscription operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) WatchTokenBidAccepted(opts *bind.WatchOpts, sink chan<- *NftkeyTokenBidAccepted, erc721Address []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenBidAccepted", erc721AddressRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenBidAccepted)
				if err := _Nftkey.contract.UnpackLog(event, "TokenBidAccepted", log); err != nil {
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

// ParseTokenBidAccepted is a log parse operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) ParseTokenBidAccepted(log types.Log) (*NftkeyTokenBidAccepted, error) {
	event := new(NftkeyTokenBidAccepted)
	if err := _Nftkey.contract.UnpackLog(event, "TokenBidAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenBidEnteredIterator is returned from FilterTokenBidEntered and is used to iterate over the raw logs and unpacked data for TokenBidEntered events raised by the Nftkey contract.
type NftkeyTokenBidEnteredIterator struct {
	Event *NftkeyTokenBidEntered // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenBidEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenBidEntered)
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
		it.Event = new(NftkeyTokenBidEntered)
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
func (it *NftkeyTokenBidEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenBidEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenBidEntered represents a TokenBidEntered event raised by the Nftkey contract.
type NftkeyTokenBidEntered struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Bid           INFTKEYMarketplaceV2Bid
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidEntered is a free log retrieval operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) FilterTokenBidEntered(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NftkeyTokenBidEnteredIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenBidEntered", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenBidEnteredIterator{contract: _Nftkey.contract, event: "TokenBidEntered", logs: logs, sub: sub}, nil
}

// WatchTokenBidEntered is a free log subscription operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) WatchTokenBidEntered(opts *bind.WatchOpts, sink chan<- *NftkeyTokenBidEntered, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenBidEntered", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenBidEntered)
				if err := _Nftkey.contract.UnpackLog(event, "TokenBidEntered", log); err != nil {
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

// ParseTokenBidEntered is a log parse operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) ParseTokenBidEntered(log types.Log) (*NftkeyTokenBidEntered, error) {
	event := new(NftkeyTokenBidEntered)
	if err := _Nftkey.contract.UnpackLog(event, "TokenBidEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenBidWithdrawnIterator is returned from FilterTokenBidWithdrawn and is used to iterate over the raw logs and unpacked data for TokenBidWithdrawn events raised by the Nftkey contract.
type NftkeyTokenBidWithdrawnIterator struct {
	Event *NftkeyTokenBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenBidWithdrawn)
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
		it.Event = new(NftkeyTokenBidWithdrawn)
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
func (it *NftkeyTokenBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenBidWithdrawn represents a TokenBidWithdrawn event raised by the Nftkey contract.
type NftkeyTokenBidWithdrawn struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Bid           INFTKEYMarketplaceV2Bid
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidWithdrawn is a free log retrieval operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) FilterTokenBidWithdrawn(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NftkeyTokenBidWithdrawnIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenBidWithdrawn", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenBidWithdrawnIterator{contract: _Nftkey.contract, event: "TokenBidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenBidWithdrawn is a free log subscription operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) WatchTokenBidWithdrawn(opts *bind.WatchOpts, sink chan<- *NftkeyTokenBidWithdrawn, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenBidWithdrawn", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenBidWithdrawn)
				if err := _Nftkey.contract.UnpackLog(event, "TokenBidWithdrawn", log); err != nil {
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

// ParseTokenBidWithdrawn is a log parse operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_Nftkey *NftkeyFilterer) ParseTokenBidWithdrawn(log types.Log) (*NftkeyTokenBidWithdrawn, error) {
	event := new(NftkeyTokenBidWithdrawn)
	if err := _Nftkey.contract.UnpackLog(event, "TokenBidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the Nftkey contract.
type NftkeyTokenBoughtIterator struct {
	Event *NftkeyTokenBought // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenBought)
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
		it.Event = new(NftkeyTokenBought)
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
func (it *NftkeyTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenBought represents a TokenBought event raised by the Nftkey contract.
type NftkeyTokenBought struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Buyer         common.Address
	Listing       INFTKEYMarketplaceV2Listing
	ServiceFee    *big.Int
	RoyaltyFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) FilterTokenBought(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int, buyer []common.Address) (*NftkeyTokenBoughtIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenBought", erc721AddressRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenBoughtIterator{contract: _Nftkey.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *NftkeyTokenBought, erc721Address []common.Address, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenBought", erc721AddressRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenBought)
				if err := _Nftkey.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ParseTokenBought is a log parse operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_Nftkey *NftkeyFilterer) ParseTokenBought(log types.Log) (*NftkeyTokenBought, error) {
	event := new(NftkeyTokenBought)
	if err := _Nftkey.contract.UnpackLog(event, "TokenBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenDelistedIterator is returned from FilterTokenDelisted and is used to iterate over the raw logs and unpacked data for TokenDelisted events raised by the Nftkey contract.
type NftkeyTokenDelistedIterator struct {
	Event *NftkeyTokenDelisted // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenDelisted)
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
		it.Event = new(NftkeyTokenDelisted)
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
func (it *NftkeyTokenDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenDelisted represents a TokenDelisted event raised by the Nftkey contract.
type NftkeyTokenDelisted struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Listing       INFTKEYMarketplaceV2Listing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenDelisted is a free log retrieval operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) FilterTokenDelisted(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NftkeyTokenDelistedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenDelisted", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenDelistedIterator{contract: _Nftkey.contract, event: "TokenDelisted", logs: logs, sub: sub}, nil
}

// WatchTokenDelisted is a free log subscription operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) WatchTokenDelisted(opts *bind.WatchOpts, sink chan<- *NftkeyTokenDelisted, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenDelisted", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenDelisted)
				if err := _Nftkey.contract.UnpackLog(event, "TokenDelisted", log); err != nil {
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

// ParseTokenDelisted is a log parse operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) ParseTokenDelisted(log types.Log) (*NftkeyTokenDelisted, error) {
	event := new(NftkeyTokenDelisted)
	if err := _Nftkey.contract.UnpackLog(event, "TokenDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftkeyTokenListedIterator is returned from FilterTokenListed and is used to iterate over the raw logs and unpacked data for TokenListed events raised by the Nftkey contract.
type NftkeyTokenListedIterator struct {
	Event *NftkeyTokenListed // Event containing the contract specifics and raw log

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
func (it *NftkeyTokenListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftkeyTokenListed)
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
		it.Event = new(NftkeyTokenListed)
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
func (it *NftkeyTokenListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftkeyTokenListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftkeyTokenListed represents a TokenListed event raised by the Nftkey contract.
type NftkeyTokenListed struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Listing       INFTKEYMarketplaceV2Listing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenListed is a free log retrieval operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) FilterTokenListed(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NftkeyTokenListedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.FilterLogs(opts, "TokenListed", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftkeyTokenListedIterator{contract: _Nftkey.contract, event: "TokenListed", logs: logs, sub: sub}, nil
}

// WatchTokenListed is a free log subscription operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) WatchTokenListed(opts *bind.WatchOpts, sink chan<- *NftkeyTokenListed, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nftkey.contract.WatchLogs(opts, "TokenListed", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftkeyTokenListed)
				if err := _Nftkey.contract.UnpackLog(event, "TokenListed", log); err != nil {
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

// ParseTokenListed is a log parse operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_Nftkey *NftkeyFilterer) ParseTokenListed(log types.Log) (*NftkeyTokenListed, error) {
	event := new(NftkeyTokenListed)
	if err := _Nftkey.contract.UnpackLog(event, "TokenListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

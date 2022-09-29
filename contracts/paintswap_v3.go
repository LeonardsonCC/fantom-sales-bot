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

// PaintSwapMarketplaceV3StateBaseBulkTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type PaintSwapMarketplaceV3StateBaseBulkTransferInfo struct {
	Nft      common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	To       common.Address
}

// PaintSwapMarketplaceV3StateBaseFullDetails is an auto generated low-level Go binding around an user-defined struct.
type PaintSwapMarketplaceV3StateBaseFullDetails struct {
	Nfts            []common.Address
	TokenIds        []*big.Int
	AmountBatches   []*big.Int
	StartTime       *big.Int
	EndTime         *big.Int
	Price           *big.Int
	MaxBid          *big.Int
	MaxBidder       common.Address
	Seller          common.Address
	IsAuction       bool
	Complete        bool
	Antisnipe       bool
	FlashAuction    bool
	Amount          *big.Int
	AmountRemaining *big.Int
	PaymentToken    common.Address
	Vault           common.Address
}

// PaintSwapMarketplaceV3StateBaseOffer is an auto generated low-level Go binding around an user-defined struct.
type PaintSwapMarketplaceV3StateBaseOffer struct {
	Nft               common.Address
	Nfts              []common.Address
	TokenIds          []*big.Int
	Quantity          *big.Int
	QuantityRemaining *big.Int
	Price             *big.Int
	Prices            []*big.Int
	From              common.Address
	OfferType         uint8
	Expires           *big.Int
	SaleId            *big.Int
}

// PaintswapMetaData contains all meta data concerning the Paintswap contract.
var PaintswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_officialNFTDiscount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplaceWhitelist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collectionRoyalties\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paintRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wftm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implBuyEditListing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implMakeAndEditOffers\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implBidsAndAcceptOffers\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implListAndComplete\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implBatchTransferNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feePercentage\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"AddedVaultFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"}],\"name\":\"CancelledSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bid\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextMinimum\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"searchKeywords\",\"type\":\"string\"}],\"name\":\"NewCollectionOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"prices\",\"type\":\"uint128[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"searchKeywords\",\"type\":\"string[]\"}],\"name\":\"NewFilteredCollectionOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"searchKeywords\",\"type\":\"string\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"antisnipe\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flashAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNSFW\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"searchKeywords\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"NewSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"marketplaceId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[][]\",\"name\":\"nfts\",\"type\":\"address[][]\"},{\"indexed\":false,\"internalType\":\"uint256[][]\",\"name\":\"tokenIds\",\"type\":\"uint256[][]\"},{\"indexed\":false,\"internalType\":\"uint256[][]\",\"name\":\"amountBatches\",\"type\":\"uint256[][]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"prices\",\"type\":\"uint128[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"durations\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"isAuctions\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"isAntisnipes\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"isNSFWs\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"searchKeywords\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"address[][]\",\"name\":\"addresses\",\"type\":\"address[][]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sellers\",\"type\":\"address[]\"}],\"name\":\"NewSaleBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"antisnipe\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flashAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNSFW\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"searchKeywords\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"NewTempListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"}],\"name\":\"OfferAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"OfferCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"OfferExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"OfferRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"RemoveExpiredListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"RemovedVaultFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"}],\"name\":\"SaleFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"UpdateEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newPrice\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"UpdateOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmountRemaining\",\"type\":\"uint256\"}],\"name\":\"UpdateQuantity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"UpdateStartTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"_searchKeywords\",\"type\":\"string\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_offerIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdOrMarketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isMarketplaceIds\",\"type\":\"bool[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_quantities\",\"type\":\"uint128[]\"},{\"internalType\":\"string[]\",\"name\":\"_searchKeywords\",\"type\":\"string[]\"}],\"name\":\"acceptOfferBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"acceptOnSaleOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountBatches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint64[]\",\"name\":\"_times\",\"type\":\"uint64[]\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"bool[]\",\"name\":\"_flags\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"_searchKeywords\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[][]\",\"name\":\"_nfts\",\"type\":\"address[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"_amountBatches\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint64[][]\",\"name\":\"_times\",\"type\":\"uint64[][]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"bool[][]\",\"name\":\"_flags\",\"type\":\"bool[][]\"},{\"internalType\":\"string[]\",\"name\":\"_searchKeywords\",\"type\":\"string[]\"},{\"internalType\":\"address[][]\",\"name\":\"_addresses\",\"type\":\"address[][]\"}],\"name\":\"addSaleBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactory\",\"type\":\"address\"}],\"name\":\"addValidVaultFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTimeIncreaseCutOff\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionGracePeriodForCancelling\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"bool\",\"name\":\"_allowFailures\",\"type\":\"bool\"}],\"name\":\"buyBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_offerIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelOfferBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"cancelSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelSaleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"completeMarketplaceEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMarketplaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_newQuantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_newOfferPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_newOfferDuration\",\"type\":\"uint256\"}],\"name\":\"editOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_offerIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_newQuantities\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_newOfferPrices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newOfferDurations\",\"type\":\"uint256[]\"}],\"name\":\"editOfferBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"}],\"name\":\"editPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_newPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_newQuantity\",\"type\":\"uint128\"}],\"name\":\"editPriceAndQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceId\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_newPrices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_newQuantities\",\"type\":\"uint128[]\"}],\"name\":\"editPriceAndQuantityBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"}],\"name\":\"editPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_newQuantity\",\"type\":\"uint128\"}],\"name\":\"editQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_newQuantities\",\"type\":\"uint128[]\"}],\"name\":\"editQuantityBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_offerIds\",\"type\":\"uint256[]\"}],\"name\":\"expireOffers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashAuctionEndTimeIncreaseCutOff\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashAuctionGracePeriodForCancelling\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offerId\",\"type\":\"uint256\"}],\"name\":\"getOffer\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"quantityRemaining\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint128[]\",\"name\":\"prices\",\"type\":\"uint128[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"enumPaintSwapMarketplaceV3StateBase.OfferType\",\"name\":\"offerType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"}],\"internalType\":\"structPaintSwapMarketplaceV3StateBase.Offer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"getSaleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxBid\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maxBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuction\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"antisnipe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flashAuction\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountRemaining\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"internalType\":\"structPaintSwapMarketplaceV3StateBase.FullDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"hasExistingActiveOffers\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_searchKeywords\",\"type\":\"string[]\"}],\"name\":\"makeAltOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_bid\",\"type\":\"uint128\"}],\"name\":\"makeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_bids\",\"type\":\"uint128[]\"}],\"name\":\"makeBidBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_searchKeywords\",\"type\":\"string\"}],\"name\":\"makeCollectionOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_searchKeywords\",\"type\":\"string[]\"}],\"name\":\"makeFilteredCollectionOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint128[]\",\"name\":\"_quantities\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[][]\",\"name\":\"_prices\",\"type\":\"uint128[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"},{\"internalType\":\"string[][]\",\"name\":\"_searchKeywords\",\"type\":\"string[][]\"}],\"name\":\"makeFilteredCollectionOfferBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"makeOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_marketplaceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_quantities\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"}],\"name\":\"makeOfferBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_quantity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_searchKeywords\",\"type\":\"string\"}],\"name\":\"makeOfferSingle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_quantities\",\"type\":\"uint128[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_searchKeywords\",\"type\":\"string[]\"}],\"name\":\"makeOfferSingleBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxActiveOfferCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxOfferDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStartTimeIncrement\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFlashAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minIncreasedBidPercent\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOfferDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVaultVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"nextMinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offerCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offerId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactory\",\"type\":\"address\"}],\"name\":\"removeInvalidVaultFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reserve1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reserve2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reserve3\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reserve4\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reserveOnlyOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeNFTTransferBulk\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structPaintSwapMarketplaceV3StateBase.BulkTransferInfo[]\",\"name\":\"_nftsInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeNFTTransferBulkOrdered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salesFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_maxBundleNumber\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setBundles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minFlashAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAuctionDuration\",\"type\":\"uint256\"}],\"name\":\"setDurations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnableSelling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dev\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_salesFee\",\"type\":\"address\"}],\"name\":\"setFeeAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feePercentage\",\"type\":\"address\"}],\"name\":\"setFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implBatchTransferNFT\",\"type\":\"address\"}],\"name\":\"setImplBatchTransferNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implBidsAndAcceptOffers\",\"type\":\"address\"}],\"name\":\"setImplBidsAndAcceptOffers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implBuyEditListing\",\"type\":\"address\"}],\"name\":\"setImplBuyEditListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implListAndComplete\",\"type\":\"address\"}],\"name\":\"setImplListAndComplete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxActiveOfferCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minVaultVersion\",\"type\":\"uint256\"}],\"name\":\"setMaxActiveOfferCountAndMinVaultVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDuration\",\"type\":\"uint256\"}],\"name\":\"setOfferDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxRoyalty\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_maxStartTimeIncrement\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_officialNFTDiscount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_paintRouter\",\"type\":\"address\"}],\"name\":\"setVarious\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_auctionGracePeriodForCancelling\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_flashAuctionGracePeriodForCancelling\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_auctionEndTimeIncreaseCutOff\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_flashAuctionEndTimeIncreaseCutOff\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_minIncreasedBidPercent\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"_modifyInitialStartTime\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_modifyPostStartTime\",\"type\":\"bool\"}],\"name\":\"setVariousTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implMakeAndEditOffers\",\"type\":\"address\"}],\"name\":\"setimplMakeAndEditOffers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"singleNFTListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAltOffersMade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userBundleOffersMade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCollectionOffersMade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userFilteredCollectionOffersMade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userOffersMade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PaintswapABI is the input ABI used to generate the binding from.
// Deprecated: Use PaintswapMetaData.ABI instead.
var PaintswapABI = PaintswapMetaData.ABI

// Paintswap is an auto generated Go binding around an Ethereum contract.
type Paintswap struct {
	PaintswapCaller     // Read-only binding to the contract
	PaintswapTransactor // Write-only binding to the contract
	PaintswapFilterer   // Log filterer for contract events
}

// PaintswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaintswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaintswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaintswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaintswapSession struct {
	Contract     *Paintswap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaintswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaintswapCallerSession struct {
	Contract *PaintswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PaintswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaintswapTransactorSession struct {
	Contract     *PaintswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PaintswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaintswapRaw struct {
	Contract *Paintswap // Generic contract binding to access the raw methods on
}

// PaintswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaintswapCallerRaw struct {
	Contract *PaintswapCaller // Generic read-only contract binding to access the raw methods on
}

// PaintswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaintswapTransactorRaw struct {
	Contract *PaintswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaintswap creates a new instance of Paintswap, bound to a specific deployed contract.
func NewPaintswap(address common.Address, backend bind.ContractBackend) (*Paintswap, error) {
	contract, err := bindPaintswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paintswap{PaintswapCaller: PaintswapCaller{contract: contract}, PaintswapTransactor: PaintswapTransactor{contract: contract}, PaintswapFilterer: PaintswapFilterer{contract: contract}}, nil
}

// NewPaintswapCaller creates a new read-only instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapCaller(address common.Address, caller bind.ContractCaller) (*PaintswapCaller, error) {
	contract, err := bindPaintswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaintswapCaller{contract: contract}, nil
}

// NewPaintswapTransactor creates a new write-only instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapTransactor(address common.Address, transactor bind.ContractTransactor) (*PaintswapTransactor, error) {
	contract, err := bindPaintswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaintswapTransactor{contract: contract}, nil
}

// NewPaintswapFilterer creates a new log filterer instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapFilterer(address common.Address, filterer bind.ContractFilterer) (*PaintswapFilterer, error) {
	contract, err := bindPaintswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaintswapFilterer{contract: contract}, nil
}

// bindPaintswap binds a generic wrapper to an already deployed contract.
func bindPaintswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaintswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paintswap *PaintswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paintswap.Contract.PaintswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paintswap *PaintswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.Contract.PaintswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paintswap *PaintswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paintswap.Contract.PaintswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paintswap *PaintswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paintswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paintswap *PaintswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paintswap *PaintswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paintswap.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Paintswap *PaintswapCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Paintswap *PaintswapSession) VERSION() (*big.Int, error) {
	return _Paintswap.Contract.VERSION(&_Paintswap.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) VERSION() (*big.Int, error) {
	return _Paintswap.Contract.VERSION(&_Paintswap.CallOpts)
}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapCaller) AuctionEndTimeIncreaseCutOff(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "auctionEndTimeIncreaseCutOff")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapSession) AuctionEndTimeIncreaseCutOff() (uint64, error) {
	return _Paintswap.Contract.AuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapCallerSession) AuctionEndTimeIncreaseCutOff() (uint64, error) {
	return _Paintswap.Contract.AuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// AuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x49ae64b0.
//
// Solidity: function auctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapCaller) AuctionGracePeriodForCancelling(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "auctionGracePeriodForCancelling")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x49ae64b0.
//
// Solidity: function auctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapSession) AuctionGracePeriodForCancelling() (uint64, error) {
	return _Paintswap.Contract.AuctionGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// AuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x49ae64b0.
//
// Solidity: function auctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapCallerSession) AuctionGracePeriodForCancelling() (uint64, error) {
	return _Paintswap.Contract.AuctionGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapCaller) CurrentMarketplaceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "currentMarketplaceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapSession) CurrentMarketplaceId() (*big.Int, error) {
	return _Paintswap.Contract.CurrentMarketplaceId(&_Paintswap.CallOpts)
}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) CurrentMarketplaceId() (*big.Int, error) {
	return _Paintswap.Contract.CurrentMarketplaceId(&_Paintswap.CallOpts)
}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapCaller) DevFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "devFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapSession) DevFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevFeeAddress(&_Paintswap.CallOpts)
}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapCallerSession) DevFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevFeeAddress(&_Paintswap.CallOpts)
}

// FlashAuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x1db84f7c.
//
// Solidity: function flashAuctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapCaller) FlashAuctionEndTimeIncreaseCutOff(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "flashAuctionEndTimeIncreaseCutOff")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FlashAuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x1db84f7c.
//
// Solidity: function flashAuctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapSession) FlashAuctionEndTimeIncreaseCutOff() (uint64, error) {
	return _Paintswap.Contract.FlashAuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// FlashAuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x1db84f7c.
//
// Solidity: function flashAuctionEndTimeIncreaseCutOff() view returns(uint64)
func (_Paintswap *PaintswapCallerSession) FlashAuctionEndTimeIncreaseCutOff() (uint64, error) {
	return _Paintswap.Contract.FlashAuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// FlashAuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x3092a388.
//
// Solidity: function flashAuctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapCaller) FlashAuctionGracePeriodForCancelling(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "flashAuctionGracePeriodForCancelling")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FlashAuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x3092a388.
//
// Solidity: function flashAuctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapSession) FlashAuctionGracePeriodForCancelling() (uint64, error) {
	return _Paintswap.Contract.FlashAuctionGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// FlashAuctionGracePeriodForCancelling is a free data retrieval call binding the contract method 0x3092a388.
//
// Solidity: function flashAuctionGracePeriodForCancelling() view returns(uint64)
func (_Paintswap *PaintswapCallerSession) FlashAuctionGracePeriodForCancelling() (uint64, error) {
	return _Paintswap.Contract.FlashAuctionGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 _offerId) view returns((address,address[],uint256[],uint128,uint128,uint128,uint128[],address,uint8,uint256,uint256))
func (_Paintswap *PaintswapCaller) GetOffer(opts *bind.CallOpts, _offerId *big.Int) (PaintSwapMarketplaceV3StateBaseOffer, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "getOffer", _offerId)

	if err != nil {
		return *new(PaintSwapMarketplaceV3StateBaseOffer), err
	}

	out0 := *abi.ConvertType(out[0], new(PaintSwapMarketplaceV3StateBaseOffer)).(*PaintSwapMarketplaceV3StateBaseOffer)

	return out0, err

}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 _offerId) view returns((address,address[],uint256[],uint128,uint128,uint128,uint128[],address,uint8,uint256,uint256))
func (_Paintswap *PaintswapSession) GetOffer(_offerId *big.Int) (PaintSwapMarketplaceV3StateBaseOffer, error) {
	return _Paintswap.Contract.GetOffer(&_Paintswap.CallOpts, _offerId)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(uint256 _offerId) view returns((address,address[],uint256[],uint128,uint128,uint128,uint128[],address,uint8,uint256,uint256))
func (_Paintswap *PaintswapCallerSession) GetOffer(_offerId *big.Int) (PaintSwapMarketplaceV3StateBaseOffer, error) {
	return _Paintswap.Contract.GetOffer(&_Paintswap.CallOpts, _offerId)
}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],uint256,uint256,uint128,uint128,address,address,bool,bool,bool,bool,uint128,uint128,address,address))
func (_Paintswap *PaintswapCaller) GetSaleDetails(opts *bind.CallOpts, _marketplaceId *big.Int) (PaintSwapMarketplaceV3StateBaseFullDetails, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "getSaleDetails", _marketplaceId)

	if err != nil {
		return *new(PaintSwapMarketplaceV3StateBaseFullDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(PaintSwapMarketplaceV3StateBaseFullDetails)).(*PaintSwapMarketplaceV3StateBaseFullDetails)

	return out0, err

}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],uint256,uint256,uint128,uint128,address,address,bool,bool,bool,bool,uint128,uint128,address,address))
func (_Paintswap *PaintswapSession) GetSaleDetails(_marketplaceId *big.Int) (PaintSwapMarketplaceV3StateBaseFullDetails, error) {
	return _Paintswap.Contract.GetSaleDetails(&_Paintswap.CallOpts, _marketplaceId)
}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],uint256,uint256,uint128,uint128,address,address,bool,bool,bool,bool,uint128,uint128,address,address))
func (_Paintswap *PaintswapCallerSession) GetSaleDetails(_marketplaceId *big.Int) (PaintSwapMarketplaceV3StateBaseFullDetails, error) {
	return _Paintswap.Contract.GetSaleDetails(&_Paintswap.CallOpts, _marketplaceId)
}

// HasExistingActiveOffers is a free data retrieval call binding the contract method 0xf968d006.
//
// Solidity: function hasExistingActiveOffers(address[] _nfts, uint256[] _tokenIds, address _user) view returns(bool[])
func (_Paintswap *PaintswapCaller) HasExistingActiveOffers(opts *bind.CallOpts, _nfts []common.Address, _tokenIds []*big.Int, _user common.Address) ([]bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "hasExistingActiveOffers", _nfts, _tokenIds, _user)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// HasExistingActiveOffers is a free data retrieval call binding the contract method 0xf968d006.
//
// Solidity: function hasExistingActiveOffers(address[] _nfts, uint256[] _tokenIds, address _user) view returns(bool[])
func (_Paintswap *PaintswapSession) HasExistingActiveOffers(_nfts []common.Address, _tokenIds []*big.Int, _user common.Address) ([]bool, error) {
	return _Paintswap.Contract.HasExistingActiveOffers(&_Paintswap.CallOpts, _nfts, _tokenIds, _user)
}

// HasExistingActiveOffers is a free data retrieval call binding the contract method 0xf968d006.
//
// Solidity: function hasExistingActiveOffers(address[] _nfts, uint256[] _tokenIds, address _user) view returns(bool[])
func (_Paintswap *PaintswapCallerSession) HasExistingActiveOffers(_nfts []common.Address, _tokenIds []*big.Int, _user common.Address) ([]bool, error) {
	return _Paintswap.Contract.HasExistingActiveOffers(&_Paintswap.CallOpts, _nfts, _tokenIds, _user)
}

// MaxActiveOfferCount is a free data retrieval call binding the contract method 0x6b67a72c.
//
// Solidity: function maxActiveOfferCount() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxActiveOfferCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxActiveOfferCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxActiveOfferCount is a free data retrieval call binding the contract method 0x6b67a72c.
//
// Solidity: function maxActiveOfferCount() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxActiveOfferCount() (*big.Int, error) {
	return _Paintswap.Contract.MaxActiveOfferCount(&_Paintswap.CallOpts)
}

// MaxActiveOfferCount is a free data retrieval call binding the contract method 0x6b67a72c.
//
// Solidity: function maxActiveOfferCount() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxActiveOfferCount() (*big.Int, error) {
	return _Paintswap.Contract.MaxActiveOfferCount(&_Paintswap.CallOpts)
}

// MaxAuctionDuration is a free data retrieval call binding the contract method 0x48c9581e.
//
// Solidity: function maxAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAuctionDuration is a free data retrieval call binding the contract method 0x48c9581e.
//
// Solidity: function maxAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxAuctionDuration(&_Paintswap.CallOpts)
}

// MaxAuctionDuration is a free data retrieval call binding the contract method 0x48c9581e.
//
// Solidity: function maxAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxAuctionDuration(&_Paintswap.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxDuration(&_Paintswap.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxDuration(&_Paintswap.CallOpts)
}

// MaxOfferDuration is a free data retrieval call binding the contract method 0xda1a200e.
//
// Solidity: function maxOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxOfferDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxOfferDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOfferDuration is a free data retrieval call binding the contract method 0xda1a200e.
//
// Solidity: function maxOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxOfferDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxOfferDuration(&_Paintswap.CallOpts)
}

// MaxOfferDuration is a free data retrieval call binding the contract method 0xda1a200e.
//
// Solidity: function maxOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxOfferDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxOfferDuration(&_Paintswap.CallOpts)
}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint64)
func (_Paintswap *PaintswapCaller) MaxStartTimeIncrement(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxStartTimeIncrement")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint64)
func (_Paintswap *PaintswapSession) MaxStartTimeIncrement() (uint64, error) {
	return _Paintswap.Contract.MaxStartTimeIncrement(&_Paintswap.CallOpts)
}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint64)
func (_Paintswap *PaintswapCallerSession) MaxStartTimeIncrement() (uint64, error) {
	return _Paintswap.Contract.MaxStartTimeIncrement(&_Paintswap.CallOpts)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MinAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinAuctionDuration(&_Paintswap.CallOpts)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinAuctionDuration(&_Paintswap.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MinDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinDuration(&_Paintswap.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinDuration(&_Paintswap.CallOpts)
}

// MinFlashAuctionDuration is a free data retrieval call binding the contract method 0x858122fd.
//
// Solidity: function minFlashAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinFlashAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minFlashAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFlashAuctionDuration is a free data retrieval call binding the contract method 0x858122fd.
//
// Solidity: function minFlashAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MinFlashAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinFlashAuctionDuration(&_Paintswap.CallOpts)
}

// MinFlashAuctionDuration is a free data retrieval call binding the contract method 0x858122fd.
//
// Solidity: function minFlashAuctionDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinFlashAuctionDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinFlashAuctionDuration(&_Paintswap.CallOpts)
}

// MinIncreasedBidPercent is a free data retrieval call binding the contract method 0xebc11956.
//
// Solidity: function minIncreasedBidPercent() view returns(uint128)
func (_Paintswap *PaintswapCaller) MinIncreasedBidPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minIncreasedBidPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinIncreasedBidPercent is a free data retrieval call binding the contract method 0xebc11956.
//
// Solidity: function minIncreasedBidPercent() view returns(uint128)
func (_Paintswap *PaintswapSession) MinIncreasedBidPercent() (*big.Int, error) {
	return _Paintswap.Contract.MinIncreasedBidPercent(&_Paintswap.CallOpts)
}

// MinIncreasedBidPercent is a free data retrieval call binding the contract method 0xebc11956.
//
// Solidity: function minIncreasedBidPercent() view returns(uint128)
func (_Paintswap *PaintswapCallerSession) MinIncreasedBidPercent() (*big.Int, error) {
	return _Paintswap.Contract.MinIncreasedBidPercent(&_Paintswap.CallOpts)
}

// MinOfferDuration is a free data retrieval call binding the contract method 0xc99e0ffc.
//
// Solidity: function minOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinOfferDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minOfferDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOfferDuration is a free data retrieval call binding the contract method 0xc99e0ffc.
//
// Solidity: function minOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MinOfferDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinOfferDuration(&_Paintswap.CallOpts)
}

// MinOfferDuration is a free data retrieval call binding the contract method 0xc99e0ffc.
//
// Solidity: function minOfferDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinOfferDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinOfferDuration(&_Paintswap.CallOpts)
}

// MinVaultVersion is a free data retrieval call binding the contract method 0xe4958c01.
//
// Solidity: function minVaultVersion() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinVaultVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minVaultVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVaultVersion is a free data retrieval call binding the contract method 0xe4958c01.
//
// Solidity: function minVaultVersion() view returns(uint256)
func (_Paintswap *PaintswapSession) MinVaultVersion() (*big.Int, error) {
	return _Paintswap.Contract.MinVaultVersion(&_Paintswap.CallOpts)
}

// MinVaultVersion is a free data retrieval call binding the contract method 0xe4958c01.
//
// Solidity: function minVaultVersion() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinVaultVersion() (*big.Int, error) {
	return _Paintswap.Contract.MinVaultVersion(&_Paintswap.CallOpts)
}

// NextMinimumBid is a free data retrieval call binding the contract method 0xf92171e8.
//
// Solidity: function nextMinimumBid(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapCaller) NextMinimumBid(opts *bind.CallOpts, _marketplaceId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "nextMinimumBid", _marketplaceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMinimumBid is a free data retrieval call binding the contract method 0xf92171e8.
//
// Solidity: function nextMinimumBid(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapSession) NextMinimumBid(_marketplaceId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.NextMinimumBid(&_Paintswap.CallOpts, _marketplaceId)
}

// NextMinimumBid is a free data retrieval call binding the contract method 0xf92171e8.
//
// Solidity: function nextMinimumBid(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) NextMinimumBid(_marketplaceId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.NextMinimumBid(&_Paintswap.CallOpts, _marketplaceId)
}

// OfferCounts is a free data retrieval call binding the contract method 0x3ef748f6.
//
// Solidity: function offerCounts(address ) view returns(uint256)
func (_Paintswap *PaintswapCaller) OfferCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "offerCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferCounts is a free data retrieval call binding the contract method 0x3ef748f6.
//
// Solidity: function offerCounts(address ) view returns(uint256)
func (_Paintswap *PaintswapSession) OfferCounts(arg0 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.OfferCounts(&_Paintswap.CallOpts, arg0)
}

// OfferCounts is a free data retrieval call binding the contract method 0x3ef748f6.
//
// Solidity: function offerCounts(address ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) OfferCounts(arg0 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.OfferCounts(&_Paintswap.CallOpts, arg0)
}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() view returns(uint256)
func (_Paintswap *PaintswapCaller) OfferId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "offerId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() view returns(uint256)
func (_Paintswap *PaintswapSession) OfferId() (*big.Int, error) {
	return _Paintswap.Contract.OfferId(&_Paintswap.CallOpts)
}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) OfferId() (*big.Int, error) {
	return _Paintswap.Contract.OfferId(&_Paintswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapSession) Owner() (common.Address, error) {
	return _Paintswap.Contract.Owner(&_Paintswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapCallerSession) Owner() (common.Address, error) {
	return _Paintswap.Contract.Owner(&_Paintswap.CallOpts)
}

// SalesFeeAddress is a free data retrieval call binding the contract method 0xf602eccd.
//
// Solidity: function salesFeeAddress() view returns(address)
func (_Paintswap *PaintswapCaller) SalesFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "salesFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SalesFeeAddress is a free data retrieval call binding the contract method 0xf602eccd.
//
// Solidity: function salesFeeAddress() view returns(address)
func (_Paintswap *PaintswapSession) SalesFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.SalesFeeAddress(&_Paintswap.CallOpts)
}

// SalesFeeAddress is a free data retrieval call binding the contract method 0xf602eccd.
//
// Solidity: function salesFeeAddress() view returns(address)
func (_Paintswap *PaintswapCallerSession) SalesFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.SalesFeeAddress(&_Paintswap.CallOpts)
}

// SingleNFTListings is a free data retrieval call binding the contract method 0x4451d01c.
//
// Solidity: function singleNFTListings(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCaller) SingleNFTListings(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "singleNFTListings", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleNFTListings is a free data retrieval call binding the contract method 0x4451d01c.
//
// Solidity: function singleNFTListings(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapSession) SingleNFTListings(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.SingleNFTListings(&_Paintswap.CallOpts, arg0, arg1, arg2)
}

// SingleNFTListings is a free data retrieval call binding the contract method 0x4451d01c.
//
// Solidity: function singleNFTListings(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) SingleNFTListings(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.SingleNFTListings(&_Paintswap.CallOpts, arg0, arg1, arg2)
}

// UserAltOffersMade is a free data retrieval call binding the contract method 0x804d3f47.
//
// Solidity: function userAltOffersMade(address ) view returns(uint256)
func (_Paintswap *PaintswapCaller) UserAltOffersMade(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "userAltOffersMade", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAltOffersMade is a free data retrieval call binding the contract method 0x804d3f47.
//
// Solidity: function userAltOffersMade(address ) view returns(uint256)
func (_Paintswap *PaintswapSession) UserAltOffersMade(arg0 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserAltOffersMade(&_Paintswap.CallOpts, arg0)
}

// UserAltOffersMade is a free data retrieval call binding the contract method 0x804d3f47.
//
// Solidity: function userAltOffersMade(address ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) UserAltOffersMade(arg0 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserAltOffersMade(&_Paintswap.CallOpts, arg0)
}

// UserBundleOffersMade is a free data retrieval call binding the contract method 0xa51377a4.
//
// Solidity: function userBundleOffersMade(address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCaller) UserBundleOffersMade(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "userBundleOffersMade", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBundleOffersMade is a free data retrieval call binding the contract method 0xa51377a4.
//
// Solidity: function userBundleOffersMade(address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapSession) UserBundleOffersMade(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.UserBundleOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserBundleOffersMade is a free data retrieval call binding the contract method 0xa51377a4.
//
// Solidity: function userBundleOffersMade(address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) UserBundleOffersMade(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.UserBundleOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserCollectionOffersMade is a free data retrieval call binding the contract method 0xa41e4776.
//
// Solidity: function userCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapCaller) UserCollectionOffersMade(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "userCollectionOffersMade", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCollectionOffersMade is a free data retrieval call binding the contract method 0xa41e4776.
//
// Solidity: function userCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapSession) UserCollectionOffersMade(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserCollectionOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserCollectionOffersMade is a free data retrieval call binding the contract method 0xa41e4776.
//
// Solidity: function userCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) UserCollectionOffersMade(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserCollectionOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserFilteredCollectionOffersMade is a free data retrieval call binding the contract method 0x47cdb1bc.
//
// Solidity: function userFilteredCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapCaller) UserFilteredCollectionOffersMade(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "userFilteredCollectionOffersMade", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFilteredCollectionOffersMade is a free data retrieval call binding the contract method 0x47cdb1bc.
//
// Solidity: function userFilteredCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapSession) UserFilteredCollectionOffersMade(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserFilteredCollectionOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserFilteredCollectionOffersMade is a free data retrieval call binding the contract method 0x47cdb1bc.
//
// Solidity: function userFilteredCollectionOffersMade(address , address ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) UserFilteredCollectionOffersMade(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paintswap.Contract.UserFilteredCollectionOffersMade(&_Paintswap.CallOpts, arg0, arg1)
}

// UserOffersMade is a free data retrieval call binding the contract method 0x66afec1f.
//
// Solidity: function userOffersMade(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCaller) UserOffersMade(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "userOffersMade", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOffersMade is a free data retrieval call binding the contract method 0x66afec1f.
//
// Solidity: function userOffersMade(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapSession) UserOffersMade(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.UserOffersMade(&_Paintswap.CallOpts, arg0, arg1, arg2)
}

// UserOffersMade is a free data retrieval call binding the contract method 0x66afec1f.
//
// Solidity: function userOffersMade(address , address , uint256 ) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) UserOffersMade(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.UserOffersMade(&_Paintswap.CallOpts, arg0, arg1, arg2)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x37b3630f.
//
// Solidity: function acceptOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _quantity, string _searchKeywords) returns()
func (_Paintswap *PaintswapTransactor) AcceptOffer(opts *bind.TransactOpts, _offerId *big.Int, _nft common.Address, _tokenId *big.Int, _quantity *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "acceptOffer", _offerId, _nft, _tokenId, _quantity, _searchKeywords)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x37b3630f.
//
// Solidity: function acceptOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _quantity, string _searchKeywords) returns()
func (_Paintswap *PaintswapSession) AcceptOffer(_offerId *big.Int, _nft common.Address, _tokenId *big.Int, _quantity *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOffer(&_Paintswap.TransactOpts, _offerId, _nft, _tokenId, _quantity, _searchKeywords)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x37b3630f.
//
// Solidity: function acceptOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _quantity, string _searchKeywords) returns()
func (_Paintswap *PaintswapTransactorSession) AcceptOffer(_offerId *big.Int, _nft common.Address, _tokenId *big.Int, _quantity *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOffer(&_Paintswap.TransactOpts, _offerId, _nft, _tokenId, _quantity, _searchKeywords)
}

// AcceptOfferBatch is a paid mutator transaction binding the contract method 0xb10b6338.
//
// Solidity: function acceptOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIdOrMarketplaceIds, bool[] _isMarketplaceIds, uint128[] _quantities, string[] _searchKeywords) returns()
func (_Paintswap *PaintswapTransactor) AcceptOfferBatch(opts *bind.TransactOpts, _offerIds []*big.Int, _nfts []common.Address, _tokenIdOrMarketplaceIds []*big.Int, _isMarketplaceIds []bool, _quantities []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "acceptOfferBatch", _offerIds, _nfts, _tokenIdOrMarketplaceIds, _isMarketplaceIds, _quantities, _searchKeywords)
}

// AcceptOfferBatch is a paid mutator transaction binding the contract method 0xb10b6338.
//
// Solidity: function acceptOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIdOrMarketplaceIds, bool[] _isMarketplaceIds, uint128[] _quantities, string[] _searchKeywords) returns()
func (_Paintswap *PaintswapSession) AcceptOfferBatch(_offerIds []*big.Int, _nfts []common.Address, _tokenIdOrMarketplaceIds []*big.Int, _isMarketplaceIds []bool, _quantities []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOfferBatch(&_Paintswap.TransactOpts, _offerIds, _nfts, _tokenIdOrMarketplaceIds, _isMarketplaceIds, _quantities, _searchKeywords)
}

// AcceptOfferBatch is a paid mutator transaction binding the contract method 0xb10b6338.
//
// Solidity: function acceptOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIdOrMarketplaceIds, bool[] _isMarketplaceIds, uint128[] _quantities, string[] _searchKeywords) returns()
func (_Paintswap *PaintswapTransactorSession) AcceptOfferBatch(_offerIds []*big.Int, _nfts []common.Address, _tokenIdOrMarketplaceIds []*big.Int, _isMarketplaceIds []bool, _quantities []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOfferBatch(&_Paintswap.TransactOpts, _offerIds, _nfts, _tokenIdOrMarketplaceIds, _isMarketplaceIds, _quantities, _searchKeywords)
}

// AcceptOnSaleOffer is a paid mutator transaction binding the contract method 0xd7e90894.
//
// Solidity: function acceptOnSaleOffer(uint256 _offerId, uint128 _quantity, uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) AcceptOnSaleOffer(opts *bind.TransactOpts, _offerId *big.Int, _quantity *big.Int, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "acceptOnSaleOffer", _offerId, _quantity, _marketplaceId)
}

// AcceptOnSaleOffer is a paid mutator transaction binding the contract method 0xd7e90894.
//
// Solidity: function acceptOnSaleOffer(uint256 _offerId, uint128 _quantity, uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) AcceptOnSaleOffer(_offerId *big.Int, _quantity *big.Int, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOnSaleOffer(&_Paintswap.TransactOpts, _offerId, _quantity, _marketplaceId)
}

// AcceptOnSaleOffer is a paid mutator transaction binding the contract method 0xd7e90894.
//
// Solidity: function acceptOnSaleOffer(uint256 _offerId, uint128 _quantity, uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) AcceptOnSaleOffer(_offerId *big.Int, _quantity *big.Int, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptOnSaleOffer(&_Paintswap.TransactOpts, _offerId, _quantity, _marketplaceId)
}

// AddSale is a paid mutator transaction binding the contract method 0x4aa8d113.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint64[] _times, uint128 _amount, bool[] _flags, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactor) AddSale(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []uint64, _amount *big.Int, _flags []bool, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "addSale", _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _searchKeywords, _addresses)
}

// AddSale is a paid mutator transaction binding the contract method 0x4aa8d113.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint64[] _times, uint128 _amount, bool[] _flags, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapSession) AddSale(_nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []uint64, _amount *big.Int, _flags []bool, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSale(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _searchKeywords, _addresses)
}

// AddSale is a paid mutator transaction binding the contract method 0x4aa8d113.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint64[] _times, uint128 _amount, bool[] _flags, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactorSession) AddSale(_nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []uint64, _amount *big.Int, _flags []bool, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSale(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _searchKeywords, _addresses)
}

// AddSaleBatch is a paid mutator transaction binding the contract method 0x180466f8.
//
// Solidity: function addSaleBatch(address[][] _nfts, uint256[][] _tokenIds, uint256[][] _amountBatches, uint128[] _prices, uint64[][] _times, uint128[] _amounts, bool[][] _flags, string[] _searchKeywords, address[][] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactor) AddSaleBatch(opts *bind.TransactOpts, _nfts [][]common.Address, _tokenIds [][]*big.Int, _amountBatches [][]*big.Int, _prices []*big.Int, _times [][]uint64, _amounts []*big.Int, _flags [][]bool, _searchKeywords []string, _addresses [][]common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "addSaleBatch", _nfts, _tokenIds, _amountBatches, _prices, _times, _amounts, _flags, _searchKeywords, _addresses)
}

// AddSaleBatch is a paid mutator transaction binding the contract method 0x180466f8.
//
// Solidity: function addSaleBatch(address[][] _nfts, uint256[][] _tokenIds, uint256[][] _amountBatches, uint128[] _prices, uint64[][] _times, uint128[] _amounts, bool[][] _flags, string[] _searchKeywords, address[][] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapSession) AddSaleBatch(_nfts [][]common.Address, _tokenIds [][]*big.Int, _amountBatches [][]*big.Int, _prices []*big.Int, _times [][]uint64, _amounts []*big.Int, _flags [][]bool, _searchKeywords []string, _addresses [][]common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSaleBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _prices, _times, _amounts, _flags, _searchKeywords, _addresses)
}

// AddSaleBatch is a paid mutator transaction binding the contract method 0x180466f8.
//
// Solidity: function addSaleBatch(address[][] _nfts, uint256[][] _tokenIds, uint256[][] _amountBatches, uint128[] _prices, uint64[][] _times, uint128[] _amounts, bool[][] _flags, string[] _searchKeywords, address[][] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactorSession) AddSaleBatch(_nfts [][]common.Address, _tokenIds [][]*big.Int, _amountBatches [][]*big.Int, _prices []*big.Int, _times [][]uint64, _amounts []*big.Int, _flags [][]bool, _searchKeywords []string, _addresses [][]common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSaleBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _prices, _times, _amounts, _flags, _searchKeywords, _addresses)
}

// AddValidVaultFactory is a paid mutator transaction binding the contract method 0x08e64e4d.
//
// Solidity: function addValidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapTransactor) AddValidVaultFactory(opts *bind.TransactOpts, _vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "addValidVaultFactory", _vaultFactory)
}

// AddValidVaultFactory is a paid mutator transaction binding the contract method 0x08e64e4d.
//
// Solidity: function addValidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapSession) AddValidVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddValidVaultFactory(&_Paintswap.TransactOpts, _vaultFactory)
}

// AddValidVaultFactory is a paid mutator transaction binding the contract method 0x08e64e4d.
//
// Solidity: function addValidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapTransactorSession) AddValidVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddValidVaultFactory(&_Paintswap.TransactOpts, _vaultFactory)
}

// Buy is a paid mutator transaction binding the contract method 0x87fbfdc7.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount) payable returns()
func (_Paintswap *PaintswapTransactor) Buy(opts *bind.TransactOpts, _marketplaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "buy", _marketplaceId, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0x87fbfdc7.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount) payable returns()
func (_Paintswap *PaintswapSession) Buy(_marketplaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.Buy(&_Paintswap.TransactOpts, _marketplaceId, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0x87fbfdc7.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount) payable returns()
func (_Paintswap *PaintswapTransactorSession) Buy(_marketplaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.Buy(&_Paintswap.TransactOpts, _marketplaceId, _amount)
}

// BuyBatch is a paid mutator transaction binding the contract method 0xe2a70fb4.
//
// Solidity: function buyBatch(uint256[] _marketplaceIds, uint128[] _amounts, bool _allowFailures) payable returns()
func (_Paintswap *PaintswapTransactor) BuyBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int, _amounts []*big.Int, _allowFailures bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "buyBatch", _marketplaceIds, _amounts, _allowFailures)
}

// BuyBatch is a paid mutator transaction binding the contract method 0xe2a70fb4.
//
// Solidity: function buyBatch(uint256[] _marketplaceIds, uint128[] _amounts, bool _allowFailures) payable returns()
func (_Paintswap *PaintswapSession) BuyBatch(_marketplaceIds []*big.Int, _amounts []*big.Int, _allowFailures bool) (*types.Transaction, error) {
	return _Paintswap.Contract.BuyBatch(&_Paintswap.TransactOpts, _marketplaceIds, _amounts, _allowFailures)
}

// BuyBatch is a paid mutator transaction binding the contract method 0xe2a70fb4.
//
// Solidity: function buyBatch(uint256[] _marketplaceIds, uint128[] _amounts, bool _allowFailures) payable returns()
func (_Paintswap *PaintswapTransactorSession) BuyBatch(_marketplaceIds []*big.Int, _amounts []*big.Int, _allowFailures bool) (*types.Transaction, error) {
	return _Paintswap.Contract.BuyBatch(&_Paintswap.TransactOpts, _marketplaceIds, _amounts, _allowFailures)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 _offerId) returns()
func (_Paintswap *PaintswapTransactor) CancelOffer(opts *bind.TransactOpts, _offerId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "cancelOffer", _offerId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 _offerId) returns()
func (_Paintswap *PaintswapSession) CancelOffer(_offerId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelOffer(&_Paintswap.TransactOpts, _offerId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 _offerId) returns()
func (_Paintswap *PaintswapTransactorSession) CancelOffer(_offerId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelOffer(&_Paintswap.TransactOpts, _offerId)
}

// CancelOfferBatch is a paid mutator transaction binding the contract method 0x9ef305d6.
//
// Solidity: function cancelOfferBatch(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapTransactor) CancelOfferBatch(opts *bind.TransactOpts, _offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "cancelOfferBatch", _offerIds)
}

// CancelOfferBatch is a paid mutator transaction binding the contract method 0x9ef305d6.
//
// Solidity: function cancelOfferBatch(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapSession) CancelOfferBatch(_offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelOfferBatch(&_Paintswap.TransactOpts, _offerIds)
}

// CancelOfferBatch is a paid mutator transaction binding the contract method 0x9ef305d6.
//
// Solidity: function cancelOfferBatch(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapTransactorSession) CancelOfferBatch(_offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelOfferBatch(&_Paintswap.TransactOpts, _offerIds)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) CancelSale(opts *bind.TransactOpts, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "cancelSale", _marketplaceId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) CancelSale(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSale(&_Paintswap.TransactOpts, _marketplaceId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) CancelSale(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSale(&_Paintswap.TransactOpts, _marketplaceId)
}

// CancelSaleBatch is a paid mutator transaction binding the contract method 0x1a745951.
//
// Solidity: function cancelSaleBatch(uint256[] _marketplaceIds) returns()
func (_Paintswap *PaintswapTransactor) CancelSaleBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "cancelSaleBatch", _marketplaceIds)
}

// CancelSaleBatch is a paid mutator transaction binding the contract method 0x1a745951.
//
// Solidity: function cancelSaleBatch(uint256[] _marketplaceIds) returns()
func (_Paintswap *PaintswapSession) CancelSaleBatch(_marketplaceIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSaleBatch(&_Paintswap.TransactOpts, _marketplaceIds)
}

// CancelSaleBatch is a paid mutator transaction binding the contract method 0x1a745951.
//
// Solidity: function cancelSaleBatch(uint256[] _marketplaceIds) returns()
func (_Paintswap *PaintswapTransactorSession) CancelSaleBatch(_marketplaceIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSaleBatch(&_Paintswap.TransactOpts, _marketplaceIds)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) CompleteMarketplaceEntry(opts *bind.TransactOpts, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "completeMarketplaceEntry", _marketplaceId)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) CompleteMarketplaceEntry(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CompleteMarketplaceEntry(&_Paintswap.TransactOpts, _marketplaceId)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) CompleteMarketplaceEntry(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CompleteMarketplaceEntry(&_Paintswap.TransactOpts, _marketplaceId)
}

// EditOffer is a paid mutator transaction binding the contract method 0x91f28218.
//
// Solidity: function editOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _newQuantity, uint128 _newOfferPrice, uint256 _newOfferDuration) payable returns()
func (_Paintswap *PaintswapTransactor) EditOffer(opts *bind.TransactOpts, _offerId *big.Int, _nft common.Address, _tokenId *big.Int, _newQuantity *big.Int, _newOfferPrice *big.Int, _newOfferDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editOffer", _offerId, _nft, _tokenId, _newQuantity, _newOfferPrice, _newOfferDuration)
}

// EditOffer is a paid mutator transaction binding the contract method 0x91f28218.
//
// Solidity: function editOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _newQuantity, uint128 _newOfferPrice, uint256 _newOfferDuration) payable returns()
func (_Paintswap *PaintswapSession) EditOffer(_offerId *big.Int, _nft common.Address, _tokenId *big.Int, _newQuantity *big.Int, _newOfferPrice *big.Int, _newOfferDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditOffer(&_Paintswap.TransactOpts, _offerId, _nft, _tokenId, _newQuantity, _newOfferPrice, _newOfferDuration)
}

// EditOffer is a paid mutator transaction binding the contract method 0x91f28218.
//
// Solidity: function editOffer(uint256 _offerId, address _nft, uint256 _tokenId, uint128 _newQuantity, uint128 _newOfferPrice, uint256 _newOfferDuration) payable returns()
func (_Paintswap *PaintswapTransactorSession) EditOffer(_offerId *big.Int, _nft common.Address, _tokenId *big.Int, _newQuantity *big.Int, _newOfferPrice *big.Int, _newOfferDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditOffer(&_Paintswap.TransactOpts, _offerId, _nft, _tokenId, _newQuantity, _newOfferPrice, _newOfferDuration)
}

// EditOfferBatch is a paid mutator transaction binding the contract method 0x28aca4c9.
//
// Solidity: function editOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIds, uint128[] _newQuantities, uint128[] _newOfferPrices, uint256[] _newOfferDurations) payable returns()
func (_Paintswap *PaintswapTransactor) EditOfferBatch(opts *bind.TransactOpts, _offerIds []*big.Int, _nfts []common.Address, _tokenIds []*big.Int, _newQuantities []*big.Int, _newOfferPrices []*big.Int, _newOfferDurations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editOfferBatch", _offerIds, _nfts, _tokenIds, _newQuantities, _newOfferPrices, _newOfferDurations)
}

// EditOfferBatch is a paid mutator transaction binding the contract method 0x28aca4c9.
//
// Solidity: function editOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIds, uint128[] _newQuantities, uint128[] _newOfferPrices, uint256[] _newOfferDurations) payable returns()
func (_Paintswap *PaintswapSession) EditOfferBatch(_offerIds []*big.Int, _nfts []common.Address, _tokenIds []*big.Int, _newQuantities []*big.Int, _newOfferPrices []*big.Int, _newOfferDurations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditOfferBatch(&_Paintswap.TransactOpts, _offerIds, _nfts, _tokenIds, _newQuantities, _newOfferPrices, _newOfferDurations)
}

// EditOfferBatch is a paid mutator transaction binding the contract method 0x28aca4c9.
//
// Solidity: function editOfferBatch(uint256[] _offerIds, address[] _nfts, uint256[] _tokenIds, uint128[] _newQuantities, uint128[] _newOfferPrices, uint256[] _newOfferDurations) payable returns()
func (_Paintswap *PaintswapTransactorSession) EditOfferBatch(_offerIds []*big.Int, _nfts []common.Address, _tokenIds []*big.Int, _newQuantities []*big.Int, _newOfferPrices []*big.Int, _newOfferDurations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditOfferBatch(&_Paintswap.TransactOpts, _offerIds, _nfts, _tokenIds, _newQuantities, _newOfferPrices, _newOfferDurations)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapTransactor) EditPrice(opts *bind.TransactOpts, _marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editPrice", _marketplaceId, _price)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapSession) EditPrice(_marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPrice(&_Paintswap.TransactOpts, _marketplaceId, _price)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapTransactorSession) EditPrice(_marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPrice(&_Paintswap.TransactOpts, _marketplaceId, _price)
}

// EditPriceAndQuantity is a paid mutator transaction binding the contract method 0x0cfc2532.
//
// Solidity: function editPriceAndQuantity(uint256 _marketplaceId, uint128 _newPrice, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapTransactor) EditPriceAndQuantity(opts *bind.TransactOpts, _marketplaceId *big.Int, _newPrice *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editPriceAndQuantity", _marketplaceId, _newPrice, _newQuantity)
}

// EditPriceAndQuantity is a paid mutator transaction binding the contract method 0x0cfc2532.
//
// Solidity: function editPriceAndQuantity(uint256 _marketplaceId, uint128 _newPrice, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapSession) EditPriceAndQuantity(_marketplaceId *big.Int, _newPrice *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceAndQuantity(&_Paintswap.TransactOpts, _marketplaceId, _newPrice, _newQuantity)
}

// EditPriceAndQuantity is a paid mutator transaction binding the contract method 0x0cfc2532.
//
// Solidity: function editPriceAndQuantity(uint256 _marketplaceId, uint128 _newPrice, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapTransactorSession) EditPriceAndQuantity(_marketplaceId *big.Int, _newPrice *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceAndQuantity(&_Paintswap.TransactOpts, _marketplaceId, _newPrice, _newQuantity)
}

// EditPriceAndQuantityBatch is a paid mutator transaction binding the contract method 0x5a005963.
//
// Solidity: function editPriceAndQuantityBatch(uint256[] _marketplaceId, uint128[] _newPrices, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapTransactor) EditPriceAndQuantityBatch(opts *bind.TransactOpts, _marketplaceId []*big.Int, _newPrices []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editPriceAndQuantityBatch", _marketplaceId, _newPrices, _newQuantities)
}

// EditPriceAndQuantityBatch is a paid mutator transaction binding the contract method 0x5a005963.
//
// Solidity: function editPriceAndQuantityBatch(uint256[] _marketplaceId, uint128[] _newPrices, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapSession) EditPriceAndQuantityBatch(_marketplaceId []*big.Int, _newPrices []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceAndQuantityBatch(&_Paintswap.TransactOpts, _marketplaceId, _newPrices, _newQuantities)
}

// EditPriceAndQuantityBatch is a paid mutator transaction binding the contract method 0x5a005963.
//
// Solidity: function editPriceAndQuantityBatch(uint256[] _marketplaceId, uint128[] _newPrices, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapTransactorSession) EditPriceAndQuantityBatch(_marketplaceId []*big.Int, _newPrices []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceAndQuantityBatch(&_Paintswap.TransactOpts, _marketplaceId, _newPrices, _newQuantities)
}

// EditPriceBatch is a paid mutator transaction binding the contract method 0x5bd0ed85.
//
// Solidity: function editPriceBatch(uint256[] _marketplaceIds, uint128[] _prices) returns()
func (_Paintswap *PaintswapTransactor) EditPriceBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int, _prices []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editPriceBatch", _marketplaceIds, _prices)
}

// EditPriceBatch is a paid mutator transaction binding the contract method 0x5bd0ed85.
//
// Solidity: function editPriceBatch(uint256[] _marketplaceIds, uint128[] _prices) returns()
func (_Paintswap *PaintswapSession) EditPriceBatch(_marketplaceIds []*big.Int, _prices []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceBatch(&_Paintswap.TransactOpts, _marketplaceIds, _prices)
}

// EditPriceBatch is a paid mutator transaction binding the contract method 0x5bd0ed85.
//
// Solidity: function editPriceBatch(uint256[] _marketplaceIds, uint128[] _prices) returns()
func (_Paintswap *PaintswapTransactorSession) EditPriceBatch(_marketplaceIds []*big.Int, _prices []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPriceBatch(&_Paintswap.TransactOpts, _marketplaceIds, _prices)
}

// EditQuantity is a paid mutator transaction binding the contract method 0x171b2172.
//
// Solidity: function editQuantity(uint256 _marketplaceId, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapTransactor) EditQuantity(opts *bind.TransactOpts, _marketplaceId *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editQuantity", _marketplaceId, _newQuantity)
}

// EditQuantity is a paid mutator transaction binding the contract method 0x171b2172.
//
// Solidity: function editQuantity(uint256 _marketplaceId, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapSession) EditQuantity(_marketplaceId *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditQuantity(&_Paintswap.TransactOpts, _marketplaceId, _newQuantity)
}

// EditQuantity is a paid mutator transaction binding the contract method 0x171b2172.
//
// Solidity: function editQuantity(uint256 _marketplaceId, uint128 _newQuantity) returns()
func (_Paintswap *PaintswapTransactorSession) EditQuantity(_marketplaceId *big.Int, _newQuantity *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditQuantity(&_Paintswap.TransactOpts, _marketplaceId, _newQuantity)
}

// EditQuantityBatch is a paid mutator transaction binding the contract method 0xd197911a.
//
// Solidity: function editQuantityBatch(uint256[] _marketplaceIds, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapTransactor) EditQuantityBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editQuantityBatch", _marketplaceIds, _newQuantities)
}

// EditQuantityBatch is a paid mutator transaction binding the contract method 0xd197911a.
//
// Solidity: function editQuantityBatch(uint256[] _marketplaceIds, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapSession) EditQuantityBatch(_marketplaceIds []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditQuantityBatch(&_Paintswap.TransactOpts, _marketplaceIds, _newQuantities)
}

// EditQuantityBatch is a paid mutator transaction binding the contract method 0xd197911a.
//
// Solidity: function editQuantityBatch(uint256[] _marketplaceIds, uint128[] _newQuantities) returns()
func (_Paintswap *PaintswapTransactorSession) EditQuantityBatch(_marketplaceIds []*big.Int, _newQuantities []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditQuantityBatch(&_Paintswap.TransactOpts, _marketplaceIds, _newQuantities)
}

// ExpireOffers is a paid mutator transaction binding the contract method 0xd58d360c.
//
// Solidity: function expireOffers(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapTransactor) ExpireOffers(opts *bind.TransactOpts, _offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "expireOffers", _offerIds)
}

// ExpireOffers is a paid mutator transaction binding the contract method 0xd58d360c.
//
// Solidity: function expireOffers(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapSession) ExpireOffers(_offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.ExpireOffers(&_Paintswap.TransactOpts, _offerIds)
}

// ExpireOffers is a paid mutator transaction binding the contract method 0xd58d360c.
//
// Solidity: function expireOffers(uint256[] _offerIds) returns()
func (_Paintswap *PaintswapTransactorSession) ExpireOffers(_offerIds []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.ExpireOffers(&_Paintswap.TransactOpts, _offerIds)
}

// MakeAltOffer is a paid mutator transaction binding the contract method 0xc0b64720.
//
// Solidity: function makeAltOffer(address[] _nfts, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeAltOffer(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeAltOffer", _nfts, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeAltOffer is a paid mutator transaction binding the contract method 0xc0b64720.
//
// Solidity: function makeAltOffer(address[] _nfts, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeAltOffer(_nfts []common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeAltOffer(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeAltOffer is a paid mutator transaction binding the contract method 0xc0b64720.
//
// Solidity: function makeAltOffer(address[] _nfts, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeAltOffer(_nfts []common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeAltOffer(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeBid is a paid mutator transaction binding the contract method 0x848b2138.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid) payable returns()
func (_Paintswap *PaintswapTransactor) MakeBid(opts *bind.TransactOpts, _marketplaceId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeBid", _marketplaceId, _bid)
}

// MakeBid is a paid mutator transaction binding the contract method 0x848b2138.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid) payable returns()
func (_Paintswap *PaintswapSession) MakeBid(_marketplaceId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBid(&_Paintswap.TransactOpts, _marketplaceId, _bid)
}

// MakeBid is a paid mutator transaction binding the contract method 0x848b2138.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeBid(_marketplaceId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBid(&_Paintswap.TransactOpts, _marketplaceId, _bid)
}

// MakeBidBatch is a paid mutator transaction binding the contract method 0x38d0096e.
//
// Solidity: function makeBidBatch(uint256[] _marketplaceIds, uint128[] _bids) payable returns()
func (_Paintswap *PaintswapTransactor) MakeBidBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int, _bids []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeBidBatch", _marketplaceIds, _bids)
}

// MakeBidBatch is a paid mutator transaction binding the contract method 0x38d0096e.
//
// Solidity: function makeBidBatch(uint256[] _marketplaceIds, uint128[] _bids) payable returns()
func (_Paintswap *PaintswapSession) MakeBidBatch(_marketplaceIds []*big.Int, _bids []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBidBatch(&_Paintswap.TransactOpts, _marketplaceIds, _bids)
}

// MakeBidBatch is a paid mutator transaction binding the contract method 0x38d0096e.
//
// Solidity: function makeBidBatch(uint256[] _marketplaceIds, uint128[] _bids) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeBidBatch(_marketplaceIds []*big.Int, _bids []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBidBatch(&_Paintswap.TransactOpts, _marketplaceIds, _bids)
}

// MakeCollectionOffer is a paid mutator transaction binding the contract method 0xc38f1fb9.
//
// Solidity: function makeCollectionOffer(address _nft, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeCollectionOffer(opts *bind.TransactOpts, _nft common.Address, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeCollectionOffer", _nft, _quantity, _price, _duration, _searchKeywords)
}

// MakeCollectionOffer is a paid mutator transaction binding the contract method 0xc38f1fb9.
//
// Solidity: function makeCollectionOffer(address _nft, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeCollectionOffer(_nft common.Address, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeCollectionOffer(&_Paintswap.TransactOpts, _nft, _quantity, _price, _duration, _searchKeywords)
}

// MakeCollectionOffer is a paid mutator transaction binding the contract method 0xc38f1fb9.
//
// Solidity: function makeCollectionOffer(address _nft, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeCollectionOffer(_nft common.Address, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeCollectionOffer(&_Paintswap.TransactOpts, _nft, _quantity, _price, _duration, _searchKeywords)
}

// MakeFilteredCollectionOffer is a paid mutator transaction binding the contract method 0xe48a2ba7.
//
// Solidity: function makeFilteredCollectionOffer(address _nft, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeFilteredCollectionOffer(opts *bind.TransactOpts, _nft common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeFilteredCollectionOffer", _nft, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeFilteredCollectionOffer is a paid mutator transaction binding the contract method 0xe48a2ba7.
//
// Solidity: function makeFilteredCollectionOffer(address _nft, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeFilteredCollectionOffer(_nft common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeFilteredCollectionOffer(&_Paintswap.TransactOpts, _nft, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeFilteredCollectionOffer is a paid mutator transaction binding the contract method 0xe48a2ba7.
//
// Solidity: function makeFilteredCollectionOffer(address _nft, uint256[] _tokenIds, uint128 _quantity, uint128[] _prices, uint256 _duration, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeFilteredCollectionOffer(_nft common.Address, _tokenIds []*big.Int, _quantity *big.Int, _prices []*big.Int, _duration *big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeFilteredCollectionOffer(&_Paintswap.TransactOpts, _nft, _tokenIds, _quantity, _prices, _duration, _searchKeywords)
}

// MakeFilteredCollectionOfferBatch is a paid mutator transaction binding the contract method 0x6f59270d.
//
// Solidity: function makeFilteredCollectionOfferBatch(address[] _nfts, uint256[][] _tokenIds, uint128[] _quantities, uint128[][] _prices, uint256[] _durations, string[][] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeFilteredCollectionOfferBatch(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds [][]*big.Int, _quantities []*big.Int, _prices [][]*big.Int, _durations []*big.Int, _searchKeywords [][]string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeFilteredCollectionOfferBatch", _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// MakeFilteredCollectionOfferBatch is a paid mutator transaction binding the contract method 0x6f59270d.
//
// Solidity: function makeFilteredCollectionOfferBatch(address[] _nfts, uint256[][] _tokenIds, uint128[] _quantities, uint128[][] _prices, uint256[] _durations, string[][] _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeFilteredCollectionOfferBatch(_nfts []common.Address, _tokenIds [][]*big.Int, _quantities []*big.Int, _prices [][]*big.Int, _durations []*big.Int, _searchKeywords [][]string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeFilteredCollectionOfferBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// MakeFilteredCollectionOfferBatch is a paid mutator transaction binding the contract method 0x6f59270d.
//
// Solidity: function makeFilteredCollectionOfferBatch(address[] _nfts, uint256[][] _tokenIds, uint128[] _quantities, uint128[][] _prices, uint256[] _durations, string[][] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeFilteredCollectionOfferBatch(_nfts []common.Address, _tokenIds [][]*big.Int, _quantities []*big.Int, _prices [][]*big.Int, _durations []*big.Int, _searchKeywords [][]string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeFilteredCollectionOfferBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xd9437799.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _quantity, uint128 _price, uint256 _duration) payable returns()
func (_Paintswap *PaintswapTransactor) MakeOffer(opts *bind.TransactOpts, _marketplaceId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeOffer", _marketplaceId, _quantity, _price, _duration)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xd9437799.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _quantity, uint128 _price, uint256 _duration) payable returns()
func (_Paintswap *PaintswapSession) MakeOffer(_marketplaceId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOffer(&_Paintswap.TransactOpts, _marketplaceId, _quantity, _price, _duration)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xd9437799.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _quantity, uint128 _price, uint256 _duration) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeOffer(_marketplaceId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOffer(&_Paintswap.TransactOpts, _marketplaceId, _quantity, _price, _duration)
}

// MakeOfferBatch is a paid mutator transaction binding the contract method 0xad11ee1b.
//
// Solidity: function makeOfferBatch(uint256[] _marketplaceIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations) payable returns()
func (_Paintswap *PaintswapTransactor) MakeOfferBatch(opts *bind.TransactOpts, _marketplaceIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeOfferBatch", _marketplaceIds, _quantities, _prices, _durations)
}

// MakeOfferBatch is a paid mutator transaction binding the contract method 0xad11ee1b.
//
// Solidity: function makeOfferBatch(uint256[] _marketplaceIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations) payable returns()
func (_Paintswap *PaintswapSession) MakeOfferBatch(_marketplaceIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferBatch(&_Paintswap.TransactOpts, _marketplaceIds, _quantities, _prices, _durations)
}

// MakeOfferBatch is a paid mutator transaction binding the contract method 0xad11ee1b.
//
// Solidity: function makeOfferBatch(uint256[] _marketplaceIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeOfferBatch(_marketplaceIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferBatch(&_Paintswap.TransactOpts, _marketplaceIds, _quantities, _prices, _durations)
}

// MakeOfferSingle is a paid mutator transaction binding the contract method 0x9676bdf8.
//
// Solidity: function makeOfferSingle(address _nft, uint256 _tokenId, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeOfferSingle(opts *bind.TransactOpts, _nft common.Address, _tokenId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeOfferSingle", _nft, _tokenId, _quantity, _price, _duration, _searchKeywords)
}

// MakeOfferSingle is a paid mutator transaction binding the contract method 0x9676bdf8.
//
// Solidity: function makeOfferSingle(address _nft, uint256 _tokenId, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeOfferSingle(_nft common.Address, _tokenId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferSingle(&_Paintswap.TransactOpts, _nft, _tokenId, _quantity, _price, _duration, _searchKeywords)
}

// MakeOfferSingle is a paid mutator transaction binding the contract method 0x9676bdf8.
//
// Solidity: function makeOfferSingle(address _nft, uint256 _tokenId, uint128 _quantity, uint128 _price, uint256 _duration, string _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeOfferSingle(_nft common.Address, _tokenId *big.Int, _quantity *big.Int, _price *big.Int, _duration *big.Int, _searchKeywords string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferSingle(&_Paintswap.TransactOpts, _nft, _tokenId, _quantity, _price, _duration, _searchKeywords)
}

// MakeOfferSingleBatch is a paid mutator transaction binding the contract method 0xcecf9c8f.
//
// Solidity: function makeOfferSingleBatch(address[] _nfts, uint256[] _tokenIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactor) MakeOfferSingleBatch(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeOfferSingleBatch", _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// MakeOfferSingleBatch is a paid mutator transaction binding the contract method 0xcecf9c8f.
//
// Solidity: function makeOfferSingleBatch(address[] _nfts, uint256[] _tokenIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapSession) MakeOfferSingleBatch(_nfts []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferSingleBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// MakeOfferSingleBatch is a paid mutator transaction binding the contract method 0xcecf9c8f.
//
// Solidity: function makeOfferSingleBatch(address[] _nfts, uint256[] _tokenIds, uint128[] _quantities, uint128[] _prices, uint256[] _durations, string[] _searchKeywords) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeOfferSingleBatch(_nfts []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _prices []*big.Int, _durations []*big.Int, _searchKeywords []string) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOfferSingleBatch(&_Paintswap.TransactOpts, _nfts, _tokenIds, _quantities, _prices, _durations, _searchKeywords)
}

// RemoveInvalidVaultFactory is a paid mutator transaction binding the contract method 0xe4286976.
//
// Solidity: function removeInvalidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapTransactor) RemoveInvalidVaultFactory(opts *bind.TransactOpts, _vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "removeInvalidVaultFactory", _vaultFactory)
}

// RemoveInvalidVaultFactory is a paid mutator transaction binding the contract method 0xe4286976.
//
// Solidity: function removeInvalidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapSession) RemoveInvalidVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.RemoveInvalidVaultFactory(&_Paintswap.TransactOpts, _vaultFactory)
}

// RemoveInvalidVaultFactory is a paid mutator transaction binding the contract method 0xe4286976.
//
// Solidity: function removeInvalidVaultFactory(address _vaultFactory) returns()
func (_Paintswap *PaintswapTransactorSession) RemoveInvalidVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.RemoveInvalidVaultFactory(&_Paintswap.TransactOpts, _vaultFactory)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paintswap.Contract.RenounceOwnership(&_Paintswap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paintswap.Contract.RenounceOwnership(&_Paintswap.TransactOpts)
}

// Reserve1 is a paid mutator transaction binding the contract method 0xd4cb6685.
//
// Solidity: function reserve1(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactor) Reserve1(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "reserve1", _data)
}

// Reserve1 is a paid mutator transaction binding the contract method 0xd4cb6685.
//
// Solidity: function reserve1(bytes _data) payable returns()
func (_Paintswap *PaintswapSession) Reserve1(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve1(&_Paintswap.TransactOpts, _data)
}

// Reserve1 is a paid mutator transaction binding the contract method 0xd4cb6685.
//
// Solidity: function reserve1(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactorSession) Reserve1(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve1(&_Paintswap.TransactOpts, _data)
}

// Reserve2 is a paid mutator transaction binding the contract method 0x61532ecc.
//
// Solidity: function reserve2(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactor) Reserve2(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "reserve2", _data)
}

// Reserve2 is a paid mutator transaction binding the contract method 0x61532ecc.
//
// Solidity: function reserve2(bytes _data) payable returns()
func (_Paintswap *PaintswapSession) Reserve2(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve2(&_Paintswap.TransactOpts, _data)
}

// Reserve2 is a paid mutator transaction binding the contract method 0x61532ecc.
//
// Solidity: function reserve2(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactorSession) Reserve2(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve2(&_Paintswap.TransactOpts, _data)
}

// Reserve3 is a paid mutator transaction binding the contract method 0x5f588c5a.
//
// Solidity: function reserve3(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactor) Reserve3(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "reserve3", _data)
}

// Reserve3 is a paid mutator transaction binding the contract method 0x5f588c5a.
//
// Solidity: function reserve3(bytes _data) payable returns()
func (_Paintswap *PaintswapSession) Reserve3(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve3(&_Paintswap.TransactOpts, _data)
}

// Reserve3 is a paid mutator transaction binding the contract method 0x5f588c5a.
//
// Solidity: function reserve3(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactorSession) Reserve3(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve3(&_Paintswap.TransactOpts, _data)
}

// Reserve4 is a paid mutator transaction binding the contract method 0xd6602b78.
//
// Solidity: function reserve4(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactor) Reserve4(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "reserve4", _data)
}

// Reserve4 is a paid mutator transaction binding the contract method 0xd6602b78.
//
// Solidity: function reserve4(bytes _data) payable returns()
func (_Paintswap *PaintswapSession) Reserve4(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve4(&_Paintswap.TransactOpts, _data)
}

// Reserve4 is a paid mutator transaction binding the contract method 0xd6602b78.
//
// Solidity: function reserve4(bytes _data) payable returns()
func (_Paintswap *PaintswapTransactorSession) Reserve4(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.Reserve4(&_Paintswap.TransactOpts, _data)
}

// ReserveOnlyOwner is a paid mutator transaction binding the contract method 0x22aa84c3.
//
// Solidity: function reserveOnlyOwner(bytes _data) returns()
func (_Paintswap *PaintswapTransactor) ReserveOnlyOwner(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "reserveOnlyOwner", _data)
}

// ReserveOnlyOwner is a paid mutator transaction binding the contract method 0x22aa84c3.
//
// Solidity: function reserveOnlyOwner(bytes _data) returns()
func (_Paintswap *PaintswapSession) ReserveOnlyOwner(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.ReserveOnlyOwner(&_Paintswap.TransactOpts, _data)
}

// ReserveOnlyOwner is a paid mutator transaction binding the contract method 0x22aa84c3.
//
// Solidity: function reserveOnlyOwner(bytes _data) returns()
func (_Paintswap *PaintswapTransactorSession) ReserveOnlyOwner(_data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.ReserveOnlyOwner(&_Paintswap.TransactOpts, _data)
}

// SafeNFTTransferBulk is a paid mutator transaction binding the contract method 0x6b39a92c.
//
// Solidity: function safeNFTTransferBulk(address[] _nfts, uint256[] _tokenIds, address[] _tos, uint256[] _amounts, uint256 _chainId, address[] _destinations, bytes _data) payable returns()
func (_Paintswap *PaintswapTransactor) SafeNFTTransferBulk(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds []*big.Int, _tos []common.Address, _amounts []*big.Int, _chainId *big.Int, _destinations []common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "safeNFTTransferBulk", _nfts, _tokenIds, _tos, _amounts, _chainId, _destinations, _data)
}

// SafeNFTTransferBulk is a paid mutator transaction binding the contract method 0x6b39a92c.
//
// Solidity: function safeNFTTransferBulk(address[] _nfts, uint256[] _tokenIds, address[] _tos, uint256[] _amounts, uint256 _chainId, address[] _destinations, bytes _data) payable returns()
func (_Paintswap *PaintswapSession) SafeNFTTransferBulk(_nfts []common.Address, _tokenIds []*big.Int, _tos []common.Address, _amounts []*big.Int, _chainId *big.Int, _destinations []common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.SafeNFTTransferBulk(&_Paintswap.TransactOpts, _nfts, _tokenIds, _tos, _amounts, _chainId, _destinations, _data)
}

// SafeNFTTransferBulk is a paid mutator transaction binding the contract method 0x6b39a92c.
//
// Solidity: function safeNFTTransferBulk(address[] _nfts, uint256[] _tokenIds, address[] _tos, uint256[] _amounts, uint256 _chainId, address[] _destinations, bytes _data) payable returns()
func (_Paintswap *PaintswapTransactorSession) SafeNFTTransferBulk(_nfts []common.Address, _tokenIds []*big.Int, _tos []common.Address, _amounts []*big.Int, _chainId *big.Int, _destinations []common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.SafeNFTTransferBulk(&_Paintswap.TransactOpts, _nfts, _tokenIds, _tos, _amounts, _chainId, _destinations, _data)
}

// SafeNFTTransferBulkOrdered is a paid mutator transaction binding the contract method 0x3574abaf.
//
// Solidity: function safeNFTTransferBulkOrdered((address,uint256[],uint256[],address)[] _nftsInfo, uint256 _chainId, address _destination, bytes _data) returns()
func (_Paintswap *PaintswapTransactor) SafeNFTTransferBulkOrdered(opts *bind.TransactOpts, _nftsInfo []PaintSwapMarketplaceV3StateBaseBulkTransferInfo, _chainId *big.Int, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "safeNFTTransferBulkOrdered", _nftsInfo, _chainId, _destination, _data)
}

// SafeNFTTransferBulkOrdered is a paid mutator transaction binding the contract method 0x3574abaf.
//
// Solidity: function safeNFTTransferBulkOrdered((address,uint256[],uint256[],address)[] _nftsInfo, uint256 _chainId, address _destination, bytes _data) returns()
func (_Paintswap *PaintswapSession) SafeNFTTransferBulkOrdered(_nftsInfo []PaintSwapMarketplaceV3StateBaseBulkTransferInfo, _chainId *big.Int, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.SafeNFTTransferBulkOrdered(&_Paintswap.TransactOpts, _nftsInfo, _chainId, _destination, _data)
}

// SafeNFTTransferBulkOrdered is a paid mutator transaction binding the contract method 0x3574abaf.
//
// Solidity: function safeNFTTransferBulkOrdered((address,uint256[],uint256[],address)[] _nftsInfo, uint256 _chainId, address _destination, bytes _data) returns()
func (_Paintswap *PaintswapTransactorSession) SafeNFTTransferBulkOrdered(_nftsInfo []PaintSwapMarketplaceV3StateBaseBulkTransferInfo, _chainId *big.Int, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.SafeNFTTransferBulkOrdered(&_Paintswap.TransactOpts, _nftsInfo, _chainId, _destination, _data)
}

// SetBundles is a paid mutator transaction binding the contract method 0xf4f1af1f.
//
// Solidity: function setBundles(uint128 _maxBundleNumber, bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetBundles(opts *bind.TransactOpts, _maxBundleNumber *big.Int, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setBundles", _maxBundleNumber, _enable)
}

// SetBundles is a paid mutator transaction binding the contract method 0xf4f1af1f.
//
// Solidity: function setBundles(uint128 _maxBundleNumber, bool _enable) returns()
func (_Paintswap *PaintswapSession) SetBundles(_maxBundleNumber *big.Int, _enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetBundles(&_Paintswap.TransactOpts, _maxBundleNumber, _enable)
}

// SetBundles is a paid mutator transaction binding the contract method 0xf4f1af1f.
//
// Solidity: function setBundles(uint128 _maxBundleNumber, bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetBundles(_maxBundleNumber *big.Int, _enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetBundles(&_Paintswap.TransactOpts, _maxBundleNumber, _enable)
}

// SetDurations is a paid mutator transaction binding the contract method 0xa4e44fdd.
//
// Solidity: function setDurations(uint256 _minDuration, uint256 _maxDuration, uint256 _minAuctionDuration, uint256 _minFlashAuctionDuration, uint256 _maxAuctionDuration) returns()
func (_Paintswap *PaintswapTransactor) SetDurations(opts *bind.TransactOpts, _minDuration *big.Int, _maxDuration *big.Int, _minAuctionDuration *big.Int, _minFlashAuctionDuration *big.Int, _maxAuctionDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setDurations", _minDuration, _maxDuration, _minAuctionDuration, _minFlashAuctionDuration, _maxAuctionDuration)
}

// SetDurations is a paid mutator transaction binding the contract method 0xa4e44fdd.
//
// Solidity: function setDurations(uint256 _minDuration, uint256 _maxDuration, uint256 _minAuctionDuration, uint256 _minFlashAuctionDuration, uint256 _maxAuctionDuration) returns()
func (_Paintswap *PaintswapSession) SetDurations(_minDuration *big.Int, _maxDuration *big.Int, _minAuctionDuration *big.Int, _minFlashAuctionDuration *big.Int, _maxAuctionDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDurations(&_Paintswap.TransactOpts, _minDuration, _maxDuration, _minAuctionDuration, _minFlashAuctionDuration, _maxAuctionDuration)
}

// SetDurations is a paid mutator transaction binding the contract method 0xa4e44fdd.
//
// Solidity: function setDurations(uint256 _minDuration, uint256 _maxDuration, uint256 _minAuctionDuration, uint256 _minFlashAuctionDuration, uint256 _maxAuctionDuration) returns()
func (_Paintswap *PaintswapTransactorSession) SetDurations(_minDuration *big.Int, _maxDuration *big.Int, _minAuctionDuration *big.Int, _minFlashAuctionDuration *big.Int, _maxAuctionDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDurations(&_Paintswap.TransactOpts, _minDuration, _maxDuration, _minAuctionDuration, _minFlashAuctionDuration, _maxAuctionDuration)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetEnableSelling(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setEnableSelling", _enable)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetEnableSelling(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEnableSelling(&_Paintswap.TransactOpts, _enable)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetEnableSelling(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEnableSelling(&_Paintswap.TransactOpts, _enable)
}

// SetFeeAddresses is a paid mutator transaction binding the contract method 0x11c84120.
//
// Solidity: function setFeeAddresses(address _dev, address _salesFee) returns()
func (_Paintswap *PaintswapTransactor) SetFeeAddresses(opts *bind.TransactOpts, _dev common.Address, _salesFee common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setFeeAddresses", _dev, _salesFee)
}

// SetFeeAddresses is a paid mutator transaction binding the contract method 0x11c84120.
//
// Solidity: function setFeeAddresses(address _dev, address _salesFee) returns()
func (_Paintswap *PaintswapSession) SetFeeAddresses(_dev common.Address, _salesFee common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetFeeAddresses(&_Paintswap.TransactOpts, _dev, _salesFee)
}

// SetFeeAddresses is a paid mutator transaction binding the contract method 0x11c84120.
//
// Solidity: function setFeeAddresses(address _dev, address _salesFee) returns()
func (_Paintswap *PaintswapTransactorSession) SetFeeAddresses(_dev common.Address, _salesFee common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetFeeAddresses(&_Paintswap.TransactOpts, _dev, _salesFee)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0x0c697fdf.
//
// Solidity: function setFeePercentage(address _feePercentage) returns()
func (_Paintswap *PaintswapTransactor) SetFeePercentage(opts *bind.TransactOpts, _feePercentage common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setFeePercentage", _feePercentage)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0x0c697fdf.
//
// Solidity: function setFeePercentage(address _feePercentage) returns()
func (_Paintswap *PaintswapSession) SetFeePercentage(_feePercentage common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetFeePercentage(&_Paintswap.TransactOpts, _feePercentage)
}

// SetFeePercentage is a paid mutator transaction binding the contract method 0x0c697fdf.
//
// Solidity: function setFeePercentage(address _feePercentage) returns()
func (_Paintswap *PaintswapTransactorSession) SetFeePercentage(_feePercentage common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetFeePercentage(&_Paintswap.TransactOpts, _feePercentage)
}

// SetImplBatchTransferNFT is a paid mutator transaction binding the contract method 0x36453ddf.
//
// Solidity: function setImplBatchTransferNFT(address _implBatchTransferNFT) returns()
func (_Paintswap *PaintswapTransactor) SetImplBatchTransferNFT(opts *bind.TransactOpts, _implBatchTransferNFT common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setImplBatchTransferNFT", _implBatchTransferNFT)
}

// SetImplBatchTransferNFT is a paid mutator transaction binding the contract method 0x36453ddf.
//
// Solidity: function setImplBatchTransferNFT(address _implBatchTransferNFT) returns()
func (_Paintswap *PaintswapSession) SetImplBatchTransferNFT(_implBatchTransferNFT common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBatchTransferNFT(&_Paintswap.TransactOpts, _implBatchTransferNFT)
}

// SetImplBatchTransferNFT is a paid mutator transaction binding the contract method 0x36453ddf.
//
// Solidity: function setImplBatchTransferNFT(address _implBatchTransferNFT) returns()
func (_Paintswap *PaintswapTransactorSession) SetImplBatchTransferNFT(_implBatchTransferNFT common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBatchTransferNFT(&_Paintswap.TransactOpts, _implBatchTransferNFT)
}

// SetImplBidsAndAcceptOffers is a paid mutator transaction binding the contract method 0xb1388dbf.
//
// Solidity: function setImplBidsAndAcceptOffers(address _implBidsAndAcceptOffers) returns()
func (_Paintswap *PaintswapTransactor) SetImplBidsAndAcceptOffers(opts *bind.TransactOpts, _implBidsAndAcceptOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setImplBidsAndAcceptOffers", _implBidsAndAcceptOffers)
}

// SetImplBidsAndAcceptOffers is a paid mutator transaction binding the contract method 0xb1388dbf.
//
// Solidity: function setImplBidsAndAcceptOffers(address _implBidsAndAcceptOffers) returns()
func (_Paintswap *PaintswapSession) SetImplBidsAndAcceptOffers(_implBidsAndAcceptOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBidsAndAcceptOffers(&_Paintswap.TransactOpts, _implBidsAndAcceptOffers)
}

// SetImplBidsAndAcceptOffers is a paid mutator transaction binding the contract method 0xb1388dbf.
//
// Solidity: function setImplBidsAndAcceptOffers(address _implBidsAndAcceptOffers) returns()
func (_Paintswap *PaintswapTransactorSession) SetImplBidsAndAcceptOffers(_implBidsAndAcceptOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBidsAndAcceptOffers(&_Paintswap.TransactOpts, _implBidsAndAcceptOffers)
}

// SetImplBuyEditListing is a paid mutator transaction binding the contract method 0x9716eafa.
//
// Solidity: function setImplBuyEditListing(address _implBuyEditListing) returns()
func (_Paintswap *PaintswapTransactor) SetImplBuyEditListing(opts *bind.TransactOpts, _implBuyEditListing common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setImplBuyEditListing", _implBuyEditListing)
}

// SetImplBuyEditListing is a paid mutator transaction binding the contract method 0x9716eafa.
//
// Solidity: function setImplBuyEditListing(address _implBuyEditListing) returns()
func (_Paintswap *PaintswapSession) SetImplBuyEditListing(_implBuyEditListing common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBuyEditListing(&_Paintswap.TransactOpts, _implBuyEditListing)
}

// SetImplBuyEditListing is a paid mutator transaction binding the contract method 0x9716eafa.
//
// Solidity: function setImplBuyEditListing(address _implBuyEditListing) returns()
func (_Paintswap *PaintswapTransactorSession) SetImplBuyEditListing(_implBuyEditListing common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplBuyEditListing(&_Paintswap.TransactOpts, _implBuyEditListing)
}

// SetImplListAndComplete is a paid mutator transaction binding the contract method 0x802807ba.
//
// Solidity: function setImplListAndComplete(address _implListAndComplete) returns()
func (_Paintswap *PaintswapTransactor) SetImplListAndComplete(opts *bind.TransactOpts, _implListAndComplete common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setImplListAndComplete", _implListAndComplete)
}

// SetImplListAndComplete is a paid mutator transaction binding the contract method 0x802807ba.
//
// Solidity: function setImplListAndComplete(address _implListAndComplete) returns()
func (_Paintswap *PaintswapSession) SetImplListAndComplete(_implListAndComplete common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplListAndComplete(&_Paintswap.TransactOpts, _implListAndComplete)
}

// SetImplListAndComplete is a paid mutator transaction binding the contract method 0x802807ba.
//
// Solidity: function setImplListAndComplete(address _implListAndComplete) returns()
func (_Paintswap *PaintswapTransactorSession) SetImplListAndComplete(_implListAndComplete common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetImplListAndComplete(&_Paintswap.TransactOpts, _implListAndComplete)
}

// SetMaxActiveOfferCountAndMinVaultVersion is a paid mutator transaction binding the contract method 0x482d3fb3.
//
// Solidity: function setMaxActiveOfferCountAndMinVaultVersion(uint256 _maxActiveOfferCount, uint256 _minVaultVersion) returns()
func (_Paintswap *PaintswapTransactor) SetMaxActiveOfferCountAndMinVaultVersion(opts *bind.TransactOpts, _maxActiveOfferCount *big.Int, _minVaultVersion *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMaxActiveOfferCountAndMinVaultVersion", _maxActiveOfferCount, _minVaultVersion)
}

// SetMaxActiveOfferCountAndMinVaultVersion is a paid mutator transaction binding the contract method 0x482d3fb3.
//
// Solidity: function setMaxActiveOfferCountAndMinVaultVersion(uint256 _maxActiveOfferCount, uint256 _minVaultVersion) returns()
func (_Paintswap *PaintswapSession) SetMaxActiveOfferCountAndMinVaultVersion(_maxActiveOfferCount *big.Int, _minVaultVersion *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxActiveOfferCountAndMinVaultVersion(&_Paintswap.TransactOpts, _maxActiveOfferCount, _minVaultVersion)
}

// SetMaxActiveOfferCountAndMinVaultVersion is a paid mutator transaction binding the contract method 0x482d3fb3.
//
// Solidity: function setMaxActiveOfferCountAndMinVaultVersion(uint256 _maxActiveOfferCount, uint256 _minVaultVersion) returns()
func (_Paintswap *PaintswapTransactorSession) SetMaxActiveOfferCountAndMinVaultVersion(_maxActiveOfferCount *big.Int, _minVaultVersion *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxActiveOfferCountAndMinVaultVersion(&_Paintswap.TransactOpts, _maxActiveOfferCount, _minVaultVersion)
}

// SetOfferDuration is a paid mutator transaction binding the contract method 0xdd991466.
//
// Solidity: function setOfferDuration(uint256 _minDuration, uint256 _maxDuration) returns()
func (_Paintswap *PaintswapTransactor) SetOfferDuration(opts *bind.TransactOpts, _minDuration *big.Int, _maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setOfferDuration", _minDuration, _maxDuration)
}

// SetOfferDuration is a paid mutator transaction binding the contract method 0xdd991466.
//
// Solidity: function setOfferDuration(uint256 _minDuration, uint256 _maxDuration) returns()
func (_Paintswap *PaintswapSession) SetOfferDuration(_minDuration *big.Int, _maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOfferDuration(&_Paintswap.TransactOpts, _minDuration, _maxDuration)
}

// SetOfferDuration is a paid mutator transaction binding the contract method 0xdd991466.
//
// Solidity: function setOfferDuration(uint256 _minDuration, uint256 _maxDuration) returns()
func (_Paintswap *PaintswapTransactorSession) SetOfferDuration(_minDuration *big.Int, _maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOfferDuration(&_Paintswap.TransactOpts, _minDuration, _maxDuration)
}

// SetVarious is a paid mutator transaction binding the contract method 0xebd13475.
//
// Solidity: function setVarious(uint256 _maxRoyalty, uint64 _maxStartTimeIncrement, address _officialNFTDiscount, uint256 _callGasLimit, address _paintRouter) returns()
func (_Paintswap *PaintswapTransactor) SetVarious(opts *bind.TransactOpts, _maxRoyalty *big.Int, _maxStartTimeIncrement uint64, _officialNFTDiscount common.Address, _callGasLimit *big.Int, _paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setVarious", _maxRoyalty, _maxStartTimeIncrement, _officialNFTDiscount, _callGasLimit, _paintRouter)
}

// SetVarious is a paid mutator transaction binding the contract method 0xebd13475.
//
// Solidity: function setVarious(uint256 _maxRoyalty, uint64 _maxStartTimeIncrement, address _officialNFTDiscount, uint256 _callGasLimit, address _paintRouter) returns()
func (_Paintswap *PaintswapSession) SetVarious(_maxRoyalty *big.Int, _maxStartTimeIncrement uint64, _officialNFTDiscount common.Address, _callGasLimit *big.Int, _paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetVarious(&_Paintswap.TransactOpts, _maxRoyalty, _maxStartTimeIncrement, _officialNFTDiscount, _callGasLimit, _paintRouter)
}

// SetVarious is a paid mutator transaction binding the contract method 0xebd13475.
//
// Solidity: function setVarious(uint256 _maxRoyalty, uint64 _maxStartTimeIncrement, address _officialNFTDiscount, uint256 _callGasLimit, address _paintRouter) returns()
func (_Paintswap *PaintswapTransactorSession) SetVarious(_maxRoyalty *big.Int, _maxStartTimeIncrement uint64, _officialNFTDiscount common.Address, _callGasLimit *big.Int, _paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetVarious(&_Paintswap.TransactOpts, _maxRoyalty, _maxStartTimeIncrement, _officialNFTDiscount, _callGasLimit, _paintRouter)
}

// SetVariousTimes is a paid mutator transaction binding the contract method 0x8c8b7f0c.
//
// Solidity: function setVariousTimes(uint64 _auctionGracePeriodForCancelling, uint64 _flashAuctionGracePeriodForCancelling, uint64 _auctionEndTimeIncreaseCutOff, uint64 _flashAuctionEndTimeIncreaseCutOff, uint128 _minIncreasedBidPercent, bool _modifyInitialStartTime, bool _modifyPostStartTime) returns()
func (_Paintswap *PaintswapTransactor) SetVariousTimes(opts *bind.TransactOpts, _auctionGracePeriodForCancelling uint64, _flashAuctionGracePeriodForCancelling uint64, _auctionEndTimeIncreaseCutOff uint64, _flashAuctionEndTimeIncreaseCutOff uint64, _minIncreasedBidPercent *big.Int, _modifyInitialStartTime bool, _modifyPostStartTime bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setVariousTimes", _auctionGracePeriodForCancelling, _flashAuctionGracePeriodForCancelling, _auctionEndTimeIncreaseCutOff, _flashAuctionEndTimeIncreaseCutOff, _minIncreasedBidPercent, _modifyInitialStartTime, _modifyPostStartTime)
}

// SetVariousTimes is a paid mutator transaction binding the contract method 0x8c8b7f0c.
//
// Solidity: function setVariousTimes(uint64 _auctionGracePeriodForCancelling, uint64 _flashAuctionGracePeriodForCancelling, uint64 _auctionEndTimeIncreaseCutOff, uint64 _flashAuctionEndTimeIncreaseCutOff, uint128 _minIncreasedBidPercent, bool _modifyInitialStartTime, bool _modifyPostStartTime) returns()
func (_Paintswap *PaintswapSession) SetVariousTimes(_auctionGracePeriodForCancelling uint64, _flashAuctionGracePeriodForCancelling uint64, _auctionEndTimeIncreaseCutOff uint64, _flashAuctionEndTimeIncreaseCutOff uint64, _minIncreasedBidPercent *big.Int, _modifyInitialStartTime bool, _modifyPostStartTime bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetVariousTimes(&_Paintswap.TransactOpts, _auctionGracePeriodForCancelling, _flashAuctionGracePeriodForCancelling, _auctionEndTimeIncreaseCutOff, _flashAuctionEndTimeIncreaseCutOff, _minIncreasedBidPercent, _modifyInitialStartTime, _modifyPostStartTime)
}

// SetVariousTimes is a paid mutator transaction binding the contract method 0x8c8b7f0c.
//
// Solidity: function setVariousTimes(uint64 _auctionGracePeriodForCancelling, uint64 _flashAuctionGracePeriodForCancelling, uint64 _auctionEndTimeIncreaseCutOff, uint64 _flashAuctionEndTimeIncreaseCutOff, uint128 _minIncreasedBidPercent, bool _modifyInitialStartTime, bool _modifyPostStartTime) returns()
func (_Paintswap *PaintswapTransactorSession) SetVariousTimes(_auctionGracePeriodForCancelling uint64, _flashAuctionGracePeriodForCancelling uint64, _auctionEndTimeIncreaseCutOff uint64, _flashAuctionEndTimeIncreaseCutOff uint64, _minIncreasedBidPercent *big.Int, _modifyInitialStartTime bool, _modifyPostStartTime bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetVariousTimes(&_Paintswap.TransactOpts, _auctionGracePeriodForCancelling, _flashAuctionGracePeriodForCancelling, _auctionEndTimeIncreaseCutOff, _flashAuctionEndTimeIncreaseCutOff, _minIncreasedBidPercent, _modifyInitialStartTime, _modifyPostStartTime)
}

// SetimplMakeAndEditOffers is a paid mutator transaction binding the contract method 0x9b43ce50.
//
// Solidity: function setimplMakeAndEditOffers(address _implMakeAndEditOffers) returns()
func (_Paintswap *PaintswapTransactor) SetimplMakeAndEditOffers(opts *bind.TransactOpts, _implMakeAndEditOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setimplMakeAndEditOffers", _implMakeAndEditOffers)
}

// SetimplMakeAndEditOffers is a paid mutator transaction binding the contract method 0x9b43ce50.
//
// Solidity: function setimplMakeAndEditOffers(address _implMakeAndEditOffers) returns()
func (_Paintswap *PaintswapSession) SetimplMakeAndEditOffers(_implMakeAndEditOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetimplMakeAndEditOffers(&_Paintswap.TransactOpts, _implMakeAndEditOffers)
}

// SetimplMakeAndEditOffers is a paid mutator transaction binding the contract method 0x9b43ce50.
//
// Solidity: function setimplMakeAndEditOffers(address _implMakeAndEditOffers) returns()
func (_Paintswap *PaintswapTransactorSession) SetimplMakeAndEditOffers(_implMakeAndEditOffers common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetimplMakeAndEditOffers(&_Paintswap.TransactOpts, _implMakeAndEditOffers)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.TransferOwnership(&_Paintswap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.TransferOwnership(&_Paintswap.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapSession) Receive() (*types.Transaction, error) {
	return _Paintswap.Contract.Receive(&_Paintswap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapTransactorSession) Receive() (*types.Transaction, error) {
	return _Paintswap.Contract.Receive(&_Paintswap.TransactOpts)
}

// PaintswapAddedVaultFactoryIterator is returned from FilterAddedVaultFactory and is used to iterate over the raw logs and unpacked data for AddedVaultFactory events raised by the Paintswap contract.
type PaintswapAddedVaultFactoryIterator struct {
	Event *PaintswapAddedVaultFactory // Event containing the contract specifics and raw log

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
func (it *PaintswapAddedVaultFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapAddedVaultFactory)
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
		it.Event = new(PaintswapAddedVaultFactory)
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
func (it *PaintswapAddedVaultFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapAddedVaultFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapAddedVaultFactory represents a AddedVaultFactory event raised by the Paintswap contract.
type PaintswapAddedVaultFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedVaultFactory is a free log retrieval operation binding the contract event 0x5ace8eeba0e775f3a9e6f61d3983c537568596af7574020e4e8320551aa312b0.
//
// Solidity: event AddedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) FilterAddedVaultFactory(opts *bind.FilterOpts) (*PaintswapAddedVaultFactoryIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "AddedVaultFactory")
	if err != nil {
		return nil, err
	}
	return &PaintswapAddedVaultFactoryIterator{contract: _Paintswap.contract, event: "AddedVaultFactory", logs: logs, sub: sub}, nil
}

// WatchAddedVaultFactory is a free log subscription operation binding the contract event 0x5ace8eeba0e775f3a9e6f61d3983c537568596af7574020e4e8320551aa312b0.
//
// Solidity: event AddedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) WatchAddedVaultFactory(opts *bind.WatchOpts, sink chan<- *PaintswapAddedVaultFactory) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "AddedVaultFactory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapAddedVaultFactory)
				if err := _Paintswap.contract.UnpackLog(event, "AddedVaultFactory", log); err != nil {
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

// ParseAddedVaultFactory is a log parse operation binding the contract event 0x5ace8eeba0e775f3a9e6f61d3983c537568596af7574020e4e8320551aa312b0.
//
// Solidity: event AddedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) ParseAddedVaultFactory(log types.Log) (*PaintswapAddedVaultFactory, error) {
	event := new(PaintswapAddedVaultFactory)
	if err := _Paintswap.contract.UnpackLog(event, "AddedVaultFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapCancelledSaleIterator is returned from FilterCancelledSale and is used to iterate over the raw logs and unpacked data for CancelledSale events raised by the Paintswap contract.
type PaintswapCancelledSaleIterator struct {
	Event *PaintswapCancelledSale // Event containing the contract specifics and raw log

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
func (it *PaintswapCancelledSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapCancelledSale)
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
		it.Event = new(PaintswapCancelledSale)
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
func (it *PaintswapCancelledSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapCancelledSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapCancelledSale represents a CancelledSale event raised by the Paintswap contract.
type PaintswapCancelledSale struct {
	MarketplaceId *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	AmountBatches []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCancelledSale is a free log retrieval operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) FilterCancelledSale(opts *bind.FilterOpts) (*PaintswapCancelledSaleIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "CancelledSale")
	if err != nil {
		return nil, err
	}
	return &PaintswapCancelledSaleIterator{contract: _Paintswap.contract, event: "CancelledSale", logs: logs, sub: sub}, nil
}

// WatchCancelledSale is a free log subscription operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) WatchCancelledSale(opts *bind.WatchOpts, sink chan<- *PaintswapCancelledSale) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "CancelledSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapCancelledSale)
				if err := _Paintswap.contract.UnpackLog(event, "CancelledSale", log); err != nil {
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

// ParseCancelledSale is a log parse operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) ParseCancelledSale(log types.Log) (*PaintswapCancelledSale, error) {
	event := new(PaintswapCancelledSale)
	if err := _Paintswap.contract.UnpackLog(event, "CancelledSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Paintswap contract.
type PaintswapNewBidIterator struct {
	Event *PaintswapNewBid // Event containing the contract specifics and raw log

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
func (it *PaintswapNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewBid)
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
		it.Event = new(PaintswapNewBid)
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
func (it *PaintswapNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewBid represents a NewBid event raised by the Paintswap contract.
type PaintswapNewBid struct {
	MarketplaceId *big.Int
	Bidder        common.Address
	Bid           *big.Int
	NextMinimum   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) FilterNewBid(opts *bind.FilterOpts, bidder []common.Address) (*PaintswapNewBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewBid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewBidIterator{contract: _Paintswap.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *PaintswapNewBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewBid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewBid)
				if err := _Paintswap.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) ParseNewBid(log types.Log) (*PaintswapNewBid, error) {
	event := new(PaintswapNewBid)
	if err := _Paintswap.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewCollectionOfferIterator is returned from FilterNewCollectionOffer and is used to iterate over the raw logs and unpacked data for NewCollectionOffer events raised by the Paintswap contract.
type PaintswapNewCollectionOfferIterator struct {
	Event *PaintswapNewCollectionOffer // Event containing the contract specifics and raw log

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
func (it *PaintswapNewCollectionOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewCollectionOffer)
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
		it.Event = new(PaintswapNewCollectionOffer)
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
func (it *PaintswapNewCollectionOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewCollectionOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewCollectionOffer represents a NewCollectionOffer event raised by the Paintswap contract.
type PaintswapNewCollectionOffer struct {
	OfferId        *big.Int
	Nft            common.Address
	From           common.Address
	Quantity       *big.Int
	Price          *big.Int
	Expires        *big.Int
	SearchKeywords string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewCollectionOffer is a free log retrieval operation binding the contract event 0xed4ba8a9a3b1024c3cac6bb66f839e0f55267e422b107dbc69f3540d4efef595.
//
// Solidity: event NewCollectionOffer(uint256 offerId, address nft, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) FilterNewCollectionOffer(opts *bind.FilterOpts, from []common.Address) (*PaintswapNewCollectionOfferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewCollectionOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewCollectionOfferIterator{contract: _Paintswap.contract, event: "NewCollectionOffer", logs: logs, sub: sub}, nil
}

// WatchNewCollectionOffer is a free log subscription operation binding the contract event 0xed4ba8a9a3b1024c3cac6bb66f839e0f55267e422b107dbc69f3540d4efef595.
//
// Solidity: event NewCollectionOffer(uint256 offerId, address nft, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) WatchNewCollectionOffer(opts *bind.WatchOpts, sink chan<- *PaintswapNewCollectionOffer, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewCollectionOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewCollectionOffer)
				if err := _Paintswap.contract.UnpackLog(event, "NewCollectionOffer", log); err != nil {
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

// ParseNewCollectionOffer is a log parse operation binding the contract event 0xed4ba8a9a3b1024c3cac6bb66f839e0f55267e422b107dbc69f3540d4efef595.
//
// Solidity: event NewCollectionOffer(uint256 offerId, address nft, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) ParseNewCollectionOffer(log types.Log) (*PaintswapNewCollectionOffer, error) {
	event := new(PaintswapNewCollectionOffer)
	if err := _Paintswap.contract.UnpackLog(event, "NewCollectionOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewFilteredCollectionOfferIterator is returned from FilterNewFilteredCollectionOffer and is used to iterate over the raw logs and unpacked data for NewFilteredCollectionOffer events raised by the Paintswap contract.
type PaintswapNewFilteredCollectionOfferIterator struct {
	Event *PaintswapNewFilteredCollectionOffer // Event containing the contract specifics and raw log

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
func (it *PaintswapNewFilteredCollectionOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewFilteredCollectionOffer)
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
		it.Event = new(PaintswapNewFilteredCollectionOffer)
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
func (it *PaintswapNewFilteredCollectionOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewFilteredCollectionOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewFilteredCollectionOffer represents a NewFilteredCollectionOffer event raised by the Paintswap contract.
type PaintswapNewFilteredCollectionOffer struct {
	OfferId        *big.Int
	Nft            common.Address
	TokenIds       []*big.Int
	From           common.Address
	Quantity       *big.Int
	Price          *big.Int
	Prices         []*big.Int
	Expires        *big.Int
	SearchKeywords []string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewFilteredCollectionOffer is a free log retrieval operation binding the contract event 0x2030e9ce00344962fd6da99a2beb0742a184908b3fb402b9ab5bd18e48a42a2f.
//
// Solidity: event NewFilteredCollectionOffer(uint256 offerId, address nft, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint128[] prices, uint256 expires, string[] searchKeywords)
func (_Paintswap *PaintswapFilterer) FilterNewFilteredCollectionOffer(opts *bind.FilterOpts, from []common.Address) (*PaintswapNewFilteredCollectionOfferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewFilteredCollectionOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewFilteredCollectionOfferIterator{contract: _Paintswap.contract, event: "NewFilteredCollectionOffer", logs: logs, sub: sub}, nil
}

// WatchNewFilteredCollectionOffer is a free log subscription operation binding the contract event 0x2030e9ce00344962fd6da99a2beb0742a184908b3fb402b9ab5bd18e48a42a2f.
//
// Solidity: event NewFilteredCollectionOffer(uint256 offerId, address nft, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint128[] prices, uint256 expires, string[] searchKeywords)
func (_Paintswap *PaintswapFilterer) WatchNewFilteredCollectionOffer(opts *bind.WatchOpts, sink chan<- *PaintswapNewFilteredCollectionOffer, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewFilteredCollectionOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewFilteredCollectionOffer)
				if err := _Paintswap.contract.UnpackLog(event, "NewFilteredCollectionOffer", log); err != nil {
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

// ParseNewFilteredCollectionOffer is a log parse operation binding the contract event 0x2030e9ce00344962fd6da99a2beb0742a184908b3fb402b9ab5bd18e48a42a2f.
//
// Solidity: event NewFilteredCollectionOffer(uint256 offerId, address nft, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint128[] prices, uint256 expires, string[] searchKeywords)
func (_Paintswap *PaintswapFilterer) ParseNewFilteredCollectionOffer(log types.Log) (*PaintswapNewFilteredCollectionOffer, error) {
	event := new(PaintswapNewFilteredCollectionOffer)
	if err := _Paintswap.contract.UnpackLog(event, "NewFilteredCollectionOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the Paintswap contract.
type PaintswapNewOfferIterator struct {
	Event *PaintswapNewOffer // Event containing the contract specifics and raw log

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
func (it *PaintswapNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewOffer)
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
		it.Event = new(PaintswapNewOffer)
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
func (it *PaintswapNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewOffer represents a NewOffer event raised by the Paintswap contract.
type PaintswapNewOffer struct {
	OfferId        *big.Int
	MarketplaceId  *big.Int
	Nfts           []common.Address
	TokenIds       []*big.Int
	From           common.Address
	Quantity       *big.Int
	Price          *big.Int
	Expires        *big.Int
	SearchKeywords string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x9701d3b731f36eb8db6e75d8b08a00d396a034dad7436293b1725f268788deeb.
//
// Solidity: event NewOffer(uint256 offerId, uint256 marketplaceId, address[] nfts, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) FilterNewOffer(opts *bind.FilterOpts, from []common.Address) (*PaintswapNewOfferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewOfferIterator{contract: _Paintswap.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x9701d3b731f36eb8db6e75d8b08a00d396a034dad7436293b1725f268788deeb.
//
// Solidity: event NewOffer(uint256 offerId, uint256 marketplaceId, address[] nfts, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *PaintswapNewOffer, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewOffer", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewOffer)
				if err := _Paintswap.contract.UnpackLog(event, "NewOffer", log); err != nil {
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

// ParseNewOffer is a log parse operation binding the contract event 0x9701d3b731f36eb8db6e75d8b08a00d396a034dad7436293b1725f268788deeb.
//
// Solidity: event NewOffer(uint256 offerId, uint256 marketplaceId, address[] nfts, uint256[] tokenIds, address indexed from, uint128 quantity, uint128 price, uint256 expires, string searchKeywords)
func (_Paintswap *PaintswapFilterer) ParseNewOffer(log types.Log) (*PaintswapNewOffer, error) {
	event := new(PaintswapNewOffer)
	if err := _Paintswap.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewSaleIterator is returned from FilterNewSale and is used to iterate over the raw logs and unpacked data for NewSale events raised by the Paintswap contract.
type PaintswapNewSaleIterator struct {
	Event *PaintswapNewSale // Event containing the contract specifics and raw log

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
func (it *PaintswapNewSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewSale)
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
		it.Event = new(PaintswapNewSale)
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
func (it *PaintswapNewSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewSale represents a NewSale event raised by the Paintswap contract.
type PaintswapNewSale struct {
	MarketplaceId  *big.Int
	Nfts           []common.Address
	TokenIds       []*big.Int
	AmountBatches  []*big.Int
	Price          *big.Int
	Duration       *big.Int
	IsAuction      bool
	Antisnipe      bool
	FlashAuction   bool
	Amount         *big.Int
	IsNSFW         bool
	SearchKeywords string
	Addresses      []common.Address
	Seller         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewSale is a free log retrieval operation binding the contract event 0xeb482f0c68d0e97481a4ccffcaf0d0021ee556e4e8fb1c3a849d60e62f2f9867.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) FilterNewSale(opts *bind.FilterOpts) (*PaintswapNewSaleIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewSale")
	if err != nil {
		return nil, err
	}
	return &PaintswapNewSaleIterator{contract: _Paintswap.contract, event: "NewSale", logs: logs, sub: sub}, nil
}

// WatchNewSale is a free log subscription operation binding the contract event 0xeb482f0c68d0e97481a4ccffcaf0d0021ee556e4e8fb1c3a849d60e62f2f9867.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) WatchNewSale(opts *bind.WatchOpts, sink chan<- *PaintswapNewSale) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewSale)
				if err := _Paintswap.contract.UnpackLog(event, "NewSale", log); err != nil {
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

// ParseNewSale is a log parse operation binding the contract event 0xeb482f0c68d0e97481a4ccffcaf0d0021ee556e4e8fb1c3a849d60e62f2f9867.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) ParseNewSale(log types.Log) (*PaintswapNewSale, error) {
	event := new(PaintswapNewSale)
	if err := _Paintswap.contract.UnpackLog(event, "NewSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewSaleBatchIterator is returned from FilterNewSaleBatch and is used to iterate over the raw logs and unpacked data for NewSaleBatch events raised by the Paintswap contract.
type PaintswapNewSaleBatchIterator struct {
	Event *PaintswapNewSaleBatch // Event containing the contract specifics and raw log

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
func (it *PaintswapNewSaleBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewSaleBatch)
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
		it.Event = new(PaintswapNewSaleBatch)
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
func (it *PaintswapNewSaleBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewSaleBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewSaleBatch represents a NewSaleBatch event raised by the Paintswap contract.
type PaintswapNewSaleBatch struct {
	MarketplaceId  []*big.Int
	Nfts           [][]common.Address
	TokenIds       [][]*big.Int
	AmountBatches  [][]*big.Int
	Prices         []*big.Int
	Durations      []*big.Int
	IsAuctions     []bool
	IsAntisnipes   []bool
	Amounts        []*big.Int
	IsNSFWs        []bool
	SearchKeywords []string
	Addresses      [][]common.Address
	Sellers        []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewSaleBatch is a free log retrieval operation binding the contract event 0x8335c4aedd471cf3a80d02f474b40448ddfc3cc65aed721d8c636c9ef472710b.
//
// Solidity: event NewSaleBatch(uint256[] marketplaceId, address[][] nfts, uint256[][] tokenIds, uint256[][] amountBatches, uint128[] prices, uint256[] durations, bool[] isAuctions, bool[] isAntisnipes, uint256[] amounts, bool[] isNSFWs, string[] searchKeywords, address[][] addresses, address[] sellers)
func (_Paintswap *PaintswapFilterer) FilterNewSaleBatch(opts *bind.FilterOpts) (*PaintswapNewSaleBatchIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewSaleBatch")
	if err != nil {
		return nil, err
	}
	return &PaintswapNewSaleBatchIterator{contract: _Paintswap.contract, event: "NewSaleBatch", logs: logs, sub: sub}, nil
}

// WatchNewSaleBatch is a free log subscription operation binding the contract event 0x8335c4aedd471cf3a80d02f474b40448ddfc3cc65aed721d8c636c9ef472710b.
//
// Solidity: event NewSaleBatch(uint256[] marketplaceId, address[][] nfts, uint256[][] tokenIds, uint256[][] amountBatches, uint128[] prices, uint256[] durations, bool[] isAuctions, bool[] isAntisnipes, uint256[] amounts, bool[] isNSFWs, string[] searchKeywords, address[][] addresses, address[] sellers)
func (_Paintswap *PaintswapFilterer) WatchNewSaleBatch(opts *bind.WatchOpts, sink chan<- *PaintswapNewSaleBatch) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewSaleBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewSaleBatch)
				if err := _Paintswap.contract.UnpackLog(event, "NewSaleBatch", log); err != nil {
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

// ParseNewSaleBatch is a log parse operation binding the contract event 0x8335c4aedd471cf3a80d02f474b40448ddfc3cc65aed721d8c636c9ef472710b.
//
// Solidity: event NewSaleBatch(uint256[] marketplaceId, address[][] nfts, uint256[][] tokenIds, uint256[][] amountBatches, uint128[] prices, uint256[] durations, bool[] isAuctions, bool[] isAntisnipes, uint256[] amounts, bool[] isNSFWs, string[] searchKeywords, address[][] addresses, address[] sellers)
func (_Paintswap *PaintswapFilterer) ParseNewSaleBatch(log types.Log) (*PaintswapNewSaleBatch, error) {
	event := new(PaintswapNewSaleBatch)
	if err := _Paintswap.contract.UnpackLog(event, "NewSaleBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewTempListingIterator is returned from FilterNewTempListing and is used to iterate over the raw logs and unpacked data for NewTempListing events raised by the Paintswap contract.
type PaintswapNewTempListingIterator struct {
	Event *PaintswapNewTempListing // Event containing the contract specifics and raw log

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
func (it *PaintswapNewTempListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewTempListing)
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
		it.Event = new(PaintswapNewTempListing)
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
func (it *PaintswapNewTempListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewTempListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewTempListing represents a NewTempListing event raised by the Paintswap contract.
type PaintswapNewTempListing struct {
	MarketplaceId  *big.Int
	Nfts           []common.Address
	TokenIds       []*big.Int
	AmountBatches  []*big.Int
	Price          *big.Int
	Duration       *big.Int
	IsAuction      bool
	Antisnipe      bool
	FlashAuction   bool
	Amount         *big.Int
	IsNSFW         bool
	SearchKeywords string
	Addresses      []common.Address
	Seller         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTempListing is a free log retrieval operation binding the contract event 0x906f9a452d885b01be0c3ff320ae205e2a7208ceb14f13c7343434b60a2e0877.
//
// Solidity: event NewTempListing(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) FilterNewTempListing(opts *bind.FilterOpts) (*PaintswapNewTempListingIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewTempListing")
	if err != nil {
		return nil, err
	}
	return &PaintswapNewTempListingIterator{contract: _Paintswap.contract, event: "NewTempListing", logs: logs, sub: sub}, nil
}

// WatchNewTempListing is a free log subscription operation binding the contract event 0x906f9a452d885b01be0c3ff320ae205e2a7208ceb14f13c7343434b60a2e0877.
//
// Solidity: event NewTempListing(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) WatchNewTempListing(opts *bind.WatchOpts, sink chan<- *PaintswapNewTempListing) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewTempListing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewTempListing)
				if err := _Paintswap.contract.UnpackLog(event, "NewTempListing", log); err != nil {
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

// ParseNewTempListing is a log parse operation binding the contract event 0x906f9a452d885b01be0c3ff320ae205e2a7208ceb14f13c7343434b60a2e0877.
//
// Solidity: event NewTempListing(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, bool antisnipe, bool flashAuction, uint256 amount, bool isNSFW, string searchKeywords, address[] addresses, address seller)
func (_Paintswap *PaintswapFilterer) ParseNewTempListing(log types.Log) (*PaintswapNewTempListing, error) {
	event := new(PaintswapNewTempListing)
	if err := _Paintswap.contract.UnpackLog(event, "NewTempListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOfferAcceptedIterator is returned from FilterOfferAccepted and is used to iterate over the raw logs and unpacked data for OfferAccepted events raised by the Paintswap contract.
type PaintswapOfferAcceptedIterator struct {
	Event *PaintswapOfferAccepted // Event containing the contract specifics and raw log

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
func (it *PaintswapOfferAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOfferAccepted)
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
		it.Event = new(PaintswapOfferAccepted)
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
func (it *PaintswapOfferAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOfferAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOfferAccepted represents a OfferAccepted event raised by the Paintswap contract.
type PaintswapOfferAccepted struct {
	OfferId       *big.Int
	Nft           common.Address
	TokenId       *big.Int
	Quantity      *big.Int
	MarketplaceId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOfferAccepted is a free log retrieval operation binding the contract event 0x88514d2c723a1277b312444baaa1f2bd2c53dd6355c2cc51d349d04998055e45.
//
// Solidity: event OfferAccepted(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) FilterOfferAccepted(opts *bind.FilterOpts) (*PaintswapOfferAcceptedIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return &PaintswapOfferAcceptedIterator{contract: _Paintswap.contract, event: "OfferAccepted", logs: logs, sub: sub}, nil
}

// WatchOfferAccepted is a free log subscription operation binding the contract event 0x88514d2c723a1277b312444baaa1f2bd2c53dd6355c2cc51d349d04998055e45.
//
// Solidity: event OfferAccepted(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) WatchOfferAccepted(opts *bind.WatchOpts, sink chan<- *PaintswapOfferAccepted) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOfferAccepted)
				if err := _Paintswap.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
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

// ParseOfferAccepted is a log parse operation binding the contract event 0x88514d2c723a1277b312444baaa1f2bd2c53dd6355c2cc51d349d04998055e45.
//
// Solidity: event OfferAccepted(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) ParseOfferAccepted(log types.Log) (*PaintswapOfferAccepted, error) {
	event := new(PaintswapOfferAccepted)
	if err := _Paintswap.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOfferCompletedIterator is returned from FilterOfferCompleted and is used to iterate over the raw logs and unpacked data for OfferCompleted events raised by the Paintswap contract.
type PaintswapOfferCompletedIterator struct {
	Event *PaintswapOfferCompleted // Event containing the contract specifics and raw log

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
func (it *PaintswapOfferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOfferCompleted)
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
		it.Event = new(PaintswapOfferCompleted)
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
func (it *PaintswapOfferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOfferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOfferCompleted represents a OfferCompleted event raised by the Paintswap contract.
type PaintswapOfferCompleted struct {
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCompleted is a free log retrieval operation binding the contract event 0x422b793202eb794642b1017af34abd55c83b22253a8323bfb35132b8cf8d5cc3.
//
// Solidity: event OfferCompleted(uint256 offerId)
func (_Paintswap *PaintswapFilterer) FilterOfferCompleted(opts *bind.FilterOpts) (*PaintswapOfferCompletedIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OfferCompleted")
	if err != nil {
		return nil, err
	}
	return &PaintswapOfferCompletedIterator{contract: _Paintswap.contract, event: "OfferCompleted", logs: logs, sub: sub}, nil
}

// WatchOfferCompleted is a free log subscription operation binding the contract event 0x422b793202eb794642b1017af34abd55c83b22253a8323bfb35132b8cf8d5cc3.
//
// Solidity: event OfferCompleted(uint256 offerId)
func (_Paintswap *PaintswapFilterer) WatchOfferCompleted(opts *bind.WatchOpts, sink chan<- *PaintswapOfferCompleted) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OfferCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOfferCompleted)
				if err := _Paintswap.contract.UnpackLog(event, "OfferCompleted", log); err != nil {
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

// ParseOfferCompleted is a log parse operation binding the contract event 0x422b793202eb794642b1017af34abd55c83b22253a8323bfb35132b8cf8d5cc3.
//
// Solidity: event OfferCompleted(uint256 offerId)
func (_Paintswap *PaintswapFilterer) ParseOfferCompleted(log types.Log) (*PaintswapOfferCompleted, error) {
	event := new(PaintswapOfferCompleted)
	if err := _Paintswap.contract.UnpackLog(event, "OfferCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOfferExpiredIterator is returned from FilterOfferExpired and is used to iterate over the raw logs and unpacked data for OfferExpired events raised by the Paintswap contract.
type PaintswapOfferExpiredIterator struct {
	Event *PaintswapOfferExpired // Event containing the contract specifics and raw log

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
func (it *PaintswapOfferExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOfferExpired)
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
		it.Event = new(PaintswapOfferExpired)
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
func (it *PaintswapOfferExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOfferExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOfferExpired represents a OfferExpired event raised by the Paintswap contract.
type PaintswapOfferExpired struct {
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferExpired is a free log retrieval operation binding the contract event 0x48e57f5f7de6743b5c1c749447ff980262d36d0ad5142c942523d48dc821a171.
//
// Solidity: event OfferExpired(uint256 offerId)
func (_Paintswap *PaintswapFilterer) FilterOfferExpired(opts *bind.FilterOpts) (*PaintswapOfferExpiredIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OfferExpired")
	if err != nil {
		return nil, err
	}
	return &PaintswapOfferExpiredIterator{contract: _Paintswap.contract, event: "OfferExpired", logs: logs, sub: sub}, nil
}

// WatchOfferExpired is a free log subscription operation binding the contract event 0x48e57f5f7de6743b5c1c749447ff980262d36d0ad5142c942523d48dc821a171.
//
// Solidity: event OfferExpired(uint256 offerId)
func (_Paintswap *PaintswapFilterer) WatchOfferExpired(opts *bind.WatchOpts, sink chan<- *PaintswapOfferExpired) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OfferExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOfferExpired)
				if err := _Paintswap.contract.UnpackLog(event, "OfferExpired", log); err != nil {
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

// ParseOfferExpired is a log parse operation binding the contract event 0x48e57f5f7de6743b5c1c749447ff980262d36d0ad5142c942523d48dc821a171.
//
// Solidity: event OfferExpired(uint256 offerId)
func (_Paintswap *PaintswapFilterer) ParseOfferExpired(log types.Log) (*PaintswapOfferExpired, error) {
	event := new(PaintswapOfferExpired)
	if err := _Paintswap.contract.UnpackLog(event, "OfferExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOfferRemovedIterator is returned from FilterOfferRemoved and is used to iterate over the raw logs and unpacked data for OfferRemoved events raised by the Paintswap contract.
type PaintswapOfferRemovedIterator struct {
	Event *PaintswapOfferRemoved // Event containing the contract specifics and raw log

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
func (it *PaintswapOfferRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOfferRemoved)
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
		it.Event = new(PaintswapOfferRemoved)
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
func (it *PaintswapOfferRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOfferRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOfferRemoved represents a OfferRemoved event raised by the Paintswap contract.
type PaintswapOfferRemoved struct {
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferRemoved is a free log retrieval operation binding the contract event 0xe78251393d13c4f887458774ea505a80414b5608a0c963bb2e75da45811919f1.
//
// Solidity: event OfferRemoved(uint256 offerId)
func (_Paintswap *PaintswapFilterer) FilterOfferRemoved(opts *bind.FilterOpts) (*PaintswapOfferRemovedIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OfferRemoved")
	if err != nil {
		return nil, err
	}
	return &PaintswapOfferRemovedIterator{contract: _Paintswap.contract, event: "OfferRemoved", logs: logs, sub: sub}, nil
}

// WatchOfferRemoved is a free log subscription operation binding the contract event 0xe78251393d13c4f887458774ea505a80414b5608a0c963bb2e75da45811919f1.
//
// Solidity: event OfferRemoved(uint256 offerId)
func (_Paintswap *PaintswapFilterer) WatchOfferRemoved(opts *bind.WatchOpts, sink chan<- *PaintswapOfferRemoved) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OfferRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOfferRemoved)
				if err := _Paintswap.contract.UnpackLog(event, "OfferRemoved", log); err != nil {
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

// ParseOfferRemoved is a log parse operation binding the contract event 0xe78251393d13c4f887458774ea505a80414b5608a0c963bb2e75da45811919f1.
//
// Solidity: event OfferRemoved(uint256 offerId)
func (_Paintswap *PaintswapFilterer) ParseOfferRemoved(log types.Log) (*PaintswapOfferRemoved, error) {
	event := new(PaintswapOfferRemoved)
	if err := _Paintswap.contract.UnpackLog(event, "OfferRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paintswap contract.
type PaintswapOwnershipTransferredIterator struct {
	Event *PaintswapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaintswapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOwnershipTransferred)
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
		it.Event = new(PaintswapOwnershipTransferred)
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
func (it *PaintswapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOwnershipTransferred represents a OwnershipTransferred event raised by the Paintswap contract.
type PaintswapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paintswap *PaintswapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaintswapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapOwnershipTransferredIterator{contract: _Paintswap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paintswap *PaintswapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaintswapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOwnershipTransferred)
				if err := _Paintswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Paintswap *PaintswapFilterer) ParseOwnershipTransferred(log types.Log) (*PaintswapOwnershipTransferred, error) {
	event := new(PaintswapOwnershipTransferred)
	if err := _Paintswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapRemoveExpiredListingIterator is returned from FilterRemoveExpiredListing and is used to iterate over the raw logs and unpacked data for RemoveExpiredListing events raised by the Paintswap contract.
type PaintswapRemoveExpiredListingIterator struct {
	Event *PaintswapRemoveExpiredListing // Event containing the contract specifics and raw log

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
func (it *PaintswapRemoveExpiredListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapRemoveExpiredListing)
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
		it.Event = new(PaintswapRemoveExpiredListing)
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
func (it *PaintswapRemoveExpiredListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapRemoveExpiredListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapRemoveExpiredListing represents a RemoveExpiredListing event raised by the Paintswap contract.
type PaintswapRemoveExpiredListing struct {
	MarketplaceId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveExpiredListing is a free log retrieval operation binding the contract event 0xbfb5649267bb589f5c7a9de1ac683ba7ccde4ecd4498163193978c99d309a977.
//
// Solidity: event RemoveExpiredListing(uint256 _marketplaceId)
func (_Paintswap *PaintswapFilterer) FilterRemoveExpiredListing(opts *bind.FilterOpts) (*PaintswapRemoveExpiredListingIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "RemoveExpiredListing")
	if err != nil {
		return nil, err
	}
	return &PaintswapRemoveExpiredListingIterator{contract: _Paintswap.contract, event: "RemoveExpiredListing", logs: logs, sub: sub}, nil
}

// WatchRemoveExpiredListing is a free log subscription operation binding the contract event 0xbfb5649267bb589f5c7a9de1ac683ba7ccde4ecd4498163193978c99d309a977.
//
// Solidity: event RemoveExpiredListing(uint256 _marketplaceId)
func (_Paintswap *PaintswapFilterer) WatchRemoveExpiredListing(opts *bind.WatchOpts, sink chan<- *PaintswapRemoveExpiredListing) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "RemoveExpiredListing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapRemoveExpiredListing)
				if err := _Paintswap.contract.UnpackLog(event, "RemoveExpiredListing", log); err != nil {
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

// ParseRemoveExpiredListing is a log parse operation binding the contract event 0xbfb5649267bb589f5c7a9de1ac683ba7ccde4ecd4498163193978c99d309a977.
//
// Solidity: event RemoveExpiredListing(uint256 _marketplaceId)
func (_Paintswap *PaintswapFilterer) ParseRemoveExpiredListing(log types.Log) (*PaintswapRemoveExpiredListing, error) {
	event := new(PaintswapRemoveExpiredListing)
	if err := _Paintswap.contract.UnpackLog(event, "RemoveExpiredListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapRemovedVaultFactoryIterator is returned from FilterRemovedVaultFactory and is used to iterate over the raw logs and unpacked data for RemovedVaultFactory events raised by the Paintswap contract.
type PaintswapRemovedVaultFactoryIterator struct {
	Event *PaintswapRemovedVaultFactory // Event containing the contract specifics and raw log

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
func (it *PaintswapRemovedVaultFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapRemovedVaultFactory)
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
		it.Event = new(PaintswapRemovedVaultFactory)
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
func (it *PaintswapRemovedVaultFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapRemovedVaultFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapRemovedVaultFactory represents a RemovedVaultFactory event raised by the Paintswap contract.
type PaintswapRemovedVaultFactory struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedVaultFactory is a free log retrieval operation binding the contract event 0x0a8c6b50709426bea2be9b24815cf77e3b6635f6a2c752c44eb90b3b59ca8b11.
//
// Solidity: event RemovedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) FilterRemovedVaultFactory(opts *bind.FilterOpts) (*PaintswapRemovedVaultFactoryIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "RemovedVaultFactory")
	if err != nil {
		return nil, err
	}
	return &PaintswapRemovedVaultFactoryIterator{contract: _Paintswap.contract, event: "RemovedVaultFactory", logs: logs, sub: sub}, nil
}

// WatchRemovedVaultFactory is a free log subscription operation binding the contract event 0x0a8c6b50709426bea2be9b24815cf77e3b6635f6a2c752c44eb90b3b59ca8b11.
//
// Solidity: event RemovedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) WatchRemovedVaultFactory(opts *bind.WatchOpts, sink chan<- *PaintswapRemovedVaultFactory) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "RemovedVaultFactory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapRemovedVaultFactory)
				if err := _Paintswap.contract.UnpackLog(event, "RemovedVaultFactory", log); err != nil {
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

// ParseRemovedVaultFactory is a log parse operation binding the contract event 0x0a8c6b50709426bea2be9b24815cf77e3b6635f6a2c752c44eb90b3b59ca8b11.
//
// Solidity: event RemovedVaultFactory(address _factory)
func (_Paintswap *PaintswapFilterer) ParseRemovedVaultFactory(log types.Log) (*PaintswapRemovedVaultFactory, error) {
	event := new(PaintswapRemovedVaultFactory)
	if err := _Paintswap.contract.UnpackLog(event, "RemovedVaultFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapSaleFinishedIterator is returned from FilterSaleFinished and is used to iterate over the raw logs and unpacked data for SaleFinished events raised by the Paintswap contract.
type PaintswapSaleFinishedIterator struct {
	Event *PaintswapSaleFinished // Event containing the contract specifics and raw log

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
func (it *PaintswapSaleFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapSaleFinished)
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
		it.Event = new(PaintswapSaleFinished)
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
func (it *PaintswapSaleFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapSaleFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapSaleFinished represents a SaleFinished event raised by the Paintswap contract.
type PaintswapSaleFinished struct {
	MarketplaceId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSaleFinished is a free log retrieval operation binding the contract event 0xb068776b457cefe34568f0e1ffb0cc2480cf799e77410c955f6e7502d654655d.
//
// Solidity: event SaleFinished(uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) FilterSaleFinished(opts *bind.FilterOpts) (*PaintswapSaleFinishedIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "SaleFinished")
	if err != nil {
		return nil, err
	}
	return &PaintswapSaleFinishedIterator{contract: _Paintswap.contract, event: "SaleFinished", logs: logs, sub: sub}, nil
}

// WatchSaleFinished is a free log subscription operation binding the contract event 0xb068776b457cefe34568f0e1ffb0cc2480cf799e77410c955f6e7502d654655d.
//
// Solidity: event SaleFinished(uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) WatchSaleFinished(opts *bind.WatchOpts, sink chan<- *PaintswapSaleFinished) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "SaleFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapSaleFinished)
				if err := _Paintswap.contract.UnpackLog(event, "SaleFinished", log); err != nil {
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

// ParseSaleFinished is a log parse operation binding the contract event 0xb068776b457cefe34568f0e1ffb0cc2480cf799e77410c955f6e7502d654655d.
//
// Solidity: event SaleFinished(uint256 marketplaceId)
func (_Paintswap *PaintswapFilterer) ParseSaleFinished(log types.Log) (*PaintswapSaleFinished, error) {
	event := new(PaintswapSaleFinished)
	if err := _Paintswap.contract.UnpackLog(event, "SaleFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapSoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the Paintswap contract.
type PaintswapSoldIterator struct {
	Event *PaintswapSold // Event containing the contract specifics and raw log

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
func (it *PaintswapSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapSold)
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
		it.Event = new(PaintswapSold)
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
func (it *PaintswapSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapSold represents a Sold event raised by the Paintswap contract.
type PaintswapSold struct {
	MarketplaceId *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	AmountBatches []*big.Int
	Price         *big.Int
	Buyer         common.Address
	Seller        common.Address
	Amount        *big.Int
	OfferId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0x13687c09723c1541422d20cf795b479c42f8dbd4d4b06f46f5b25bd66192ea00.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount, uint256 offerId)
func (_Paintswap *PaintswapFilterer) FilterSold(opts *bind.FilterOpts) (*PaintswapSoldIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return &PaintswapSoldIterator{contract: _Paintswap.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0x13687c09723c1541422d20cf795b479c42f8dbd4d4b06f46f5b25bd66192ea00.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount, uint256 offerId)
func (_Paintswap *PaintswapFilterer) WatchSold(opts *bind.WatchOpts, sink chan<- *PaintswapSold) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapSold)
				if err := _Paintswap.contract.UnpackLog(event, "Sold", log); err != nil {
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

// ParseSold is a log parse operation binding the contract event 0x13687c09723c1541422d20cf795b479c42f8dbd4d4b06f46f5b25bd66192ea00.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount, uint256 offerId)
func (_Paintswap *PaintswapFilterer) ParseSold(log types.Log) (*PaintswapSold, error) {
	event := new(PaintswapSold)
	if err := _Paintswap.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateEndTimeIterator is returned from FilterUpdateEndTime and is used to iterate over the raw logs and unpacked data for UpdateEndTime events raised by the Paintswap contract.
type PaintswapUpdateEndTimeIterator struct {
	Event *PaintswapUpdateEndTime // Event containing the contract specifics and raw log

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
func (it *PaintswapUpdateEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateEndTime)
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
		it.Event = new(PaintswapUpdateEndTime)
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
func (it *PaintswapUpdateEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateEndTime represents a UpdateEndTime event raised by the Paintswap contract.
type PaintswapUpdateEndTime struct {
	MarketplaceId *big.Int
	EndTime       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateEndTime is a free log retrieval operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) FilterUpdateEndTime(opts *bind.FilterOpts) (*PaintswapUpdateEndTimeIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateEndTime")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateEndTimeIterator{contract: _Paintswap.contract, event: "UpdateEndTime", logs: logs, sub: sub}, nil
}

// WatchUpdateEndTime is a free log subscription operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) WatchUpdateEndTime(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateEndTime) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateEndTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateEndTime)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateEndTime", log); err != nil {
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

// ParseUpdateEndTime is a log parse operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) ParseUpdateEndTime(log types.Log) (*PaintswapUpdateEndTime, error) {
	event := new(PaintswapUpdateEndTime)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateOfferIterator is returned from FilterUpdateOffer and is used to iterate over the raw logs and unpacked data for UpdateOffer events raised by the Paintswap contract.
type PaintswapUpdateOfferIterator struct {
	Event *PaintswapUpdateOffer // Event containing the contract specifics and raw log

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
func (it *PaintswapUpdateOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateOffer)
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
		it.Event = new(PaintswapUpdateOffer)
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
func (it *PaintswapUpdateOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateOffer represents a UpdateOffer event raised by the Paintswap contract.
type PaintswapUpdateOffer struct {
	OfferId  *big.Int
	Nft      common.Address
	TokenId  *big.Int
	Quantity *big.Int
	NewPrice *big.Int
	Expires  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOffer is a free log retrieval operation binding the contract event 0x5f63f69e4dbc29f15e8c72167eaa48adf30c77a889c1bea4dfe05403edaaee8f.
//
// Solidity: event UpdateOffer(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint128 newPrice, uint256 expires)
func (_Paintswap *PaintswapFilterer) FilterUpdateOffer(opts *bind.FilterOpts) (*PaintswapUpdateOfferIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateOffer")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateOfferIterator{contract: _Paintswap.contract, event: "UpdateOffer", logs: logs, sub: sub}, nil
}

// WatchUpdateOffer is a free log subscription operation binding the contract event 0x5f63f69e4dbc29f15e8c72167eaa48adf30c77a889c1bea4dfe05403edaaee8f.
//
// Solidity: event UpdateOffer(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint128 newPrice, uint256 expires)
func (_Paintswap *PaintswapFilterer) WatchUpdateOffer(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateOffer) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateOffer)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateOffer", log); err != nil {
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

// ParseUpdateOffer is a log parse operation binding the contract event 0x5f63f69e4dbc29f15e8c72167eaa48adf30c77a889c1bea4dfe05403edaaee8f.
//
// Solidity: event UpdateOffer(uint256 offerId, address nft, uint256 tokenId, uint128 quantity, uint128 newPrice, uint256 expires)
func (_Paintswap *PaintswapFilterer) ParseUpdateOffer(log types.Log) (*PaintswapUpdateOffer, error) {
	event := new(PaintswapUpdateOffer)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the Paintswap contract.
type PaintswapUpdatePriceIterator struct {
	Event *PaintswapUpdatePrice // Event containing the contract specifics and raw log

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
func (it *PaintswapUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdatePrice)
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
		it.Event = new(PaintswapUpdatePrice)
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
func (it *PaintswapUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdatePrice represents a UpdatePrice event raised by the Paintswap contract.
type PaintswapUpdatePrice struct {
	MarketplaceId *big.Int
	Price         *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*PaintswapUpdatePriceIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdatePriceIterator{contract: _Paintswap.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *PaintswapUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdatePrice)
				if err := _Paintswap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) ParseUpdatePrice(log types.Log) (*PaintswapUpdatePrice, error) {
	event := new(PaintswapUpdatePrice)
	if err := _Paintswap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateQuantityIterator is returned from FilterUpdateQuantity and is used to iterate over the raw logs and unpacked data for UpdateQuantity events raised by the Paintswap contract.
type PaintswapUpdateQuantityIterator struct {
	Event *PaintswapUpdateQuantity // Event containing the contract specifics and raw log

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
func (it *PaintswapUpdateQuantityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateQuantity)
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
		it.Event = new(PaintswapUpdateQuantity)
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
func (it *PaintswapUpdateQuantityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateQuantityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateQuantity represents a UpdateQuantity event raised by the Paintswap contract.
type PaintswapUpdateQuantity struct {
	MarketplaceId      *big.Int
	Nfts               []common.Address
	TokenIds           []*big.Int
	NewAmount          *big.Int
	NewAmountRemaining *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuantity is a free log retrieval operation binding the contract event 0x767e006df6f79f6f4c33316e352d91e3cc914891a321b1f03a98ea6335893e75.
//
// Solidity: event UpdateQuantity(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256 newAmount, uint256 newAmountRemaining)
func (_Paintswap *PaintswapFilterer) FilterUpdateQuantity(opts *bind.FilterOpts) (*PaintswapUpdateQuantityIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateQuantity")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateQuantityIterator{contract: _Paintswap.contract, event: "UpdateQuantity", logs: logs, sub: sub}, nil
}

// WatchUpdateQuantity is a free log subscription operation binding the contract event 0x767e006df6f79f6f4c33316e352d91e3cc914891a321b1f03a98ea6335893e75.
//
// Solidity: event UpdateQuantity(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256 newAmount, uint256 newAmountRemaining)
func (_Paintswap *PaintswapFilterer) WatchUpdateQuantity(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateQuantity) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateQuantity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateQuantity)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateQuantity", log); err != nil {
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

// ParseUpdateQuantity is a log parse operation binding the contract event 0x767e006df6f79f6f4c33316e352d91e3cc914891a321b1f03a98ea6335893e75.
//
// Solidity: event UpdateQuantity(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256 newAmount, uint256 newAmountRemaining)
func (_Paintswap *PaintswapFilterer) ParseUpdateQuantity(log types.Log) (*PaintswapUpdateQuantity, error) {
	event := new(PaintswapUpdateQuantity)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateQuantity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateStartTimeIterator is returned from FilterUpdateStartTime and is used to iterate over the raw logs and unpacked data for UpdateStartTime events raised by the Paintswap contract.
type PaintswapUpdateStartTimeIterator struct {
	Event *PaintswapUpdateStartTime // Event containing the contract specifics and raw log

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
func (it *PaintswapUpdateStartTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateStartTime)
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
		it.Event = new(PaintswapUpdateStartTime)
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
func (it *PaintswapUpdateStartTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateStartTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateStartTime represents a UpdateStartTime event raised by the Paintswap contract.
type PaintswapUpdateStartTime struct {
	MarketplaceId *big.Int
	StartTime     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateStartTime is a free log retrieval operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) FilterUpdateStartTime(opts *bind.FilterOpts) (*PaintswapUpdateStartTimeIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateStartTime")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateStartTimeIterator{contract: _Paintswap.contract, event: "UpdateStartTime", logs: logs, sub: sub}, nil
}

// WatchUpdateStartTime is a free log subscription operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) WatchUpdateStartTime(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateStartTime) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateStartTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateStartTime)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateStartTime", log); err != nil {
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

// ParseUpdateStartTime is a log parse operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) ParseUpdateStartTime(log types.Log) (*PaintswapUpdateStartTime, error) {
	event := new(PaintswapUpdateStartTime)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateStartTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

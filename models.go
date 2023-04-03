///////////////////////////////////////////////////////////////
// File containing the model structs for the POD definition,
// and blockchain inputs and outputs
///////////////////////////////////////////////////////////////

package fractionalise

import (
	"github.com/shopspring/decimal"
)

/*---------------------------------------------------------------------------
SMART CONTRACT MODELS FOR FRACTINALISE NFT
-----------------------------------------------------------------------------*/

// Define our struct to store the fractionalised Media on Blockchain //
type Fractionalise struct {
	TokenSymbol  string          `json:"TokenSymbol"`
	OwnerAddress string          `json:"OwnerAddress"`
	Fraction     decimal.Decimal `json:"Fraction"`
	BuyBackPrice decimal.Decimal `json:"BuyBackPrice"`
	InitialPrice decimal.Decimal `json:"InitialPrice"`
	FundingToken string          `json:"FundingToken"`
	InterestRate decimal.Decimal `json:"InterestRate"`
}

type FractionaliseStreaming struct {
	TokenSymbol string `json:"TokenSymbol"`
	Address     string `json:"Address"`
	StreamingId string `json:"StreamingId"`
}

// Model to buy POD tokens from a buying market order //
type BuyingOffer struct {
	OrderId     string          `json:"OrderId"`
	TokenSymbol string          `json:"TokenSymbol"`
	PodAddress  string          `json:"PodAddress"`
	BAddress    string          `json:"BAddress"`
	Amount      decimal.Decimal `json:"Amount"`
	Token       string          `json:"Token"`
	Price       decimal.Decimal `json:"Price"`
}

// Model to buy POD tokens from a buying market order //
type SellingOffer struct {
	OrderId     string          `json:"OrderId"`
	TokenSymbol string          `json:"TokenSymbol"`
	PodAddress  string          `json:"PodAddress"`
	SAddress    string          `json:"SAddress"`
	Amount      decimal.Decimal `json:"Amount"`
	Token       string          `json:"Token"`
	Price       decimal.Decimal `json:"Price"`
}

type PodInstantiateRequest struct {
	OfferList []SellingOffer `json:"OfferList"`
	TxnId     string         `json:"TxnId"`
}

type DeleteOrderRequest struct {
	RequesterAddress string `json:"RequesterAddress"`
	TokenSymbol      string `json:"TokenSymbol"`
	OrderId          string `json:"OrderId"`
}

type BuyFractionRequest struct {
	TokenSymbol  string          `json:"TokenSymbol"`
	SAddress     string          `json:"SAddress"`
	OrderId      string          `json:"OrderId"`
	Amount       decimal.Decimal `json:"Amount"`
	BuyerAddress string          `json:"BuyerAddress"`
}

type SellFractionRequest struct {
	TokenSymbol   string          `json:"TokenSymbol"`
	BAddress      string          `json:"BAddress"`
	OrderId       string          `json:"OrderId"`
	Amount        decimal.Decimal `json:"Amount"`
	SellerAddress string          `json:"SellerAddress"`
}

type BuyBackRequest struct {
	TokenSymbol string `json:"TokenSymbol"`
	Address     string `json:"Address"`
}

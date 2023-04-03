///////////////////////////////////////////////////////////////////
// File containing the constants for the Cache Coin Smart Contract
///////////////////////////////////////////////////////////////////

package fractionalise

import "github.com/shopspring/decimal"

/* -------------------------------------------------------
-------------------------------------------------------- */

const (
	BuyingOffersIndex           = "FRACTIONALISE_BUYING"
	SellingOffersIndex          = "FRACTIONALISE_SELLING"
	FractionaliseIndex          = "FRACTIONALISE"
	FractionaliseStreamingIndex = "FRACTIONALISE_STREAMING"
)

var (
	YearSeconds = decimal.NewFromInt(365 * 24 * 60 * 60)
)

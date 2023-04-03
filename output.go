package fractionalise

import (
	"github.com/Get-Cache/Privi/contracts/coinbalance"
)

type Output struct {
	UpdateFractionalise map[string]Fractionalise         `json:"UpdateFractionalise"`
	UpdateBuyingOffers  map[string]BuyingOffer           `json:"UpdateBuyingOffers"`
	UpdateSellingOffers map[string]SellingOffer          `json:"UpdateSellingOffers"`
	UpdateStreamings    map[string]coinbalance.Streaming `json:"UpdateStreamings"`
	UpdateTokens        map[string]coinbalance.Token     `json:"UpdateTokens"`
	Transactions        []coinbalance.Transfer           `json:"Transactions"`
}

func (o *Output) WithFractionalises(args ...Fractionalise) *Output {
	if o.UpdateFractionalise == nil {
		o.UpdateFractionalise = make(map[string]Fractionalise)
	}
	for _, e := range args {
		o.UpdateFractionalise[e.TokenSymbol] = e
	}
	return o
}

func (o *Output) WithBuyingOffers(args ...BuyingOffer) *Output {
	if o.UpdateBuyingOffers == nil {
		o.UpdateBuyingOffers = make(map[string]BuyingOffer)
	}
	for _, e := range args {
		o.UpdateBuyingOffers[e.OrderId] = e
	}
	return o
}

func (o *Output) WithSellingOffers(args ...SellingOffer) *Output {
	if o.UpdateSellingOffers == nil {
		o.UpdateSellingOffers = make(map[string]SellingOffer)
	}
	for _, e := range args {
		o.UpdateSellingOffers[e.OrderId] = e
	}
	return o
}

func (o *Output) WithStreamings(args ...coinbalance.Streaming) *Output {
	if o.UpdateStreamings == nil {
		o.UpdateStreamings = make(map[string]coinbalance.Streaming)
	}
	for _, e := range args {
		o.UpdateStreamings[e.StreamingId] = e
	}
	return o
}

func (o *Output) WithTokens(args ...coinbalance.Token) *Output {
	if o.UpdateTokens == nil {
		o.UpdateTokens = make(map[string]coinbalance.Token)
	}
	for _, e := range args {
		o.UpdateTokens[e.Symbol] = e
	}
	return o
}

func (o *Output) WithTransactions(args ...coinbalance.Transfer) *Output {
	o.Transactions = append(o.Transactions, args...)
	return o
}

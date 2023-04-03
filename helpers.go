/*--------------------------------------------------------------------------
----------------------------------------------------------------------------
   HELPER FUNCTIONS CALLED SEVERAL TIMES ON MAIN SMART CONTRACT FUNCIOTNS
----------------------------------------------------------------------------
-------------------------------------------------------------------------- */

package fractionalise

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Get-Cache/Privi/contracts/coinbalance"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/shopspring/decimal"
)

/* -------------------------------------------------------------------------------------------------
registerAddress: this function register a balance
------------------------------------------------------------------------------------------------- */

func registerAddress(stub shim.ChaincodeStubInterface,
	address, addressType string) /*error*/ {

	// // Invoke CoinBalance Chaincode //
	// invoke_call := []string{"registerAddress"}
	// invoke_call = append(invoke_call, address, addressType)
	// multiChainCodeArgs := ToChaincodeArgs(invoke_call)
	// response := stub.InvokeChaincode(COIN_BALANCE_CHAINCODE, multiChainCodeArgs,
	// 	CHANNEL_NAME)
	// if response.Status != shim.OK {
	// 	return errors.New(response.Message)
	// }
	// return nil
}

/* -------------------------------------------------------------------------------------------------
multiTransfer: this function computes all the transfers taking place on the smart contract
------------------------------------------------------------------------------------------------- */

func multiTransfer(stub shim.ChaincodeStubInterface,
	multitransfers []string) /*([]Transfer, error)*/ {

	// // Invoke CoinBalance Chaincode to perform multitransfer //
	// multiChainCodeArgs := ToChaincodeArgs(multitransfers)
	// response := stub.InvokeChaincode(COIN_BALANCE_CHAINCODE, multiChainCodeArgs,
	// 	CHANNEL_NAME)
	// if response.Status != shim.OK {
	// 	return []Transfer{}, errors.New(response.Message)
	// }
	// output := Output{}
	// json.Unmarshal(response.Payload, &output)
	// return output.Transactions, nil
}

/* -------------------------------------------------------------------------------------------------
registerSellingOffer: this function registers a selling offer on blockchain
------------------------------------------------------------------------------------------------- */

func registerSellingOffer(stub shim.ChaincodeStubInterface, sellingOffer SellingOffer) error {
	// Store selling offer on Blockchain //
	err := sellingOffer.SaveState(stub)
	if err != nil {
		return errors.New("ERROR: REGISTRATING SELLING OFFER " + sellingOffer.OrderId +
			" ON BLOCKCHAIN. " + err.Error())
	}
	return nil
}

/* -------------------------------------------------------------------------------------------------
registerBuyingOffer: this function registers a buying offer on blockchain
------------------------------------------------------------------------------------------------- */

func registerBuyingOffer(stub shim.ChaincodeStubInterface, buyingOffer *BuyingOffer) error {
	// Store selling offer on Blockchain //
	err := buyingOffer.SaveState(stub)
	if err != nil {
		return errors.New("ERROR: REGISTRATING BUYING OFFER " + buyingOffer.OrderId +
			" ON BLOCKCHAIN. " + err.Error())
	}
	return nil
}

/* -------------------------------------------------------------------------------------------------
updateSellingOffer: this function updates a selling offer ( or deletes it if empty )
------------------------------------------------------------------------------------------------- */

func updateSellingOffer(stub shim.ChaincodeStubInterface, offer SellingOffer) (SellingOffer, error) {
	if offer.Amount.GreaterThan(decimal.Zero) {
		err := offer.SaveState(stub)
		if err != nil {
			return SellingOffer{}, errors.New(err.Error())
		}
		return offer, nil
	}
	err := offer.DeleteState(stub)
	if err != nil {
		return SellingOffer{}, errors.New(err.Error())
	}
	return offer, nil
}

/* -------------------------------------------------------------------------------------------------
updateBuyingOffer: this function updates a buying offer ( or deletes it if empty )
------------------------------------------------------------------------------------------------- */

func updateBuyingOffer(stub shim.ChaincodeStubInterface, offer BuyingOffer) (BuyingOffer, error) {
	if offer.Amount.GreaterThan(decimal.Zero) {
		err := offer.SaveState(stub)
		if err != nil {
			return BuyingOffer{}, errors.New(err.Error())
		}
		return offer, nil
	}
	err := offer.DeleteState(stub)
	if err != nil {
		return BuyingOffer{}, errors.New(err.Error())
	}
	return offer, nil
}

/* -------------------------------------------------------------------------------------------------
getSellingOffers: this function returns selling offers of a fractionalise with given fractionalise address
------------------------------------------------------------------------------------------------- */

func GetSellingOffers(stub shim.ChaincodeStubInterface, tokensymbol string) ([]SellingOffer, error) {
	// Retrieve offers from Blockchain //
	it, err := stub.GetStateByPartialCompositeKey(SellingOffersIndex, []string{tokensymbol})
	if err != nil {
		return nil, errors.New("ERROR: unable to get an iterator over the offers")
	}

	defer it.Close()
	var offers []SellingOffer
	for it.HasNext() {
		response, err1 := it.Next()
		if err1 != nil {
			message := fmt.Sprintf("unable to get the next element: %s", err1.Error())
			return nil, errors.New(message)
		}
		var offer SellingOffer
		if err = json.Unmarshal(response.Value, &offer); err != nil {
			message := fmt.Sprintf("ERROR: unable to parse the response: %s", err.Error())
			return nil, errors.New(message)
		}
		offers = append(offers, offer)
	}
	return offers, nil
}

/* -------------------------------------------------------------------------------------------------
getBuyingOffers: this function returns buying offers of a fractionalise with given fractionalise address
------------------------------------------------------------------------------------------------- */

func GetBuyingOffers(stub shim.ChaincodeStubInterface, tokensymbol string) ([]BuyingOffer, error) {
	// Retrieve offers from Blockchain //
	it, err := stub.GetStateByPartialCompositeKey(BuyingOffersIndex, []string{tokensymbol})
	if err != nil {
		return nil, errors.New("ERROR: unable to get an iterator over the offers")
	}

	defer it.Close()
	var offers []BuyingOffer
	for it.HasNext() {
		response, err1 := it.Next()
		if err1 != nil {
			message := fmt.Sprintf("unable to get the next element: %s", err1.Error())
			return nil, errors.New(message)
		}
		var offer BuyingOffer
		if err = json.Unmarshal(response.Value, &offer); err != nil {
			message := fmt.Sprintf("ERROR: unable to parse the response: %s", err.Error())
			return nil, errors.New(message)
		}
		offers = append(offers, offer)
	}
	return offers, nil
}

/* -------------------------------------------------------------------------------------------------
getUserOffers: this function retrieves all the Offers of a user on NFT Pods
------------------------------------------------------------------------------------------------- */

func GetUserOffers(stub shim.ChaincodeStubInterface, address string) (*Output, error) {

	//find selling offers
	queryString := fmt.Sprintf(`{"selector":{"SAddress":"%s"}}`, address)

	itSellers, err := stub.GetQueryResult(queryString)
	if err != nil {
		// return shim.Error("ERROR: unable to get an iterator over the selling offers")
		return nil, err
	}
	defer itSellers.Close()
	var sellingoffers []SellingOffer
	for itSellers.HasNext() {
		response, err := itSellers.Next()
		if err != nil {
			// message := fmt.Sprintf("unable to get the next element: %s", err.Error())
			// return shim.Error(message)
			return nil, err
		}
		var offer SellingOffer
		if err = json.Unmarshal(response.Value, &offer); err != nil {
			// message := fmt.Sprintf("ERROR: unable to parse the response: %s", err.Error())
			// return shim.Error(message)
			return nil, err
		}
		sellingoffers = append(sellingoffers, offer)
	}

	// find buying offers
	queryString = fmt.Sprintf(`{"selector":{"BAddress":"%s"}}`, address)
	itBuyers, err := stub.GetQueryResult(queryString)
	if err != nil {
		// return shim.Error("ERROR: unable to get an iterator over the buying offers")
		return nil, err
	}
	defer itBuyers.Close()
	var buyingoffers []BuyingOffer
	for itBuyers.HasNext() {
		response, err := itBuyers.Next()
		if err != nil {
			// message := fmt.Sprintf("unable to get the next element: %s", err.Error())
			// return shim.Error(message)
			return nil, err
		}
		var offer BuyingOffer
		if err = json.Unmarshal(response.Value, &offer); err != nil {
			// message := fmt.Sprintf("ERROR: unable to parse the response: %s", err.Error())
			// return shim.Error(message)
			return nil, err
		}
		buyingoffers = append(buyingoffers, offer)
	}

	// Output of the result //
	return new(Output).
		WithSellingOffers(sellingoffers...).
		WithBuyingOffers(buyingoffers...), nil
}

/* -------------------------------------------------------------------------------------------------
getAttachedAddress: this function returns the address attached to a given userID.
------------------------------------------------------------------------------------------------- */

func getAttachedAddress(stub shim.ChaincodeStubInterface, address string) /*(string, error)*/ {

	// /// Invoke DataProtocol Chaincode //
	// invoke_call := []string{"getUser"}
	// invoke_call = append(invoke_call, address)
	// multiChainCodeArgs := ToChaincodeArgs(invoke_call)
	// response := stub.InvokeChaincode(DATA_PROTOCOL_CHAINCODE, multiChainCodeArgs,
	// 	CHANNEL_NAME)
	// if response.Status != shim.OK {
	// 	return "", errors.New(response.Message)
	// }

	// if response.Status != shim.OK {
	// 	return "", errors.New("ERROR: GETTING THE ADDRESS OF " + address +
	// 		" ON BLOCKCHAIN. " + response.Message)
	// }
	// actor := Actor{}
	// json.Unmarshal(response.Payload, &actor)
	// return actor.PublicAddress, nil
}

/* -------------------------------------------------------------------------------------------------
getBalance: this function gets the balance from an Address
------------------------------------------------------------------------------------------------- */

func getBalance(stub shim.ChaincodeStubInterface, userAddress string,
	token string) /*(decimal.Decimal, error)*/ {

	// // Invoke CoinBalance Chaincode //
	// invokeCall := []string{"balanceOf", userAddress, token}
	// multiChainCodeArgs := ToChaincodeArgs(invokeCall)
	// response := stub.InvokeChaincode(COIN_BALANCE_CHAINCODE, multiChainCodeArgs,
	// 	CHANNEL_NAME)
	// if response.Status != shim.OK {
	// 	return decimal.Zero, errors.New(response.Message)
	// }

	// // Get balance //
	// balance := Balance{}
	// err := json.Unmarshal(response.Payload, &balance)
	// if err != nil {
	// 	return decimal.Zero, errors.New("ERROR: GETTING BALANCE INFO. " + err.Error())
	// }
	// return balance.Amount, nil
}

/* -------------------------------------------------------------------------------------------------
ACCOUNTING STREAMINGS
------------------------------------------------------------------------------------------------- */

// interestFromAmount returns the annual interest for a fraction given the current portion
func interestFromAmount(fractionalise *Fractionalise, amount decimal.Decimal) decimal.Decimal {
	return amount.Div(fractionalise.Fraction).Mul(fractionalise.InterestRate)
}

// amountFromInterest returns the holders amount of a fraction given his annual interest
func amountFromInterest(fractionalise *Fractionalise, interest decimal.Decimal) decimal.Decimal {
	return interest.Div(fractionalise.InterestRate).Mul(fractionalise.Fraction)
}

// setStreaming updates the fractionalise interest streaming for the given address
func setStreaming(stub shim.ChaincodeStubInterface, address, tokensymbol, streamingtoken string, amount decimal.Decimal) ([]coinbalance.TransferRequest, []coinbalance.Streaming, error) {
	// output
	var (
		transfers  []coinbalance.TransferRequest
		streamings []coinbalance.Streaming
	)

	// get fractionalise info
	fractionalise, err := GetFractionaliseByTokenSymbol(stub, tokensymbol)
	if err != nil {
		return nil, nil, err
	}

	// get user possible streamings
	streams, err := GetFractionalisedStreamings(stub, tokensymbol, address)
	if err != nil {
		return nil, nil, err
	}

	// user streams
	var (
		stream    FractionaliseStreaming
		streaming coinbalance.Streaming
	)

	// create new streaming
	if len(streams) == 0 {
		streaming = coinbalance.Streaming{
			SenderAddress:   address,
			Type:            "Fractionalise-Streaming-" + tokensymbol,
			ReceiverAddress: fractionalise.OwnerAddress,
			StreamingToken:  streamingtoken,
			Frequency:       decimal.NewFromInt(1),
			EndingDate:      int64(1e18),
		}

		stream = FractionaliseStreaming{
			TokenSymbol: tokensymbol,
			Address:     address,
		}
	}

	// get existing streaming and cancel it
	if len(streams) == 1 {
		stream = streams[0]

		// get existing streaming
		if streaming, err = coinbalance.GetStreaming(stub, stream.StreamingId); err != nil {
			return nil, nil, err
		}

		// stop existing streaming
		if transfers, err = coinbalance.StopStreamings(stub, stream.StreamingId); err != nil {
			return nil, nil, err
		}

		streamings = append(streamings, streaming)
	}

	// get streaming amount
	streaming.AmountPerPeriod = interestFromAmount(fractionalise, amount).Div(YearSeconds)

	// get streaming date
	var date int64
	if time, err := stub.GetTxTimestamp(); err != nil {
		return nil, nil, err
	} else {
		date = time.GetSeconds()
	}

	// create new streaming
	streaming.StartingDate = date
	if id, err := coinbalance.CreateStreaming(stub, &streaming); err != nil {
		return nil, nil, err
	} else {
		streaming.StreamingId = id
		stream.StreamingId = id

		streamings = append(streamings, streaming)
	}

	// udpate user stream state
	if err := stream.SaveState(stub); err != nil {
		return nil, nil, err
	}

	return transfers, streamings, nil
}

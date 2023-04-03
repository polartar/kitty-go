package fractionalise

import (
	"errors"
	"github.com/Get-Cache/Privi/contracts/coinbalance"
	"github.com/Get-Cache/Privi/signature"
	"github.com/Get-Cache/Privi/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/shopspring/decimal"
)

type SmartContract struct{}

func (*SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()
	if len(args) > 0 && args[0] == "UPGRADE" {
		return shim.Success(nil)
	}
	return shim.Success(nil)
}

func (*SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve function and arguments //
	function, args := stub.GetFunctionAndParameters()

	// Call the proper function //
	switch function {

	// public functions
	case `getFractionaliseInfo`:
		// check args lenght
		if err := utils.ValidateArgsLen(args, 1); err != nil {
			return shim.Error(err.Error())
		}
		// decode input from args
		var (
			tokensymbol = args[0]
		)
		// invoke function
		if payload, err := GetFractionaliseByTokenSymbol(stub, tokensymbol); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `getSellingOffers`:
		// check args lenght
		if err := utils.ValidateArgsLen(args, 1); err != nil {
			return shim.Error(err.Error())
		}
		// decode input from args
		var (
			tokensymbol = args[0]
		)
		// invoke function
		if payload, err := GetSellingOffers(stub, tokensymbol); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case "getBuyingOffers":
		// check args lenght
		if err := utils.ValidateArgsLen(args, 1); err != nil {
			return shim.Error(err.Error())
		}
		// decode input from args
		var (
			tokensymbol = args[0]
		)
		// invoke function
		if payload, err := GetBuyingOffers(stub, tokensymbol); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `getUserOffers`:
		// check args lenght
		if err := utils.ValidateArgsLen(args, 1); err != nil {
			return shim.Error(err.Error())
		}
		// decode input from args
		var (
			address = args[0]
		)
		// invoke function
		if payload, err := GetUserOffers(stub, address); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	// state functions
	case `fractionalise`:
		// decode input from args
		var input Fractionalise
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.OwnerAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE OWNER.")
		}
		// invoke function
		if payload, err := FractionaliseToken(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `newBuyOrder`:
		// Retrieve the input information and check signature //
		var input BuyingOffer
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.BAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE BUYER.")
		}
		// invoke function
		if payload, err := NewBuyOrder(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `deleteBuyOrder`:
		// Retrieve the input information and check signature //
		var input DeleteOrderRequest
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.RequesterAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE REQUESTER.")
		}
		// invoke function
		if payload, err := DeleteBuyOrder(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `newSellOrder`:
		// Retrieve the input information and check signature //
		var input SellingOffer
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.SAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE SELLER.")
		}
		// invoke function
		if payload, err := NewSellOrder(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `deleteSellOrder`:
		// Retrieve the input information and check signature //
		var input DeleteOrderRequest
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.RequesterAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE REQUESTER.")
		}
		// invoke function
		if payload, err := DeleteSellOrder(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `buyFraction`:
		// decode input from args
		var input BuyFractionRequest
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.BuyerAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE BUYER.")
		}
		// invoke function
		if payload, err := BuyFraction(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case "sellFraction":

		// decode input from args
		var input SellFractionRequest
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.SellerAddress {
			return shim.Error("ERROR: SIGNER HAS TO BE THE SELLER.")
		}
		// invoke function
		if payload, err := SellFraction(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}

	case `buyFractionalisedBack`:
		// decode input from args
		var input BuyBackRequest
		address, err := signature.SecureUnmarshal(stub, &input)
		if err != nil {
			return shim.Error(err.Error())
		}
		// Verify actor address with address from signature //
		if address != input.Address {
			return shim.Error("ERROR: SIGNER HAS TO BE THE SELLER.")
		}
		// invoke function
		if payload, err := BuyFractionalisedBack(stub, &input); err != nil {
			return shim.Error(err.Error())
		} else {
			return utils.MarshalSuccess(payload)
		}
	}

	return utils.NotFound(function)
}

func FractionaliseToken(stub shim.ChaincodeStubInterface, input *Fractionalise) (*Output, error) {
	// check user token balance
	if balance, err := coinbalance.BalanceOf(stub, input.OwnerAddress, input.TokenSymbol); err != nil {
		return nil, err
	} else if balance.Amount.LessThan(decimal.NewFromInt(1)) {
		return nil, errors.New("user does not hold the token")
	}

	// conver to nft. TODO: in common state.

	// Register on order book initial offer
	offer := SellingOffer{
		OrderId:     stub.GetTxID(),
		TokenSymbol: input.TokenSymbol,
		SAddress:    input.OwnerAddress,
		Amount:      input.Fraction,
		Token:       input.FundingToken,
		Price:       input.InitialPrice,
	}

	// Create pool for trading
	if addr, err := utils.GenerateAddress(stub, offer); err != nil {
		return nil, err
	} else {
		offer.PodAddress = addr
	}

	if err := coinbalance.RegisterAddress(stub, offer.PodAddress, coinbalance.LiquidityPoolAddressType); err != nil {
		return nil, err
	}

	// save fractionalised state
	if err := input.SaveState(stub); err != nil {
		return nil, err
	}

	// Transfer to AMM the fraction
	transfer := coinbalance.TransferRequest{
		Type:           "fractionalise_media",
		Token:          input.TokenSymbol,
		From:           input.OwnerAddress,
		To:             offer.PodAddress,
		AvoidCheckTo:   true,
		AvoidCheckFrom: true,
		Amount:         input.Fraction,
	}

	r, err := coinbalance.Multitransfer(stub, transfer)
	if err != nil {
		return nil, err
	}

	return new(Output).
		WithTransactions(r.Transactions...).
		WithFractionalises(*input).
		WithSellingOffers(offer), nil
}

func NewBuyOrder(stub shim.ChaincodeStubInterface, offer *BuyingOffer) (*Output, error) {

	offer.OrderId = stub.GetTxID()

	// Transfer Funds of Offer to POD //
	transfer := coinbalance.TransferRequest{
		Type:   "Fractionalise_Buy_Offer",
		Amount: offer.Amount.Mul(offer.Price),
		Token:  offer.Token,
		From:   offer.BAddress,
		To:     offer.PodAddress,
	}

	if err := registerBuyingOffer(stub, offer); err != nil {
		return nil, err
	}

	// Update balances with the transfers by invoking CoinBalance Chaincode //
	r, err := coinbalance.Multitransfer(stub, transfer)
	if err != nil {
		return nil, err
	}

	// Output of the result //
	return new(Output).
		WithTransactions(r.Transactions...).
		WithBuyingOffers(*offer), nil
}

func NewSellOrder(stub shim.ChaincodeStubInterface, offer *SellingOffer) (*Output, error) {

	// Get fractionalise info //
	fractionalise, err := GetFractionaliseByTokenSymbol(stub, offer.TokenSymbol)
	if err != nil {
		return nil, err
	}

	// Transfer Pod tokens to POD //
	transfer := coinbalance.TransferRequest{
		Type:   "Fractionalise_Sell_Offer",
		Amount: offer.Amount,
		Token:  fractionalise.TokenSymbol,
		From:   offer.SAddress,
		To:     offer.PodAddress,
	}

	offer.OrderId = stub.GetTxID()
	if err = registerSellingOffer(stub, *offer); err != nil {
		return nil, err
	}

	// Update balances with the transfers by invoking CoinBalance Chaincode //
	r, err := coinbalance.Multitransfer(stub, transfer)
	if err != nil {
		return nil, err
	}

	// Output of the result //
	return new(Output).
		WithTransactions(r.Transactions...).
		WithSellingOffers(*offer), nil
}

func DeleteBuyOrder(stub shim.ChaincodeStubInterface, request *DeleteOrderRequest) (*Output, error) {

	// Get Buying Offer By Id //
	offer := BuyingOffer{
		TokenSymbol: request.TokenSymbol,
		BAddress:    request.RequesterAddress,
		OrderId:     request.OrderId,
	}

	// check if order exists
	if exists, err := offer.LoadState(stub); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("such offer does not exists")
	}

	// money recover transactions
	var transactions []coinbalance.TransferRequest

	// recover funding tokens
	balance, err := coinbalance.BalanceOf(stub, offer.PodAddress, offer.Token)
	if err != nil {
		return nil, err
	}

	if balance.Amount.GreaterThan(decimal.Zero) {
		transactions = append(transactions, coinbalance.TransferRequest{
			Type:   "NFT_Pod_Buy_Order_Delete",
			Amount: balance.Amount,
			Token:  balance.Token,
			From:   offer.PodAddress,
			To:     offer.BAddress,
		})
	}

	// recover media tokens in case of any
	balance, err = coinbalance.BalanceOf(stub, offer.PodAddress, offer.TokenSymbol)
	if err != nil {
		return nil, err
	}

	if balance.Amount.GreaterThan(decimal.Zero) {
		transactions = append(transactions, coinbalance.TransferRequest{
			Type:   "NFT_Pod_Buy_Order_Delete",
			Amount: balance.Amount,
			Token:  balance.Token,
			From:   offer.PodAddress,
			To:     offer.BAddress,
		})
	}

	// update balances with the transfers
	r, err := coinbalance.Multitransfer(stub, transactions...)
	if err != nil {
		return nil, err
	}

	// delete offer state
	if err := offer.DeleteState(stub); err != nil {
		return nil, err
	}

	return new(Output).
		WithTransactions(r.Transactions...), nil
}

func DeleteSellOrder(stub shim.ChaincodeStubInterface, request *DeleteOrderRequest) (*Output, error) {
	// Retrieve selling Offer with Id //
	offer := SellingOffer{
		TokenSymbol: request.TokenSymbol,
		SAddress:    request.RequesterAddress,
		OrderId:     request.OrderId,
	}

	// check if offer exists
	if exists, err := offer.LoadState(stub); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("such offer not exists")
	}

	// money recover transactions
	var transactions []coinbalance.TransferRequest

	// recover funding tokens
	balance, err := coinbalance.BalanceOf(stub, offer.PodAddress, offer.Token)
	if err != nil {
		return nil, err
	}

	if balance.Amount.GreaterThan(decimal.Zero) {
		transactions = append(transactions, coinbalance.TransferRequest{
			Type:   "NFT_Pod_Sell_Offer_Delete",
			Amount: balance.Amount,
			Token:  balance.Token,
			From:   offer.PodAddress,
			To:     offer.SAddress,
		})
	}

	// recover media tokens in case of any
	balance, err = coinbalance.BalanceOf(stub, offer.PodAddress, offer.TokenSymbol)
	if err != nil {
		return nil, err
	}

	if balance.Amount.GreaterThan(decimal.Zero) {
		transactions = append(transactions, coinbalance.TransferRequest{
			Type:   "NFT_Pod_Sell_Offer_Delete",
			Amount: balance.Amount,
			Token:  balance.Token,
			From:   offer.PodAddress,
			To:     offer.SAddress,
		})
	}

	// update balances with the transfers
	r, err := coinbalance.Multitransfer(stub, transactions...)
	if err != nil {
		return nil, err
	}

	// delete offer state
	if err := offer.DeleteState(stub); err != nil {
		return nil, err
	}

	// output of the result //
	return new(Output).
		WithTransactions(r.Transactions...), nil
}

func BuyFraction(stub shim.ChaincodeStubInterface, input *BuyFractionRequest) (*Output, error) {

	// Retrieve selling offer //
	offer := SellingOffer{
		OrderId:     input.OrderId,
		TokenSymbol: input.TokenSymbol,
		SAddress:    input.SAddress,
	}

	// check if offer exists
	if exists, err := offer.LoadState(stub); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("such offer not exists")
	}

	// calculate buy price
	amountprice := input.Amount.Mul(offer.Price)

	// check if buyer has sufficient funds
	if balance, err := coinbalance.BalanceOf(stub, input.BuyerAddress, offer.Token); err != nil {
		return nil, err
	} else if balance.Amount.LessThan(amountprice) {
		return nil, errors.New("not enough funds to perform this opperation")
	}

	// check if offer has sufficient funds
	if amount, err := saveSubstraction(offer.Amount, input.Amount); err != nil {
		return nil, errors.New("offer has not enough funds to satisfy request")
	} else {
		offer.Amount = amount
	}

	// load fractionalise
	fractionalise, err := GetFractionaliseByTokenSymbol(stub, offer.TokenSymbol)
	if err != nil {
		return nil, err
	}

	// update streamings
	var (
		streamings   []coinbalance.Streaming
		transactions []coinbalance.TransferRequest
	)

	if input.BuyerAddress != fractionalise.OwnerAddress {
		balance, err := coinbalance.BalanceOf(stub, input.BuyerAddress, input.TokenSymbol)
		if err != nil {
			return nil, err
		}

		if transfers, streams, err := setStreaming(stub, input.BuyerAddress, input.TokenSymbol, offer.Token, balance.Amount.Add(input.Amount)); err != nil {
			return nil, err
		} else {
			transactions = append(transactions, transfers...)
			streamings = append(streamings, streams...)
		}
	}

	if offer.SAddress != fractionalise.OwnerAddress {
		balance, err := coinbalance.BalanceOf(stub, offer.SAddress, input.TokenSymbol)
		if err != nil {
			return nil, err
		}

		if transfers, streams, err := setStreaming(stub, offer.SAddress, input.TokenSymbol, offer.Token, balance.Amount.Sub(input.Amount)); err != nil {
			return nil, err
		} else {
			transactions = append(transactions, transfers...)
			streamings = append(streamings, streams...)
		}
	}

	// make transfers
	buytransfer := coinbalance.TransferRequest{
		Type:   "Fractionalise_Buying",
		Token:  offer.Token,
		Amount: amountprice,
		From:   input.BuyerAddress,
		To:     input.SAddress,
	}

	selltransfer := coinbalance.TransferRequest{
		Type:   "Fractionalise_Buying",
		Token:  fractionalise.TokenSymbol,
		Amount: input.Amount,
		From:   offer.PodAddress,
		To:     input.BuyerAddress,
	}

	transactions = append(transactions, buytransfer, selltransfer)

	r, err := coinbalance.Multitransfer(stub, transactions...)
	if err != nil {
		return nil, err
	}

	// update offer if not empty, delete if empty
	offer, err = updateSellingOffer(stub, offer)
	if err != nil {
		return nil, err
	}

	// output of the result
	return new(Output).
		WithTransactions(r.Transactions...).
		WithStreamings(streamings...).
		WithSellingOffers(offer), nil
}

func SellFraction(stub shim.ChaincodeStubInterface, input *SellFractionRequest) (*Output, error) {

	// Retrieve Buying Offer from Id //
	offer := BuyingOffer{
		OrderId:     input.OrderId,
		TokenSymbol: input.TokenSymbol,
		BAddress:    input.BAddress,
	}

	// check if offer exists
	if exists, err := offer.LoadState(stub); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("such offer not exists")
	}

	// check if seller has sufficient funds
	if balance, err := coinbalance.BalanceOf(stub, input.SellerAddress, offer.TokenSymbol); err != nil {
		return nil, err
	} else if balance.Amount.LessThan(input.Amount) {
		return nil, errors.New("not enough funds to perform this opperation")
	}

	// Check if offer has sufficient funds //
	if amount, err := saveSubstraction(offer.Amount, input.Amount); err != nil {
		return nil, errors.New("offer has not enough funds to satisfy request")
	} else {
		offer.Amount = amount
	}

	// Load fractionalise info //
	fractionalise, err := GetFractionaliseByTokenSymbol(stub, offer.TokenSymbol)
	if err != nil {
		return nil, err
	}

	// update streamings
	var (
		streamings   []coinbalance.Streaming
		transactions []coinbalance.TransferRequest
	)

	if offer.BAddress != fractionalise.OwnerAddress {
		balance, err := coinbalance.BalanceOf(stub, offer.BAddress, input.TokenSymbol)
		if err != nil {
			return nil, err
		}

		if transfers, streams, err := setStreaming(stub, offer.BAddress, input.TokenSymbol, offer.Token, balance.Amount.Add(input.Amount)); err != nil {
			return nil, err
		} else {
			transactions = append(transactions, transfers...)
			streamings = append(streamings, streams...)
		}
	}

	if input.SellerAddress != fractionalise.OwnerAddress {
		balance, err := coinbalance.BalanceOf(stub, input.SellerAddress, input.TokenSymbol)
		if err != nil {
			return nil, err
		}

		if transfers, streams, err := setStreaming(stub, input.SellerAddress, input.TokenSymbol, offer.Token, balance.Amount.Sub(input.Amount)); err != nil {
			return nil, err
		} else {
			transactions = append(transactions, transfers...)
			streamings = append(streamings, streams...)
		}
	}

	// Transfer to the Buyer of the offer from the Seller //
	selling := coinbalance.TransferRequest{
		Type:   "Fractionalise_Selling",
		Token:  fractionalise.TokenSymbol,
		Amount: input.Amount,
		From:   input.SellerAddress,
		To:     input.BAddress,
	}

	// Transfer to the Seller the selling amount from the Pod //
	buying := coinbalance.TransferRequest{
		Type:   "Fractionalise_Selling",
		Token:  offer.Token,
		Amount: input.Amount.Mul(offer.Price),
		From:   offer.PodAddress,
		To:     input.SellerAddress,
	}

	transactions = append(transactions, selling, buying)

	// Update balances with the transfers by invoking CoinBalance Chaincode //
	r, err := coinbalance.Multitransfer(stub, transactions...)
	if err != nil {
		return nil, err
	}

	// Update offer if not empty, delete if empty //
	offer, err = updateBuyingOffer(stub, offer)
	if err != nil {
		return nil, err
	}

	// Output of the result //
	return new(Output).
		WithTransactions(r.Transactions...).
		WithStreamings(streamings...).
		WithBuyingOffers(offer), nil
}

func BuyFractionalisedBack(stub shim.ChaincodeStubInterface, input *BuyBackRequest) (*Output, error) {
	// recover fractionalise
	fractionalise, err := GetFractionaliseByTokenSymbol(stub, input.TokenSymbol)
	if err != nil {
		return nil, err
	}

	// check operation is bein called by the fractionalise owner
	if fractionalise.OwnerAddress != input.Address {
		return nil, errors.New("operation not allowed for the given address")
	}

	// check fractionalise owner has the funds to reconstruct the token
	if balance, err := coinbalance.BalanceOf(stub, fractionalise.OwnerAddress, fractionalise.FundingToken); err != nil {
		return nil, err
	} else if balance.Amount.LessThan(fractionalise.BuyBackPrice) {
		return nil, errors.New("insuficient amount to reconstruct the fractionalised token")
	}

	// get fractionalise token holders
	holders, err := coinbalance.GetTokenHolderList(stub, input.TokenSymbol)
	if err != nil {
		return nil, err
	}

	var (
		streamings   []string
		transactions []coinbalance.TransferRequest
	)

	// generate transactions for the owner and the token holder
	for _, holder := range holders {
		// check holder is not the owner
		if holder == fractionalise.OwnerAddress {
			continue
		}

		// get balance of token holder
		balance, err := coinbalance.BalanceOf(stub, holder, fractionalise.TokenSymbol)
		if err != nil {
			return nil, err
		}

		// get fractionalised streaming
		s := FractionaliseStreaming{
			TokenSymbol: fractionalise.TokenSymbol,
			Address:     holder,
		}

		if loaded, err := s.LoadState(stub); err != nil {
			return nil, err
		} else if loaded {
			streamings = append(streamings, s.StreamingId)
		}

		// generate transfers
		selling := coinbalance.TransferRequest{
			Type:   "Buy-Back-Fractionalise-" + input.TokenSymbol,
			Token:  fractionalise.TokenSymbol,
			From:   holder,
			To:     fractionalise.OwnerAddress,
			Amount: balance.Amount,
		}

		buying := coinbalance.TransferRequest{
			Type:   "Buy-Back-Fractionalise-" + input.TokenSymbol,
			Token:  fractionalise.FundingToken,
			From:   fractionalise.OwnerAddress,
			To:     holder,
			Amount: balance.Amount.Mul(fractionalise.BuyBackPrice),
		}

		transactions = append(transactions, selling, buying)
	}

	// stop active streamings
	if transfers, err := coinbalance.StopStreamings(stub, streamings...); err != nil {
		return nil, err
	} else {
		transactions = append(transactions, transfers...)
	}

	// make transfers
	r, err := coinbalance.Multitransfer(stub, transactions...)
	if err != nil {
		return nil, err
	}

	return new(Output).
		WithTransactions(r.Transactions...), nil
}

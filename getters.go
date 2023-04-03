package fractionalise

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetFractionaliseByTokenSymbol(stub shim.ChaincodeStubInterface, tokensymbol string) (*Fractionalise, error) {
	fractionalise := Fractionalise{
		TokenSymbol: tokensymbol,
	}

	if loaded, err := fractionalise.LoadState(stub); err != nil {
		return nil, err
	} else if !loaded {
		return nil, errors.New("fractionalised token not found")
	}

	return &fractionalise, nil
}

// GetFractionalisedStreamings returns the list of the fractionalised streamings that can take the following keys [tokensymbol, address]
func GetFractionalisedStreamings(stub shim.ChaincodeStubInterface, keys ...string) ([]FractionaliseStreaming, error) {

	it, err := stub.GetStateByPartialCompositeKey(FractionaliseStreamingIndex, keys)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var result []FractionaliseStreaming

	for it.HasNext() {
		state, err := it.Next()
		if err != nil {
			return nil, err
		}

		var data FractionaliseStreaming
		if err := json.Unmarshal(state.Value, &data); err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

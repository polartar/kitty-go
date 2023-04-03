package fractionalise

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (obj *FractionaliseStreaming) ToLedgerValue() ([]byte, error) {
	return json.Marshal(obj)
}

func (obj *FractionaliseStreaming) ToCompositeKey(stub shim.ChaincodeStubInterface) (string, error) {
	attributes := []string{obj.TokenSymbol, obj.Address}
	return stub.CreateCompositeKey(FractionaliseStreamingIndex, attributes)
}

func (obj *FractionaliseStreaming) SaveState(stub shim.ChaincodeStubInterface) error {
	key, err := obj.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	state, err := obj.ToLedgerValue()
	if err != nil {
		return err
	}

	return stub.PutState(key, state)
}

func (obj *FractionaliseStreaming) LoadState(stub shim.ChaincodeStubInterface) (bool, error) {
	key, err := obj.ToCompositeKey(stub)
	if err != nil {
		return false, err
	}

	state, err := stub.GetState(key)
	if err != nil {
		return false, err
	} else if state == nil {
		return false, nil
	}

	return true, json.Unmarshal(state, &obj)
}

func (obj *FractionaliseStreaming) DeleteState(stub shim.ChaincodeStubInterface) error {
	key, err := obj.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	return stub.DelState(key)
}

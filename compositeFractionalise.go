package fractionalise

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (obj *Fractionalise) ToLedgerValue() ([]byte, error) {
	return json.Marshal(obj)
}

func (obj *Fractionalise) ToCompositeKey(stub shim.ChaincodeStubInterface) (string, error) {
	attributes := []string{obj.TokenSymbol}
	return stub.CreateCompositeKey(FractionaliseIndex, attributes)
}

func (obj *Fractionalise) SaveState(stub shim.ChaincodeStubInterface) error {
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

// returns false if an Account object wasn't found in the ledger; otherwise returns true
func (obj *Fractionalise) LoadState(stub shim.ChaincodeStubInterface) (bool, error) {
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

func (obj *Fractionalise) DeleteState(stub shim.ChaincodeStubInterface) error {
	key, err := obj.ToCompositeKey(stub)
	if err != nil {
		return err
	}

	return stub.DelState(key)
}
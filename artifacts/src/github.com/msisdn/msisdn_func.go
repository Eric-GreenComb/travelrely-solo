package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *MsisdnChaincode) subscribe(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	msisdn := args[0]
	assetID := args[1]
	eki2 := args[2]
	bcuserID := args[3]
	bcuserKey := args[4]

	_keyM := t.generateMsisdnKey(msisdn)
	// Get the state from the ledger
	if _avalBytes, err := stub.GetState(_keyM); err == nil && _avalBytes != nil {
		var _msisdn Msisdn
		err = json.Unmarshal(_avalBytes, &_msisdn)
		if err != nil {
			jsonResp := "{\"error\":\"msisdn json error\"}"
			return shim.Error(jsonResp)
		}
		if _msisdn.Status == 1 {
			jsonResp := "{\"error\":\"msisdn has subscribe\"}"
			return shim.Error(jsonResp)
		}
	}

	// Write the state back to the ledger
	var _Msisdn Msisdn
	_Msisdn.Msisdn = msisdn
	_Msisdn.AssetID = assetID
	_Msisdn.UserID = bcuserID
	_Msisdn.UserKey = bcuserKey
	_Msisdn.Status = 1
	_bytesM, err := json.Marshal(_Msisdn)

	err = stub.PutState(_keyM, _bytesM)
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}

	var _Asset Asset
	_Asset.Msisdn = msisdn
	_Asset.AssetID = assetID
	_Asset.Eki2 = eki2
	_Asset.Status = 1
	_bytesA, err := json.Marshal(_Asset)

	_keyA := t.generateAssetKey(assetID)
	err = stub.PutState(_keyA, _bytesA)
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(nil)
}

func (t *MsisdnChaincode) unsubscribe(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke

	if len(args) != 4 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	msisdn := args[0]
	assetID := args[1]
	// bcuserID := args[2]
	// bcuserKey := args[3]

	// Get the state from the ledger
	_keyM := t.generateMsisdnKey(msisdn)
	// Get the state from the ledger
	_avalBytes, err := stub.GetState(_keyM)
	if err != nil || _avalBytes == nil {
		jsonResp := "{\"error\":\"msisdn is null\"}"
		return shim.Error(jsonResp)
	}
	var _msisdn Msisdn
	err = json.Unmarshal(_avalBytes, &_msisdn)
	if err != nil {
		jsonResp := "{\"error\":\"msisdn json error\"}"
		return shim.Error(jsonResp)
	}
	if _msisdn.Status != 1 {
		jsonResp := "{\"error\":\"msisdn hasn't subscribe\"}"
		return shim.Error(jsonResp)
	}

	// Write the state back to the ledger
	var _Msisdn Msisdn
	_Msisdn.Msisdn = msisdn
	_Msisdn.AssetID = ""
	_Msisdn.UserID = ""
	_Msisdn.UserKey = ""
	_Msisdn.Status = 0
	_bytesM, err := json.Marshal(_Msisdn)
	err = stub.PutState(_keyM, _bytesM)
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}

	var _Asset Asset
	_Asset.Msisdn = ""
	_Asset.AssetID = assetID
	_Asset.Eki2 = ""
	_Asset.Status = 0
	_bytesA, err := json.Marshal(_Asset)

	_keyA := t.generateAssetKey(msisdn)
	err = stub.PutState(_keyA, _bytesA)
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *MsisdnChaincode) msisdnState(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	msisdn := args[0]
	_keyM := t.generateMsisdnKey(msisdn)
	// Get the state from the ledger
	_avalBytes, err := stub.GetState(_keyM)

	if err != nil {
		return shim.Success(nil)
	}
	return shim.Success(_avalBytes)
}

// Query callback representing the query of a chaincode
func (t *MsisdnChaincode) getMsisdnHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	msisdn := args[0]
	_keyM := t.generateMsisdnKey(msisdn)

	resultsIterator, err := stub.GetHistoryForKey(_keyM)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var histories []TxHistory
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		var history TxHistory
		history.TxID = response.TxId
		if response.IsDelete {
			history.Value = "null"
		} else {
			history.Value = string(response.Value)
		}
		history.Timestamp = time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String()
		history.IsDelete = strconv.FormatBool(response.IsDelete)

		histories = append(histories, history)
	}

	bytesData, err := json.Marshal(histories)
	if err != nil {
		jsonResp := "{\"error\":\"json error\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(bytesData)
}

// Query callback representing the query of a chaincode
func (t *MsisdnChaincode) assetInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	assetID := args[0]
	_keyA := t.generateAssetKey(assetID)
	// Get the state from the ledger
	_avalBytes, err := stub.GetState(_keyA)

	if err != nil {
		return shim.Success(nil)
	}
	return shim.Success(_avalBytes)
}

// generateMsisdnKey generateMsisdnKey
func (t *MsisdnChaincode) generateMsisdnKey(uuid string) string {
	return fmt.Sprintf("msisdn_%s", uuid)
}

// generateAssetKey generateAssetKey
func (t *MsisdnChaincode) generateAssetKey(uuid string) string {
	return fmt.Sprintf("asset_%s", uuid)
}

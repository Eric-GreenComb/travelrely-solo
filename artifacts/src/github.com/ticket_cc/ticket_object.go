package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *TicketChaincode) createObj(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var _uuid string
	var _base64Obj string
	var err error

	if len(args) != 2 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" || args[1] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	_uuid = args[0]
	_base64Obj = args[1]

	// Get the state from the ledger
	_key := t.generateObjKey(_uuid)

	if _avalBytes, err := stub.GetState(_key); err == nil && _avalBytes != nil {
		jsonResp := "{\"error\":\"Obj has existed\"}"
		return shim.Error(jsonResp)
	}

	// Write the state back to the ledger
	err = stub.PutState(_key, []byte(_base64Obj))
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(nil)
}

func (t *TicketChaincode) updateObj(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var _uuid string
	var _base64Obj string
	var err error

	if len(args) != 2 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" || args[1] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	_uuid = args[0]
	_base64Obj = args[1]

	// Get the state from the ledger
	_key := t.generateObjKey(_uuid)

	// Write the state back to the ledger
	err = stub.PutState(_key, []byte(_base64Obj))
	if err != nil {
		jsonResp := "{\"error\":\"PutState error\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *TicketChaincode) queryObj(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	_uuid := args[0]

	_key := t.generateObjKey(_uuid)
	_byte, err := stub.GetState(_key)

	if err != nil {
		logger.Info("GetState error")
		jsonResp := "{\"error\":\"get state error\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(_byte)
}

// Query callback representing the query of a chaincode
func (t *TicketChaincode) queryObjs(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	_uuids := strings.Split(args[0], ",")

	var objs []string

	for _, _uuid := range _uuids {
		_key := t.generateObjKey(_uuid)
		_byte, err := stub.GetState(_key)
		if err != nil {
			objs = append(objs, "")
			continue
		}
		objs = append(objs, string(_byte))
	}

	_json, err := json.Marshal(objs)
	if err != nil {
		jsonResp := "{\"error\":\"json marshal error\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(_json)
}

// Query callback representing the query of a chaincode
func (t *TicketChaincode) getObjHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		jsonResp := "{\"error\":\"Incorrect number of arguments\"}"
		return shim.Error(jsonResp)
	}
	if args[0] == "" {
		jsonResp := "{\"error\":\"Arguments is nil\"}"
		return shim.Error(jsonResp)
	}

	_uuid := args[0]

	_key := t.generateObjKey(_uuid)

	resultsIterator, err := stub.GetHistoryForKey(_key)
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

// generateObjKey Generate ObjKey
func (t *TicketChaincode) generateObjKey(uuid string) string {
	return fmt.Sprintf("obj_%s", uuid)
}

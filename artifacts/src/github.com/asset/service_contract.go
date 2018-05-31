package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (asset *Asset) createServiceContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	_scsn := args[0]
	_scname := args[1]
	_account := args[2]
	_rate := args[3]
	_memo := args[4]

	_scsnKey := genServiceContractKey(_scsn)

	if _avalBytes, err := stub.GetState(_scsnKey); err == nil && _avalBytes != nil {
		jsonResp := "{\"error\":\"service contract has existed\"}"
		return shim.Error(jsonResp)
	}

	_contract := ServiceContract{}
	_contract.SCSN = _scsn
	_contract.SCName = _scname
	_contract.ServiceAccount = _account
	_contract.Rate, _ = strconv.Atoi(_rate)
	_contract.Memo = _memo

	_json, err := json.Marshal(_contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_scsnKey, _json)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) queryServiceContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	_scsn := args[0]
	_scsnKey := genServiceContractKey(_scsn)

	_accountBytes, err := stub.GetState(_scsnKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(_accountBytes)
}

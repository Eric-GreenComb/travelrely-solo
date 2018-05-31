package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (asset *Asset) loan(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	_scsn := args[4]
	_scsnKey := genServiceContractKey(_scsn)

	_scBytes, err := stub.GetState(_scsnKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	var _serviceContract ServiceContract
	err = json.Unmarshal(_scBytes, &_serviceContract)
	if err != nil {
		return shim.Error(err.Error())
	}

	_amount, _ := strconv.Atoi(args[3])
	_fee := _amount * _serviceContract.Rate / 100

	err = asset.loantrans(stub, args[0], args[1], args[2], args[3], _serviceContract.ServiceAccount, strconv.Itoa(_fee))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) loantrans(stub shim.ChaincodeStubInterface, symbol, from, to, amount, contract, fee string) error {
	_fromKey := genERC20AccountKey(symbol, from)
	_toKey := genERC20AccountKey(symbol, to)
	_contractKey := genERC20AccountKey(symbol, contract)
	_amount, _ := strconv.Atoi(amount)
	_fee, _ := strconv.Atoi(fee)

	if _amount <= 0 {
		return errors.New("Incorrect number of amount")
	}

	_memo := from + ":" + to + ":" + amount + ":" + contract + ":" + fee

	fromBytes, err := stub.GetState(_fromKey)
	if err != nil {
		return err
	}
	fromAccount := Account{}
	err = json.Unmarshal(fromBytes, &fromAccount)
	if err != nil {
		return err
	}
	if fromAccount.Frozen == true {
		return errors.New("from account is frozened")
	}

	if fromAccount.Balance < _amount+_fee {
		return errors.New("from account'balance < amount")
	}

	toAccount := Account{}
	toAccount.Frozen = false
	toBytes, err := stub.GetState(_toKey)
	if err == nil && toBytes != nil {
		err = json.Unmarshal(toBytes, &toAccount)
		if err != nil {
			return err
		}
	}
	if toAccount.Frozen == true {
		return errors.New("to account is frozened")
	}

	toContract := Account{}
	toContract.Frozen = false
	contractBytes, err := stub.GetState(_contractKey)
	if err == nil && contractBytes != nil {
		err = json.Unmarshal(contractBytes, &toContract)
		if err != nil {
			return err
		}
	}
	if toContract.Frozen == true {
		return errors.New("contract account is frozened")
	}

	fromAccount.Balance -= _amount
	fromAccount.Balance -= _fee

	fromAccount.Memo = _memo
	fromJSON, err := json.Marshal(fromAccount)
	if err != nil {
		return err
	}
	err = stub.PutState(_fromKey, fromJSON)
	if err != nil {
		return err
	}

	toAccount.Balance += _amount
	toAccount.Memo = _memo
	toJSON, err := json.Marshal(toAccount)
	if err != nil {
		return err
	}
	err = stub.PutState(_toKey, toJSON)
	if err != nil {
		return err
	}

	toContract.Balance += _fee
	toContract.Memo = _memo
	contractJSON, err := json.Marshal(toContract)
	if err != nil {
		return err
	}
	err = stub.PutState(_contractKey, contractJSON)
	if err != nil {
		return err
	}

	return nil
}

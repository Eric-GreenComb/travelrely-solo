package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (asset *Asset) coinbase(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	symbol := args[0]

	erc20Key := genERC20Key(symbol)

	if _avalBytes, err := stub.GetState(erc20Key); err == nil && _avalBytes != nil {
		jsonResp := "{\"error\":\"asset has existed\"}"
		return shim.Error(jsonResp)
	}

	name := args[1]
	supply, _ := strconv.Atoi(args[2])

	erc20 := &ERC20{
		Owner:       "coinbase",
		TotalSupply: supply,
		Name:        name,
		Symbol:      symbol,
		Lock:        false}

	tokenBytes, _ := json.Marshal(erc20)

	err := stub.PutState(erc20Key, tokenBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	_coinbase := &Account{
		Balance: supply,
		Frozen:  false}

	_accountBytes, _ := json.Marshal(_coinbase)

	_accountKey := genERC20AccountKey(symbol, "coinbase")
	err = stub.PutState(_accountKey, _accountBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) lock(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	erc20Key := genERC20Key(args[0])

	tokenBytes, err := stub.GetState(erc20Key)
	if err != nil {
		return shim.Error(err.Error())
	}

	erc20 := ERC20{}

	json.Unmarshal(tokenBytes, &erc20)
	_bool, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("args[1] is not bool true or false")
	}

	erc20.Lock = _bool

	_json, err := json.Marshal(erc20)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(erc20Key, _json)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) account(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	_fromKey := genERC20AccountKey(args[0], args[1])

	if _avalBytes, err := stub.GetState(_fromKey); err == nil && _avalBytes != nil {
		jsonResp := "{\"error\":\"account has existed\"}"
		return shim.Error(jsonResp)
	}

	_fromAccount := Account{}
	_fromAccount.Balance = 0
	_fromAccount.Frozen = false
	fromJSON, err := json.Marshal(_fromAccount)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_fromKey, fromJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	err := asset.trans(stub, args[0], args[1], args[2], args[3])
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) trans(stub shim.ChaincodeStubInterface, symbol, from, to, amount string) error {
	_fromKey := genERC20AccountKey(symbol, from)
	_toKey := genERC20AccountKey(symbol, to)
	_amount, _ := strconv.Atoi(amount)

	if _amount <= 0 {
		return errors.New("Incorrect number of amount")
	}

	_memo := from + ":" + to + ":" + amount

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

	if fromAccount.Balance < _amount {
		return errors.New("from account'balance < amount")
	}

	fromAccount.Balance -= _amount

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
	return nil
}

func (asset *Asset) frozen(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	_fromKey := genERC20AccountKey(args[0], args[1])

	var status bool
	if args[2] == "true" {
		status = true
	} else {
		status = false
	}

	fromBytes, err := stub.GetState(_fromKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	fromAccount := Account{}
	err = json.Unmarshal(fromBytes, &fromAccount)
	if err != nil {
		return shim.Error(err.Error())
	}

	fromAccount.Frozen = status

	fromJSON, err := json.Marshal(fromAccount)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_fromKey, fromJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) mint(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	symbol := args[0]
	_amount, _ := strconv.Atoi(args[1])

	erc20Key := genERC20Key(symbol)

	_avalBytes, err := stub.GetState(erc20Key)
	if err != nil {
		return shim.Error(err.Error())
	}

	var erc20 ERC20
	err = json.Unmarshal(_avalBytes, &erc20)
	if err != nil {
		return shim.Error(err.Error())
	}
	erc20.TotalSupply += _amount

	tokenJSON, _ := json.Marshal(erc20)

	err = stub.PutState(erc20Key, tokenJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	_accountKey := genERC20AccountKey(symbol, "coinbase")
	_accountBytes, err := stub.GetState(_accountKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	var _account Account
	err = json.Unmarshal(_accountBytes, &_account)
	if err != nil {
		return shim.Error(err.Error())
	}
	_account.Balance += _amount
	_accountJSON, _ := json.Marshal(_account)
	err = stub.PutState(_accountKey, _accountJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (asset *Asset) balance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	_accountKey := genERC20AccountKey(args[0], args[1])
	_accountBytes, err := stub.GetState(_accountKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(_accountBytes)
}

func (asset *Asset) history(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	_accountKey := genERC20AccountKey(args[0], args[1])

	resultsIterator, err := stub.GetHistoryForKey(_accountKey)
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

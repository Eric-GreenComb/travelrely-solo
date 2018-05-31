/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("ticket_cc")

// MsisdnChaincode msisdn Chaincode implementation
type MsisdnChaincode struct {
}

// Init Init
func (t *MsisdnChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### MsisdnChaincode Init ###########")

	return shim.Success(nil)

}

// Invoke Transaction
func (t *MsisdnChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### MsisdnChaincode Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "subscribe" {
		// subscribe : msisdn,asset_id,eki2,bcuser_id,bcuser_key
		return t.subscribe(stub, args)
	}
	if function == "unsubscribe" {
		// unsubscribe : msisdn,asset_id,bcuser_id,bcuser_key
		return t.unsubscribe(stub, args)
	}
	if function == "msisdn_state" {
		// msisdn_state : msisdn
		return t.msisdnState(stub, args)
	}
	if function == "get_msisdn_history" {
		// get_msisdn_history : msisdn
		return t.getMsisdnHistory(stub, args)
	}
	if function == "asset_info" {
		// asset_info : asset_id
		return t.assetInfo(stub, args)
	}

	logger.Errorf("Unknown action, got: %v", function)
	return shim.Error(fmt.Sprintf("Unknown action, got: %v", function))
}

func main() {
	err := shim.Start(new(MsisdnChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}

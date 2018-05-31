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

// TicketChaincode ticket Chaincode implementation
type TicketChaincode struct {
}

// Init Init
func (t *TicketChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### ticket_cc Init ###########")

	return shim.Success(nil)

}

// Invoke Transaction
func (t *TicketChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### ticket_cc Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "createObj" {
		// create obj : uuid(type_version_id),obj(base64)
		return t.createObj(stub, args)
	}
	if function == "updateObj" {
		// update obj : uuid(type_version_id),obj(base64)
		return t.updateObj(stub, args)
	}
	if function == "queryObj" {
		// query obj : uuid(type_version_id)
		return t.queryObj(stub, args)
	}
	if function == "queryObjs" {
		// queries an entity state
		return t.queryObjs(stub, args)
	}
	if function == "getObjHistory" {
		// get an obj's history
		return t.getObjHistory(stub, args)
	}

	logger.Errorf("Unknown action, got: %v", function)
	return shim.Error(fmt.Sprintf("Unknown action, got: %v", function))
}

func main() {
	err := shim.Start(new(TicketChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}

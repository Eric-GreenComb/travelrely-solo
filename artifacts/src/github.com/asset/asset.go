/*
Copyright FiFu Corp. 2018 All Rights Reserved.

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

// Asset the Asset Smart Contract structure
type Asset struct {
}

// Init Init
func (asset *Asset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke Invoke
func (asset *Asset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()

	// invoke
	if function == "coinbase" {
		return asset.coinbase(stub, args)
	} else if function == "lock" {
		return asset.lock(stub, args)
	} else if function == "account" {
		return asset.account(stub, args)
	} else if function == "transfer" {
		return asset.transfer(stub, args)
	} else if function == "frozen" {
		return asset.frozen(stub, args)
	} else if function == "mint" {
		return asset.mint(stub, args)
	}
	// Query
	if function == "balance" {
		return asset.balance(stub, args)
	}
	if function == "history" {
		return asset.history(stub, args)
	}

	// service contract
	if function == "createServiceContract" {
		return asset.createServiceContract(stub, args)
	}
	// Query
	if function == "queryServiceContract" {
		return asset.queryServiceContract(stub, args)
	}

	if function == "loan" {
		return asset.loan(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(Asset))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

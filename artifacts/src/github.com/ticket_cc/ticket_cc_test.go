package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestTicketCC_Init(t *testing.T) {
	tcc := new(TicketChaincode)
	stub := shim.NewMockStub("ticket_cc", tcc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	checkInvokeCreateObj(t, stub, [][]byte{[]byte("createObj"), []byte("8d1f65547e284dadb4e42372fc879985"), []byte("ZGRk")})

	checkInvokeQueryObj(t, stub, [][]byte{[]byte("queryObj"), []byte("8d1f65547e284dadb4e42372fc879985")})

}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("init", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvokeCreateObj(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("createObj", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvokeQueryObj(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("queryObj", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

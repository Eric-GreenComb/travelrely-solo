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
)

// ERC20 ERC20 token
type ERC20 struct {
	Owner       string `json:"Owner"`
	Symbol      string `json:"Symbol"`
	Name        string `json:"Name"`
	TotalSupply int    `json:"TotalSupply"`
	Lock        bool   `json:"Lock"`
}

// Account account
type Account struct {
	Balance int    `json:"BalanceOf"`
	Frozen  bool   `json:"Frozen"`
	Memo    string `json:"memo"`
}

// ServiceContract ServiceContract
type ServiceContract struct {
	SCSN           string `json:"scsn"`    // 服务合同编号
	SCName         string `json:"scname"`  // 服务合同名称
	ServiceAccount string `json:"account"` // 服务账户
	Rate           int    `json:"rate"`    // 服务费率 1 -> 1%
	Memo           string `json:"memo"`
}

// TxHistory txid History
type TxHistory struct {
	TxID      string `json:"txid,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	IsDelete  string `json:"isdelete,omitempty"`
}

func genERC20Key(symbol string) string {
	return fmt.Sprintf("erc20_%s", symbol)
}

func genERC20AccountKey(symbol, owner string) string {
	return fmt.Sprintf("erc20_%s_%s", symbol, owner)
}

func genServiceContractKey(scsn string) string {
	return fmt.Sprintf("scsn_%s", scsn)
}

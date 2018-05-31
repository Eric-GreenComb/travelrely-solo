package main

// TxHistory txid History
type TxHistory struct {
	TxID      string `json:"txid,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	IsDelete  string `json:"isdelete,omitempty"`
}

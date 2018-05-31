package main

// TxHistory txid History
type TxHistory struct {
	TxID      string `json:"txid,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	IsDelete  string `json:"isdelete,omitempty"`
}

// Msisdn Msisdn
type Msisdn struct {
	Msisdn  string `json:"msisdn,omitempty"`
	AssetID string `json:"asset_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	UserKey string `json:"user_key,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// Asset Asset
type Asset struct {
	AssetID string `json:"asset_id,omitempty"`
	Msisdn  string `json:"msisdn,omitempty"`
	Eki2    string `json:"eki2,omitempty"`
	Status  int    `json:"status,omitempty"`
}

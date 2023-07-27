package entity

type Balance struct {
	Free string `json:"free"`
}

type TxRecords struct {
	Rows  []TxRecord `json:"rows"`
	Total int        `json:"total"`
}

type TxRecord struct {
	Amount    string `json:"amount"`
	Asset     string `json:"asset"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	TxId      int64  `json:"txId"`
	Type      string `json:"type"`
}

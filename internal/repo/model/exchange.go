package model

import "noneland/backend/interview/internal/entity"

type Balance struct {
	Free string `json:"free"`
}

func BalanceModelToEntity(input *Balance) *entity.Balance {
	return &entity.Balance{
		Free: input.Free,
	}
}

func BalanceEntityToModel(input *entity.Balance) *Balance {
	return &Balance{
		Free: input.Free,
	}
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

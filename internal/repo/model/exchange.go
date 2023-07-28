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

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

type ExchangeInfo struct {
	Timezone   string      `json:"timezone"`
	ServerTime int64       `json:"serverTime"`
	RateLimits []RateLimit `json:"rateLimits"`
}

func ExchangeInfoModelToEntity(input *ExchangeInfo) *entity.ExchangeInfo {
	var rateLimits []entity.RateLimit
	for _, rl := range input.RateLimits {
		rateLimits = append(rateLimits, entity.RateLimit{
			RateLimitType: rl.RateLimitType,
			Interval:      rl.Interval,
			IntervalNum:   rl.IntervalNum,
			Limit:         rl.Limit,
		})
	}

	return &entity.ExchangeInfo{
		Timezone:   input.Timezone,
		ServerTime: input.ServerTime,
		RateLimits: rateLimits,
	}
}

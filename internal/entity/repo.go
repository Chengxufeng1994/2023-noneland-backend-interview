package entity

type Repository interface {
	// User
	GetUsers() (users []User, err error)
	// Exchange
	GetSpotExchangeInfo() (exchangeInfo ExchangeInfo, err error)
	GetFuturesExchangeInfo() (exchangeInfo ExchangeInfo, err error)
	GetSpotBalance() (balance Balance, err error)
	GetFuturesBalance() (balance Balance, err error)
	GetTxRecords(arg GetTxRecordsArg) (txRecords TxRecords, err error)
}

type GetTxRecordsArg struct {
	StartTime int64
	EndTime   int64
	Current   int64
	Size      int64
}

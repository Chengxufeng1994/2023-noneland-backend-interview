package entity

type Repository interface {
	// User
	GetUsers() (users []User, err error)
	// Exchange
	GetSpotBalance() (balance Balance, err error)
	GetFuturesBalance() (balance Balance, err error)
	GetTxRecords() (txRecords TxRecords, err error)
}

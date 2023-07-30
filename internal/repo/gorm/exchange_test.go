package gorm

import (
	"github.com/stretchr/testify/require"
	"noneland/backend/interview/internal/entity"
	"testing"
	"time"
)

func TestGetSpotBalance(t *testing.T) {
	result, err := repo.GetSpotBalance()
	require.NoError(t, err)
	require.NotEmpty(t, result)

	answer := entity.Balance{
		Free: "10.12345",
	}
	require.Equal(t, result, answer)
}

func TestGetFuturesBalance(t *testing.T) {
	result, err := repo.GetFuturesBalance()
	require.NoError(t, err)
	require.NotEmpty(t, result)

	answer := entity.Balance{
		Free: "10.12345",
	}
	require.Equal(t, result, answer)
}

func TestGetTxRecords(t *testing.T) {
	end := time.Now()
	start := end.AddDate(0, -1, 0)
	args := entity.GetTxRecordsArg{
		Current:   1,
		Size:      10,
		StartTime: start.Unix(),
		EndTime:   end.Unix(),
	}
	result, err := repo.GetTxRecords(args)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	answer := entity.TxRecords{
		Rows: []entity.TxRecord{
			{
				Amount:    "0.10000000",
				Asset:     "BNB",
				Status:    "CONFIRMED",
				Timestamp: 1566898617,
				TxId:      5240372201,
				Type:      "IN",
			},
			{
				Amount:    "5.00000000",
				Asset:     "USDT",
				Status:    "CONFIRMED",
				Timestamp: 1566888436,
				TxId:      5239810406,
				Type:      "OUT",
			},
			{
				Amount:    "1.00000000",
				Asset:     "EOS",
				Status:    "CONFIRMED",
				Timestamp: 1566888403,
				TxId:      5239808703,
				Type:      "IN",
			},
		},
		Total: 3,
	}

	require.Equal(t, result.Total, answer.Total)
	for i, row := range result.Rows {
		require.Equal(t, answer.Rows[i].Amount, row.Amount)
		require.Equal(t, answer.Rows[i].Asset, row.Asset)
		require.Equal(t, answer.Rows[i].Status, row.Status)
		require.Equal(t, answer.Rows[i].Timestamp, row.Timestamp)
		require.Equal(t, answer.Rows[i].TxId, row.TxId)
		require.Equal(t, answer.Rows[i].Type, row.Type)
	}
}

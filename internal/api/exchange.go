package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/internal/di"
	"noneland/backend/interview/internal/entity"
)

type getBalanceResponse struct {
	OK              bool   `json:"ok"`
	SpotBalance     string `json:"spot_balance"`
	FeaturesBalance string `json:"futures_balance"`
}

func GetBalance(c *gin.Context) {
	_ = c.Param("userId")
	repo, err := di.NewRepo()
	if err != nil {
		log.Printf("create repository failed: %s", err.Error())
		errResponse(c)
		return
	}

	spotBalance, err := repo.GetSpotBalance()
	if err != nil {
		log.Printf("get spot balance failed: %s", err.Error())
		errResponse(c)
		return
	}
	futuresBalance, err := repo.GetFuturesBalance()
	if err != nil {
		log.Printf("get future balance failed: %s", err.Error())
		errResponse(c)
		return
	}

	c.JSON(http.StatusOK,
		getBalanceResponse{
			OK:              true,
			SpotBalance:     spotBalance.Free,
			FeaturesBalance: futuresBalance.Free,
		},
	)
}

type getTxRecordsRequest struct {
	StartTime int64 `form:"startTime"`
	EndTime   int64 `form:"endTime"`
	Current   int64 `form:"current"`
	Size      int64 `form:"size" `
}

type getTxRecordsResponse struct {
	OK        bool             `json:"ok"`
	TxRecords entity.TxRecords `json:"data"`
}

func GetTxRecords(c *gin.Context) {
	var req getTxRecordsRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("request bind failed: %s", err.Error())
		errResponse(c)
		return
	}

	repo, err := di.NewRepo()
	if err != nil {
		log.Printf("create repository failed: %s", err.Error())
		errResponse(c)
		return
	}
	records, err := repo.GetTxRecords()
	c.JSON(http.StatusOK, getTxRecordsResponse{
		OK:        true,
		TxRecords: records,
	})
}

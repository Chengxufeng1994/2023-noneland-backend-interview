package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

type getTxRecordsResponse struct {
	OK        bool             `json:"ok"`
	TxRecords entity.TxRecords `json:"data"`
}

func GetTxRecords(c *gin.Context) {
	startTimeStr := c.DefaultQuery("startTime", "0")
	endTimeStr := c.DefaultQuery("endTime", "0")
	currentStr := c.DefaultQuery("current", "1")
	sizeStr := c.DefaultQuery("size", "10")

	startTime, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		log.Printf("parse start time failed: %s", err.Error())
		c.Error(NewError(http.StatusBadRequest, err.Error()))
		return
	}
	endTime, err := strconv.ParseInt(endTimeStr, 10, 64)
	if err != nil {
		log.Printf("parse end time failed: %s", err.Error())
		c.Error(NewError(http.StatusBadRequest, err.Error()))
		return
	}
	current, err := strconv.ParseInt(currentStr, 10, 64)
	if err != nil {
		log.Printf("parse current failed: %s", err.Error())
		c.Error(NewError(http.StatusBadRequest, err.Error()))
		return
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		log.Printf("parse size failed: %s", err.Error())
		c.Error(NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if size < 10 || size > 100 {
		msg := fmt.Sprintf("size must greater then 10 and less than 100")
		c.Error(NewError(http.StatusBadRequest, msg))
		return
	}

	var args entity.GetTxRecordsArg
	args.StartTime = startTime
	args.EndTime = endTime
	args.Current = current
	args.Size = size

	repo, err := di.NewRepo()
	if err != nil {
		log.Printf("create repository failed: %s", err.Error())
		errResponse(c)
		return
	}
	records, err := repo.GetTxRecords(args)
	c.JSON(http.StatusOK, getTxRecordsResponse{
		OK:        true,
		TxRecords: records,
	})
}

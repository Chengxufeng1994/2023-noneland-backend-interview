package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"noneland/backend/interview/internal/limiter"
	"strconv"
	"time"

	"noneland/backend/interview/internal/di"
	"noneland/backend/interview/internal/entity"
)

var spotReqWeightRateLimiter limiter.RateLimiter
var spotRawReqRateLimiter limiter.RateLimiter
var futuresReqWeightRateLimiter limiter.RateLimiter
var futuresRawReqRateLimiter limiter.RateLimiter

func calculateRate(interval string, intervalNum int, limit int) float64 {
	var seconds int
	if interval == "SECONDS" {
		seconds = 1
	} else if interval == "MINUTE" {
		seconds = 60
	}

	return float64(limit / seconds / intervalNum)
}

func init() {
	repo, err := di.NewRepo()
	if err != nil {
		log.Printf("create repository failed: %s", err.Error())
	}
	spotExchangeInfo, err := repo.GetSpotExchangeInfo()
	if err != nil {
		log.Printf("get spot exchange info failed: %s", err.Error())
	}
	futuresExchangeInfo, err := repo.GetFuturesExchangeInfo()
	if err != nil {
		log.Printf("get futures exchange info failed: %s", err.Error())
	}
	for _, rl := range spotExchangeInfo.RateLimits {
		rate := calculateRate(rl.Interval, rl.IntervalNum, rl.Limit)
		switch rl.RateLimitType {
		case "REQUEST_WEIGHT":
			spotReqWeightRateLimiter = limiter.New(rl.Limit, rate)
		case "RAW_REQUESTS":
			spotRawReqRateLimiter = limiter.New(rl.Limit, rate)
		}
	}
	for _, rl := range futuresExchangeInfo.RateLimits {
		rate := calculateRate(rl.Interval, rl.IntervalNum, rl.Limit)
		switch rl.RateLimitType {
		case "REQUEST_WEIGHT":
			futuresReqWeightRateLimiter = limiter.New(rl.Limit, rate)
		case "RAW_REQUESTS":
			futuresRawReqRateLimiter = limiter.New(rl.Limit, rate)
		}
	}
}

type getBalanceResponse struct {
	OK              bool   `json:"ok"`
	SpotBalance     string `json:"spot_balance"`
	FeaturesBalance string `json:"futures_balance"`
}

func GetBalance(c *gin.Context) {
	repo, err := di.NewRepo()
	if err != nil {
		log.Printf("create repository failed: %s", err.Error())
		errResponse(c)
		return
	}

	if !spotReqWeightRateLimiter.Allow() {
		log.Printf("spot request weight exceed")
		c.Error(NewError(http.StatusTooManyRequests, "spot request weight exceed"))
		return
	}
	if !spotRawReqRateLimiter.Allow() {
		log.Printf("spot raw request exceed")
		c.Error(NewError(http.StatusTooManyRequests, "spot raw request exceed"))
		return
	}
	if !futuresReqWeightRateLimiter.Allow() {
		log.Printf("futures request weight exceed")
		c.Error(NewError(http.StatusTooManyRequests, "future request weight exceed"))
		return
	}
	if !futuresRawReqRateLimiter.Allow() {
		log.Printf("futures request weight exceed")
		c.Error(NewError(http.StatusTooManyRequests, "future raw request exceed"))
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
	now := time.Now()
	endTs := time.Now().Unix()
	startTs := now.AddDate(-6, 0, 0).Unix()
	startTimeStr := c.DefaultQuery("startTime", strconv.FormatInt(startTs, 10))
	endTimeStr := c.DefaultQuery("endTime", strconv.FormatInt(endTs, 10))
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
	if startTime > endTime {
		msg := fmt.Sprintf("end time must greater than start time")
		c.Error(NewError(http.StatusBadRequest, msg))
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

	if !spotReqWeightRateLimiter.Allow() {
		log.Printf("spot request weight exceed")
		c.Error(NewError(http.StatusTooManyRequests, "spot request weight exceed"))
		return
	}
	if !spotRawReqRateLimiter.Allow() {
		log.Printf("spot raw request exceed")
		c.Error(NewError(http.StatusTooManyRequests, "spot raw request exceed"))
		return
	}
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

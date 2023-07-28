package gorm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/repo/model"
)

const spotUrl = "https://exchange.com/api"
const featuresUrl = "https://exhcange.com/api"

func (repo *repository) GetSpotExchangeInfo() (exchangeInfo entity.ExchangeInfo, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := spotUrl + "/exchangeInfo"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("?api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
		response, err = http.Get(baseUrl)
		if err != nil {
			return exchangeInfo, err
		}
	}

	if response != nil {
		reader = response.Body
	} else {
		reader, err = os.Open("test/spot_exchange_info.json")
		if err != nil {
			return exchangeInfo, err
		}
	}

	rawData, err := io.ReadAll(reader)
	if err != nil {
		return exchangeInfo, err
	}

	data := &model.ExchangeInfo{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return exchangeInfo, err
	}

	exchangeInfo = *model.ExchangeInfoModelToEntity(data)

	return
}
func (repo *repository) GetFuturesExchangeInfo() (exchangeInfo entity.ExchangeInfo, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := spotUrl + "/exchangeInfo"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("?api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
		response, err = http.Get(baseUrl)
		if err != nil {
			return exchangeInfo, err
		}
	}

	if response != nil {
		reader = response.Body
	} else {
		reader, err = os.Open("test/spot_exchange_info.json")
		if err != nil {
			return exchangeInfo, err
		}
	}

	rawData, err := io.ReadAll(reader)
	if err != nil {
		return exchangeInfo, err
	}

	data := &model.ExchangeInfo{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return exchangeInfo, err
	}

	exchangeInfo = *model.ExchangeInfoModelToEntity(data)

	return
}

func (repo *repository) GetSpotBalance() (balance entity.Balance, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := spotUrl + "/spot/balance"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("?api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
		response, err = http.Get(baseUrl)
		if err != nil {
			return balance, err
		}
	}

	if response != nil {
		reader = response.Body
	} else {
		reader, err = os.Open("test/spot_balance.json")
		if err != nil {
			return balance, err
		}
	}

	rawData, err := io.ReadAll(reader)
	if err != nil {
		return balance, err
	}

	data := &model.Balance{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return balance, err
	}

	balance = *model.BalanceModelToEntity(data)

	return
}

func (repo *repository) GetFuturesBalance() (balance entity.Balance, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := featuresUrl + "/futures/balance"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("?api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
		response, err = http.Get(baseUrl)
		if err != nil {
			return balance, err
		}
	}

	if response != nil {
		reader = response.Body
	} else {
		reader, err = os.Open("test/futures_balance.json")
		if err != nil {
			return balance, err
		}
	}

	rawData, err := io.ReadAll(reader)
	if err != nil {
		return balance, err
	}

	data := &model.Balance{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		return balance, err
	}

	balance = *model.BalanceModelToEntity(data)

	return
}

func (repo *repository) GetTxRecords(args entity.GetTxRecordsArg) (txRecords entity.TxRecords, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := spotUrl + "/spot/transfer/records?"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("current=%d", args.Current)
		baseUrl += fmt.Sprintf("&size=%d", args.Size)
		if args.StartTime != 0 {
			baseUrl += fmt.Sprintf("&startTime=%d", args.StartTime)
		}
		if args.EndTime != 0 {
			baseUrl += fmt.Sprintf("&endTime=%d", args.EndTime)
		}
		baseUrl += fmt.Sprintf("&api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
		response, err = http.Get(baseUrl)
		if err != nil {
			return
		}
	}

	if response != nil {
		reader = response.Body
	} else {
		reader, err = os.Open("test/spot_transfer_records.json")
		if err != nil {
			return
		}
	}

	rawData, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	if err = json.Unmarshal(rawData, &txRecords); err != nil {
		return
	}

	return
}

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

const domain = "https://xx-exchange.com"

func (repo *repository) GetSpotBalance() (balance entity.Balance, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := domain + "/spot/balance"
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
	baseUrl := domain + "/spot/balance"
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

func (repo *repository) GetTxRecords() (txRecords entity.TxRecords, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := domain + "/spot/transfer/records"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		baseUrl += fmt.Sprintf("?api_key=%s&api_secret=%s", repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret)
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

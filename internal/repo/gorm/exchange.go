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
const featuresUrl = "https://exchange.com/api"

func makeAPIRequest(client *http.Client, apiKey, apiSecret, apiURL string) (*http.Response, error) {
	apiURL += fmt.Sprintf("?api_key=%s&api_secret=%s", apiKey, apiSecret)
	// Create an HTTP request with the appropriate headers, body, etc.
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request using the client with rate limiting.
	return client.Do(req)
}

func (repo *repository) GetSpotExchangeInfo() (exchangeInfo entity.ExchangeInfo, err error) {
	var reader io.Reader
	var response *http.Response
	baseUrl := spotUrl + "/exchangeInfo"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		client := &http.Client{}
		response, err = makeAPIRequest(client, repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret, baseUrl)
		if err != nil {
			return exchangeInfo, err
		}
	}

	if response != nil {
		defer response.Body.Close()
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
	baseUrl := featuresUrl + "/exchangeInfo"
	if repo.config.XXExchange.ApiKey != "" && repo.config.XXExchange.ApiSecret != "" {
		client := &http.Client{}
		response, err = makeAPIRequest(client, repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret, baseUrl)
		if err != nil {
			return exchangeInfo, err
		}
	}

	if response != nil {
		defer response.Body.Close()
		reader = response.Body
	} else {
		reader, err = os.Open("test/futures_exchange_info.json")
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
		client := &http.Client{}
		response, err = makeAPIRequest(client, repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret, baseUrl)
		if err != nil {
			return balance, err
		}
	}

	if response != nil {
		defer response.Body.Close()
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
		client := &http.Client{}
		response, err = makeAPIRequest(client, repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret, baseUrl)
		if err != nil {
			return balance, err
		}
	}

	if response != nil {
		defer response.Body.Close()
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
		baseUrl += fmt.Sprintf("&startTime=%d", args.StartTime)
		baseUrl += fmt.Sprintf("&endTime=%d", args.EndTime)
		client := &http.Client{}
		response, err = makeAPIRequest(client, repo.config.XXExchange.ApiKey, repo.config.XXExchange.ApiSecret, baseUrl)
		if err != nil {
			return
		}
	}

	if response != nil {
		defer response.Body.Close()
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

package MarketManagers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
	"trading-automation-system/api/internal/MarketManagers/clients"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type MockedMarketManager struct{}

func (m *MockedMarketManager) Get(dateFrom, dateTo *time.Time, timeFrame constants.TimeFrame) ([]domain.CandleStick, error) {
	jsonFile, err := os.Open("./api/internal/mocks/historical_BTCUSDT_1h.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data [][]interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	binanceClient := clients.BinanceApi{}

	var candleStickCollection []domain.CandleStick
	for _, r := range data {
		candleStickCollection = append(candleStickCollection, binanceClient.ParseResponse(r))
	}

	return candleStickCollection, nil
}

func (m *MockedMarketManager) FullBuy(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return &MarketOperation{
		Quantity:   quantity,
		EntryPrice: price,
	}, nil
}

func (m *MockedMarketManager) FullSell(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return &MarketOperation{
		Quantity:   quantity,
		EntryPrice: price,
	}, nil
}

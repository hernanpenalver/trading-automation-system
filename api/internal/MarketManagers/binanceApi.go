package MarketManagers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type BinanceApi struct {
}

func NewBinanceApi() *BinanceApi {
	return &BinanceApi{}
}

const limit = 1000

func (b *BinanceApi) Get(symbol constants.Symbol, interval constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) ([]domain.CandleStick, error) {
	var candleStickList []domain.CandleStick

	for dateFrom.Before(*dateTo) {
		aux, err := b.get(symbol, interval, dateFrom, dateTo)
		if err != nil {
			return nil, err
		}

		if len(aux) > 0 {
			candleStickList = append(candleStickList, aux...)
		}

		closeDateTime := time.Unix(0, int64(aux[len(aux)-1].CloseTime)*int64(time.Millisecond))
		dateFrom = &closeDateTime
	}

	return candleStickList, nil
}

func (b *BinanceApi) get(symbol constants.Symbol, interval constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) ([]domain.CandleStick, error) {
	uri := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=%s&limit=%d", symbol, interval, limit)

	if dateFrom != nil && dateTo != nil {
		dateFromMillis := dateFrom.UnixNano() / int64(time.Millisecond)
		dateToMillis := dateTo.UnixNano() / int64(time.Millisecond)

		uri += fmt.Sprintf("&startTime=%d&endTime=%d", dateFromMillis, dateToMillis)
	}

	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var binanceResponse [][]interface{}
	err = json.Unmarshal(responseData, &binanceResponse)
	if err != nil {
		return nil, err
	}

	var candleStickCollection []domain.CandleStick
	for _, r := range binanceResponse {
		candleStickCollection = append(candleStickCollection, b.ParseResponse(r))
	}

	return candleStickCollection, nil
}

func (b *BinanceApi) ParseResponse(data []interface{}) domain.CandleStick {
	openTime := data[0].(float64)
	closeTime := data[6].(float64)
	closePrice, _ := strconv.ParseFloat(data[4].(string), 64)
	open, _ := strconv.ParseFloat(data[1].(string), 64)
	max, _ := strconv.ParseFloat(data[2].(string), 64)
	min, _ := strconv.ParseFloat(data[3].(string), 64)

	openDateTime := time.Unix(0, int64(openTime)*int64(time.Millisecond))

	return domain.CandleStick{
		OpenTime:     int64(openTime),
		CloseTime:    int64(closeTime),
		Close:        closePrice,
		Open:         open,
		Max:          max,
		Min:          min,
		OpenDateTime: openDateTime.String(),
	}
}

func (b *BinanceApi) FullBuy(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return &MarketOperation{
		Quantity:   quantity,
		EntryPrice: price,
	}, nil
}

func (b *BinanceApi) FullSell(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return &MarketOperation{
		Quantity:   quantity,
		EntryPrice: price,
	}, nil
}

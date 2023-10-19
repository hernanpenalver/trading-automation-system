package MarketManagers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

const binanceBasePath = "https://api.binance.com"
const binanceNewOrderUri = "/api/v3/order/test"
const defaultLimit = 1000

func (b *BinanceApi) Get(symbol constants.Symbol, interval constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) ([]domain.CandleStick, error) {
	var candleStickList []domain.CandleStick

	for dateFrom.Before(*dateTo) {
		aux, err := b.get(symbol, interval, dateFrom, dateTo, 0)
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

func (b *BinanceApi) GetLasts(symbol constants.Symbol, interval constants.TimeFrame, limit int) ([]domain.CandleStick, error) {
	return b.get(symbol, interval, nil, nil, limit)
}

func (b *BinanceApi) get(symbol constants.Symbol, interval constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time, limit int) ([]domain.CandleStick, error) {
	uri := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=%s", symbol, interval)

	if dateFrom != nil && dateTo != nil {
		dateFromMillis := dateFrom.UnixNano() / int64(time.Millisecond)
		dateToMillis := dateTo.UnixNano() / int64(time.Millisecond)

		uri += fmt.Sprintf("&startTime=%d&endTime=%d", dateFromMillis, dateToMillis)
	}

	if limit != 0 {
		uri += fmt.Sprintf("&limit=%d", limit)
	} else {
		uri += fmt.Sprintf("&limit=%d", defaultLimit)
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

func (b *BinanceApi) Buy(symbol string, quantity float64) (*BuyResponse, error) {
	uri := fmt.Sprint(binanceBasePath, binanceNewOrderUri)

	postBody := buyOrder{
		Symbol:    symbol,
		Side:      "BUY",
		Type:      "MARKET",
		Quantity:  quantity,
		Timestamp: time.Now().UnixMilli(),
	}
	jsonData, err := json.Marshal(postBody)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var response BuyResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	return &response, nil
}

type buyOrder struct {
	Symbol    string  `json:"symbol"`
	Side      string  `json:"side"` // BUY, SELL
	Type      string  `json:"type"`
	Quantity  float64 `json:"quantity"`
	Timestamp int64   `json:"timestamp"`
}

type BuyResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

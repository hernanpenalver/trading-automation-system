package MarketManagers

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
	"trading-automation-system/api/internal/domain"
)

// const url = "wss://testnet.binance.vision/stream?streams=btcusdt@kline_5m"
const url = "wss://testnet.binance.vision/stream?streams=btcusdt@kline_1m"

type BinanceStream struct {
	streamConn *websocket.Conn
}

func NewBinanceStream() *BinanceStream {
	return &BinanceStream{}
}

func (b *BinanceStream) StartConnection() error {
	streamConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	b.streamConn = streamConn
	return err
}

func (b *BinanceStream) GetNextResponse() *domain.CandleStick {
	var kline KlineRequestResponse
	if err := b.streamConn.ReadJSON(&kline); err != nil {
		log.Fatal(err)
	}

	openTime := kline.Data.K.StartTime
	closeTime := kline.Data.K.CloseTime
	closePrice, _ := strconv.ParseFloat(kline.Data.K.ClosePrice, 64)
	open, _ := strconv.ParseFloat(kline.Data.K.OpenPrice, 64)
	//max, _ := strconv.ParseFloat(data[2].(string), 64)
	//min, _ := strconv.ParseFloat(data[3].(string), 64)

	openDateTime := time.Unix(0, int64(openTime)*int64(time.Millisecond))

	return &domain.CandleStick{
		OpenTime:  int64(openTime),
		CloseTime: int64(closeTime),
		Close:     closePrice,
		Open:      open,
		//Max:          max,
		//Min:          min,
		OpenDateTime: openDateTime.String(),
		IsClose:      kline.Data.K.KlineIsClosed,
	}
}

func (b *BinanceStream) CloseConnection() error {
	return b.streamConn.Close()
}

type KlineRequestResponse struct {
	Data KlineEventResponse `json:"data"`
}

type KlineEventResponse struct {
	EventType string `json:"e"`
	EventTime int    `json:"E"`
	Symbol    string `json:"s"`
	K         struct {
		StartTime  int    `json:"t"`
		CloseTime  int    `json:"T"`
		Interval   string `json:"i"`
		OpenPrice  string `json:"o"`
		ClosePrice string `json:"c"`
		HighPrice  string `json:"h"`
		//LowPrice      string `json:"l"`
		KlineIsClosed bool `json:"x"`
	} `json:"k"`
}

package livetest

import (
	"log"
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/strategies"
)

type Service struct {
	Market    *MarketManagers.BinanceStream
	MarketApi *MarketManagers.BinanceApi
}

func NewService(market *MarketManagers.BinanceStream) *Service {
	return &Service{Market: market}
}

func (l *Service) Execute() (interface{}, error) {
	strategy := strategies.CrossingSimpleMovingAverages{
		Name:    strategies.CrossingSimpleMovingAverageName,
		FastSma: indicators.NewSimpleMovingAverage(37, indicators.CloseSource),
		SlowSma: indicators.NewSimpleMovingAverage(40, indicators.CloseSource),
	}

	if err := l.Market.StartConnection(); err != nil {
		log.Fatal(err)
	}

	var candleStick *domain.CandleStick
	//var candleStickList []domain.CandleStick

	candleStickList, err := l.MarketApi.GetLasts("BTCUSDT", "1m", 40)
	if err != nil {
		return nil, err
	}

	log.Printf("Lista de klines encontrada [len: %d]", len(candleStickList))

	for {
		if candleStick = l.Market.GetNextResponse(); candleStick == nil || !candleStick.IsClose {
			continue
		}

		candleStickList = append(candleStickList, *candleStick)
		log.Print(candleStick)
		log.Print(len(candleStickList))

		if operation := strategy.GetOperation(candleStickList); operation != nil || true {
			log.Printf("Operacion encontrada")
			log.Print(operation)

			response, err := l.MarketApi.Buy("BTCUSDT", 0.00000001)
			if err != nil {
				log.Print("Error: ", err)
			}

			log.Print("Response: ", response)
		}

	}
}

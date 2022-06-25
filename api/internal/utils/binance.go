package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"trading-automation-system/api/internal/domain"
)

func GetCandleStickOfMock(mockPath string) ([]domain.CandleStick, error) {
	jsonFile, err := os.Open(mockPath)
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

	var candleStickCollection []domain.CandleStick
	for _, r := range data {
		candleStickCollection = append(candleStickCollection, ParseBinanceResponse(r))
	}

	return candleStickCollection, nil
}

func ParseBinanceResponse(data []interface{}) domain.CandleStick {
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

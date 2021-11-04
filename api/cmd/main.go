package main

import (
	"fmt"
	"time"
	"trading-automation-system/api/internal/MarketManagers/clients"
)

func main() {
	bapi := clients.BinanceApi{}

	symbol := "BNBBTC"
	interval := "5m"
	dateFrom := time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	dateTo := time.Date(2021, 2, 1, 1, 1, 1, 1, time.UTC)
	a, _ := bapi.Get(symbol, interval, &dateFrom, &dateTo)
	fmt.Print(a[0])
}

package datadog

import (
	"fmt"
	"github.com/DataDog/datadog-go/v5/statsd"
	"log"
)

type DatadogCustom struct {
	Prefix string
	Client statsd.ClientInterface
}

var client *DatadogCustom

func New() {
	statsd, err := statsd.New("127.0.0.1:8125", statsd.WithNamespace("trading_automation_system"))
	if err != nil {
		log.Fatal(err)
	}

	err = statsd.Count("example_metric.increment", 1, []string{"environment:dev"}, 1)

	if err != nil {
		log.Fatal(err)
	}
	client = &DatadogCustom{Prefix: "trading_automation_system", Client: statsd}
}

func MetricStrategyResult(investmentBalance float64)  {
	tags := []string{fmt.Sprintf("investment_balance:%f",investmentBalance)}

	client.Increment(client.Prefix+".strategy_result", tags)
}

func (d *DatadogCustom) Increment(name string, tags []string)  {
	err := d.Client.Incr(name, tags, 1)
	if err != nil {
		log.Fatal(err)
	}
}
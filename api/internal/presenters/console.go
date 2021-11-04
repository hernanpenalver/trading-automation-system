package presenters

import (
	"log"
	"trading-automation-system/api/internal/domain"
)

type ConsolePresenter struct {

}

func (c *ConsolePresenter) Execute(strategyResult *domain.StrategyExecutorResult)  {
	var netBalance float64
	for _, co := range strategyResult.ClosedOperations {
		netBalance += co.GetNetBalance()
	}

	log.Printf("Net Balance: %f", netBalance)
}
package presenters

import (
	"log"
	"trading-automation-system/api/internal/domain"
)

type ConsolePresenter struct {
}

func NewConsolePresenter() *ConsolePresenter {
	return &ConsolePresenter{}
}

func (c *ConsolePresenter) Execute(strategyResult []*domain.StrategyExecutorResult)  {

	log.Printf("RESULTADOS")
	for _, s := range strategyResult {
		if s.GetStrategyBalance() > 0 {
			c.execute(s)
		}
	}
}

func (c *ConsolePresenter) execute(strategyResult *domain.StrategyExecutorResult)  {
	//log.Printf("Potential operations: %d", len(strategyResult.PotentialOperations))
	//log.Printf("Operations: %d", len(strategyResult.OpenedOperations))
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Losers: %d", strategyResult.GetLosersQuantity())
	//log.Printf("Complete Balance: %f", strategyResult.GetCompleteBalance())
	log.Printf("Strategy Balance: %f", strategyResult.GetStrategyBalance())
	log.Printf("Strategy Percent Balance: %f", strategyResult.GetStrategyPercentBalance())
	log.Printf("Strategy Investment Balance: %f", strategyResult.GetInvestmentBalance(100))
}
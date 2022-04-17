package presenters

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type ConsolePresenter struct {
}

func NewConsolePresenter() *ConsolePresenter {
	return &ConsolePresenter{}
}

func (c *ConsolePresenter) Execute(strategyContext *strategies_context.StrategyContext, strategyResult []*domain.StrategyExecutorResult) {
	log.Print("=====================================")
	log.Printf("Executing console_presenter")
	log.Printf("Estrategia: %s", strategyContext.Strategy.GetName())
	log.Printf("From: %s / To: %s", strategyContext.DateFrom.String(), strategyContext.DateTo.String())

	for _, s := range strategyResult {
		c.execute(strategyContext, s)
	}
}

func (c *ConsolePresenter) execute(strategyContext *strategies_context.StrategyContext, strategyResult *domain.StrategyExecutorResult) {
	//log.Printf("Potential operations: %d", len(strategyResult.PotentialOperations))
	//log.Printf("Operations: %d", len(strategyResult.OpenedOperations))
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Losers: %d", strategyResult.GetLosersQuantity())
	//log.Printf("Complete Balance: %f", strategyResult.GetCompleteBalance())
	//log.Printf("Strategy Balance: %f", strategyResult.GetStrategyBalance())
	log.Printf("Strategy Percent Balance: %f", strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount))
	investmentBalance := strategyResult.GetInvestmentBalance(strategyContext.Investment.Amount)
	log.Printf("Strategy Investment Balance: %f", investmentBalance)
	log.Print("=====================================")
}

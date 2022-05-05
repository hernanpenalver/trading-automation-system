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

func (c *ConsolePresenter) Execute(strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *domain.StrategyExecutorResult) {
	c.execute(strategyContext, strategyResult)
}

func (c *ConsolePresenter) execute(strategyContext *strategies_context.StrategyContext, strategyResult *domain.StrategyExecutorResult) {
	//log.Printf("Potential operations: %d", len(strategyResult.PotentialOperations))
	//log.Printf("Operations: %d", len(strategyResult.OpenedOperations))
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Losers: %d", strategyResult.GetLosersQuantity())
	//log.Printf("Complete Balance: %f", strategyResult.GetCompleteBalance())
	//log.Printf("Strategy Balance: %f", strategyResult.GetStrategyBalance())
	log.Printf("Closed operations by close condition: %d", strategyResult.GetQuantityOperationsClosedBy(domain.CloseConditionReason))
	log.Printf("Closed operations by stop loss: %d", strategyResult.GetQuantityOperationsClosedBy(domain.StopLossReason))
	log.Printf("Closed operations by force close: %d", strategyResult.GetQuantityOperationsClosedBy(domain.ForceCloseReason))
	log.Printf("Strategy Percent Balance: %f", strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount))
	investmentBalance := strategyResult.GetInvestmentBalance(strategyContext.Investment.Amount)
	log.Printf("Strategy Investment Balance: %f", investmentBalance)
	log.Print("=====================================")
}

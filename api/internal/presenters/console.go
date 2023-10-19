package presenters

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/strategies_context"
)

type ConsolePresenter struct {
}

func NewConsolePresenter() *ConsolePresenter {
	return &ConsolePresenter{}
}

func (c *ConsolePresenter) Execute(strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *executors.StrategyExecutorResult) {
	c.execute(strategy, strategyContext, strategyResult)
}

func (c *ConsolePresenter) execute(strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *executors.StrategyExecutorResult) {
	//log.Printf("Potential operations: %d", len(strategyResult.PotentialOperations))
	//log.Printf("Operations: %d", len(strategyResult.OpenedOperations))
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Winners: %d", strategyResult.GetWinnersQuantity())
	//log.Printf("Losers: %d", strategyResult.GetLosersQuantity())
	//log.Printf("Complete Balance: %f", strategyResult.GetCompleteBalance())
	//log.Printf("Strategy Balance: %f", strategyResult.GetStrategyBalance())

	//for _, v := range strategyResult.ClosedOperations {
	//	fmt.Printf("EntryPrice: %f\n", v.EntryPrice)
	//	fmt.Printf("EntryDate: %s\n", v.GetEntryDate().String())
	//	fmt.Printf("CloseData.Reason: %s\n", v.CloseData.Reason)
	//	fmt.Printf("CloseData.Price: %f \n", v.CloseData.Price)
	//}

	if strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount) > 1.5 {
		log.Printf("ConsolePresenter.execute - Execute with parameters [%s]: ", strategyResult.Strategy.ToString())
		log.Printf("Closed operations by close condition: %d", strategyResult.GetQuantityOperationsClosedBy(domain.CloseConditionReason))
		log.Printf("Closed operations by stop loss: %d\n", strategyResult.GetQuantityOperationsClosedBy(domain.StopLossReason))
		log.Printf("Closed operations by force close: %d\n", strategyResult.GetQuantityOperationsClosedBy(domain.ForceCloseReason))
		log.Printf("Strategy Percent Balance: %f\n", strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount))
		investmentBalance := strategyResult.GetStrategyBalance(strategyContext.Investment.Amount)
		log.Printf("Strategy Investment Balance: %f\n", investmentBalance)
		//for _, op := range strategyResult.ClosedOperations {
		//	log.Printf(op.ToString())
		//}
		log.Print("=====================================\n")
	}
}

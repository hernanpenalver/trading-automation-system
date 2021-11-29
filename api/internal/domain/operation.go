package domain

import "trading-automation-system/api/internal/utils"

type Operation struct {
	ID         string
	Operation  Action
	Amount     float64
	EntryPrice float64
	StopLoss   float64
	TakeProfit float64
	CloseData  *CloseData
}

type CloseData struct {
	Price  float64
	Reason CloseReason
}

func (o *Operation) IsBuy() bool {
	return o.Operation == BuyAction
}

func (o *Operation) IsSell() bool {
	return o.Operation == SellAction
}

func (o *Operation) HaveToClosePosition(lastCandleStick *CandleStick) bool {
	return o.HaveToTakeProfit(lastCandleStick) || o.HaveToStopLoss(lastCandleStick)
}

func (o *Operation) HaveToTakeProfit(lastCandleStick *CandleStick) bool {
	if o.Operation == BuyAction {
		return o.TakeProfit <= lastCandleStick.Close
	}

	if o.Operation == SellAction {
		return o.TakeProfit >= lastCandleStick.Close
	}

	return false
}

func (o *Operation) HaveToStopLoss(lastCandleStick *CandleStick) bool {
	if o.IsBuy() {
		return o.StopLoss >= lastCandleStick.Close
	}

	if o.IsSell() {
		return o.StopLoss <= lastCandleStick.Close
	}

	return false
}

func (o *Operation) GetNetBalance() float64 {
	if o.CloseData == nil {
		return 0
	}

	if o.IsBuy() {
		// (x*closePrice) - (x*entryPrice)
		return (o.Amount * o.CloseData.Price) - (o.Amount * o.EntryPrice)
	}

	if o.IsSell() {
		// (x*entryPrice) - (x*closePrice)
		return (o.Amount * o.EntryPrice) - (o.Amount * o.CloseData.Price)
	}

	return 0
}

func (o *Operation) GetPercentNetBalance() float64 {
	if o.CloseData == nil {
		return 0
	}

	if o.IsBuy() {
		return utils.GetPercentageOf(o.EntryPrice, (o.Amount * o.CloseData.Price) - (o.Amount * o.EntryPrice))
	}

	if o.IsSell() {
		return utils.GetPercentageOf(o.EntryPrice, (o.Amount * o.EntryPrice) - (o.Amount * o.CloseData.Price))
	}

	return 0
}

type Action string

const (
	BuyAction  Action = "buy"
	SellAction Action = "sell"
	NoAction   Action = "nothing"
)

type CloseReason string

const (
	StopLossReason          CloseReason = "stop_loss"
	TakeProfitReason        CloseReason = "take_profit"
	ForceCloseReason        CloseReason = "force_close"
	StrategyConditionReason CloseReason = "strategy_condition"
)

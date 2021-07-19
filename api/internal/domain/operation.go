package domain

type Operation struct {
	ID         string
	Operation  Action
	Price      float64
	StopLoss   float64
	TakeProfit float64
}

func (o *Operation) IsBuy() bool {
	return o.Operation == BuyAction
}

func (o *Operation) IsSell() bool {
	return o.Operation == SellAction
}

type Action string

const (
	BuyAction  Action = "buy"
	SellAction Action = "sell"
	NoAction   Action = "nothing"
)

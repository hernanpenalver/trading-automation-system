package maths

func PlusPercentage(price, p float64) float64 {
	return price + GetAmountByPercentage(price, p)
}

func MinusPercentage(price, p float64) float64 {
	return price - GetAmountByPercentage(price, p)
}

func GetAmountByPercentage(price, p float64) float64 {
	return (price / 100) * p
}

func GetPercentageOf(price100, priceX float64) float64 {
	return (100 / price100) * priceX
}

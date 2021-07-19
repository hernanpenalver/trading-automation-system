package utils

import "trading-automation-system/api/internal/domain"

func CrossOver(serieA []float64, serieB []float64) bool {
	length := 4
	seriesQty := len(serieA)
	crossover := false

	if serieA[seriesQty-length] < serieB[seriesQty-length] &&
		serieA[seriesQty-(length-1)] < serieB[seriesQty-(length-1)] &&
		serieA[seriesQty-(length-2)] > serieB[seriesQty-(length-2)] &&
		serieA[seriesQty-(length-3)] > serieB[seriesQty-(length-3)] {

		crossover = true
	}

	return crossover
}

func CrossUnder(serieA []float64, serieB []float64) bool {
	length := 4
	seriesQty := len(serieA)
	crossunder := false

	if serieA[seriesQty-length] > serieB[seriesQty-length] &&
		serieA[seriesQty-(length-1)] > serieB[seriesQty-(length-1)] &&
		serieA[seriesQty-(length-2)] < serieB[seriesQty-(length-2)] &&
		serieA[seriesQty-(length-3)] < serieB[seriesQty-(length-3)] {

		crossunder = true
	}

	return crossunder
}

func CollectClosePrices(series []domain.CandleStick) []float64 {
	collection := make([]float64, len(series))

	for i, p := range series {
		collection[i] = p.Close
	}

	return collection
}
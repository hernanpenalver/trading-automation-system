package series

import "trading-automation-system/api/internal/domain"

func CrossOver(serieA []float64, serieB []float64) bool {
	length := 4
	seriesQty := len(serieA)
	crossover := false

	if len(serieA) < length || len(serieB) < length {
		return crossover
	}

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

	if len(serieA) < length || len(serieB) < length {
		return crossunder
	}

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

func RemoveOrderedByID(slice []*domain.Operation, ID string) []*domain.Operation {
	if i, found := getIndexByID(slice, ID); found {
		return append(slice[:i], slice[i+1:]...)
	}

	return slice
}

func getIndexByID(slice []*domain.Operation, ID string) (int, bool) {
	for i, o := range slice {
		if o.ID == ID {
			return i, true
		}
	}
	return 0, false
}
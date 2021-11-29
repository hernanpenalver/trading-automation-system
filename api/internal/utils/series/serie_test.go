package series

import (
	"fmt"
	"testing"
	"trading-automation-system/api/internal/domain"
)
import   "github.com/stretchr/testify/assert"


func TestCrossOver(t *testing.T) {

	seriesA := []float64{1, 2, 5, 6}
	seriesB := []float64{2, 3, 4, 5}

	assert.True(t, CrossOver(seriesA, seriesB))
}

func TestRemoveOrderedByID(t *testing.T) {
	operations := []*domain.Operation{
		{
			ID:         "1",
		},
		{
			ID:         "2",
		},
	}

	operations = RemoveOrderedByID(operations, "1")

	assert.Len(t, operations, 1)
	fmt.Printf("%#v", operations[0])
}
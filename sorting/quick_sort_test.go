package sorting

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/ghazninattarshah/data-structures-algorithms/utils"
)

func TestQuickSort(t *testing.T) {

	randomSlice, sortedSlice := utils.GetTestArrayForSorting(int(rand.Int31n(100)))

	tests := []struct {
		Name     string
		Input    []int
		Expected []int
	}{
		{
			"Nothing",
			[]int{},
			[]int{},
		},
		{
			"PositiveNumbers",
			[]int{5, 7, 12, 3, 8, 1, 4, 20},
			[]int{1, 3, 4, 5, 7, 8, 12, 20},
		},
		{
			"NegativeNumbers",
			[]int{-5, -7, -12, -3, -8, -1, -4, -20},
			[]int{-20, -12, -8, -7, -5, -4, -3, -1},
		},
		{
			"MixedNumbers",
			[]int{5, 7, 12, -3, -8, 0, -4, -20},
			[]int{-20, -8, -4, -3, 0, 5, 7, 12},
		},
		{
			"SizeOfOne",
			[]int{5},
			[]int{5},
		},
		{
			"RandomNumbers",
			randomSlice,
			sortedSlice,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := QuickSort(test.Input)
			success := reflect.DeepEqual(result, test.Expected)
			if !success {
				t.Fatalf("Expected %#v, got %#v", test.Expected, result)
			}
		})
	}
}

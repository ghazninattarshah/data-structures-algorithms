package utils

import (
	"math/rand"
	"sort"
	"time"
)

// GenerateRandomSlice ...
func GenerateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// SortIntSlice ...
func SortIntSlice(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr[:], arr[:])

	sort.Ints(sortedArr[:])
	return sortedArr
}

// GetTestArrayForSorting ..
func GetTestArrayForSorting(size int) ([]int, []int) {

	var arr []int
	arr = GenerateRandomSlice(size)

	return arr, SortIntSlice(arr)
}

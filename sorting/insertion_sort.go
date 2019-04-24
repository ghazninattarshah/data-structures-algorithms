package sorting

// InsertionSort ...
// https://www.tutorialspoint.com/data_structures_algorithms/insertion_sort_algorithm.htm
func InsertionSort(arr []int) []int {

	var position, value int

	size := len(arr)
	for i := 0; i < size; i++ {

		value = arr[i]
		position = i

		for position > 0 && arr[position-1] > value {

			arr[position] = arr[position-1]
			position = position - 1
		}

		arr[position] = value
	}
	return arr
}

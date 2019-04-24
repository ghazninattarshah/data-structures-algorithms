package sorting

// SelectionSort ...
// https://www.tutorialspoint.com/data_structures_algorithms/selection_sort_algorithm.htm
func SelectionSort(arr []int) []int {

	var min int

	size := len(arr)
	for i := 0; i < size-1; i++ {

		min = i
		for j := i + 1; j < size; j++ {

			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			temp := arr[min]
			arr[min] = arr[i]
			arr[i] = temp
		}
	}
	return arr
}

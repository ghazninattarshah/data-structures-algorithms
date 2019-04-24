package sorting

// QuickSort ...
// https://www.tutorialspoint.com/data_structures_algorithms/bubble_sort_algorithm.htm
func QuickSort(arr []int) []int {

	size := len(arr)
	for i := 0; i < size-1; i++ {

		swapped := false
		for j := 0; j < size-1; j++ {

			if arr[j] > arr[j+1] {

				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}

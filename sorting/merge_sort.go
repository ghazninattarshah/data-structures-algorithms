package sorting

// MergeSort ...
// https://www.tutorialspoint.com/data_structures_algorithms/merge_sort_algorithm.htm
func MergeSort(arr []int) []int {

	size := len(arr)
	if size <= 1 {
		return arr
	}

	middle := size / 2

	arr1 := MergeSort(arr[:middle])
	arr2 := MergeSort(arr[middle:])

	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {

	arr := make([]int, 0, len(arr1)+len(arr2))

	for len(arr1) > 0 || len(arr2) > 0 {

		if len(arr1) == 0 {
			return append(arr, arr2...)
		}
		if len(arr2) == 0 {
			return append(arr, arr1...)
		}

		if arr1[0] <= arr2[0] {
			arr = append(arr, arr1[0])
			arr1 = arr1[1:]
		} else {
			arr = append(arr, arr2[0])
			arr2 = arr2[1:]
		}
	}

	return arr
}

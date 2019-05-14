package problems

func OddOccurance(A []int) int {

// write your code in Go 1.4
	odds := make(map[int]bool)
	for i := 0; i < len(A); i++ {

		aval := A[i]
		if exist, ok := odds[aval]; ok {
			if exist {
				delete(odds, aval)
			}
		} else {
			odds[aval] = true
		}	
	}
	
	for k, _ := range odds {
		return k
	}
	return 0
}
package problems

func CyclicRotation(A []int, K int) []int {

	si := len(A)
	if si == 0 {
		return A
	}

	for j := 0; j < K; j++ {

		first := A[si-1]
		for i := si-2; i >= 0; i-- {
			A[i+1] = A[i]				
		}
		A[0] = first
	}
	return A
}
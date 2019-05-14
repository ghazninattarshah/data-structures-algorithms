package problems

import "testing"
import "reflect"

func TestCyclicRotation(t *testing.T) {

	tests := []struct{
		Input []int
		Rotation int
		Result []int
	}{
		{
			[]int{1,2,3,4}, 
			1, 
			[]int{4,1,2,3},
		},
		{
			[]int{1,2,3,4}, 
			4, 
			[]int{1,2,3,4},
		},
		{
			[]int{}, 
			1, 
			[]int{},
		},
	}

	for _, test := range tests {

		r := CyclicRotation(test.Input, test.Rotation)
		if !sliceEqual(r, test.Result) {
			t.Fatalf("expected %v, got %v", test.Result, r)		
		}
	}
}

func sliceEqual(e, a []int) bool {

	for i, v := range e {

		if !reflect.DeepEqual(a[i], v) {
			return false
		}
	}
	return true
}

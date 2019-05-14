package problems

import "testing"
func TestOddOccurances(t *testing.T) {

	tests := []struct{
		Input []int
		Result int
	}{
		{[]int{1,2,3,4,4,3,2,1,5}, 5},
		{[]int{1,2,3,4,4,3,2,1,7,7,9}, 9},
	}

	for _, test := range tests {

		r := OddOccurance(test.Input)
		if r != test.Result {
			t.Fatalf("expected %v, got %v", test.Result, r)		
		}
	}
}
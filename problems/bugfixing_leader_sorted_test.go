package problems

import "testing"

func TestBugfixLeaderSorted(t *testing.T) {

	tests := []struct {
		Input  []int
		Result int
	}{
		{
			[]int{2, 2, 2, 2, 2, 3, 4, 4, 4, 6},
			-1,
		},
		{
			[]int{2, 2, 2, 2, 2, 3, 4, 2, 4, 6},
			1,
		},
	}

	for _, test := range tests {
		r := BugfixLeaderSorted(test.Input)
		if r != test.Result {
			t.Fatalf("expected %v, got %v", test.Result, r)
		}
	}
}

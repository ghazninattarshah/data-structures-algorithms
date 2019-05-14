package problems

import "testing"
func TestBinaryGap(t *testing.T) {

	tests := []struct{
		Input int
		Result int
	}{
		{1041, 5},
		{328, 2},
		{51712, 2},
		{805306373, 25},
	}

	for _, test := range tests {

		r := BinaryGap(test.Input)
		if r != test.Result {
			t.Fatalf("expected %v, got %v", test.Result, r)		
		}
	}
}
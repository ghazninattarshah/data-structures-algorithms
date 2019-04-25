package utils

import (
	"testing"
)

func TestRandomArray(t *testing.T) {

	act, exp := GetTestArrayForSorting(5)
	if len(act) != len(exp) {
		t.Fail()
	}
}

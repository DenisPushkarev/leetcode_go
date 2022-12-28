package leetcode

import "testing"

func TestFindNumbers(t *testing.T) {
	input := []int{12, 345, 2, 6, 7896}
	result := FindNumbers(input)
	if result != 2 {
		t.Logf("FAIL. Expected 2, got %v", result)
		t.Fail()
	}
}

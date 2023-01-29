package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	input, expected := []int{9, 9, 9}, []int{1, 0, 0, 0}
	result := PlusOne(input)
	if !Equal(result, expected) {
		t.Logf("Expected result is %v, but got %v", expected, result)
		t.Fail()
	}

	input, expected = []int{8, 9, 9}, []int{9, 0, 0}
	result = PlusOne(input)
	if !Equal(result, expected) {
		t.Logf("Expected result is %v, but got %v", expected, result)
		t.Fail()
	}

}

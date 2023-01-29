package leetcode

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	input1, input2, expected := []int{1, 2, 2, 2, 4}, []int{2, 2}, []int{2, 2}
	result := Intersect(input1, input2)
	if !Equal(expected, result) {
		t.Logf("For %v and %v expected result is %v, but got %v\n", input1, input2, expected, result)
		t.Fail()
	}

	input1, input2, expected = []int{1, 2, 2, 2, 4}, []int{}, []int{}
	result = Intersect(input1, input2)
	if !Equal(expected, result) {
		t.Logf("For %v and %v expected result is %v, but got %v\n", input1, input2, expected, result)
		t.Fail()
	}

	input1, input2, expected = []int{1, 1, 2, 2, 3}, []int{4, 5}, []int{}
	result = Intersect(input1, input2)
	if !Equal(expected, result) {
		t.Logf("For %v and %v expected result is %v, but got %v\n", input1, input2, expected, result)
		t.Fail()
	}

}

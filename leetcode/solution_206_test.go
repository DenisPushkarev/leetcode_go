package leetcode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var input, expected = []int{1, 2, 3, 4}, []int{4, 3, 2, 1}
	resultList := reverseList(Make(input))
	resultArray := ToArray(resultList)
	if !Equal(expected, resultArray) {
		t.Logf("For input %v result expected to be %v, but got %v", input, expected, resultArray)
		t.Fail()
	}

	input, expected = []int{1}, []int{1}
	resultList = reverseList(Make(input))
	resultArray = ToArray(resultList)
	if !Equal(expected, resultArray) {
		t.Logf("For input %v result expected to be %v, but got %v", input, expected, resultArray)
		t.Fail()
	}

	input, expected = []int{}, []int{}
	resultList = reverseList(Make(input))
	resultArray = ToArray(resultList)
	if !Equal(expected, resultArray) {
		t.Logf("For input %v result expected to be %v, but got %v", input, expected, resultArray)
		t.Fail()
	}

}

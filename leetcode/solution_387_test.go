package leetcode

import "testing"

func TestFirstUniqChar(t *testing.T) {
	var input string
	var result, expectedResult int

	input, expectedResult = "leetcode", 0
	result = firstUniqChar(input)
	if result != expectedResult {
		t.Logf("For %v expected result is %v, but got %v", input, expectedResult, result)
		t.Fail()
	}

	input, expectedResult = "loveleetcode", 2
	result = firstUniqChar(input)
	if result != expectedResult {
		t.Logf("For %v expected result is %v, but got %v", input, expectedResult, result)
		t.Fail()
	}

	input, expectedResult = "aabb", -1
	result = firstUniqChar(input)
	if result != expectedResult {
		t.Logf("For %v expected result is %v, but got %v", input, expectedResult, result)
		t.Fail()
	}
	input, expectedResult = "dddccdbba", 8
	result = firstUniqChar(input)
	if result != expectedResult {
		t.Logf("For %v expected result is %v, but got %v", input, expectedResult, result)
		t.Fail()
	}
}

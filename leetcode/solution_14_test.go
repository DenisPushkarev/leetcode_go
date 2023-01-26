package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected, result := "fl", longestCommonPrefix(strs)
	if result != expected {
		t.Logf("for %v result must be %v, but got %v", strs, expected, result)
		t.Fail()
	}

}

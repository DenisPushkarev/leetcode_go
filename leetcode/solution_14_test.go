package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected, result := "fl", longestCommonPrefix(strs)
	if result != expected {
		t.Logf("for %v result must be %v, but got %v", strs, expected, result)
		t.Fail()
	}

	strs = []string{"a"}
	expected, result = "a", longestCommonPrefix(strs)
	if result != expected {
		t.Logf("for %v result must be %v, but got %v", strs, expected, result)
		t.Fail()
	}

	strs = []string{"", "a"}
	expected, result = "", longestCommonPrefix(strs)
	if result != expected {
		t.Logf("for %v result must be %v, but got %v", strs, expected, result)
		t.Fail()
	}

	strs = []string{"abcd", "abc", "ab"}
	expected, result = "ab", longestCommonPrefix(strs)
	if result != expected {
		t.Logf("for %v result must be %v, but got %v", strs, expected, result)
		t.Fail()
	}

}

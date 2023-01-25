package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	input := [2]string{"sadbutsad", "sad"}
	expected, result := 0, StrStr(input[0], input[1])
	if expected != result {
		t.Errorf("For haystack: '%v', needle: '%v', result must be %v, but got %v", input[0], input[1], expected, result)
		t.Fail()
	}

	input = [2]string{"a", "a"}
	expected, result = 0, StrStr(input[0], input[1])
	if expected != result {
		t.Errorf("For haystack: '%v', needle: '%v', result must be %v, but got %v", input[0], input[1], expected, result)
		t.Fail()
	}

	input = [2]string{"leetcode", "leeto"}
	expected, result = -1, StrStr(input[0], input[1])
	if expected != result {
		t.Errorf("For haystack: '%v', needle: '%v', result must be %v, but got %v", input[0], input[1], expected, result)
		t.Fail()
	}

	input = [2]string{"vavagno", "vag"}
	expected, result = 2, StrStr(input[0], input[1])
	if expected != result {
		t.Errorf("For haystack: '%v', needle: '%v', result must be %v, but got %v", input[0], input[1], expected, result)
		t.Fail()
	}

	input = [2]string{"mississippi", "issip"}
	expected, result = 4, StrStr(input[0], input[1])
	if expected != result {
		t.Errorf("For haystack: '%v', needle: '%v', result must be %v, but got %v", input[0], input[1], expected, result)
		t.Fail()
	}
}

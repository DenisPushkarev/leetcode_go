package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := RemoveDuplicates(input)
	if result != 5 {
		t.Logf("FAIL. Expected 5, got %v", result)
		t.Fail()
	}

	result, expected := RemoveDuplicates([]int{0, 0}), 1
	if result != expected {
		t.Logf("Expected %v, got %v", expected, result)
		t.Fail()
	}

	result, expected = RemoveDuplicates([]int{0, 0, 1, 2, 2, 2, 2, 2}), 3
	if result != expected {
		t.Logf("Expected %v, got %v", expected, result)
		t.Fail()
	}

}

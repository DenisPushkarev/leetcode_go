package leetcode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	expected := "[1,1,2,3,4,4]"
	result := MergeTwoLists("[1,2,4]", "[1,3,4]")
	if result != expected {
		t.Logf("Expected %v, got %v", expected, result)
		t.Fail()
	}
}

package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildList(t *testing.T) {
	var node *ListNode
	node = BuildList("[1,2,3,4]")
	var expectedResult = [4]int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		if node == nil {
			t.Logf("%v'th node is nil", i)
			t.Fail()
		} else {
			fmt.Printf("checked node : %v\n", node.Val)
		}
		if node.Val != expectedResult[i] {
			t.Logf("Expected %v, got %v", expectedResult[i], node.Val)
			t.Fail()
		}
		node = node.Next
	}
}
func TestToString(t *testing.T) {
	expectedResult := "[1,2,3,4]"
	var node *ListNode = BuildList(expectedResult)
	var result string = ToString(node)
	t.Logf("comparing %v and %v\n", expectedResult, result)
	if result != expectedResult {
		t.Logf("FAIL. Expected %v, but got %v", expectedResult, result)
		t.Fail()
	}
}

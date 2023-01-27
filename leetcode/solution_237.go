package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	prev := node
	for node.Next != nil {
		next := node.Next
		node.Val = next.Val
		prev = node
		node = node.Next
	}
	prev.Next = nil
}

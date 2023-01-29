package leetcode

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}

func deleteNode2(node *ListNode) {
	prev := node
	for node.Next != nil {
		next := node.Next
		node.Val = next.Val
		prev = node
		node = node.Next
	}
	prev.Next = nil
}

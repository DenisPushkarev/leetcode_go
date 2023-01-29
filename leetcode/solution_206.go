package leetcode

func reverseListIteratively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur *ListNode = head.Next
	var tmp *ListNode = cur.Next
	var prev *ListNode = head
	cur.Next = prev
	head.Next = nil
	for tmp != nil {
		prev = cur
		cur = tmp
		tmp = tmp.Next
		cur.Next = prev

	}
	return cur
}

func reverseList(head *ListNode) *ListNode {
	// return reverseListIteratively(head)
	return reverseListRecursively(head)
}

func reverseListRecursively(head *ListNode) *ListNode {
	return helper(nil, head)
}

func helper(n1, n2 *ListNode) *ListNode {
	if n2 != nil {
		r := helper(n2, n2.Next)
		n2.Next = n1
		return r
	}
	return n1
}

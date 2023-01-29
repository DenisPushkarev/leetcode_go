package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var curNode *ListNode = head
	var nth *ListNode = head
	for curNode.Next != nil {
		if n == 0 {
			nth = nth.Next
		} else {
			n--
		}
		curNode = curNode.Next
	}
	if n == 1 {
		return head.Next
	}
	if nth.Next == nil {
		return nil
	} else {
		nth.Next = nth.Next.Next
		return head
	}
}

func removeNthUsingStack(head *ListNode, n *int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNthUsingStack(head.Next, n)
	*n = *n - 1
	if *n == 0 {
		return head.Next
	} else {
		return head
	}
}

package leetcode

func isPalindromeList(head *ListNode) bool {
	slowPtr, fastPtr := head, head
	if fastPtr.Next == nil {
		return true
	}
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	fastPtr = nil
	var tmp *ListNode
	for slowPtr != nil {
		tmp = slowPtr.Next
		slowPtr.Next = fastPtr
		fastPtr = slowPtr
		slowPtr = tmp
	}
	fastPtr, slowPtr = head, fastPtr
	for slowPtr != nil {
		if slowPtr.Val != fastPtr.Val {
			return false
		} else {
			slowPtr = slowPtr.Next
			fastPtr = fastPtr.Next
		}
	}
	return true
}

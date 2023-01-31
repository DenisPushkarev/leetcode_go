package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowPtr, fastPtr := head, head.Next
	for slowPtr != nil && fastPtr != nil && fastPtr.Next != nil {
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

	}
	return false
}

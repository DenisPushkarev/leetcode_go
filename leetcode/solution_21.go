package leetcode

/**
 * Definition for singly-linked
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeTwoLists(str1 string, str2 string) string {
	list1 := BuildList(str1)
	list2 := BuildList(str2)
	mergedList := mergeTwoLists(list1, list2)
	return ToString(mergedList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head ListNode = ListNode{}
	var prev *ListNode = &head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return head.Next
}

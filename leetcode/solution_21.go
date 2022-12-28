package leetcode

import (
	"solutions/main/leetcode/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *list.ListNode
 * }
 */

func MergeTwoLists(str1 string, str2 string) string {
	list1 := list.BuildList(str1)
	list2 := list.BuildList(str2)
	mergedList := mergeTwoLists(list1, list2)
	return list.ToString(mergedList)
}

func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	var head list.ListNode = list.ListNode{}
	var prev *list.ListNode = &head
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

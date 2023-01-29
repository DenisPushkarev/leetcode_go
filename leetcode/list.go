package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToArray(head *ListNode) []int {
	var a []int = make([]int, 0)
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	return a
}

func Make(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	} else {
		return &ListNode{
			Val:  values[0],
			Next: Make(values[1:]),
		}
	}
}

func BuildList(str string) *ListNode {
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	values := strings.Split(str, ",")
	var rootNode ListNode
	var prevNode *ListNode = &rootNode
	for i := 0; i < len(values); i++ {
		fmt.Printf("adding, %v\n", values[i])
		value, _ := strconv.Atoi(values[i])
		prevNode.Next = &ListNode{Val: value}
		prevNode = prevNode.Next
	}
	return rootNode.Next
}

func PrintList(list *ListNode) {
	node := list
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

func ToString(list *ListNode) string {
	str := ""
	for list != nil {
		fmt.Printf("val: %v", list.Val)
		str += strconv.Itoa(list.Val) + ","
		list = list.Next
	}
	str = "[" + strings.TrimSuffix(str, ",") + "]"
	return str
}

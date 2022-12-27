package list

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
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

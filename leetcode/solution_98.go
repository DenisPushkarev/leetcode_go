package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTR(root, math.MinInt, math.MaxInt)
}

func isValidBSTR(node *TreeNode, l, r int) bool {
	if node == nil {
		return true
	}
	return node.Val > l &&
		node.Val < r &&
		isValidBSTR(node.Left, l, node.Val) &&
		isValidBSTR(node.Right, node.Val, r)
}

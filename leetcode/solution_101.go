package leetcode

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	return node1 != nil && node2 != nil && node1.Val == node2.Val &&
		isSymmetricHelper(node1.Left, node2.Right) &&
		isSymmetricHelper(node1.Right, node2.Left)
}

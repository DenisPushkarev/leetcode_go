package leetcode

func maxDepth(root *TreeNode) int {
	return maxDepthTreeDFS(root, 1)
}

func maxDepthTreeDFS(node *TreeNode, depth int) int {
	if node == nil {
		return depth - 1
	} else {
		return maxInt(maxDepthTreeDFS(node.Left, depth+1), maxDepthTreeDFS(node.Right, depth+1))
	}
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

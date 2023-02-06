package leetcode

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level > len(res) {
			res = append(res, []int{node.Val})
		} else {
			res[level-1] = append(res[level-1], node.Val)
		}

		helper(node.Left, level+1)
		helper(node.Right, level+1)
	}

	helper(root, 1)

	return res
}

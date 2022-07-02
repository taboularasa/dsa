package problems

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func IsBalanced(tree *Node) bool {
	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.left)
		rightHeight := dfs(node.right)
		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}
		if diff(leftHeight, rightHeight) > 1 {
			return -1
		}

		return max(leftHeight, rightHeight) + 1
	}
	return dfs(tree) != -1
}

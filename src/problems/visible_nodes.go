package problems

type Node struct {
	val   int
	left  *Node
	right *Node
}

func VisibleTreeNode(root *Node) int {
	var dfs func(node *Node, max int) int
	dfs = func(node *Node, max_val int) int {
		if node == nil {
			return 0
		}
		visible := 0
		if node.val >= max_val {
			visible++
			max_val = node.val
		}
		return visible + dfs(node.left, max_val) + dfs(node.right, max_val)
	}

	return dfs(root, root.val)
}

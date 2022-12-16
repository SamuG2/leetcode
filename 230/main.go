package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, max_depth int) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, inorder(root.Left, max_depth)...)
	if len(res) >= max_depth {
		return res
	}
	res = append(res, root.Val)
	if len(res) >= max_depth {
		return res
	}
	res = append(res, inorder(root.Right, max_depth)...)
	return res
}

func kthSmallest(root *TreeNode, k int) int {
	nodes := inorder(root, k)
	return nodes[k-1]
}

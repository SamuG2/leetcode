package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if lower < node.Val && upper > node.Val {
		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
	}
	return false
}

func main() {
	//root := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root))

}

package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func minNode(left, right *TreeNode) *TreeNode {
// 	if left.Val > right.Val {
// 		return right
// 	}
// 	return left
// }
// func maxNode(left, right *TreeNode) *TreeNode {
// 	if left.Val < right.Val {
// 		return right
// 	}
// 	return left
// }

// func findSecondMinimumValue(root *TreeNode) int {
// 	if root.Left != nil {
// 		min_node := minNode(root.Left, root.Right)
// 		max_node := maxNode(root.Left, root.Right)
// 		if min_node.Left == nil {
// 			if max_node.Val == root.Val {
// 				return -1
// 			}
// 			return max_node.Val
// 		}
// 		return findSecondMinimumValue(min_node)
// 	}
// 	return -1
// }

func dfs(root *TreeNode, vals []int) []int {
	if root.Left != nil {
		vals = append(vals, dfs(root.Left, vals)...)
		vals = append(vals, dfs(root.Right, vals)...)
	}
	return append(vals, root.Val)
}

func findSecondMinimumValue(root *TreeNode) int {
	vals := dfs(root, []int{})
	sort.Ints(vals)
	smallest, second_smallest := vals[len(vals)-1], vals[len(vals)-1]
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] < smallest {
			second_smallest = smallest
			smallest = vals[i]
		}
	}
	if second_smallest == smallest {
		return -1
	}
	return second_smallest
}

func main() {

}

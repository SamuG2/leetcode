package main

/*Diferentes formas de explorar un arbol binario*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Enorden: Izquierda -> raíz -> derecha
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

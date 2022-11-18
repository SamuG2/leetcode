package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkNils(root *TreeNode, max_depth, cur_depth int, cur_nils *int) bool {

	if root == nil { //comprobamos si el nodo existe, si no aumentamos el numero de nils y devolvemos false
		*cur_nils++
		return false
	}
	if cur_depth == max_depth { //si existe y estamos en el último nivel es que encontramos todos los nils, devolvemos true
		return true
	}
	found := checkNils(root.Right, max_depth, cur_depth+1, cur_nils) //comprobamos siempre primero el de la derecha
	if found {
		return true
	}
	found = checkNils(root.Left, max_depth, cur_depth+1, cur_nils) // si el de la derecha no está, comprobamos el de la izquierda

	return found
}

func countNodes(root *TreeNode) int {
	if root == nil { // si root es nill ya sabemos que  la respuesta es 0
		return 0
	}
	depth := 0
	cur := root
	for cur != nil { //calculamos la profundidad máxima
		cur = cur.Left
		depth++
	}

	var nils int
	checkNils(root, depth-1, 0, &nils) //una vez tenemos la profundidad máxima miramos cuantos faltan en el último nivel

	return int(math.Pow(2, float64(depth))) - 1 - nils //la respuesta  es 2**profundiad máxima -1 - los que faltan del último nivel
}

func main() {

}

package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

//********** Método 1: Idea original, bastante lenta

// func findLowest(lists []*ListNode) int {

// 	var idx int
// 	min := math.MaxInt32
// 	for i, node := range lists {
// 		if node != nil {
// 			if node.Val < min {
// 				min = node.Val
// 				idx = i
// 			}
// 		}
// 	}
// 	lists[idx] = lists[idx].Next
// 	return min
// }

// func checkNils(lists []*ListNode) bool {
// 	for _, v := range lists {
// 		if v != nil {
// 			return true
// 		}
// 	}
// 	return false
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if !checkNils(lists) {
// 		return nil
// 	}
// 	head := &ListNode{}
// 	cur := head
// 	for checkNils(lists) {
// 		cur.Val = findLowest(lists)
// 		if checkNils(lists) {
// 			cur.Next = &ListNode{}
// 			cur = cur.Next
// 		}
// 	}
// 	return head
// }

//********** Método 2: respuesta mejor rendimiento

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	} else if len(lists) == 1 {
// 		return lists[0]
// 	} else if len(lists) == 2 {
// 		return mergeLists(lists[0], lists[1])
// 	}
// 	headA := mergeKLists(lists[:len(lists)/2])
// 	headB := mergeKLists(lists[len(lists)/2:])
// 	return mergeLists(headA, headB)

// }

// func mergeLists(l1, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	head := &ListNode{}
// 	cur := head
// 	for l1 != nil && l2 != nil {
// 		if l2.Val > l1.Val {
// 			cur.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			cur.Next = l2
// 			l2 = l2.Next
// 		}
// 		cur = cur.Next
// 	}
// 	if l1 != nil {
// 		cur.Next = l1
// 	}
// 	if l2 != nil {
// 		cur.Next = l2
// 	}
// 	return head.Next
// }

// ********** Método 3: Probar con sort.Ints-> casi igual de rápido que método 2, pero utiliza bastante mas memoria

func mergeKLists(lists []*ListNode) *ListNode {
	var values []int
	for _, node := range lists {
		for node != nil {
			values = append(values, node.Val)
			node = node.Next
		}
	}
	sort.Ints(values)
	head := createNodeList(values)
	return head
}

func createNodeList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head
	l := len(values) - 1
	for i, v := range values {
		cur.Val = v
		if i != l {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

func main() {

}

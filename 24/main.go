package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNodes(vals []int) *ListNode {
	head := new(ListNode)
	cur := head
	for i, v := range vals {
		cur.Val = v
		if i < len(vals)-1 {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		aux := &ListNode{Val: cur.Val, Next: cur.Next.Next}
		*cur = *cur.Next
		cur.Next = aux
		cur = aux.Next
	}
	return head
}

func main() {
	nodes := createNodes([]int{1, 2, 3, 4})

	nodes = swapPairs(nodes)
	for nodes != nil {
		fmt.Println(nodes)
		nodes = nodes.Next
	}
}

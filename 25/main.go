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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i := 0; i <= len(values)-k; i += k {
		for t, j := i, i+k-1; j > t; t, j = t+1, j-1 {
			values[t], values[j] = values[j], values[t]
		}
	}

	return createNodes(values)
}

func main() {
	nodes := createNodes([]int{1, 2, 3, 4, 5})
	k := 2
	nodes = reverseKGroup(nodes, k)

	for nodes != nil {
		fmt.Println(nodes)
		nodes = nodes.Next
	}

}

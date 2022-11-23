package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListNode(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for i, num := range nums {
		cur.Val = num
		if i != len(nums)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	var list_nodes []*ListNode
	for node != nil {
		list_nodes = append(list_nodes, node)
		node = node.Next
	}
	n = len(list_nodes) - n
	if n == 0 {
		return head.Next

	} else if n == len(list_nodes)-1 {
		list_nodes[n-1].Next = nil
		return head
	}
	list_nodes[n-1].Next = list_nodes[n+1]
	return head

}

func main() {
	head := createListNode([]int{1, 2})
	head = removeNthFromEnd(head, 1)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

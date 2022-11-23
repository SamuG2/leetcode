package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNodeList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head
	for i, v := range values {
		cur.Val = v
		if i != len(values)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			cur.Val = list1.Val
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				cur.Val = list1.Val
				list1 = list1.Next

			} else if list2.Val < list1.Val {
				cur.Val = list2.Val
				list2 = list2.Next

			} else if list1.Val == list2.Val {
				cur.Val = list1.Val
				cur.Next = &ListNode{}
				cur = cur.Next
				list1 = list1.Next
				cur.Val = list2.Val
				list2 = list2.Next
			}
		}

		if list1 != nil || list2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return head
}

func main() {
	list1 := createNodeList([]int{1})
	list2 := createNodeList([]int{2})

	fmt.Println(mergeTwoLists(list1, list2))
}

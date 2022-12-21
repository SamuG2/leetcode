package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var vals_l1, vals_l2 []int
	for l1 != nil {
		vals_l1 = append(vals_l1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		vals_l2 = append(vals_l2, l2.Val)
		l2 = l2.Next
	}
	if len(vals_l1) > len(vals_l2) {
		vals_l2 = append(make([]int, len(vals_l1)-len(vals_l2)), vals_l2...)
	} else if len(vals_l1) < len(vals_l2) {
		vals_l1 = append(make([]int, len(vals_l2)-len(vals_l1)), vals_l1...)
	}
	var extra int
	for i := len(vals_l1) - 1; i >= 0; i-- {
		vals_l1[i] = vals_l1[i] + vals_l2[i] + extra
		if vals_l1[i] > 9 {
			extra = 1
			vals_l1[i] = vals_l1[i] - 10
		} else {
			extra = 0
		}
	}
	if extra > 0 {
		vals_l1 = append([]int{extra}, vals_l1...)
	}
	head := &ListNode{}
	cur := head
	for i, v := range vals_l1 {
		cur.Val = v
		if i < len(vals_l1)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

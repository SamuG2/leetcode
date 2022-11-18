package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	extra := 0

	for {
		sum := extra
		extra = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			extra = 1
			sum -= 10
		}
		res.Val = sum
		if l1 != nil || l2 != nil || extra != 0 {
			res.Next = &ListNode{}
			res = res.Next
		} else {
			break
		}
	}
	res = nil

	return head
}

func main() {

}

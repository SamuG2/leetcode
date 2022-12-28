package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left_operand, right_operand := make([]int, l), make([]int, l)
	left_operand[0], right_operand[l-1] = 1, 1

	for i := 1; i < l; i++ {
		left_operand[i] = left_operand[i-1] * nums[i-1]
		right_operand[l-i-1] = right_operand[l-i] * nums[l-i]
	}
	for i := range nums {
		nums[i] = left_operand[i] * right_operand[i]
	}
	return nums
}

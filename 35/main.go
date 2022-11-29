package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		idx := left + (right-left)/2
		if nums[idx] > target {
			right = idx
		} else if nums[idx] < target {
			left = idx + 1
		} else {
			return idx
		}
	}
	return left
}

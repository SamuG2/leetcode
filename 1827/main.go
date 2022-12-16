package main

func minOperations(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			res += 1 + (nums[i-1] - nums[i])
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}

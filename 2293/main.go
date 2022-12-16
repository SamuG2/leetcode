package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	new_nums := make([]int, len(nums)/2)
	for i := range new_nums {
		if i%2 == 0 {
			new_nums[i] = min(nums[2*i], nums[2*i+1])
		} else {
			new_nums[i] = max(nums[2*i], nums[2*i+1])
		}
	}
	return minMaxGame(new_nums)
}

func main() {
	nums := []int{1, 3, 5, 2, 4, 8, 2, 2}
	minMaxGame(nums)
}

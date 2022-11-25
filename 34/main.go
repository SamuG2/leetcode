package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	idx := -1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			idx = mid
			break
		}
	}
	if idx == -1 {
		return []int{-1, -1}
	}
	i, j := idx, idx
	for nums[i] == target || nums[j] == target {
		if nums[i] == target && i > 0 {
			i--
		}
		if nums[j] == target && j < len(nums)-1 {
			j++
		}
	}
	return []int{i + 1, j - 1}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

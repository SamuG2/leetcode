package main

func search(nums []int, target int) int {
	n := len(nums)

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot := left
	right = pivot - 1 + n

	for left <= right {
		mid := left + (right-left)/2
		v := nums[mid%n]
		if v > target {
			right = mid - 1
		} else if v < target {
			left = mid + 1
		} else {
			return mid % n
		}
	}
	return -1

}

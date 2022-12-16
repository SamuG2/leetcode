package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func constrainedSubsetSum(nums []int, k int) int {

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			if nums[i]-nums[i+1] > k {
				if nums[i] < nums[i+1]*-1 {
					nums = append(nums[:i+1], nums[i+2:]...)
				} else {
					nums = append(nums[:i], nums[i+1:]...)
				}
				i--
			}
		}
	}
	var res int
	fmt.Println(nums)
	for _, v := range nums {
		res += v
	}
	return res
}

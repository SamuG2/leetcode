package main

import (
	"fmt"
	"sort"
)

// func nextPermutation(nums []int) []int {

// 	for i := len(nums) - 1; i > 0; i-- {
// 		for j := i - 1; j >= 0; j-- {
// 			if nums[i] > nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 				return nums
// 			} else {
// 				i = j
// 			}
// 		}
// 	}
// 	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	return nums
// }

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			min_idx := i
			for j := i; j < len(nums); j++ {
				if nums[j] <= nums[min_idx] && nums[j] > nums[i-1] {
					min_idx = j
					if nums[min_idx] == nums[i-1]+1 {
						break
					}
				}
			}
			nums[i-1], nums[min_idx] = nums[min_idx], nums[i-1]
			aux := nums[i:]
			sort.Ints(aux)
			for j := 0; j+i < len(nums)-1; j++ {
				nums[j+i] = aux[j]
			}
			return
		}
	}
	sort.Ints(nums)

	return
}

func main() {
	nums := []int{2, 1, 2, 2, 2, 2, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

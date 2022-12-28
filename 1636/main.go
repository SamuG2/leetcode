package main

import (
	"sort"
)

func frequencySort(nums []int) []int {
	frequencyMap := make(map[int]int)
	for _, v := range nums {
		frequencyMap[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if frequencyMap[nums[i]] == frequencyMap[nums[j]] {
			return nums[i] > nums[j]
		} else {
			return frequencyMap[nums[i]] < frequencyMap[nums[j]]
		}
	})
	return nums
}

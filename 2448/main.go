package main

import (
	"math"
	"sort"
)

// func changeButOne(nums, cost []int, idx int) int64 {
// var res int64
// for i := range nums {
// 	if i != idx {
// 		res += int64(cost[i]) * int64(math.Abs(float64(nums[idx])-float64(nums[i])))
// 	}
// }
// return res
// }

func minCost(nums []int, cost []int) int64 {
	arr := make([][2]int, len(nums))
	var total, count, target int
	var res int64
	for i := range nums {
		arr[i][0], arr[i][1] = nums[i], cost[i]
		total += cost[i]
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	for i := range arr {
		count += arr[i][1]
		if count > total/2 {
			target = arr[i][0]
			break
		}
	}
	for i := range arr {
		res += int64(arr[i][1]) * int64(math.Abs(float64(arr[i][0]-target)))
	}
	return res
}

func main() {
	nums := []int{1, 3, 5, 2}
	cost := []int{2, 3, 1, 14}
	minCost(nums, cost)

}

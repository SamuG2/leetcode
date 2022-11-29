package main

import "fmt"

// func jump(nums []int) int {
// 	return dfs(nums, 0, 0, len(nums)-1)
// }

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

// func dfs(nums []int, idx int, jumps int, min_jumps int) int {

// 	if jumps >= min_jumps {
// 		return min_jumps
// 	}
// 	if idx >= len(nums)-1 {
// 		return jumps
// 	}
// 	for i := nums[idx]; i > 0; i-- {
// 		min_jumps = min(min_jumps, dfs(nums, i+idx, jumps+1, min_jumps))
// 	}
// 	return min_jumps
//}

func jump(nums []int) int {
	curJump, farthestJump, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthestJump {
			farthestJump = i + nums[i]
		}

		if i == curJump {
			jumps, curJump = jumps+1, farthestJump
			if curJump >= len(nums)-1 {
				return jumps
			}
		}
	}
	return 0
}

func main() {
	nums := []int{5, 8, 1, 8, 9, 8, 7, 1, 7, 5, 8, 6, 5, 4, 7, 3, 9, 9, 0, 6, 6, 3, 4, 8, 0, 5, 8, 9, 5, 3, 7, 2, 1, 8, 2, 3, 8, 9, 4, 7, 6, 2, 5, 2, 8, 2, 7, 9, 3, 7, 6, 9, 2, 0, 8, 2, 7, 8, 4, 4, 1, 1, 6, 4, 1, 0, 7, 2, 0, 3, 9, 8, 7, 7, 0, 6, 9, 9, 7, 3, 6, 3, 4, 8, 6, 4, 3, 3, 2, 7, 8, 5, 8, 6, 0}
	fmt.Println(jump(nums))
}

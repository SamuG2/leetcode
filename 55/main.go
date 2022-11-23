package main

import "fmt"

//*** Demasiado lento para muestras muy grandes
// func dfs(idx int, nums []int, checked *[]int) bool {
// 	if idx == len(nums)-1 {
// 		return true
// 	}
// 	for _, v := range *checked {
// 		if v == idx {
// 			return false
// 		}
// 	}
// 	for i := nums[idx]; i > 0; i-- {
// 		if i+idx < len(nums) {
// 			if dfs(idx+i, nums, checked) {
// 				return true
// 			}
// 		}
// 	}
// 	*checked = append(*checked, idx)
// 	return false

// }

// func canJump(nums []int) bool {
// 	var checked []int
// 	return dfs(0, nums, &checked)
// }

// *** y yendo desde atrás?

func checkJump(nums []int, idx int, checked *[]int) bool {
	if idx == 0 {
		return true
	}
	for _, v := range *checked {
		if v == idx {
			return false
		}
	}
	for i := idx - 1; i >= 0; i-- {
		if i+nums[i] >= idx {
			res := checkJump(nums, i, checked)
			if res {
				return true
			}
		}
	}
	*checked = append(*checked, idx)
	return false
}

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	var checked []int
	return checkJump(nums, len(nums)-1, &checked)
}

func main() {
	nums := []int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6}
	fmt.Println(canJump(nums))
}

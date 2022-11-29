package main

import (
	"fmt"
	"time"
)

func permuteRec(cur, left []int, res *[][]int) {
	if 0 == len(left) {
		*res = append(*res, cur)
		return
	}
	for idx, l := range left {
		permuteRec(append(cur, l), append(append([]int{}, left[:idx]...), left[idx+1:]...), res)
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res

}

func main() {
	nums := []int{1, 2, 3}
	// defer exeTime("main")()
	start := time.Now()
	aux := permute(nums)
	//permute(nums)
	time := time.Since(start)
	fmt.Println(len(aux))
	fmt.Println(time)
}

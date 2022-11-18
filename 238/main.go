package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}
	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix = prefix * nums[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
	}
	return res

}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

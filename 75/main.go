package main

import "fmt"

func sortColors(nums []int) {
	ceroes, ones, twos := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			ceroes++
		} else if v == 1 {
			ones++
		} else {
			twos++
		}
	}
	for i := 0; i < ceroes; i++ {
		nums[i] = 0
	}
	for i := ceroes; i < ones+ceroes; i++ {
		nums[i] = 1
	}
	for i := ceroes + ones; i < twos+ceroes+ones; i++ {
		nums[i] = 2
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)
	fmt.Println(nums)
}

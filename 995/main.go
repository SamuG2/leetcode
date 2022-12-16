package main

func minKBitFlips(nums []int, k int) int {
	bools := make([]bool, len(nums))
	for i, v := range nums {
		bools[i] = 1 == v
	}
	transformations, i := 0, 0
	for {
		found := false
		for j, v := range bools[:i] {
			if !v {
				i = j
				found = true
				break
			}
		}
		if !found {
			return transformations
		}
		if i+k > len(nums) {
			return -1
		}
		for j := i; j < i+k; j++ {
			bools[j] = !bools[j]
		}
		transformations++
	}
}

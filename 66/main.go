package main

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

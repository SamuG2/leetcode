package main

func mySqrt(x int) int {

	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		m := mid * mid
		if m < x {
			left = mid + 1
		} else if m > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left - 1
}

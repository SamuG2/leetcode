package main

func addDigits(num int) int {
	if num != 0 {
		return (num-1)%9 + 1
	}
	return 0
}

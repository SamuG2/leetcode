package main

import "fmt"

var number int

func guess(n int) int {
	if n == number {
		return 0
	} else if n > number {
		return -1
	} else {
		return 1
	}
}

func guessNumberRec(max_number, min_number, n int) int {
	aux := guess(n)
	if aux == 0 {
		return n
	}
	if aux == -1 {
		max_number = n
		return guessNumberRec(max_number, min_number, min_number+(max_number-min_number)/2)
	} else {
		min_number = n
		return guessNumberRec(max_number, min_number, min_number+(max_number-min_number)/2)
	}
}

func guessNumber(n int) int {
	return guessNumberRec(n, 0, n)
}

func main() {
	number = 5
	fmt.Println(guessNumber(9))
}

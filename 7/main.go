package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	neg := 1
	x_s := []rune(strconv.Itoa(x))

	if x < 0 {
		neg = -1
		x_s = x_s[1:]
	}
	l := len(x_s)
	for i := 0; i < l/2; i++ {
		x_s[i], x_s[l-i-1] = x_s[l-i-1], x_s[i]
	}
	res, _ := strconv.ParseInt(string(x_s), 10, 64)
	if res > math.MaxInt32 {
		return 0
	}
	return int(res) * neg

}

func main() {

}

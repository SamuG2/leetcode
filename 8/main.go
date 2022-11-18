package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	neg := 1
	s_r := []rune(s)
	for i := 0; i < len(s_r); {
		if s_r[i] >= '0' && s_r[i] <= '9' {
			break
		} else if s_r[i] == '+' {
			s_r = s_r[1:]
			break
		} else if s_r[i] == '-' {
			s_r = s_r[1:]
			neg = -1
			break
		} else if s_r[i] == ' ' {
			s_r = s_r[1:]
		} else {
			return 0
		}
	}
	var res_s []int
	for _, v := range s_r {
		if v >= '0' && v <= '9' {
			res_s = append(res_s, int(v-'0'))
		} else {
			break
		}
	}
	var res int
	for i, v := range res_s {

		if neg == 1 && float64(res)+float64(v)*(math.Pow(10, float64(len(res_s)-i-1))) > math.MaxInt32 {
			return math.MaxInt32
		} else if neg == -1 && (float64(res)+float64(v)*(math.Pow(10, float64(len(res_s)-i-1))))*float64(-1) < math.MinInt32 {
			return math.MinInt32
		}
		res += v * int(math.Pow(10, float64(len(res_s)-i-1)))
	}
	res = res * neg

	return res
}

func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}

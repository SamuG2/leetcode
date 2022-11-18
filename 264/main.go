package main

import (
	"fmt"
	"math"
)

// func isUgly(n int) bool {
// 	for n%5 == 0 {
// 		n = n / 5
// 	}
// 	for n%3 == 0 {
// 		n = n / 3
// 	}
// 	for n%2 == 0 {
// 		n = n / 2
// 	}

// 	return n == 1
// }

// func nthUglyNumber(n int) int {
// 	var bad_primes []int
// 	res := 0
// 	i := 1
// 	for ; res < n; i++ {
// 		candidate := true
// 		for _, p := range bad_primes {
// 			if i%p == 0 {
// 				candidate = false
// 				break
// 			}
// 		}
// 		if candidate {
// 			if isUgly(i) {
// 				res++
// 			} else {
// 				bad_primes = append(bad_primes, i)
// 			}
// 		}

// 	}
// 	return i - 1
// }

func nthUglyNumber(n int) int {
	ugly_nums := make([]int, n)
	ugly_nums[0] = 1
	i2, i3, i5, f2, f3, f5 := 0, 0, 0, 2, 3, 5

	for i := 1; i < n; i++ {
		min := int(math.Min(math.Min(float64(f2), float64(f3)), float64(f5)))
		if ugly_nums[i-1] != min {
			ugly_nums[i] = min
		} else {
			i--
		}
		if f2 == min {
			i2++
			f2 = 2 * ugly_nums[i2]

		} else if f3 == min {
			i3++
			f3 = 3 * ugly_nums[i3]

		} else if f5 == min {
			i5++
			f5 = 5 * ugly_nums[i5]

		}
	}

	return ugly_nums[n-1]

}

func main() {
	fmt.Println(nthUglyNumber(10))
}

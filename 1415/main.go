package main

import "math"

func getHappyString(n int, k int) string {
	posibles := 3 * int(math.Pow(2, float64(n-1)))
	if n == 0 || k > posibles {
		return ""
	}
	var res []rune
	var lower, upper int
	if k >= 1 && k <= posibles/3 {
		res = append(res, 'a')
		lower = 1
		upper = posibles / 3
	} else if k > posibles/3 && k <= posibles-posibles/3 {
		res = append(res, 'b')
		lower = posibles/3 + 1
		upper = posibles - posibles/3
	} else if k > posibles-posibles/3 {
		res = append(res, 'c')
		lower = posibles - posibles/3 + 1
		upper = posibles
	}
	for i := 1; i < n; i++ {
		half := lower + (upper-lower)/2
		if k >= lower && k <= half {
			res = append(res, upperHalf(res[len(res)-1]))
			upper = half
		} else {
			res = append(res, lowerHalf(res[len(res)-1]))
			lower = half + 1
		}
	}
	return string(res)
}

func upperHalf(c rune) rune {
	if c == 'a' {
		return 'b'
	}
	return 'a'
}

func lowerHalf(c rune) rune {
	if c == 'c' {
		return 'b'
	}
	return 'c'
}

func main() {
	n, k := 3, 9

	getHappyString(n, k)

}

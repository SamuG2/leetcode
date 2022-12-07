package main

func myPow(x float64, n int) float64 {
	var n_a int
	if n < 0 {
		n_a = n * -1
	} else if n > 0 {
		n_a = n
	} else {
		return 1
	}
	res := 1.0
	for n_a > 0 {
		if n_a%2 == 0 {
			n_a = n_a / 2
			x = x * x
		} else {
			n_a = n_a - 1
			res = res * x
		}
	}
	if n < 0 {
		res = 1 / res
	}
	return res

}

func main() {
	myPow(2.0, 10)
}

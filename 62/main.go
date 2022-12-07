package main

import "math/big"

// uint63 no llega, habría que usar esta librería pero es un rollo

func factorial(x *big.Int, i *big.Int) big.Int {

	if x.Cmp(i) == -1 {
		return *big.NewInt(1)
	}
	return x * factorial(x.Sub(x, big.NewInt(-1)), i)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func uniquePaths(m int, n int) int {
	i, j := big.NewInt(m), big.NewInt(n)
	if max(m, n) == m {

		return int(factorial(*i.Add(i, j), *i) / factorial(j-1, 1))
	}
	return int(factorial(i+j-2, j) / factorial(i-1, 1))
}

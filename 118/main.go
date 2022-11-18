package main

import "fmt"

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = row
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}

package main

import "fmt"

func getRow(rowIndex int) []int {
	row := []int{1}

	for i := 1; i < rowIndex; i++ {
		new_row := make([]int, i+1)
		new_row[0] = 1
		new_row[i] = 1
		for j := 1; j < i; j++ {
			new_row[j] = row[j-1] + row[j]
		}
		row = new_row
	}

	return row
}

func main() {
	fmt.Println(getRow(5))
}

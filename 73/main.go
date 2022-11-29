package main

func setZeroes(matrix [][]int) {
	rows, columns := make(map[int]bool, len(matrix)), make(map[int]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
				// matrix[i][0] = 0
				// matrix[0][j] = 0
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if rows[i] || columns[j] {
				matrix[i][j] = 0
			}
		}
	}
}

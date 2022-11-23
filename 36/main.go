package main

import "fmt"

func selectSubBox(i, j int) (res int) {
	switch i {
	case 0, 1, 2:
		switch j {
		case 0, 1, 2:
			res = 0
		case 3, 4, 5:
			res = 1
		case 6, 7, 8:
			res = 2
		}
	case 3, 4, 5:
		switch j {
		case 0, 1, 2:
			res = 3
		case 3, 4, 5:
			res = 4
		case 6, 7, 8:
			res = 5
		}
	case 6, 7, 8:
		switch j {
		case 0, 1, 2:
			res = 6
		case 3, 4, 5:
			res = 7
		case 6, 7, 8:
			res = 8
		}
	}
	return res
}

func isValidSudoku(board [][]byte) bool {
	var sub_boxes = map[int][]byte{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}}
	var rows = map[int][]byte{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}}
	var columns = map[int][]byte{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}}

	for i := range board {
		for j, v := range board[i] {
			if v != '.' {
				//rows
				for _, v1 := range rows[i] {
					if v == v1 {
						return false
					}
				}

				rows[i] = append(rows[i], v)
				//columns
				for _, v1 := range columns[j] {
					if v == v1 {
						return false
					}
				}
				columns[j] = append(columns[j], v)
				//sub_boxes
				for _, v1 := range sub_boxes[selectSubBox(i, j)] {
					if v == v1 {
						return false
					}
				}
				sub_boxes[selectSubBox(i, j)] = append(sub_boxes[selectSubBox(i, j)], v)
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board))

}

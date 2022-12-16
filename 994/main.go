package main

func rot(grid [][]int, rotten_oranges [][]int, max_path, number_rotten_oranges int) (int, int) {
	var new_rots [][]int
	for _, rotten := range rotten_oranges {
		if rotten[0] > 0 && grid[rotten[0]-1][rotten[1]] == 1 {
			new_rots = append(new_rots, []int{rotten[0] - 1, rotten[1]})
			grid[rotten[0]-1][rotten[1]] = 2
			number_rotten_oranges++
		}
		if rotten[1] > 0 && grid[rotten[0]][rotten[1]-1] == 1 {
			new_rots = append(new_rots, []int{rotten[0], rotten[1] - 1})
			grid[rotten[0]][rotten[1]-1] = 2
			number_rotten_oranges++
		}
		if rotten[0] < len(grid)-1 && grid[rotten[0]+1][rotten[1]] == 1 {
			new_rots = append(new_rots, []int{rotten[0] + 1, rotten[1]})
			grid[rotten[0]+1][rotten[1]] = 2
			number_rotten_oranges++

		}
		if rotten[1] < len(grid[rotten[0]])-1 && grid[rotten[0]][rotten[1]+1] == 1 {
			new_rots = append(new_rots, []int{rotten[0], rotten[1] + 1})
			grid[rotten[0]][rotten[1]+1] = 2
			number_rotten_oranges++
		}
	}
	if len(new_rots) == 0 {
		return max_path, number_rotten_oranges
	}

	return rot(grid, new_rots, max_path+1, number_rotten_oranges)
}

func orangesRotting(grid [][]int) int {

	var fresh, rotten [][]int

	for row := range grid {
		for col, v := range grid[row] {
			if v == 1 {
				fresh = append(fresh, []int{row, col})
			} else if v == 2 {
				rotten = append(rotten, []int{row, col})
			}
		}
	}
	max_path, n_rotten_oranges := rot(grid, rotten, 0, len(rotten))
	if n_rotten_oranges == len(fresh)+len(rotten) {
		return max_path
	}
	return -1
}

func main() {
	//grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	grid := [][]int{{0, 2}}
	orangesRotting(grid)
}

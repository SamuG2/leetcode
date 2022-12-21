package main

var north = [2]int{0, 1}
var south = [2]int{0, -1}
var east = [2]int{1, 0}
var west = [2]int{-1, 0}

func changeDir(cur_dir [2]int, command int) [2]int {
	switch cur_dir {
	case north:
		if command == -1 {
			cur_dir = east
		} else {
			cur_dir = west
		}
	case south:
		if command == -1 {
			cur_dir = west
		} else {
			cur_dir = east
		}
	case east:
		if command == -1 {
			cur_dir = south
		} else {
			cur_dir = north
		}
	case west:
		if command == -1 {
			cur_dir = north
		} else {
			cur_dir = south
		}
	}
	return cur_dir
}

func robotSim(commands []int, obstacles [][]int) int {
	obstacles_map := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstacles_map[[2]int{obstacle[0], obstacle[1]}] = true
	}
	cur_pos := [2]int{0, 0}
	cur_dir := [2]int{0, 1}
	var max_distance int
	for _, command := range commands {
		if command == -1 || command == -2 {
			cur_dir = changeDir(cur_dir, command)
		} else {
			for i := 1; i <= command; i++ {
				next := [2]int{cur_pos[0] + cur_dir[0], cur_pos[1] + cur_dir[1]}
				if !obstacles_map[next] {
					cur_pos = next
				}
			}
			distance := cur_pos[0]*cur_pos[0] + cur_pos[1]*cur_pos[1]
			if distance > max_distance {
				max_distance = distance
			}
		}
	}

	return max_distance
}

func main() {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	robotSim(commands, obstacles)
}

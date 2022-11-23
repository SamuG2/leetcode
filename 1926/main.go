package main

import "fmt"

func checkNextSteps(maze [][]byte, pos []int) (res [][]int) {
	if pos[0] < len(maze)-1 && maze[pos[0]+1][pos[1]] == '.' {
		res = append(res, []int{pos[0] + 1, pos[1]})
	}
	if pos[0] > 0 && maze[pos[0]-1][pos[1]] == '.' {
		res = append(res, []int{pos[0] - 1, pos[1]})
	}
	if pos[1] < len(maze[pos[0]])-1 && maze[pos[0]][pos[1]+1] == '.' {
		res = append(res, []int{pos[0], pos[1] + 1})
	}
	if pos[1] > 0 && maze[pos[0]][pos[1]-1] == '.' {
		res = append(res, []int{pos[0], pos[1] - 1})
	}
	return
}

func bfs(maze [][]byte, to_explore [][]int, count int) int {
	var to_explore_next [][]int
	for _, pos := range to_explore {
		if maze[pos[0]][pos[1]] == '.' {
			if pos[0] == 0 || pos[1] == 0 || pos[0] == len(maze)-1 || pos[1] == len(maze[pos[0]])-1 {
				return count
			} else {
				maze[pos[0]][pos[1]] = '-'
				next_steps := checkNextSteps(maze, pos)
				for _, n := range next_steps {
					already_in := false
					for _, e := range to_explore_next {
						if n[0] == e[0] && n[1] == e[1] {
							already_in = true
							break
						}
					}
					if !already_in {
						to_explore_next = append(to_explore_next, n)
					}
				}

			}
		}
	}
	if len(to_explore_next) == 0 {
		return -1
	}
	return bfs(maze, to_explore_next, count+1)
}

func nearestExit(maze [][]byte, entrance []int) int {
	maze[entrance[0]][entrance[1]] = '-'
	next_steps := checkNextSteps(maze, entrance)
	return bfs(maze, next_steps, 1)
}

func main() {
	// maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	// entrance := []int{1, 2}
	maze := [][]byte{{'.', '.'}}
	entrance := []int{0, 1}

	fmt.Println(nearestExit(maze, entrance))

}

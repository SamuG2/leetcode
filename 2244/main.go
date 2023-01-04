package main

func minimumRounds(tasks []int) int {
	tasks_map := make(map[int]int)
	for _, task := range tasks {
		tasks_map[task]++
	}
	var rounds int
	for _, val := range tasks_map {
		if val == 1 {
			return -1
		} else if val%3 == 0 {
			rounds += val / 3
		} else {
			rounds += val/3 + 1
		}
	}
	return rounds
}

package main

import "sort"

func findWinners(matches [][]int) [][]int {
	losers := make(map[int]int)
	winners := make(map[int]int)
	for _, match := range matches {
		losers[match[1]]++
		winners[match[0]]++

	}
	res := make([][]int, 2)
	for key := range winners {
		if losers[key] == 0 {
			res[0] = append(res[0], key)
		}
	}
	for key, loses := range losers {
		if loses == 1 {
			res[1] = append(res[1], key)
		}

	}
	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}

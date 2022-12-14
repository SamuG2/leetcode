package main

import "sort"

type v_i struct {
	value, index int
}

func earliestFullBloom(plantTime []int, growTime []int) int {

	growTimeIdx := make([]v_i, len(growTime))
	for i, v := range growTime {
		growTimeIdx[i].value, growTimeIdx[i].index = v, i
	}
	sort.Slice(growTimeIdx, func(i, j int) bool {
		return growTimeIdx[i].value > growTimeIdx[j].value
	})

	var total_time, max int
	for _, pair := range growTimeIdx {
		total_time += plantTime[pair.index]
		if pair.value+total_time > max {
			max = pair.value + total_time
		}
	}
	return max
}

func main() {
	plantTime := []int{1, 2, 3, 2}
	growTime := []int{2, 1, 2, 1}

	earliestFullBloom(plantTime, growTime)

}

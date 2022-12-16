package main

func mostFrequentEven(nums []int) int {
	evens_map := make(map[int]int)
	res := -1
	evens_map[res] = 0
	for _, v := range nums {
		if v%2 == 0 || v == 0 {
			evens_map[v]++
			if evens_map[v] > evens_map[res] || (evens_map[v] == evens_map[res] && v < res) {
				res = v
			}
		}
	}
	return res
}

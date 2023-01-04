package main

func minimumCardPickup(cards []int) int {
	cards_idx := make(map[int]int)
	min_distance := len(cards) + 1
	for idx, card := range cards {
		if cards_idx[card] != 0 {
			candidate := idx - cards_idx[card] + 2
			if candidate < min_distance {
				min_distance = candidate
			}
		}
		cards_idx[card] = idx + 1
	}

	if min_distance == len(cards)+1 {
		return -1
	}
	return min_distance
}

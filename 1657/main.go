package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	for _, v1 := range word1 {
		var exists bool
		for _, v2 := range word2 {
			if v1 == v2 {
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}

	map_word1, map_word2 := make(map[rune]int), make(map[rune]int)

	for _, v := range word1 {
		map_word1[v]++
	}

	for _, v := range word2 {
		map_word2[v]++
	}

	for key, val := range map_word1 {
		if map_word2[key] == val {
			delete(map_word2, key)
			delete(map_word1, key)
		} else {
			for k, v := range map_word2 {
				if v == val {
					delete(map_word2, k)
					delete(map_word1, key)
					break
				}
			}
		}
	}

	return len(map_word2) == 0 && len(map_word1) == 0

}

func main() {
	fmt.Println(closeStrings("abbzccca", "babzzczc"))
}

package main

func possibleBipartition(n int, dislikes [][]int) bool {

	dislikes_map, groups := make(map[int][]int), make(map[int]int)
	for _, pair := range dislikes {
		dislikes_map[pair[0]] = append(dislikes_map[pair[0]], pair[1])
		dislikes_map[pair[1]] = append(dislikes_map[pair[1]], pair[0])
	}
	checked := make(map[int]bool)
	for key := range dislikes_map {
		if !checked[key] {
			queue := []int{key}
			checked[key] = true
			for len(queue) != 0 {
				if groups[queue[0]] == 0 {
					groups[queue[0]] = 1
					for _, v := range dislikes_map[queue[0]] {
						groups[v] = 2
						if !checked[v] {
							queue = append(queue, v)
							checked[v] = true
						}
					}
				} else if groups[queue[0]] == 1 {
					for _, v := range dislikes_map[queue[0]] {
						if groups[v] == 1 {
							return false
						}
						groups[v] = 2
						if !checked[v] {
							queue = append(queue, v)
							checked[v] = true
						}
					}
				} else if groups[queue[0]] == 2 {
					for _, v := range dislikes_map[queue[0]] {
						if groups[v] == 2 {
							return false
						}
						groups[v] = 1
						if !checked[v] {
							queue = append(queue, v)
							checked[v] = true
						}
					}
				}
				queue = queue[1:]
			}

		}
	}
	return true
}

func main() {
	n := 3
	dislikes := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	possibleBipartition(n, dislikes)

}

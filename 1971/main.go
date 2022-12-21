package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	if n == 1 {
		return true
	}
	var stack []int
	for i := 0; i < len(edges); i++ {
		if edges[i][0] == source || edges[i][1] == source {
			if edges[i][0] == destination || edges[i][1] == destination {
				return true
			} else {
				if edges[i][0] == source {
					stack = append(stack, edges[i][1])
				} else {
					stack = append(stack, edges[i][0])
				}
			}
			edges = append(edges[:i], edges[i+1:]...)
			i--
		}
	}
	for _, s := range stack {
		if validPath(n, edges, s, destination) {
			return true
		}
	}
	return false
}

package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	visited[0] = true
	keys := rooms[0]
	for len(keys) != 0 {
		if !visited[keys[0]] {
			keys = append(keys, rooms[keys[0]]...)
			visited[keys[0]] = true
		}
		keys = keys[1:]
	}
	return len(visited) == len(rooms)
}

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	canVisitAllRooms(rooms)
}

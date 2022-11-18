package main

import (
	"fmt"
)

//usando dfs
func removeStones(stones [][]int) int {
	// El numero máximo de piedras que se pueden retirar es igual al total menos el número de grupos que
	// puedas hacer siendo un grupo un conjunto de piedras que están enlazadas por una fila o columna.

	//guardamos en los indices en stones de cada piedra en mapas de filas y columnas, con las claves siendo la fila o columna concreta
	cols, rows := make(map[int][]int), make(map[int][]int)
	for i := 0; i < len(stones); i++ {
		rows[stones[i][0]] = append(rows[stones[i][0]], i)
		cols[stones[i][1]] = append(cols[stones[i][1]], i)

	}
	visited := make([]bool, len(stones))
	var groups int // de cada grupo se retiran todas menos una
	for i := 0; i < len(stones); i++ {
		if !visited[i] {
			visited[i] = true
			groups++
			dfs(rows, cols, visited, stones, i)
		}
	}
	return len(stones) - groups

}

func dfs(rows, cols map[int][]int, visited []bool, stones [][]int, curr int) {
	for _, rowNeighbour := range rows[stones[curr][0]] {
		if !visited[rowNeighbour] {
			visited[rowNeighbour] = true
			dfs(rows, cols, visited, stones, rowNeighbour)
		}
	}
	for _, colNeighbour := range cols[stones[curr][1]] {
		if !visited[colNeighbour] {
			visited[colNeighbour] = true
			dfs(rows, cols, visited, stones, colNeighbour)
		}
	}
}

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	fmt.Println(removeStones(stones))

}

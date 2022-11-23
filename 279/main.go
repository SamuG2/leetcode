package main

import (
	"fmt"
)

//encontrar suma de cuadrados

// func numSquares(n int) int {
// 	dp := make([]int, n+1)
// 	for i := range dp {
// 		dp[i] = i
// 	}

// 	for i := 0; i < n+1; i++ {
// 		for j := 1; j*j <= i; j++ {
// 			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-(j*j)]+1)))
// 		}
// 	}
// 	return dp[n]
// }

func numSquares(n int) int {
	var num_squares int
	var squares []int
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)

	}
	queue := []int{n}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		num_squares++
		cur_length := len(queue)
		for _, popped := range queue {
			for _, square := range squares {
				if popped == square {
					return num_squares
				}

				remainder := popped - square
				if remainder > 0 && !visited[remainder] {
					visited[remainder] = true
					queue = append(queue, remainder)
				}
			}
		}
		queue = queue[cur_length:]
	}
	return -1
}

func main() {
	fmt.Println(numSquares(888888))
}

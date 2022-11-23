package main

import "fmt"

func maxProfit(prices []int) int {
	var max_profit int
	min, max := prices[0], prices[0]

	for i := range prices {
		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min {
			min = prices[i]
			max = min
		}
		if max-min > max_profit {
			max_profit = max - min
		}
	}

	return max_profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

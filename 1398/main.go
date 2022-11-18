package main

import (
	"fmt"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, []int{0, h}...)
	verticalCuts = append(verticalCuts, []int{w, 0}...)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	max_h, max_v := 0, 0

	for i := 1; i < len(horizontalCuts); i++ {
		aux := horizontalCuts[i] - horizontalCuts[i-1]
		if aux > max_h {
			max_h = aux
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		aux := verticalCuts[i] - verticalCuts[i-1]
		if aux > max_v {
			max_v = aux
		}
	}
	return (max_h * max_v) % 1000000007

}

func main() {
	fmt.Println(maxArea(5, 4, []int{3, 1}, []int{1}))
}

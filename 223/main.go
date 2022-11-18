package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	res := (ay2-ay1)*(ax2-ax1) + (by2-by1)*(bx2-bx1)

	if ax1 < bx2 && bx1 < ax2 && ay1 < by2 && by1 < ay2 {
		var cx1, cx2, cy1, cy2 int
		if ax2 < bx2 {
			cx2 = ax2
		} else {
			cx2 = bx2
		}
		if ax1 > bx1 {
			cx1 = ax1
		} else {
			cx1 = bx1
		}

		if ay2 < by2 {
			cy2 = ay2
		} else {
			cy2 = by2
		}
		if ay1 > by1 {
			cy1 = ay1
		} else {
			cy1 = by1
		}
		res = res - (cx2-cx1)*(cy2-cy1)
	}
	return res
}

func main() {
	fmt.Println(computeArea(0, 0, 0, 0, -1, -1, 1, 1))
}

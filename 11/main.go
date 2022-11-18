package main

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func maxArea(height []int) int {
	max_v := 0
	for i, j := 0, len(height)-1; i < j; {
		min_h := min(height[i], height[j])
		cur_v := min_h * (j - i)
		if cur_v > max_v {
			max_v = cur_v
		}
		if min_h == height[i] {
			i++
		} else {
			j--
		}
	}
	return max_v
}

func main() {

}

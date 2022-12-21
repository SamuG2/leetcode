package main

func findDisappearedNumbers(nums []int) []int {
	found := make(map[int]bool)
	for _, n := range nums {
		found[n] = true
	}
	var res []int
	for i := 1; i <= len(nums); i++ {
		if !found[i] {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(nums)
}

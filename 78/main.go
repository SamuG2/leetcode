package main

import "fmt"

func subsets(nums []int) [][]int {
	l := len(nums) - 1

	var new_iteration, cur_iteration, res [][]int
	for i := 0; i <= l; i++ {
		cur_iteration = append(cur_iteration, []int{i})
		res = append(res, []int{i})
	}
	for i := 1; i <= l; i++ {
		for _, subset := range cur_iteration {
			if subset[len(subset)-1] != l {
				for i := subset[len(subset)-1] + 1; i <= l; i++ {
					var aux []int
					aux = append(aux, append(subset, i)...)
					// aux = append(aux, subset...)
					// aux = append(aux, i)
					new_iteration = append(new_iteration, aux)
				}
			}
		}

		res = append(res, new_iteration...)
		cur_iteration = new_iteration
		new_iteration = [][]int{}
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = nums[res[i][j]]
		}
	}
	res = append(res, []int{})
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(subsets(nums))
}

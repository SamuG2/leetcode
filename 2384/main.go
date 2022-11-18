package main

import "fmt"

func largestPalindromic(num string) string {

	frequency := make(map[rune]int, 10)
	var max_non_zero_frequency int
	var idx rune

	for _, v := range num {
		frequency[v]++
		if v != '0' && ((frequency[v] > max_non_zero_frequency) || (frequency[v] == max_non_zero_frequency && v > idx)) {
			max_non_zero_frequency = frequency[v]
			idx = v

		}
	}
	if max_non_zero_frequency == 0 {
		return "0"
	} else if max_non_zero_frequency == 1 {
		return string(idx)
	}

	var res []rune

	for i := '9'; i >= '0'; i-- {
		if frequency[i]%2 != 0 {
			res = append(res, i)
			frequency[i]--
			break
		}
	}

	for i := '0'; i <= '9'; i++ {
		aux := make([]rune, frequency[i]/2)
		for x := range aux {
			aux[x] = i
		}
		res = append(aux, res...)
		res = append(res, aux...)
	}

	return string(res)
}

func main() {
	fmt.Println(largestPalindromic("444947137"))
}

package main

import "fmt"

func strStr(haystack string, needle string) int {
	needle_idx := 0
	haystack_idx := 0
	for i := 0; i < len(haystack); i++ {
		if needle_idx == len(needle) {
			break
		}
		if haystack[i] == needle[needle_idx] {
			haystack_idx = i - needle_idx
			needle_idx++
		} else if needle_idx > 0 {
			i = haystack_idx
			needle_idx = 0
			haystack_idx = 0
		}
	}
	if needle_idx == len(needle) {
		return haystack_idx
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

package main

import "fmt"

func checkPalindrome(s string) bool {
	l := len(s)

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	for max := len(s) - 1; max > 0; max-- {
		for i := 0; i+max < len(s); i++ {
			if checkPalindrome(s[i : max+i+1]) {
				return s[i : max+i+1]
			}
		}
	}
	return s[0:1]

}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

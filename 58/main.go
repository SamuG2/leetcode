package main

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			for j := i; j >= 0; j-- {
				if s[j] == ' ' {
					break
				}
				res++
			}
			break
		}
	}
	return res
}

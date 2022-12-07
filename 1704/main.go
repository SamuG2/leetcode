package main

func halvesAreAlike(s string) bool {
	var n int
	vowels := "aAeEiIoOuU"
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		for t := range vowels {
			if s[i] == vowels[t] {
				n++
			}
			if s[j] == vowels[t] {
				n--
			}
		}
	}
	return n == 0
}

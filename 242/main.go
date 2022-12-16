package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_map, t_map := make(map[byte]int), make(map[byte]int)
	for i := range s {
		s_map[s[i]]++
		t_map[t[i]]++
	}
	for k, v := range s_map {
		if v != t_map[k] {
			return false
		}
	}
	return true
}

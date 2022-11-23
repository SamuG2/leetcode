package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	var prefix []byte
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, str := range strs {
			if len(str)-1 < i || str[i] != c {
				return string(prefix)
			}
		}
		prefix = append(prefix, c)
	}

	return string(prefix)
}

func main() {}

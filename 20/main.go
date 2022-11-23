package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var stack []rune

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func main() {

}

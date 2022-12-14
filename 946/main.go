package main

func validateStackSequences(pushed []int, popped []int) bool {
	for stack := []int{}; len(pushed) > 0 || len(popped) > 0; {
		if len(stack) != 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]

		} else if len(pushed) > 0 {
			stack = append(stack, pushed[0])
			pushed = pushed[1:]
		} else {
			return false
		}
	}
	return true
}

func main()

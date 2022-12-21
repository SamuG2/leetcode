package main

type CustomStack struct {
	stack []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{stack: make([]int, 0, maxSize)}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) != cap(this.stack) {
		this.stack = append(this.stack, x)
	}

}

func (this *CustomStack) Pop() int {
	if len(this.stack) > 0 {
		res := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return res
	}
	return -1
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(k, len(this.stack)); i++ {
		this.stack[i] += val
	}
}

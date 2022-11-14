/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	NormalStack, MinimalStack *Stack
}

func Constructor() MinStack {
	return MinStack{
		NormalStack:  &Stack{},
		MinimalStack: &Stack{},
	}
}

func (this *MinStack) Push(val int) {
	this.NormalStack.Push(val)
	if this.MinimalStack.Empty() || val <= this.MinimalStack.Top() {
		this.MinimalStack.Push(val)
	}
}

func (this *MinStack) Pop() {
	temp := this.NormalStack.Pop()
	if temp == this.MinimalStack.Top() {
		this.MinimalStack.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.NormalStack.Top()
}

func (this *MinStack) GetMin() int {
	return this.MinimalStack.Top()
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end


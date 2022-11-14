/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	A, B *Stack
}

func Constructor() MyQueue {
	return MyQueue{&Stack{}, &Stack{}}
}

func (this *MyQueue) Push(x int) {
	this.A.Push(x)
}

func (this *MyQueue) Pop() int {
	if !this.B.Empty() {
		return this.B.Pop()
	} else {
		for !this.A.Empty() {
			this.B.Push(this.A.Pop())
		}
		return this.B.Pop()
	}
}

func (this *MyQueue) Peek() int {
	if !this.B.Empty() {
		return this.B.Peek()
	} else {
		for !this.A.Empty() {
			this.B.Push(this.A.Pop())
		}
		return this.B.Peek()
	}
}

func (this *MyQueue) Empty() bool {
	return this.A.Empty() && this.B.Empty()
}

type Stack []int

func (ins *Stack) Push(x int) {
	*ins = append(*ins, x)
}

func (ins *Stack) Pop() int {
	temp := (*ins)[len(*ins)-1]
	*ins = (*ins)[:len(*ins)-1]
	return temp
}

func (ins *Stack) Empty() bool {
	return len(*ins) == 0
}

func (ins *Stack) Peek() int {
	return (*ins)[len(*ins)-1]
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end


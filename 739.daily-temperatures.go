/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// package leetcode

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{0}
	}
	ans := make([]int, len(temperatures))
	stack := &Stack{}
	stack.Push([]int{temperatures[0], 0})
	for i := 1; i < len(temperatures); i++ {
		for !stack.Empty() && stack.Top()[0] < temperatures[i] {
			ans[stack.Top()[1]] = i - stack.Top()[1]
			stack.Pop()
		}
		stack.Push([]int{temperatures[i], i})
	}
	return ans
}

type Stack [][]int

func (s *Stack) Push(x []int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() []int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *Stack) Top() []int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// @lc code=end


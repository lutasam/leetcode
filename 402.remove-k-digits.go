/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := &Stack{}
	i := 0
	for i < len(num) {
		if stack.Empty() || num[i] > stack.Top() {
			stack.Push(num[i])
		} else {
			for !stack.Empty() && stack.Top() > num[i] && k > 0 {
				stack.Pop()
				k--
			}
			stack.Push(num[i])
		}
		i++
	}
	for k > 0 {
		stack.Pop()
		k--
	}
	list := make([]byte, len(*stack))
	for i := len(list) - 1; i >= 0; i-- {
		list[i] = stack.Pop()
	}
	i = 0
	if len(list) == 0 || len(list) == 1 && list[0] == '0' {
		return "0"
	}
	for i < len(list) && list[i] == '0' {
		i++
	}
	if i >= len(list) {
		return "0"
	}
	return string(list[i:])
}

type Stack []byte

func (s *Stack) Push(x byte) {
	*s = append(*s, x)
}

func (s *Stack) Pop() byte {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *Stack) Top() byte {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// @lc code=end


/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := &Stack{}
	m := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else if stack.Empty() {
			return false
		} else {
			if m[stack.Peek()] == s[i] {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.Empty()
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

func (s *Stack) Peek() byte {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// @lc code=end


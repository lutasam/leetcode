/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
func calculate(s string) int {
	stack := &Stack{}
	stack.Push(0)
	var preSign byte = '+'
	num := 0
	for i := 0; i < len(s); i++ {
		if IsDigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '(' {
			j := FindRightBracket(s[i:])
			num = calculate(s[i+1 : i+j])
			i += j
		}
		if IsOp(s[i]) || i == len(s)-1 {
			switch preSign {
			case '+':
				stack.Push(num)
			case '-':
				stack.Push(-num)
			case '*':
				stack.Push(stack.Pop() * num)
			case '/':
				stack.Push(stack.Pop() / num)
			}
			preSign = s[i]
			num = 0
		}
	}
	ans := 0
	for !stack.Empty() {
		ans += stack.Pop()
	}
	return ans
}

func FindRightBracket(s string) int {
	i, level := 0, 0
	for i = 0; i < len(s); i++ {
		if s[i] == '(' {
			level++
		}
		if s[i] == ')' {
			level--
			if level == 0 {
				break
			}
		}
	}
	return i
}

func IsDigit(c byte) bool {
	return strings.Contains("0123456789", string(c))
}

func IsOp(c byte) bool {
	return strings.Contains("+-*/", string(c))
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

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// @lc code=end


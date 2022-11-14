/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// package leetcode

// import (
// 	"strconv"
// 	"strings"
// )

// @lc code=start
func decodeString(s string) string {
	stack := &Stack{}
	pos := 0
	for pos < len(s) {
		if IsNum(s[pos]) {
			var numStr string
			numStr, pos = GetNum(s, pos)
			stack.Push(numStr)
		} else if IsLetter(s[pos]) || s[pos] == '[' {
			stack.Push(string(s[pos]))
			pos++
		} else {
			pos++
			list := []string{}
			str := stack.Pop()
			for str != "[" {
				list = append(list, str)
				str = stack.Pop()
			}
			numStr := stack.Pop()
			repeatTime, _ := strconv.Atoi(numStr)
			decodeStr := TranslateStr(list, repeatTime)
			stack.Push(decodeStr)
		}
	}
	ans := GetAns(stack)
	return ans
}

type Stack []string

func (ins *Stack) Push(x string) {
	*ins = append(*ins, x)
}

func (ins *Stack) Pop() string {
	temp := (*ins)[len(*ins)-1]
	*ins = (*ins)[:len(*ins)-1]
	return temp
}

func IsNum(c byte) bool {
	return strings.Contains("0123456789", string(c))
}

func GetNum(s string, pos int) (string, int) {
	cs := []byte{}
	for IsNum(s[pos]) {
		cs = append(cs, s[pos])
		pos++
	}
	return string(cs), pos
}

func IsLetter(c byte) bool {
	return strings.Contains("abcdefghijklmnopqrstuvwxyz", string(c))
}

func TranslateStr(list []string, repeatTime int) string {
	s := ""
	for i := len(list) - 1; i >= 0; i-- {
		s += list[i]
	}
	ret := ""
	for i := 0; i < repeatTime; i++ {
		ret += s
	}
	return ret
}

func GetAns(s *Stack) string {
	ans := ""
	for i := 0; i < len(*s); i++ {
		ans += (*s)[i]
	}
	return ans
}

// @lc code=end

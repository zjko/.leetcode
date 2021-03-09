/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

type stack struct {
	list []uint8
	top  int
	size int
}

func (s *stack) init(length int) {
	s.list = make([]uint8, length)
	s.size = 0
	s.top = -1
}

func (s *stack) push(val uint8) bool {

	if s.top > -1 {
		c := s.list[s.top] == '[' && val == ']' || s.list[s.top] == '(' && val == ')' || s.list[s.top] == '{' && val == '}'
		if c {
			s.top--
			s.size--
			return false
		}

	}
	s.top++
	s.list[s.top] = val
	s.size++
	return false
}

func isValid(s string) bool {
	a := &stack{}
	length := len(s)
	a.init(length)
	for i := 0; i < length; i++ {
		a.push(s[i])
	}
	return a.top == -1
}

// @lc code=end


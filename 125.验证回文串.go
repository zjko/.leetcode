/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start

func isPalindrome(s string) bool {
	l := len(s)
	p := 0
	q := l - 1
	for p < q && p < l && q >= 0 {

		for p <= q && p < l-1 && !filter(s[p]) {
			p++
		}
		for p <= q && q > 0 && !filter(s[q]) {
			q--
		}
		if q == 0 && p == l-1 {
			return true
		} else {
			return false
		}

		// fmt.Printf("%c %c \n", s[p], s[q])
		if convert(s[p]) == convert(s[q]) {
			p++
			q--

		} else {
			return false
		}
	}
	return true
}
func convert(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c
}
func filter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || c >= 'a' && c <= 'z'
}

// @lc code=end

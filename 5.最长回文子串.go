/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	maxL := 0
	ls := 0
	le := 0
	for i := 0; i < len(s); i++ {
		l := 0
		for j := 1; j+i < len(s); j++ {
			if i-j+1 < 0 || i+j >= len(s) || s[i+j] != s[i-j+1] {
				break
			}
			l += 2
			if maxL < l {
				maxL = l
				ls = i - j + 1
				le = i + j
			}
		}

		l = 1
		for j := 1; j < len(s); j++ {
			if i-j < 0 || i+j >= len(s) || s[i+j] != s[i-j] {
				break
			}
			//println(string(s[i-1]),string(s[i]),string(s[i+1]))
			l += 2
			if maxL < l {
				maxL = l
				ls = i - j
				le = i + j
			}

		}
	}

	ns := make([]byte, le-ls+1)
	c := 0
	//println(ls,le)
	for i := ls; i <= le; i++ {
		ns[c] = s[i]
		c++
	}
	return string(ns)
}

// @lc code=end

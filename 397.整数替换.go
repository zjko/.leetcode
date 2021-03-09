/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	k := 0
	for n > 1 {
		switch n {
		case 2:
			return k + 1
		case 3:
			return k + 2
		}
		k++
		if n%2 == 1 {
			a := maxPower(n + 1)
			b := maxPower(n - 1)
			if a > b {
				n++
			} else {
				n--
			}
		} else {
			n /= 2
		}
	}
	return k
}
func maxPower(n int) int {
	c := 2
	for n%c == 0 {
		c *= 2
	}
	return c
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, min(len(s), numRows))
	curRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	var ret string
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		ret += row
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
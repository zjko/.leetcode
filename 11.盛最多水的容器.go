/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

func maxArea(height []int) int {
	s := 0
	e := len(height) - 1
	max := 0
	for s < e {
		r := min(height[s], height[e]) * (e - s)
		if r > max {
			max = r
			//println(s,e,r)
		}
		if height[s] <= height[e] {
			s++
		} else {
			e--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

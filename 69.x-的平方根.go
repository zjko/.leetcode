import "math"

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// @lc code=end

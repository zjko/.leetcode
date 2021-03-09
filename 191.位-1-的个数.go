/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	c := 0

	for num != 0 {
		if num%2 == 1 {
			c++
		}

		num /= 2
	}

	return c
}

// @lc code=end


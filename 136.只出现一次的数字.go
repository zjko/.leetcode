/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	t := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		t ^= nums[i]
	}
	return t
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	delCnt := 0
	arrLen := len(nums)
	for i := 0; i < arrLen; i++ {
		if val == nums[i] {
			delCnt++
		} else {
			nums[i-delCnt] = nums[i]
		}
	}

	return arrLen - delCnt
}

// @lc code=end


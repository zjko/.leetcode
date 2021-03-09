/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	delCnt := 0
	arrLen := len(nums)
	for i := 1; i < arrLen; i++ {
		if nums[i-1] == nums[i] {
			delCnt++
		} else {
			nums[i-delCnt] = nums[i]
		}
	}
	return arrLen - delCnt
}

// @lc code=end


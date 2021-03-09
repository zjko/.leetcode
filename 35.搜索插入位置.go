/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start

func searchInsert(nums []int, target int) int {
	nLen := len(nums)
	for i := 0; i < nLen; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return nLen
}

// @lc code=end


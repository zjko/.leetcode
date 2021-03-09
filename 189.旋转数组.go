/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	numsLen := len(nums)
	tmp := 0
	targetPosition := 0
	for i := 0; i < numsLen; i++ {
		targetPosition = (i + k) % numsLen
		tmp = nums[targetPosition]

	}
}

// @lc code=end


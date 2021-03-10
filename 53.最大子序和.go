/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)
	dp[0] = nums[0]
	// 动态规划 以i作为结尾的最大连续子序列和dp[i] = max{dp[i-1]+nums[i],nums[i]}
	for i := 1; i < numsLen; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	//找出 nums 中最大的连续子序列和
	maxRes := dp[0]
	for i := 1; i < numsLen; i++ {
		if dp[i] > maxRes {
			maxRes = dp[i]
		}
	}
	return maxRes
}

func max(a ...int) int {
	m := -999999
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}

// @lc code=end

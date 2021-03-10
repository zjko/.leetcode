/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen < 2 {
		return 0
	}
	nums := make([]int, pricesLen-1)
	for i := 0; i < len(nums); i++ {
		nums[i] = prices[i+1] - prices[i]
	}
	// fmt.Println(nums)
	//计算最大连续子序列 dp[i] = max{num[i] ,dp[i - 1] + nums[i]}
	// dp := make([]int, len(nums))
	// dp[0] = nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	if dp[i-1] > 0 {
	// 		dp[i] = dp[i-1] + nums[i]
	// 	} else {
	// 		dp[i] = nums[i]
	// 	}
	// }

	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			sum += nums[i]
		}
	}
	return sum
}

// @lc code=end

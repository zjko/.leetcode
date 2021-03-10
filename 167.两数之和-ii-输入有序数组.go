/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	numMap := make(map[int]int, 100)
	l := len(numbers)
	for i := 0; i < l; i++ {
		numMap[numbers[i]] = i
	}

	for i := 0; i < l; i++ {
		if _, ok := numMap[target-numbers[i]]; ok {
			return []int{i + 1, numMap[target-numbers[i]]}
		}
	}

	return nil

}

// @lc code=end

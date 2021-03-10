/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	countMap := make(map[int]int, 10000)
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	count := l / 2
	for i := 0; i < l; i++ {
		// fmt.Println(countMap)
		_, ok := countMap[nums[i]]
		if ok {
			if countMap[nums[i]]+1 > count {
				return nums[i]
			}
			countMap[nums[i]]++
		} else {
			countMap[nums[i]] = 1
		}
	}
	return 0

}

// @lc code=end

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	// 用倒排索引解决问题
	var indexArr [255]int

	length := len(strs)
	// 找出长度最长的字符串
	maxStrLen := 0
	for i := 0; i < length; i++ {
		strLength := len(strs[i])
		if strLength > maxStrLen {
			maxStrLen = strLength
		}
	}
	publicStrArr := make([]byte, maxStrLen)
	resStr := make([]byte, maxStrLen)

	t := 0
	res := 0
	// resE := 0
	for i := 0; i < length; i++ {
		cLength := len(strs[i])
		for j := 0; j < cLength; j++ {
			indexArr[strs[i][j]]++
			if indexArr[strs[i][j]] == 3 {
				publicStrArr[t] = strs[i][j]
				t++
				resStr[t] = strs[i][j]
			} else {

				if t > res {
					res = t
					t = 0
				}
			}
		}
	}

	return (string)(publicStrArr)
}

// @lc code=end


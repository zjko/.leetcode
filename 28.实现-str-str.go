/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}
	if haystackLen < needleLen {
		return -1
	}

	for i := 0; i < haystackLen; i++ {
		var j int
		for j = 0; j < needleLen; j++ {
			if i+j > haystackLen-1 || haystack[i+j] != needle[j] {
				break
			}

		}
		// fmt.Printf("%d %d\n", j, needleLen)
		if j == needleLen {
			return i
		}

	}
	return -1
}

// @lc code=end


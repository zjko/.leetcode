/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	minLength := 99999999
	arrLength := len(strs)
	if arrLength == 1 {
		return strs[0]
	}
	for i := 0; i < arrLength; i++ {
		strLength := len(strs[i])
		if strLength < minLength {
			minLength = strLength
		}
	}
	resStr := make([]byte, minLength)
	t := 0
	for i := 0; i < minLength; i++ {
		k := 0

		for j := 1; j < arrLength; j++ {
			// fmt.Printf("%d %d %d %c %c\n", i, j, k, strs[j][i], strs[0][i])
			// fmt.Println(resStr)

			if strs[j][i] == strs[0][i] {
				k++
			}
		}
		if k == arrLength-1 {
			resStr[t] = strs[0][i]
			t++
		} else {
			return byteString(resStr)
		}

	}
	return byteString(resStr)
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

// @lc code=end


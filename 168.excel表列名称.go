/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

func convertToTitle(n int) string {
	str := make([]byte, 5)

	for i := 0; n != 0; i++ {
		str[i] = (byte)(n%26 + 64)
		n /= 26
		// fmt.Println(str[i])
	}

	resStr := make([]byte, 5)
	k := 0
	for i := 4; i >= 0; i-- {
		if str[i] == 0 {
			continue
		}
		resStr[k] = str[i]
		k++

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

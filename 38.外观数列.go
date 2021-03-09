/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return function(n-1, "1")

}

func function(n int, str string) string {
	if n == 0 {
		return str
	}
	strLen := len(str)
	arr := make([]byte, strLen*2)

	t := 1
	p := 0
	for i := 0; i < strLen; i++ {
		if i+1 >= strLen || str[i] != str[i+1] {
			arr[p] = (byte)(t + 48)
			arr[p+1] = str[i]
			// fmt.Printf("%c*%c*%d\t", arr[p], str[i], t)
			t = 1
			p += 2
		} else {
			t++
		}

	}
	// fmt.Println(byteString(arr))
	return function(n-1, byteString(arr))

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


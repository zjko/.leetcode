/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	newArr := make([]byte, aLen+bLen)
	carry := false
	sum := (byte)(0)
	i := 0
	for ; i < aLen || i < bLen; i++ {
		if i >= aLen {
			sum = b[bLen-i-1] - '0'
		} else if i >= bLen {
			sum = a[aLen-i-1] - '0'
		} else {
			sum = a[aLen-i-1] - '0' + b[bLen-i-1] - '0'
		}
		if carry {
			sum++
			carry = false
		}

		if sum > 1 {
			sum %= 2
			carry = true
		}

		newArr[i] = sum + '0'
	}

	if carry {
		newArr[i] = 1 + '0'
		i++
	}
	res := make([]byte, i)
	for p := 0; p < i; p++ {
		res[p] = newArr[i-p-1]
	}

	return (string)(res)

}

// @lc code=end

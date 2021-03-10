/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digitsLen := len(digits)
	digits[digitsLen-1]++
	carry := false
	for i := digitsLen - 1; i > -1; i-- {
		if carry {
			digits[i]++
			carry = false
		}
		if digits[i] > 9 {
			digits[i] %= 10
			carry = true
		}
	}
	if carry {
		newArr := make([]int, digitsLen+1)
		newArr[0] = 1
		for i := 1; i < digitsLen+1; i++ {
			newArr[i] = digits[i-1]
		}
		digits = newArr
	}
	return digits
}

// @lc code=end

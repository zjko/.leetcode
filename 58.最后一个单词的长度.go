/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	wordLen := 0
	flag := true
	// flag  true表示上一个字符 是空格
	if s[0] != ' ' {
		wordLen = 1
		flag = false
	}
	for i := 1; i < sLen; i++ {
		if s[i] != ' ' {
			// 这一个为字母
			if flag {
				// 上一个为空格
				wordLen = 1
				flag = false
			} else {
				//上一个为字母
				wordLen++
			}
		} else {

			flag = true
		}
	}
	return wordLen
}

// @lc code=end

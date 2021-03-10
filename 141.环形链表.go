/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val == -999999 {
			return true
		}
		head.Val = -999999
		head = head.Next
	}
	return false
}

// @lc code=end

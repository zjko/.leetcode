/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	if p == nil {
		return nil
	}
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}
	return head
}

// @lc code=end

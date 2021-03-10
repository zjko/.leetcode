/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB
	for p != nil && q != nil {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next

		if p == nil && q != nil {
			p = headB
		}

		if q == nil && p != nil {
			q = headA
		}

	}
	return nil
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var res *ListNode
	// var pre *ListNode

	res = &ListNode{}
	head = res
	for l1 != nil || l2 != nil {
		// fmt.Printf("%d %d \n", l1.Val, l2.Val)
		if l1.Val > l2.Val {
			res.Val = l2.Val
			l2 = l2.Next
			if l2 == nil {
				// fmt.Printf("l2 end preVal %d", pre.Val)
				res.Next = l1
				break
			}
		} else {
			res.Val = l1.Val
			l1 = l1.Next
			if l1 == nil {
				// fmt.Printf("l1 end preVal %d", pre.Val)
				res.Next = l2
				break
			}
		}
		// pre = res
		res.Next = &ListNode{}
		res = res.Next
	}

	return head

}

// @lc code=end


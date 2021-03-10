/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// print1(head)
	var currentMaxPoint, prePoint *ListNode

	for p := head; p != nil; p = p.Next {
		if p.Val >= x {
			// 找到一个>x的节点
			currentMaxPoint = p
			// fmt.Println(p.Val,"// 将这个节点后面所有 < x的节点挪在前面")
			// 将这个节点后面所有 < x的节点挪在前面
			pre2 := currentMaxPoint
			for t := currentMaxPoint.Next; t != nil; t = t.Next {
				// fmt.Println(t.Val)

				if t.Val < x {
					// fmt.Println(t.Val,"找到一个小的，准备换位置")
					// 找比 当前大节点小的，放在其 前方
					if prePoint == nil {
						head = t
						if pre2 != nil {
							pre2.Next = t.Next
						}

						t.Next = currentMaxPoint
						prePoint = t
					} else {
						prePoint.Next = t
						if pre2 != nil {
							pre2.Next = t.Next
						}
						t.Next = currentMaxPoint
						prePoint = t
					}

				}
				pre2 = t

			}

		}
		// 记录上一个节点
		prePoint = p
	}
	// print1(head)
	return head
}

// @lc code=end


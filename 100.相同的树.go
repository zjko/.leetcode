/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}

	if p.Val == q.Val {
		a := isSameTree(p.Left, q.Left)
		b := isSameTree(p.Right, q.Right)
		if a && b {
			return true
		}
	}
	return false
}

// @lc code=end

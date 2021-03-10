/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// fmt.Println(root.Val)
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	// fmt.Printf("%d %d %d\n", a, b, root.Val)
	max := a
	if max < b {
		max = b
	}

	return max + 1
}

// @lc code=end

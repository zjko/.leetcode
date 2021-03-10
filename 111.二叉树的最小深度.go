/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return search(root, 1, 9999999)
}

// root 当前子树
// level 当前层级
// minLevel 当前最小层级，-1表示暂时为找到最小层级
// return int 返回最小层级
func search(root *TreeNode, level int, minLevel int) int {

	if level >= minLevel {
		return minLevel
	}

	// 找到叶子节点，则直接返回
	if root.Left == nil && root.Right == nil {
		return level
	}
	// 没有找到叶子节点，则需要继续深入
	res := 9999999
	if root.Left != nil{

		a := search(root.Left, level+1, minLevel)
		if res > a {
			res = a
		}
	}
	if root.Right != nil {
		a := search(root.Right, level+1, minLevel)
		if res > a{
			res = a
		}
	}


	return res

}

// 
@lc code=end

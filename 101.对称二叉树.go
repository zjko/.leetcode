import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	a := leftSearch(root)
	fmt.Println()
	b := rightSearch(root)
	fmt.Println()
	i := 0
	for ; i < len(a); i++ {
		if a[i] != b[i] {
			break
		}
	}
	if i == len(a) {
		return true
	}
	return false
}

func leftSearch(root *TreeNode) string {
	if root == nil {
		return "*"
	}
	fmt.Print(root.Val)
	s := strconv.Itoa(root.Val+97) + leftSearch(root.Left)
	return s + strconv.Itoa(root.Val+97) + leftSearch(root.Right)
}

func rightSearch(root *TreeNode) string {
	if root == nil {
		return "*"
	}
	fmt.Print(root.Val)
	s := strconv.Itoa(root.Val+97) + rightSearch(root.Right)
	return s + strconv.Itoa(root.Val+97) + rightSearch(root.Left)

}

// @lc code=end

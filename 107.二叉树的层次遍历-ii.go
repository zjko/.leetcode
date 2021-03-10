/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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

const SIZE = 1000

func levelOrderBottom(root *TreeNode) [][]int {
	resArr := make([][]int, SIZE)
	nodeArr := make([]*TreeNode, SIZE)

	if root == nil {
		return [][]int{}
	}
	nodeArr[0] = root
	resArr[0] = []int{root.Val}

	// level 二叉树层级
	l := 0

	for true {
		// 每一层每个节点的数值 其中nums[0]表示这个数组的长度
		nums := make([]int, SIZE)
		// 下一层的节点集合
		nextNodeArr := make([]*TreeNode, SIZE)
		k := 0
		i := 0
		for ; nodeArr[i] != nil; i++ {
			nums[i+1] = nodeArr[i].Val
			if nodeArr[i].Right != nil {
				nextNodeArr[k] = nodeArr[i].Right
				k++
			}
			if nodeArr[i].Left != nil {
				nextNodeArr[k] = nodeArr[i].Left
				k++
			}
		}
		nums[0] = i
		resArr[l] = nums
		l++
		if i > 0 {
			nodeArr = nextNodeArr
			i = 0
		} else {
			break
		}
	}
	l--

	// for i := 0; i < l; i++ {
	// 	fmt.Println(resArr[i])
	// }
	// 节点数组排序
	res := make([][]int, l)

	for i := 0; i <= l-1; i++ {
		pArr := resArr[l-1-i]
		rArr := make([]int, pArr[0])

		// fmt.Println(pArr)
		for j := 0; j < pArr[0]; j++ {
			rArr[j] = pArr[pArr[0]-j]
		}
		res[i] = rArr
		// fmt.Println(rArr)
	}
	// fmt.Println(res)
	return res
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start

func merge(nums1 []int, m int, nums2 []int, n int) {

	newArr := make([]int, m+n)
	k := 0
	i := 0
	j := 0
	for i < m || j < n {
		a := 0
		if i >= m {
			a = nums2[j]
			j++
		} else if j >= n {
			a = nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			a = nums2[j]
			j++
		} else {
			a = nums1[i]
			i++
		}

		newArr[k] = a
		k++

	}

	for i := 0; i < m+n; i++ {
		nums1[i] = newArr[i]
	}
}

// @lc code=end

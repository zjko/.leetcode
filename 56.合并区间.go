/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
import (
	"fmt"
	"sort"
)

type ints [][]int

func (s ints) Len() int           { return len(s) }
func (s ints) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s ints) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func merge(intervals [][]int) [][]int {

	// STEP 1 先排序
	sort.Sort(ints(intervals))
	if len(intervals) < 2 {
		return intervals
	}
	// 承载新数组的列表
	newList := make([][]int, 1, len(intervals))
	p := 0
	newList[0] = intervals[0]
	for i := 1; i < len(intervals); i++ {
		// fmt.Println(newList)
		currentS, currentE := intervals[i][0], intervals[i][1]
		if currentS <= newList[p][1] {
			if currentE >= newList[p][1] {
				newList[p][1] = currentE
			}
		} else {
			p++
			newList = append(newList, intervals[i])
			//newList[p] = intervals[i]
		}
	}
	//fmt.Println(newList)

	return newList
}

// @lc code=end


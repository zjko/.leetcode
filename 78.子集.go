import "fmt"

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	numsLen := (uint)(len(nums))
	bitSet := (uint)(1 << (uint)(numsLen))
	res := make([][]int, bitSet)
	fmt.Println(bitSet)
	for i := (uint)(0); i < bitSet; i++ {
		tmpRes := make([]int, numsLen)
		// 记录有1的数量
		t := 0
		//记录当前位图的位置
		j := (uint)(1)
		for c := (uint)(0); c < numsLen; c++ {
			p := i & j
			// fmt.Printf("i%d j%d p%d \n", i, j, p)
			if p != 0 {
				tmpRes[t] = nums[c]
				t++
			}
			j = 1 << (c + 1)
		}
		// fmt.Println(tmpRes)
		res[i] = tmpRes[0:t]
	}

	return res

}

// @lc code=end

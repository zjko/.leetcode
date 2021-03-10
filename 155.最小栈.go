/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	list    []int
	minList []int
	top     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		list:    make([]int, 10000),
		minList: make([]int, 10000),
		top:     -1,
	}
}

func (this *MinStack) Push(x int) {
	this.top++
	this.list[this.top] = x
	if this.top == 0 {
		this.minList[0] = x
	} else if x > this.minList[this.top-1] {
		this.minList[this.top] = this.minList[this.top-1]
	} else {
		this.minList[this.top] = x
	}
}

func (this *MinStack) Pop() {
	this.top--
}

func (this *MinStack) Top() int {
	return this.list[this.top]
}

func (this *MinStack) GetMin() int {
	return this.minList[this.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

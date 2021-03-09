import "fmt"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 */

// @lc code=start
type DNode struct {
	//数据
	val int
	//索引
	index int

	//后一个
	next *DNode
	//前一个节点
	prev *DNode
}

type LRUCache struct {

	// 索引
	indexMap map[int]*DNode

	// 头节点
	head *DNode
	// 尾结点
	tail *DNode

	// 链表节点数量
	count int

	// 容量
	capacity int
}

func Constructor(capacity int) LRUCache {
	fmt.Println("初始化")
	p := &DNode{}
	return LRUCache{
		indexMap: make(map[int]*DNode, capacity*2),
		head:     p,
		tail:     p,
		count:    -1,
		capacity: capacity,
	}
}
func (this *LRUCache) print() {
	// fmt.Println(this.indexMap)
	re1 := make([]int, this.count)
	re2 := make([]int, this.count)
	p := this.head
	for i := 0; i < this.count && p != nil; i++ {
		fmt.Print(p.index, "->")
		re1[i] = p.index
		// fmt.Print(p.index, ",", p.val, "->")
		p = p.next

	}
	p = this.tail
	fmt.Println("\n------------")
	fmt.Println()
	for i := 0; i < this.count && p != nil; i++ {
		fmt.Print(p.index, "->")
		re2[this.count-i-1] = p.index
		// fmt.Print(p.index, ",", p.val, "->")
		p = p.prev
	}

	for i := 0; i < this.count; i++ {
		if re1[i] != re2[i] {
			fmt.Println("*************")
		}
	}

	fmt.Println()
	fmt.Println("count", this.count)
	fmt.Println("capacity", this.capacity)
	fmt.Println("------------\n")

}

func (this *LRUCache) Get(key int) int {
	fmt.Println("GET key", key)
	t, e := this.indexMap[key]
	if !e {
		fmt.Println(key, "不存在 -1")
		return -1
	}
	if t == this.head {
		return this.head.val
	}
	if t.prev != nil {
		if t == this.tail {
			this.tail = t.prev
		}
		t.prev.next = t.next
	}
	if t.next != nil {
		t.next.prev = t.prev
	}

	// 将该节点放置到head后面
	t.next = this.head
	this.head.prev = t
	this.head = t
	t.prev = t.next
	return this.head.val
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("PUT key", key, "val", value)
	// 是否为第一个元素
	if this.count == -1 {
		this.indexMap[key] = this.head
		this.head.val = value
		this.head.index = key
		this.head.prev = nil
		this.head.next = nil
		this.count = 1

		return
	}

	t, e := this.indexMap[key]
	if e {
		// 若存在，则只需修改数据，并将节点前置
		t.val = value
		// 若前置节点不存在，则表示当前就是第一个元素，不需要做任何变化
		if t.prev == nil || t == this.head {
			return
		}
		//若当前节点的前置节点存在
		// 若是尾结点，则需要更新尾结点到前置节点
		if t == this.tail || t.next == nil {
			fmt.Println("我是尾结点")
			this.tail = t.prev
			this.tail.next = nil
		}

		t.prev.next = t.next

		if t.next != nil {
			t.next.prev = t.prev
		}
		this.head.prev = t
		t.prev = nil
		t.next = this.head
		this.head = t
		return
	}

	// 若该节点不存在，且总容量未满

	p := &DNode{
		val:  value,
		next: this.head,
		prev: nil,
	}
	this.head.prev = p
	this.head = p
	this.indexMap[key] = p
	this.head.index = key
	//容量未满，加完不管
	if this.count+1 <= this.capacity {
		this.count++
		return
	}
	// 容量已满，加完需要删除后置节点
	delete(this.indexMap, this.tail.index)

	this.tail = this.tail.prev
	this.tail.next = nil
	return
}

// @lc code=end


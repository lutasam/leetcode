/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// package leetcode

// @lc code=start
type LRUCache struct {
	CacheMap       map[int]*DListNode
	Head, Tail     *DListNode
	Size, Capacity int
}

type DListNode struct {
	Key, Val  int
	Pre, Next *DListNode
}

func (this *LRUCache) Add(node *DListNode) {
	this.CacheMap[node.Key] = node
	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	node.Pre = this.Head
	this.Head.Next = node
	this.Size++
}

func (this *LRUCache) Remove(node *DListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Pre = nil
	node.Next = nil
	delete(this.CacheMap, node.Key)
	this.Size--
}

func (this *LRUCache) Update(node *DListNode, newVal int) {
	this.Remove(node)
	node.Val = newVal
	this.Add(node)
}

func Constructor(capacity int) LRUCache {
	obj := LRUCache{
		CacheMap: make(map[int]*DListNode, 0),
		Head:     &DListNode{},
		Tail:     &DListNode{},
		Size:     0,
		Capacity: capacity,
	}
	obj.Head.Next = obj.Tail
	obj.Tail.Pre = obj.Head
	return obj
}

func (this *LRUCache) Get(key int) int {
	if _, exist := this.CacheMap[key]; !exist {
		return -1
	}
	temp := this.CacheMap[key].Val
	this.Update(this.CacheMap[key], temp)
	return temp
}

func (this *LRUCache) Put(key int, value int) {
	if _, exist := this.CacheMap[key]; exist {
		this.Update(this.CacheMap[key], value)
	} else {
		if this.Size >= this.Capacity {
			this.Remove(this.Tail.Pre)
			this.Add(&DListNode{key, value, nil, nil})
		} else {
			this.Add(&DListNode{key, value, nil, nil})
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

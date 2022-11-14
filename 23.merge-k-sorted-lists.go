/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := &Heap{}
	for _, list := range lists {
		curr := list
		for curr != nil {
			h.Push(curr)
			curr = curr.Next
		}
	}
	dummy := &ListNode{}
	curr := dummy
	for !h.Empty() {
		curr.Next = h.Pop()
		curr = curr.Next
	}
	curr.Next = nil
	return dummy.Next
}

type Heap []*ListNode

func (h *Heap) Push(node *ListNode) {
	*h = append(*h, node)
	h.AdjustUp()
}

func (h *Heap) Pop() *ListNode {
	temp := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.AdjustDown()
	return temp
}

func (h *Heap) Empty() bool {
	return len(*h) == 0
}

func (h *Heap) AdjustUp() {
	curr := len(*h) - 1
	for {
		parent := (curr - 1) / 2
		if parent < 0 {
			break
		}
		if (*h)[parent].Val > (*h)[curr].Val {
			(*h)[parent], (*h)[curr] = (*h)[curr], (*h)[parent]
		} else {
			break
		}
		curr = parent
	}
}

func (h *Heap) AdjustDown() {
	curr := 0
	leftChild := 2*curr + 1
	for leftChild < len(*h) {
		rightChild := leftChild + 1
		if rightChild < len(*h) && (*h)[rightChild].Val < (*h)[leftChild].Val {
			leftChild = rightChild
		}
		if (*h)[curr].Val > (*h)[leftChild].Val {
			(*h)[curr], (*h)[leftChild] = (*h)[leftChild], (*h)[curr]
		} else {
			break
		}
		curr = leftChild
		leftChild = 2*curr + 1
	}
}

// @lc code=end


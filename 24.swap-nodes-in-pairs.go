/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := &ListNode{0, nil}, &ListNode{0, nil}
	l, r, curr := left, right, head
	flag := true
	for curr != nil {
		if flag {
			l.Next = curr
			l = l.Next
		} else {
			r.Next = curr
			r = r.Next
		}
		curr = curr.Next
		flag = !flag
	}
	l.Next, r.Next = nil, nil
	dummy := &ListNode{0, nil}
	l, r, curr = left.Next, right.Next, dummy
	flag = true
	for l != nil && r != nil {
		if flag {
			curr.Next = r
			r = r.Next
		} else {
			curr.Next = l
			l = l.Next
		}
		curr = curr.Next
		flag = !flag
	}
	if l != nil {
		curr.Next = l
	}
	if r != nil {
		curr.Next = r
	}
	return dummy.Next
}

// @lc code=end


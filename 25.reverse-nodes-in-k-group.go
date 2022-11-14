/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre, tail := dummy, dummy
	var next *ListNode
out:
	for tail.Next != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				break out
			}
		}
		next = tail.Next
		tail.Next = nil
		reverseHead := Reverse(head)
		pre.Next = reverseHead
		tail = reverseHead
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// @lc code=end


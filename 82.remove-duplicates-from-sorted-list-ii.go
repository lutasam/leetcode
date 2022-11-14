/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	curr := dummy

	for curr != nil && curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			t := curr.Next.Val
			temp := curr.Next.Next
			for temp.Next != nil && temp.Next.Val == t {
				temp = temp.Next
			}
			curr.Next = temp.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

// @lc code=end


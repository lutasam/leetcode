/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	return QuickSort(head)
}

func QuickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, pivot, right := Partition(head)
	left = QuickSort(left)
	right = QuickSort(right)

	if left != nil {
		tail := left
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = pivot
		pivot.Next = right

		return left
	} else {
		pivot.Next = right
		return pivot
	}
}

func Partition(head *ListNode) (*ListNode, *ListNode, *ListNode) {
	pivot := head
	left, right := &ListNode{}, &ListNode{}
	curr, posLeft, posRight := head.Next, left, right
	for curr != nil {
		if curr.Val < pivot.Val {
			posLeft.Next = curr
			posLeft = posLeft.Next
		} else {
			posRight.Next = curr
			posRight = posRight.Next
		}
		curr = curr.Next
	}
	posLeft.Next = nil
	posRight.Next = nil
	pivot.Next = nil
	return left.Next, pivot, right.Next
}

// @lc code=end


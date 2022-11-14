/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := &Queue{}
	q.Push(root)
	root.Val = 0
	maxWidth := 0
	for !q.Empty() {
		size := len(*q)
		maxWidth = Max(maxWidth, (*q)[len(*q)-1].Val-(*q)[0].Val+1)
		for i := 0; i < size; i++ {
			node := q.Pop()
			if node.Left != nil {
				node.Left.Val = 2*node.Val + 1
				q.Push(node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 2
				q.Push(node.Right)
			}
		}
	}
	return maxWidth
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Queue []*TreeNode

func (q *Queue) Push(x *TreeNode) {
	*q = append(*q, x)
}

func (q *Queue) Pop() *TreeNode {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// @lc code=end


/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := &Queue{}
	q.Push(root)
	for !q.Empty() {
		size := len(*q)
		for i := 0; i < size; i++ {
			node := q.Pop()
			if i == size-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return ans
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


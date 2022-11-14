/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	needReverse := false
	q := &Queue{}
	q.Push(root)
	for !q.Empty() {
		size := len(*q)
		list := []int{}
		for i := 0; i < size; i++ {
			node := q.Pop()
			list = append(list, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		if needReverse {
			reverseList := []int{}
			for i := len(list) - 1; i >= 0; i-- {
				reverseList = append(reverseList, list[i])
			}
			list = reverseList
		}
		ans = append(ans, list)
		needReverse = !needReverse
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


/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next {
		node.Next = &Node{node.Val, node.Next, nil}
	}
	for node := head; node != nil; node = node.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = node.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}
	return newHead
}

// @lc code=end


/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return ser(root, "")
}

func ser(root *TreeNode, s string) string {
	if root == nil {
		return s + "#,"
	}
	s += strconv.Itoa(root.Val) + ","
	s += ser(root.Left, s) + ","
	s += ser(root.Right, s) + ","
	return s
}

var list []string

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list = strings.Split(data, ",")
	return deser(list[:len(list)-1])
}

func deser() *TreeNode {
	if list[0] == "#" {
		list = list[1:]
		return nil
	}
	root := &TreeNode{list[0], nil, nil}
	list = list[1:]
	root.Left = deser(list)
	root.Right = deser(list)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end


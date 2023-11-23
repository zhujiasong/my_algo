/*
589. N-ary Tree Preorder Traversal
Easy
Topics
Companies
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)
*/

package leetcode

// n-ary tree
type NNode struct {
	Val      int
	Children []*NNode
}

func preorder(root *NNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	ret = append(ret, root.Val)
	for _, node := range root.Children {
		ret = append(ret, preorder(node)...)
	}
	return ret
}

func preorderV2(root *NNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := []*NNode{root}
	for len(stack) != 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, t.Val)
		for i := len(t.Children) - 1; i >= 0; i-- {
			stack = append(stack, t.Children[i])
		}
	}

	return ret
}

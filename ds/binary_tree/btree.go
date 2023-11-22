package binary_tree

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// example input, pre-order, input manually
// 1 2 3 -1 4 -1 -1 5 -1 -1 6 7 -1 -1 8 9 -1 -1 -1
func CreataBTtree(root **TreeNode) {
	var data int
	fmt.Scanf("%d", &data)

	if data == -1 {
		*(root) = nil
	} else {
		(*root) = new(TreeNode)
		(*root).Val = data
		CreataBTtree(&((*root).Left))
		CreataBTtree(&((*root).Right))
	}
}

// example input, pre-order, init by array
// 1 2 3 -1 4 -1 -1 5 -1 -1 6 7 -1 -1 8 9 -1 -1 -1
var _iter_ int

func CreataBTtreeByArray(root **TreeNode, nums []int) {
	_iter_ = -1
	creataBTtreeByArray(root, nums)
}

func creataBTtreeByArray(root **TreeNode, nums []int) {
	_iter_++

	if nums[_iter_] == -1 {
		*(root) = nil
	} else {
		(*root) = new(TreeNode)
		(*root).Val = nums[_iter_]
		creataBTtreeByArray(&((*root).Left), nums)
		creataBTtreeByArray(&((*root).Right), nums)
	}
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	st := make([]*TreeNode, 0)
	st = append(st, root)

	for len(st) != 0 {
		t := st[len(st)-1]
		st = st[:len(st)-1]

		fmt.Printf("%d ", t.Val)

		if t.Right != nil {
			st = append(st, t.Right)
		}
		if t.Left != nil {
			st = append(st, t.Left)
		}
	}

	fmt.Printf("\n")
}

func MidOrder(root *TreeNode) {
	if root == nil {
		return
	}

	st := make([]*TreeNode, 0)
	t := root

	for len(st) != 0 || t != nil {
		for t != nil {
			st = append(st, t)
			t = t.Left
		}

		if len(st) != 0 {
			t = st[len(st)-1]
			st = st[:len(st)-1]

			fmt.Printf("%d ", t.Val)

			t = t.Right
		}
	}

	fmt.Printf("\n")
}

func LastOrder(root *TreeNode) {
	if root == nil {
		return
	}

	st := make([]*TreeNode, 0)
	t := root
	var flg *TreeNode

	for len(st) != 0 || t != nil {
		for t != nil {
			st = append(st, t)
			t = t.Left
		}

		if len(st) != 0 {
			t = st[len(st)-1]
			if t.Right != nil && t.Right != flg {
				t = t.Right
			} else {
				st = st[:len(st)-1]
				fmt.Printf("%d ", t.Val)
				flg = t
				t = nil
			}
		}
	}

	fmt.Printf("\n")
}

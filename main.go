package main

import (
	"fmt"
	"my_algo/algo/leetcode"
	tree "my_algo/ds/binary_tree"
)

func main() {
	// **** list ****/
	// l := list.New([]interface{}{1, 2, 3, 4, 5, 6, 7})
	// l.AddNodeHead(&list.ListNode{
	// 	Elem: 8,
	// })

	// l.DeleteNode(l.Tail())
	// l.Print()
	// l.PrintReverse()

	// **** lru cache ****/
	//cache := algo.Constructor(1)
	//cache.Put(1, 11)
	// cache.Put(2, 1)
	// fmt.Println(cache.Get(2))
	// cache.Put(3, 33)
	// fmt.Println(cache.Get(2))
	// cache.Put(4, 44)
	// fmt.Println(cache.Get(1))
	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.Get(4))

	// **** lru cache ****/
	// 380
	// set := leetcode.Constructor()
	// fmt.Println(set.Insert(1))
	// fmt.Println(set.Remove(2))
	// fmt.Println(set.Insert(2))
	// fmt.Println(set.GetRandom())
	// fmt.Println(set.Remove(1))
	// fmt.Println(set.Insert(2))
	// fmt.Println(set.GetRandom())

	// **** binary search ****/
	// fmt.Println(algo.BinarySearchLeft([]int{1, 3, 5, 7, 7, 7, 7, 8, 9}, 7))

	// **** leetcode 567 ****/
	//fmt.Println(leetcode.CheckInclusion("ao", "eidboaoo"))

	// **** leetcode 109 ****/
	// l := list.New([]interface{}{-10, -3, 0, 5, 9})
	// root := leetcode.SortedListToBST(l.Front())
	// tree.PreOrder(root)

	// **** leetcode BTree ****/
	var root *tree.TreeNode
	//initArray := []int{5, 5, -1, -1, 3, 2, -1, -1, 2, -1, -1}
	initArray := []int{5, 4, 11, 7, -1, -1, 2, -1, -1, -1, 8, 13, -1, -1, 4, 5, -1, -1, 1, -1, -1}
	tree.CreataBTtreeByArray(&root, initArray)
	fmt.Println(leetcode.PathSum(root, 22))
}

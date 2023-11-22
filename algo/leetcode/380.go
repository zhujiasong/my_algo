/*
380. Insert Delete GetRandom O(1)
Medium
8.2K
465
Companies
Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.
*/

package leetcode

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m map[int]int
	s []int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		s: make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	index := s.getIndex(val)
	if index != -1 {
		return false
	}

	s.s = append(s.s, val)
	s.m[val] = len(s.s) - 1

	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	index := s.getIndex(val)
	if index == -1 {
		return false
	}

	endIndex := len(s.s) - 1
	s.m[s.s[endIndex]] = index
	s.s[index], s.s[endIndex] = s.s[endIndex], s.s[index]
	s.s = s.s[:endIndex]
	delete(s.m, val)

	return true
}

func (s *RandomizedSet) GetRandom() int {
	if len(s.s) == 0 {
		return -1
	}

	index := rand.Int() % len(s.s)
	return s.s[index]
}

func (s *RandomizedSet) getIndex(key int) int {
	index, ok := s.m[key]
	if !ok {
		return -1
	}
	return index
}

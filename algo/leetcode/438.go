/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

package leetcode

func FindAnagrams(s2 string, s1 string) []int {
	ret := make([]int, 0)
	if len(s1) == len(s2) && len(s1) == 0 {
		return []int{}
	}
	if len(s1) > len(s2) {
		return []int{}
	}

	var vaild int
	windowS1 := make(map[byte]int, len(s1))
	windowS2 := make(map[byte]int, len(s1))
	for _, b := range s1 {
		windowS1[byte(b)]++
	}

	left, right := 0, 0
	for right < len(s2) {
		c := s2[right]
		right++

		if _, ok := windowS1[c]; ok {
			windowS2[c]++
			if windowS1[c] == windowS2[c] {
				vaild++
			}
		}

		if right < len(s1) {
			continue
		}
		if vaild == len(windowS1) {
			ret = append(ret, left)
		}

		d := s2[left]
		left++
		if _, ok := windowS1[d]; ok {
			if windowS1[d] == windowS2[d] {
				vaild--
			}
			windowS2[d]--
		}
	}

	return ret
}

/*
20. Valid Parentheses
Easy
Topics
Companies
Hint
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package leetcode

func isValid(s string) bool {
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}

	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]byte, len(s))
	top := -1
	for _, b := range s {
		if top == -1 {
			top++
			stack[top] = byte(b)
			continue
		}

		ele := stack[top]
		if m[ele] == byte(b) {
			top--
		} else {
			top++
			stack[top] = byte(b)
		}
	}

	return top == -1
}

package algo

import (
	"fmt"
	"sort"
	"strconv"
)

func SortStr(strs []string) {
	sli := strSlice(strs)
	sort.Sort(sli)
}

type strSlice []string

func (x strSlice) Len() int {
	return len(x)
}

func (x strSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x strSlice) Less(i, j int) bool {
	return compare(x[j], x[i])
}

func compare(s1, s2 string) bool {
	idx1, idx2 := 0, 0
	var numStr1, numStr2 string

	for idx1 < len(s1) && idx2 < len(s2) {
		if isNumChar(s1[idx1]) && isNumChar(s2[idx2]) {
			for idx1 < len(s1) && isNumChar(s1[idx1]) {
				numStr1 = fmt.Sprintf("%s%c", numStr1, s1[idx1])
				idx1++
			}
			for idx2 < len(s2) && isNumChar(s2[idx2]) {
				numStr2 = fmt.Sprintf("%s%c", numStr2, s2[idx2])
				idx2++
			}

			num1, _ := strconv.Atoi(numStr1)
			num2, _ := strconv.Atoi(numStr2)
			if num1 != num2 {
				return num1 > num2
			}
			numStr1, numStr2 = "", ""

		} else {
			if s1[idx1] != s2[idx2] {
				return s1[idx1] > s2[idx2]
			}
			idx1++
			idx2++
		}
	}

	if idx2 == len(s2) {
		return false
	}

	return true
}

func isNumChar(b byte) bool {
	return b >= '0' && b <= '9'
}

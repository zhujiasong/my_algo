package leetcode

func getIndex(sli []int, val int) int {
	for i, v := range sli {
		if v == val {
			return i
		}
	}
	return -1
}

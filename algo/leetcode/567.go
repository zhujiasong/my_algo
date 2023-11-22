package leetcode

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) == len(s2) && len(s1) == 0 {
		return true
	}
	if len(s1) > len(s2) {
		return false
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
			return true
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

	return false
}

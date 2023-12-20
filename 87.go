package leetcode_go

func isScramble(s1 string, s2 string) bool {
	m := make(map[string]bool)

	startIdx := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			startIdx = i
			break
		}

	}

	endIdx := 0
	for i := len(s1) - 1; i >= 0; i-- {
		if s1[i] != s2[i] {
			endIdx = i
			break
		}
	}
	if process87(s1[startIdx:endIdx+1], s2[startIdx:endIdx+1], m) {
		return true
	}
	return process87(s1, s2, m)
}

// s1=eebaacbcbcadaaedceaaacadcwdabc
// s2=eadcaacabaddaceacbceaabecsdcba
func process87(s1 string, s2 string, mm map[string]bool) bool {
	if len(s1) == 1 {
		if s1 == s2 {
			return true
		} else {
			return false
		}
	}

	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	// 加一个看 两边的字母数 是否一致
	ints := make(map[uint8]int)

	for i := 0; i < len(s1); i++ {
		ints[s1[i]]++
		ints[s2[i]]--
	}

	for _, v := range ints {
		if v != 0 {
			mm[s1+"_"+s2] = false
			return false
		}
	}

	if v, ok := mm[s1+"_"+s2]; ok {
		return v
	}

	for i := 1; i < len(s1); i++ {
		// 没有交换的场景下 左侧
		if process87(s1[:i], s2[0:i], mm) && process87(s1[i:], s2[i:], mm) || process87(s1[i:], s2[:len(s2)-i], mm) && process87(s1[:i], s2[len(s2)-i:], mm) {
			mm[s1+"_"+s2] = true
			return true
		}

	}
	mm[s1+"_"+s2] = false
	return false

}

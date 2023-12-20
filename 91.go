package leetcode_go

func numDecodings(s string) int {
	mm := make(map[int]int)
	return process91(s, 0, mm)
}

func process91(s string, idx int, mm map[int]int) int {
	n := len(s)

	if val, found := mm[idx]; found {
		return val
	}

	// 如果当前索引已经到达字符串末尾，意味着找到了一种解码方式
	if idx == n {
		return 1
	}

	// 如果当前字符是 '0'，则无法解码
	if s[idx] == '0' {
		return 0
	}

	// 只解码一个字符
	way := process91(s, idx+1, mm)
	// 如果还有足够的空间解码两个字符，并且两个字符组成的数字小于或等于26
	if idx < n-1 && ((s[idx] == '1') || (s[idx] == '2' && s[idx+1] < '7')) {
		way += process91(s, idx+2, mm)
	}
	mm[idx] = way
	return way
}

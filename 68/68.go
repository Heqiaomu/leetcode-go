package _8

func fullJustify(words []string, maxWidth int) []string {

	//对整个数组 进行分割
	res := make([]string, 0)
	first := 0
	lastLen := 0

	for i := range words {
		tmp := lastLen + len(words[i])
		if tmp == maxWidth { // 恰好等于
			res = append(res, put(words[first:i+1], maxWidth))

			if i == len(words)-1 {
				return res
			}

			first = i + 1
			lastLen = 0
		} else if tmp > maxWidth {
			res = append(res, put(words[first:i], maxWidth))
			first = i
			lastLen = len(words[i]) + 1
		} else {
			lastLen = tmp + 1
		}
	}
	res = append(res, last(words[first:], maxWidth))
	return res
}

// 先将给定的字符的按照 空格 平凑下
func put(words []string, maxWidth int) string {
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0] + space(maxWidth-len(words[0]))
	}
	// 先计算当前的单词的总数在多少
	var wordLen int
	var res string
	for _, word := range words {
		wordLen += len(word)
	}

	// 在看看我需要多少个 空格
	totalSpace := maxWidth - wordLen
	countSpace := totalSpace / (len(words) - 1) // 先看看每个位置放多少个 空格
	count := totalSpace % (len(words) - 1)      // 每个位置放置后，剩余多少个，然后在在之前的位置中，每个位置在放一个

	for i := 0; i < len(words)-1; i++ {
		res += words[i] + space(countSpace)
		if count > 0 {
			res += " "
			count--
		}
	}
	res = res + words[len(words)-1]
	return res

}

func last(words []string, maxWidth int) string {
	if len(words) < 0 {
		return ""
	}
	// var wordLen int
	var res string
	for _, word := range words {
		res = res + word + " "
	}

	res = res + space(maxWidth-len(res))
	return res[:maxWidth]

}
func space(n int) string {
	var sp string
	for i := 0; i < n; i++ {
		sp += " "
	}
	return sp
}

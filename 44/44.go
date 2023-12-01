package _4

/*
给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符序列（包括空字符序列）。
判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。
*/
func isMatch(s string, p string) bool {
	return process2(s, p)
}

// 递归的做法
func process(s string, p string, i, j int) bool {
	if i > len(s) {
		return false
	}

	if j > len(p) {
		return false
	}

	if i != len(s) && j == len(p) {
		return false
	}

	if i == len(s) && j != len(p) {
		for k := j; k < len(p); k++ {
			if p[k] != '*' {
				return false
			}
		}
	}
	if i == len(s) && j == len(p) {
		return true
	}

	if p[j] == '?' {
		return process(s, p, i+1, j+1)
	}

	if p[j] == '*' {
		if !process(s, p, i+1, j+1) {
			if !process(s, p, i, j+1) {
				return process(s, p, i+1, j)
			}
		}
		return true
	}
	if p[j] != '?' && p[j] != '*' {
		if s[i] == p[j] {
			return process(s, p, i+1, j+1)
		}
		return false
	}
	return false
}

// 动态规划
func process2(s string, p string) bool {
	dp := make([][]bool, len(s)+1) // 确定行
	for k := range dp {
		dp[k] = make([]bool, len(p)+1) // 确定列
	}
	dp[len(s)][len(p)] = true

	// 初始化dp数组处理模式p末尾的 *
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != '*' {
			break
		}
		dp[len(s)][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if p[j] == '*' {
				dp[i][j] = dp[i+1][j] || dp[i][j+1]
			} else if p[j] == '?' || s[i] == p[j] {
				dp[i][j] = dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

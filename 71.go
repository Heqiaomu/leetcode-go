package leetcode_go

import (
	"path/filepath"
)

// "/home/" , "/../" , "/home//foo/"
func simplifyPath(path string) string {

	abs, _ := filepath.Abs(path)

	return abs
	//
	//split := strings.Split(path, "/")
	//
	//var stack []string
	//
	//for _, part := range split {
	//	switch part {
	//	case "", ".": // 忽略空字符串和当前目录标记
	//		continue
	//	case "..": // 返回上级目录
	//		if len(stack) > 0 {
	//			stack = stack[:len(stack)-1]
	//		}
	//	default: // 正常的目录名
	//		stack = append(stack, part)
	//	}
	//}
	//return "/" + strings.Join(stack, "/")
}

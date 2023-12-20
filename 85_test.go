package leetcode_go

import "testing"

func Test_maximalRectangle(t *testing.T) {
	data := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	maximalRectangle(data)
}

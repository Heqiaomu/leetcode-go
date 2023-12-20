package leetcode_go

import "bytes"

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

func exist(board [][]byte, word string) bool {

	if len(word) == 0 {
		return false
	}

	if len(word) == 1 {

	}
}

func findByte(board [][]byte, wd []byte) (bool, []int) {
	// 先横着
	for i := range board {
		index := bytes.Index(board[i][0:len(board[i])], wd)
		if index >= 0 {
			return true, []int{i, i + len(wd) - 1}
		}
	}
	// 在竖着

}

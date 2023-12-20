package leetcode_go

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	solveSudoku(board)

	fmt.Println(board)
}

// solveSudoku
func solveSudoku(board [][]byte) {

	var nums []map[int]struct{}
	for i := 0; i < 9; i++ {
		nums = append(nums, make(map[int]struct{}))
	}

	// 存储每个位置，当前可能值选项
	m := make(map[int]map[byte]struct{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				v := values(board, i, j)
				m[key(i, j)] = v
				nums[len(v)-1][key(i, j)] = struct{}{}
			}
		}
	}

	solveSudokuInit(board, nums, m)

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			boxes[i/3][j/3][index]++
		}
	}
	matrix := make([][]byte, 9)
	for i := range matrix {
		matrix[i] = make([]byte, 9)
	}

	for i, data := range board {
		for j, datum := range data {
			matrix[i][j] = datum
		}
	}

	process, _ := sdProcess(matrix, m, 0, 0)
	printBoardAsChars(process)

	for i, data := range process {
		for j, datum := range data {
			board[i][j] = datum
		}
	}

	return
}

var count int

func sdProcess(tmp [][]byte, m map[int]map[byte]struct{}, i, j int) ([][]byte, bool) {
	count++
	if i == 9 {
		return tmp, true
	}

	var nexti, nextj int
	if j == 8 {
		nexti = i + 1
		nextj = 0
	} else {
		nexti = i
		nextj = j + 1
	}

	_, ok := m[key(i, j)]
	if !ok {
		return sdProcess(tmp, m, nexti, nextj)
	}

	for c, _ := range m[key(i, j)] {
		if UpdateValid(i, j, c) {
			tmp[i][j] = c
			t, b := sdProcess(tmp, m, nexti, nextj)
			if !b {
				tmp[i][j] = '.'
				LookBack(i, j, c)
				continue
			} else {
				return t, true
			}
		}
	}
	return tmp, false
}

func UpdateValid(i, j int, c byte) bool {
	index := c - '1'
	if rows[i][index] == 1 || columns[j][index] == 1 || boxes[i/3][j/3][index] == 1 {
		return false
	}

	rows[i][index]++
	columns[j][index]++
	boxes[i/3][j/3][index]++

	return true
}

var rows, columns [9][9]int
var boxes [3][3][9]int

func LookBack(i, j int, c byte) {
	index := c - '1'
	rows[i][index]--
	columns[j][index]--
	boxes[i/3][j/3][index]--
}

func solveSudokuInit(board [][]byte, nums []map[int]struct{}, m map[int]map[byte]struct{}) {
	for {
		k := -1
		for k, _ = range nums[0] {
			break
		}
		if k == -1 {
			break
		}

		i, j := keyParse(k)

		if len(m[k]) != 1 {
			panic("f")
		}
		for mKey, _ := range m[k] {
			board[i][j] = mKey
			delete(nums[0], k)
			delete(m, k)

			break
		}

		d := board[i][j]
		for a := 0; a < 9; a++ {
			update(board, nums, m, d, key(i, a))

		}
		for b := 0; b < 9; b++ {
			update(board, nums, m, d, key(b, j))
		}

		row := i / 3 * 3
		column := j / 3 * 3
		for a := 0; a < 3; a++ {
			for b := 0; b < 3; b++ {
				update(board, nums, m, d, key(row+a, column+b))
			}
		}
	}

	return
}

func key(i, j int) int {
	return i*10 + j
}

func keyParse(d int) (int, int) {
	i := d / 10
	j := d % 10
	return i, j
}

func update(board [][]byte, nums []map[int]struct{}, m map[int]map[byte]struct{}, d byte, k int) {
	v := m[k]
	if v == nil {
		return
	}
	_, ok := v[d]
	if ok {
		count := len(v)

		delete(nums[count-1], k)

		delete(v, d)

		nums[count-2][k] = struct{}{}
	}
	return
}

func values(board [][]byte, i, j int) map[byte]struct{} {
	m := map[byte]struct{}{}
	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			m[board[i][k]] = struct{}{}
		}
	}

	for k := 0; k < 9; k++ {
		if board[k][j] != '.' {
			m[board[k][j]] = struct{}{}
		}
	}

	row := i / 3 * 3
	column := j / 3 * 3

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if board[row+a][column+b] != '.' {
				m[board[row+a][column+b]] = struct{}{}
			}
		}
	}

	mm := map[byte]struct{}{}
	for n := '1'; n <= '9'; n++ {
		_, ok := m[byte(n)]
		if !ok {
			mm[byte(n)] = struct{}{}
		}
	}
	return mm
}

func printBoardAsChars(board [][]byte) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%c   ", val)
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------------------------")
}

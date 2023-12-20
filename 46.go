package leetcode_go

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return [][]int{}
	}
	process46(nums, []int{}, &result)
	return result
}

func process46(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
	}
	// 1 2 3 4 5 6
	for idx, num := range nums {
		next := append([]int{}, nums[:idx]...)
		next = append(next, nums[idx+1:]...)          // 选择 idx 位置的数 2，3，4，5，6
		process46(next, append(current, num), result) // 选择 1 进入当前的 current
	}
}

// 1 2 3 4 5 6

// 1 2 3 4 5 6, 2 1 3 4 5 6, 2 3 1 4 5 6, 2 3 4 1 5 6, 2 3 4 5 1 6, 2 3 4 5 6 1
// 2 1 3 4 5 6, 1 2 3 4 5 6, 1 3 2 4 5 6, 1 3 4 2 5 6, 1 3 4 5 2 6, 1 3 4 5 6 2
// 3 1 2 4 5 6, 1 3 2 4 5 6, 1 2 3 4 5 6, 1 2 4 3 5 6, 1 2 4 5 3 6, 1 2 4 5 6 3
// 4 1 2 3 5 6, 1 4 2 3 5 6, 1 2 4 3 5 6, 1 2 3 4 5 6, 1 2 3 5 4 6, 1 2 3 5 6 4
// 5 1 2 3 4 6, 1 5 2 3 4 6, 1 2 5 3 4 6, 1 2 3 5 4 6, 1 2 3 4 5 6, 1 2 3 4 6 5
// 6 1 2 3 4 5, 1 6 2 3 4 5, 1 2 6 3 4 5, 1 2 3 6 4 5, 1 2 3 4 6 5, 1 2 3 4 5 6

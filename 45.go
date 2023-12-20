package leetcode_go

func jump(nums []int) int {
	var steps int = 0    // 当前步数
	var position int = 0 // 每次都要跳，记录跳过去之后的位置，记录一个最大的下标位置

	var end int = 0 // 记录我当前能到的最远的点

	if len(nums) < 2 {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		position = max(position, i+nums[i])
		if i == end {
			steps++
			end = position
		}

	}
	return steps
}

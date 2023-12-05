package _5

func canJump(nums []int) bool {
	maxPos := 0
	for i, num := range nums {
		if i > maxPos {
			return false
		}
		if i+num > maxPos {
			maxPos = i + num
		}
		if maxPos >= len(nums)-1 {
			return true
		}
	}
	return maxPos >= len(nums)-1
}

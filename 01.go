package leetcode_go

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}
	return res
}

package leetcode_go

// [-2,1,-3,4,-1,2,1,-5,4]
func maxSubArray(nums []int) int {
	var curVal, maxNum int
	for _, num := range nums {
		if curVal < 0 {
			curVal = num
		} else {
			curVal += num
		}

		maxNum = max(curVal, num)

	}
	return maxNum
}

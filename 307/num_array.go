package numarray

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) Update(index int, val int) {
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return sum(this.nums[left:right])
}

func sum(nums []int) int {
	// 区分是否是偶数
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 0 {
		return 0
	}

	return sum(nums[1 : len(nums)-2])

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

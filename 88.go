package leetcode_go

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {

	}

	longNums := []int{}
	shortNums := []int{}

	if m >= n {
		longNums = nums1
		shortNums = nums2
	} else {
		longNums = nums2
		shortNums = nums1
	}
	ints := make([]int, m+n)
	lIdx := 0
	sIdx := 0
	rIdx := 0
	for rIdx < max(m, n) {
		if longNums[lIdx] < shortNums[sIdx] {
			ints[rIdx] = longNums[lIdx]
			rIdx++
			lIdx++
		} else if longNums[lIdx] < shortNums[sIdx] {
			ints[rIdx] = shortNums[lIdx]
			rIdx++
			sIdx++
		} else {
			ints[rIdx] = shortNums[lIdx]
			rIdx++
			lIdx++
			sIdx++
		}
	}

	nums1 = ints
}

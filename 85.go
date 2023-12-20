package leetcode_go

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	row := make([]int, n)

	maxV := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				row[j] += 1
			} else {
				row[j] = 0
			}
		}
		maxV = max(maxV, largestRectangleArea(row))
	}
	return maxV
}

//
//func largestRectangleArea(heights []int) int {
//	stack := []int{}
//	maxArea := 0
//
//	for i, h := range heights {
//		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
//			height := heights[stack[len(stack)-1]]
//			stack = stack[:len(stack)-1]
//			width := i
//			if len(stack) > 0 {
//				width = i - stack[len(stack)-1] - 1
//			}
//			maxArea = max(maxArea, height*width)
//		}
//		stack = append(stack, i)
//	}
//
//	for len(stack) > 0 {
//		height := heights[stack[len(stack)-1]]
//		stack = stack[:len(stack)-1]
//		width := len(heights)
//		if len(stack) > 0 {
//			width = len(heights) - stack[len(stack)-1] - 1
//		}
//		maxArea = max(maxArea, height*width)
//	}
//
//	return maxArea
//}

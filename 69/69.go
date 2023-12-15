package _9

import "math"

// f(x) = x2 -C
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 第一轮fuzhi

	// 令 x0 = C
	C, x0 := float64(x), float64(x)

	for {

		xi := 0.5 * (x0 + C/x0)

		if math.Abs(xi-x0) < 1e-7 {
			break
		}

		x0 = xi
	}
	return int(x0)
}

/**

f(x) = x ^ 2 -C
f`(x) = 2x
所以

(x0 ^ 2 -C - 0 ) / (x0 - x1) = 2 * x0

所以 x1 = 1/2 * (x0 + c / x0)



*/

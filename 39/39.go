package _39

type res struct {
	seles [][]int
}

func combinationSum(candidates []int, target int) [][]int {
	r := new(res)
	combinationSumProcess(r, candidates, target, 0, []int{}, 0)
	return r.seles
}

/*
 *
 */
func combinationSumProcess(r *res, candidates []int, target int, calTotal int, sel []int, postion int) {
	if target == calTotal {
		selTmp := make([]int, len(sel))
		copy(selTmp, sel)
		r.seles = append(r.seles, selTmp)
		return
	}

	if len(candidates) == postion || len(candidates) == 0 {
		return
	}

	if calTotal > target {
		return
	}

	// 这里看看，剩下来的需要计算多少次 01,02,4 target=10
	// 如果是 01 的话，我可以计算 10 次 先看第一个 01
	count := (target - calTotal) / candidates[postion]

	for i := 0; i <= count; i++ {
		var tmpSel []int = sel
		for j := 0; j < i; j++ {
			tmpSel = append(tmpSel, candidates[postion])
		}
		combinationSumProcess(r, candidates, target, calTotal+candidates[postion]*i, tmpSel, postion+1)
	}
}

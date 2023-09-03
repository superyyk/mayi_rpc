package tool

func GetMinNum(ary []int) int {
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal > ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal
}

func GetMaxNum(ary []int) int {
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal < ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal
}

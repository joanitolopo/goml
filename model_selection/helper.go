package modelselection

func Arange(vals ...int) []int {
	var start, stop int = 0, 0
	var result []int

	switch len(vals) {
	case 1:
		start = vals[0]
		for i := 0; i < int(start); i++ {
			result = append(result, i)
		}
	case 2:
		start = vals[0]
		stop = vals[1]
		for i := start; i < stop; i++ {
			result = append(result, i)
		}
	}

	return result

}

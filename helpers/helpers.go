package helpers

func filter(ss []int, test func(int) bool) (ret []int) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func indexOf(val string, ss []string) int {
	for i, s := range ss {
		if s == val {
			return i
		}
	}
	return -1 //not found.
}

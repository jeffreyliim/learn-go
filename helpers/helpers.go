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

func hasOwnProperty(m map[string]string, p string) bool {
	for i := range m {
		if i == p {
			return true
		}
	}
	return false
}

//	push to back of array
func push(arr []string, str string) []string {
	arr = append(arr, str)
	return arr
}

//	Push to front of array
func unshift(arr []string, str string) []string {
	arr = append([]string{str}, arr...)
	return arr
}

//	remove front of array
func shift(arr []string, str string) []string {
	if len(arr) < 1 {
		arr = arr[:0]
	} else {
		arr = arr[1:]
	}
	return arr
}

// remove from back of array
func pop(arr []string, str string) []string {
	if len(arr) < 1 {
		arr = arr[:0]
	} else {
		arr = arr[0 : len(arr)-1]
	}
	return arr
}

// remove duplicates in array
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	Solution(arr, 4)
}

func Solution(A []int, K int) []int {

	A = filter(A, func(s int) bool {
		return s < 3
	})
	fmt.Print(A)
	return A
}

func filter(strArr []int, test func(int) bool) (ret []int) {
	for _, element := range strArr {
		if test(element) {
			ret = append(ret, element)
		}
	}
	return ret
}

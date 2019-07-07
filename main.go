package main

func main() {

	Solution([]int{0, 1, 1, 0})
}

func Solution(A []int) int {
	counter := 0

	for i := range A {
		if i < len(A)-1 {
			if A[i] == A[i+1] {
				counter++
				if A[len(A)-1] != A[len(A)-2] {
					counter++
				}
			}
		}
	}
	return counter
}

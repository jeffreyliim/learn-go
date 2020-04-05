package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

//	Given two strings, determine if they share a common substring. A substring may be as small as one character.
//
//	For example, the words "a", "and", "art" share the common substring . The words "be" and "cat" do not share a substring.
//
//	Function Description
//
//	Complete the function twoStrings in the editor below. It should return a string, either YES or NO based on whether the strings share a common substring.
//
//	twoStrings has the following parameter(s):
//
//	s1, s2: two strings to analyze .
//	Input Format
//
//	The first line contains a single integer , the number of test cases.
//
//	The following  pairs of lines are as follows:
//
//	The first line contains string .
//	The second line contains string .
//	Constraints
//
// 	and  consist of characters in the range ascii[a-z].
//	Output Format
//
//	For each pair of strings, return YES or NO.
//
//	Sample Input
//
//	2
//	hello
//	world
//	hi
//	world
//	Sample Output
//
//	YES
//	NO

//func twoStrings(s1 string, s2 string) string {
//	s1 = strings.ToLower(s1)
//	s2 = strings.ToLower(s2)
//
//	m := make(map[string]int)
//	for i := range s1 {
//		if indexOf(m, string(s1[i])) == "" {
//			m[string(s1[i])] = 1
//		}
//	}
//
//	for i := range s2 {
//		if indexOf(m, string(s2[i])) == string(s2[i]) {
//			m[string(s2[i])]++
//		}
//	}
//
//	status := ""
//	for i := range m {
//		if m[i] > 1 {
//			status = "YES"
//			break
//		} else {
//			status = "NO"
//		}
//	}
//	return status
//}

//func indexOf(str map[string]int, char string) string {
//	for i := range str {
//		if i == char {
//			return i
//		}
//	}
//	return ""
//}

func findOcurrences(text string, first string, second string) []string {
	arrT := strings.Split(text, " ")
	var s []string
	firstString := ""
	secondString := ""
	if index1 := indexOf(arrT, first); index1 != -1 {
		firstString = arrT[index1+2]
		arrT = arrT[index1+2:]
	}
	if index2 := indexOf(arrT, second); index2 != -1 {
		secondString = arrT[index2+1]
	}
	s = append(s, firstString)
	s = append(s, secondString)

	return s
}

func indexOf(arr []string, str string) int {
	for i, element := range arr {
		if element == str {
			return i
		}
	}
	return -1
}

//	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
//	Example 1:
//
//	Input: 121
//	Output: true
//	Example 2:
//
//	Input: -121
//	Output: false
//	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//	Example 3:
//
//	Input: 10
//	Output: false
//	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

func isPalindrome(x int) bool {
	newX := strconv.Itoa(x)
	buf := bytes.Buffer{}

	for i := len(newX) - 1; i >= 0; i-- {
		buf.WriteString(string(newX[i]))
	}

	if buf.String() == newX {
		return true
	}

	return false
}

func longestCommonPrefix(strs []string) string {
	//buf := bytes.Buffer{}
	//for i := range strs {
	//	for ch := range strs[i] {
	//
	//	}
	//}
	//return buf
	s := make(interface{})
	s = string(1)
	fmt.Println(&s)

	str := "asdasdasdas"
	fmt.Println(str[3])
	//res := strings.Index(str, "s")
	fmt.Println(str[:len(str)-1])
	//fmt.Println(res)
	//fmt.Println(string(str[0]))
	return "asd"
}

func main() {
	//twoStrings("hackerrankcommunity", "cdecdecdecde")
	//findOcurrences("we will we will rock you", "we", "will")
	//isPalindrome(10)
	longestCommonPrefix([]string{"flower", "flow", "flight"})
}

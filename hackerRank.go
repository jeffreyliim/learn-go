package main

import (
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

func twoStrings(s1 string, s2 string) string {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	m := make(map[string]int)
	for i := range s1 {
		if indexOf(m, string(s1[i])) == "" {
			m[string(s1[i])] = 1
		}
	}

	for i := range s2 {
		if indexOf(m, string(s2[i])) == string(s2[i]) {
			m[string(s2[i])]++
		}
	}

	status := ""
	for i := range m {
		if m[i] > 1 {
			status = "YES"
			break
		} else {
			status = "NO"
		}
	}
	return status
}

func indexOf(str map[string]int, char string) string {
	for i := range str {
		if i == char {
			return i
		}
	}
	return ""
}



func main() {
	twoStrings("hackerrankcommunity", "cdecdecdecde")
}

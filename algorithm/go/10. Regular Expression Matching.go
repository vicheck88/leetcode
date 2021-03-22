// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:
// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

// Example 1:
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
// Example 2:
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
// Example 3:
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
// Example 4:
// Input: s = "aab", p = "c*a*b"
// Output: true
// Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
// Example 5:
// Input: s = "mississippi", p = "mis*is*p*."
// Output: false

// Constraints:
// 0 <= s.length <= 20
// 0 <= p.length <= 30
// s contains only lowercase English letters.
// p contains only lowercase English letters, '.', and '*'.
// It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.

package main

func isMatch(s string, p string) bool {
	strLength := len(s)
	pattLength := len(p)
	var arr [30][40]bool
	arr[strLength][pattLength] = true
	for i := strLength; i >= 0; i-- {
		for j := pattLength - 1; j >= 0; j-- {
			firstMatch := i < strLength && (p[j] == '.' || p[j] == s[i])
			if j+1 < pattLength && p[j+1] == '*' {
				arr[i][j] = (firstMatch && arr[i+1][j]) || arr[i][j+2]
			} else {
				arr[i][j] = firstMatch && arr[i+1][j+1]
			}
		}
	}
	return arr[0][0]
}

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	var str string
// 	var pattern string
// 	cnt := 0
// 	for {
// 		cnt++
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		if cnt&1 == 1 {
// 			str = string(line)
// 		} else {
// 			pattern = string(line)
// 			fmt.Println(isMatch(str, pattern))
// 		}
// 	}
// }

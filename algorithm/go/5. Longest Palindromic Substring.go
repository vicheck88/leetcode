// Given a string s, return the longest palindromic substring in s.
// Example 1:
// Input: s = "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// Example 2:
// Input: s = "cbbd"
// Output: "bb"
// Example 3:
// Input: s = "a"
// Output: "a"
// Example 4:
// Input: s = "ac"
// Output: "a"
// Constraints:
// 1 <= s.length <= 1000
// s consist of only digits and English letters (lower-case and/or upper-case)

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// func checkPalindrom(s string) bool {
// 	i := 0
// 	j := len(s) - 1
// 	for {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 		if i > j {
// 			return true
// 		}
// 	}
// }
// func findLongestPalindrome(maxLen int, s string) string {
// 	res := ""
// 	maxIdx := len(s) - maxLen
// 	for j := 0; j < maxIdx; j++ {
// 		if checkPalindrom(s[j:]) {
// 			res = s[j:]
// 			break
// 		}
// 	}
// 	return res
// }

// func longestPalindrome(s string) string {
// 	maxLen := 0
// 	maxStr := ""
// 	for i := 0; i < len(s); i++ {
// 		str := findLongestPalindrome(maxLen, s[:i+1])
// 		if len(str) > maxLen {
// 			maxLen = len(str)
// 			maxStr = str
// 		}
// 	}
// 	return maxStr
// }

//Manacher Algorithm
//link: https://www.crocus.co.kr/1075
func longestPalindrome(s string) string {
	var newStrBuf bytes.Buffer
	for i := 0; i < len(s); i++ {
		newStrBuf.WriteString("|" + string(s[i]))
		//문자의 사이사이마다 특수 문자 삽입: 각 문자+문자의 사이에서 루프를 시작해야 함
		//문자의 사이에 관련로직이 없을경우 bb와 같은 문자는 찾을 수 없음
	}
	newStrBuf.WriteString("|")
	newStr := newStrBuf.String()
	r := -1                            //찾아낸 가장 긴 Palindrome의 오른쪽 끝 index
	p := -1                            //찾아낸 가장 긴 Palindrome의 가운데 index
	strArr := make([]int, len(newStr)) //각 index를 중심으로 하는 Palindrome의 반지름 크기
	for i := 0; i < len(newStr); i++ {
		if r < i { //i의 위치가 관측된 가장 긴 Palindrome의 바깥에 있을 경우: 0
			strArr[i] = 0
		} else {
			iprime := 2*p - i //현재 가장 긴 Palindrome에서 가운데를 기준으로 i와 대칭되는 index(왼쪽)
			if strArr[iprime] < r-i {
				strArr[i] = strArr[iprime]
				//iprime을 중심으로 하는 가장 긴 Palindrome 값이 현재 가장 긴 Palendrome 안에 있을 경우
			} else {
				strArr[i] = r - i
				//iprime을 중심으로 하는 가장 긴 Palindrome 값이 현재 가장 긴 Palendrome 밖에 있을 경우
			}
		}
		for i+strArr[i] < len(newStr)-1 && i-strArr[i] > 0 {
			if newStr[i+strArr[i]+1] == newStr[i-strArr[i]-1] {
				strArr[i]++ //범위를 하나씩 늘리며 가장 긴 Palindrome 조회
			} else {
				break
			}
		}
		if p < 0 || strArr[i] > strArr[p] { //현재 index에서 구한 Palindrome 크기가 기존 것보다 클 경우
			p = i
			r = strArr[i] + i
		}
	}
	res := newStr[p-strArr[p] : strArr[p]+1+p]
	return strings.ReplaceAll(res, "|", "")
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		str := string(line)
		fmt.Println(longestPalindrome(str))
	}
}

// Given a string s, find the length of the longest substring without repeating characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// Example 4:
// Input: s = ""
// Output: 0
// Constraints:
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

package main

func lengthOfLongestSubstring(s string) int {
	strMap := make(map[byte]int, 1000) //key: char 값, value: char 위치한 idx
	maxLen := 0                        //최대 subseq 크기: 결과
	startIdx := 0                      //subseq 시작 index
	for i := 0; i < len(s); i++ {
		_, exists := strMap[s[i]] //현재 index값이 map에 저장되어있는지(이미 존재하는 값인지) 확인
		if exists {
			if startIdx <= strMap[s[i]] {
				startIdx = strMap[s[i]] + 1
				//이미 존재하는 문자가 startIdx보다 클경우 startIdx를 존재 index 다음값으로 설정
				//겹치는 문자 다음부터 subseq를 계산
			}
		}
		strMap[s[i]] = i //현재 index 설정
		if maxLen < i-startIdx+1 {
			maxLen = i - startIdx + 1 //최대크기값 업데이트
		}
	}
	return maxLen
}

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		str := string(line)
// 		fmt.Println(lengthOfLongestSubstring(str))
// 	}
// }

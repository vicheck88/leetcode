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
	strMap := make(map[byte]int, 1000)
	maxLen := 0
	startIdx := 0
	for i := 0; i < len(s); i++ {
		_, exists := strMap[s[i]]
		if exists {
			if startIdx <= strMap[s[i]] {
				startIdx = strMap[s[i]] + 1
			}
		}
		strMap[s[i]] = i
		if maxLen < i-startIdx+1 {
			maxLen = i - startIdx + 1
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

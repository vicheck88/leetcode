// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:
// Input: digits = ""
// Output: []
// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]
// Constraints:
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].

package main

func findAllPossibleCombiation(digits string, path string, letterMap map[string]string, ans *[]string) {
	if len(digits) == 0 {
		if len(path) > 0 {
			*ans = append(*ans, path)
		}

	} else {
		letters := letterMap[string(digits[0])]
		for _, l := range letters {
			findAllPossibleCombiation(digits[1:], path+string(l), letterMap, ans)
		}
	}
}
func letterCombinations(digits string) []string {
	letterMap := make(map[string]string, 8)
	letterMap["2"] = "abc"
	letterMap["3"] = "def"
	letterMap["4"] = "ghi"
	letterMap["5"] = "jkl"
	letterMap["6"] = "mno"
	letterMap["7"] = "pqrs"
	letterMap["8"] = "tuv"
	letterMap["9"] = "wxyz"
	ans := make([]string, 0)
	findAllPossibleCombiation(digits, "", letterMap, &ans)
	return ans
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
// 		fmt.Printf("%v", letterCombinations(str))
// 	}
// }

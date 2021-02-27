// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]
// Constraints:
// 2 <= nums.length <= 103
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

package main

func twoSum(nums []int, target int) []int {
	table := make(map[int]int, 1000)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, exists := table[num]
		if exists {
			return []int{table[num], i}
		}
		table[num] = i
		table[target-num] = i
	}
	return []int{-1, -1}
}

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	cnt := 0
// 	var ints []int
// 	var target int
// 	for {
// 		cnt++
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		str := string(line)
// 		if cnt&1 == 1 {
// 			_ = json.Unmarshal([]byte(str), &ints)

// 		} else {
// 			target, _ = strconv.Atoi(str)
// 			fmt.Println(twoSum(ints, target))
// 		}
// 	}
// }

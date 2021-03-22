// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
// Notice that the solution set must not contain duplicate triplets.
// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Example 2:
// Input: nums = []
// Output: []
// Example 3:
// Input: nums = [0]
// Output: []
package main

import (
	"strconv"
)

func radixSort(nums []int) []int {
	tmp := nums
	for i := 0; i < 32; i += 8 {
		cnt := make([]int, len(nums))
		cumCnt := make([]int, len(nums))
		result := make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			ord := (tmp[j] << (32 - i)) >> 24
			cnt[ord]++
		}
		cumCnt[0] = cnt[0]
		for k := 1; k < len(nums); k++ {
			cumCnt[k] = cnt[k] + cumCnt[k-1]
		}
		for j := len(nums) - 1; j >= 0; j-- {
			result[cumCnt[tmp[j]]] = tmp[j]
			cumCnt[tmp[j]]--
		}
		tmp = result
	}
	return tmp
}
func sort(nums []int) []int {
	res := nums
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	eltMap := make(map[int]int)
	uniqueCheckingMap := make(map[string]bool)
	ansList := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		eltMap[nums[i]]++
	}
	for i := 0; i < len(nums)-1; i++ {
		eltMap[nums[i]]--
		for j := i + 1; j < len(nums); j++ {
			eltMap[nums[j]]--
			if eltMap[-nums[i]-nums[j]] > 0 {
				arr := sort([]int{nums[i], nums[j], -nums[i] - nums[j]})
				stringInt := strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2])
				if uniqueCheckingMap[stringInt] == false {
					ansList = append(ansList, arr)
					uniqueCheckingMap[stringInt] = true
				}
			}
			eltMap[nums[j]]++
		}
		eltMap[nums[i]]++
	}
	return ansList
}

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	var ints []int
// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		_ = json.Unmarshal([]byte(line), &ints)
// 		fmt.Printf("%v", threeSum(ints))
// 	}
// }

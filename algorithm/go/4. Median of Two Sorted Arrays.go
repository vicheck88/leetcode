// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
// Example 3:
// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000
// Example 4:
// Input: nums1 = [], nums2 = [1]
// Output: 1.00000
// Example 5:
// Input: nums1 = [2], nums2 = []
// Output: 2.00000
// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if m <= n {
		return findMedianSortedArrays2(nums1, nums2)
	} else {
		return findMedianSortedArrays2(nums2, nums1)
	}
}
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	isEven := (n+m)&1 == 0
	imin := 0
	imax := n
	median := 0.0
	var i int
	var j int
	if m == 0 {
		if isEven {
			return float64(nums1[n/2]+nums1[n/2-1]) / 2
		} else {
			return float64(nums1[n/2])
		}
	}
	for imin <= imax {
		i = (imin + imax) / 2
		j = (m+n+1)/2 - i
		if j > m || (i < n && j > 0 && nums1[i] < nums2[j-1]) {
			imin = i + 1
		} else if j < 0 || (i > 0 && j < m && nums1[i-1] > nums2[j]) {
			imax = i - 1
		} else {
			break
		}
	}
	max1 := 0
	max2 := 0
	if i == 0 {
		max1 = nums2[j-1]
	} else if j == 0 {
		max1 = nums1[i-1]
	} else {
		if nums1[i-1] > nums2[j-1] {
			max1 = nums1[i-1]
		} else {
			max1 = nums2[j-1]
		}
	}
	median = float64(max1)
	if isEven {
		if i == n {
			max2 = nums2[j]
		} else if j == m {
			max2 = nums1[i]
		} else {
			if nums1[i] > nums2[j] {
				max2 = nums2[j]
			} else {
				max2 = nums1[i]
			}
		}
		median = (median + float64(max2)) / 2
	}
	return median
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	cnt := 0
	var ints1 []int
	var ints2 []int
	for {
		cnt++
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		str := string(line)
		if cnt&1 == 1 {
			_ = json.Unmarshal([]byte(str), &ints1)

		} else {
			_ = json.Unmarshal([]byte(str), &ints2)
			fmt.Println(findMedianSortedArrays(ints1, ints2))
		}
	}
}

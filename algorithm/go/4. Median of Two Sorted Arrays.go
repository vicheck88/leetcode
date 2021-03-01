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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//계산의 편의를 위해 무조건 길이가 긴 input값이 먼저 오도록 설정
	n := len(nums1)
	m := len(nums2)
	if m <= n {
		return findMedianSortedArrays2(nums1, nums2)
	} else {
		return findMedianSortedArrays2(nums2, nums1)
	}
}
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	/*
		1. 각 array에서 다음 조건을 만족하는 i,j값을 계산
		1) i+j=ceiling((n+m)/2)
		2) nums1[i-1]>=nums2[j] && nums2[j-1]>=nums1[i]
		3) i의 값은 binary search로 구하며, j의 값은 1)의 조건에서 직접 계산
		2. 값이 구해질 경우, i,j 값을 이용하여 median 바로 계산
	*/
	n := len(nums1)
	m := len(nums2)
	isEven := (n+m)&1 == 0
	imin := 0 //초기 i의 min값
	imax := n //초기 i의 max값
	median := 0.0
	var i int
	var j int
	if m == 0 { //한 array의 값이 비어있을 경우: 무조건 첫번째 array의 median을 출력
		if isEven {
			return float64(nums1[n/2]+nums1[n/2-1]) / 2
		} else {
			return float64(nums1[n/2])
		}
	}
	//binary search 수행
	for imin <= imax {
		i = (imin + imax) / 2 //i: imin과 imax의 중간값으로 계산
		j = (m+n+1)/2 - i
		if j > m || (i < n && j > 0 && nums1[i] < nums2[j-1]) {
			imin = i + 1 //조건을 만족하지 않거나 i의 값이 너무 작은 경우(j의 값이 너무 클 경우)
		} else if j < 0 || (i > 0 && j < m && nums1[i-1] > nums2[j]) {
			imax = i - 1 //조건을 만족하지 않거나 i의 값이 너무 큰 경우(j값이 음수가 되는 경우)
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

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	cnt := 0
// 	var ints1 []int
// 	var ints2 []int
// 	for {
// 		cnt++
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		str := string(line)
// 		if cnt&1 == 1 {
// 			_ = json.Unmarshal([]byte(str), &ints1)

// 		} else {
// 			_ = json.Unmarshal([]byte(str), &ints2)
// 			fmt.Println(findMedianSortedArrays(ints1, ints2))
// 		}
// 	}
// }

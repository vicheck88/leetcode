// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
// Notice that you may not slant the container.
// Example 1
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
//In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:
// Input: height = [1,1]
// Output: 1
// Example 3:
// Input: height = [4,3,2,1,4]
// Output: 16
// Example 4:
// Input: height = [1,2,1]
// Output: 2

package main

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		minH := min(height[left], height[right])
		area := minH * (right - left)
		if maxArea < area {
			maxArea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
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
// 		println(maxArea(ints))
// 	}
// }

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Follow up: Could you do this in one pass?
// Example 1:
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:
// Input: head = [1], n = 1
// Output: []
// Example 3:
// Input: head = [1,2], n = 1
// Output: [1]
// Constraints:
// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

//input값을 listNode로 변환
func initialize(nums []int) *ListNode {
	header := new(ListNode) //header node(dummy)
	tmp := header
	for _, num := range nums {
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = num
	}
	return header.Next
}
func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	var nums []int
	var n int
	cnt := 0
	var list *ListNode
	for {
		cnt++
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		str := string(line)
		if cnt&1 == 1 {
			_ = json.Unmarshal([]byte(str), &nums)
		} else {
			n, _ = strconv.Atoi(str)
			list = initialize(nums)
			var res = removeNthFromEnd(list, n)
			for res != nil {
				fmt.Printf("%d ", res.Val)
			}
			fmt.Printf("\n")
		}
	}
}

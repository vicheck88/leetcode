// You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//Example
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Input: l1 = [0], l2 = [0]
// Output: [0]
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

package main

type ListNode struct { //single linked list 생성
	Val  int
	Next *ListNode //다음 list struct
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	roundNum := 0        //올림값
	res := new(ListNode) //최종 결과 list
	tmp := res           //계산 위한 list
	for {
		tmpp := tmp
		tmp.Next = new(ListNode) //Next값 생성
		tmp = tmp.Next           //계산용 list로 이동
		tmp.Val = roundNum       //올림값 미리 반영
		if l1 == nil && l2 == nil {
			//두 값이 전부 없을 경우: 작업 종료
			if roundNum == 0 {
				tmpp.Next = nil //next값 삭제
			}
			break
		}
		if l1 != nil {
			//listNode1 값이 있을 경우 listNode1의 값 연산
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			//listNode2 값이 있을 경우 listNode2의 값 연산
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		//올림값 설정
		if tmp.Val >= 10 {
			roundNum = 1
			tmp.Val -= 10
		} else {
			roundNum = 0
		}
	}
	return res.Next
}

// func main() {
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	cnt := 0
// 	var num1 *ListNode
// 	var num2 *ListNode
// 	var ints []int
// 	for {
// 		cnt++
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 			break
// 		}
// 		str := string(line)
// 		_ = json.Unmarshal([]byte(str), &ints)
// 		if cnt&1 == 1 {
// 			num1 = initialize(ints)
// 		} else {
// 			num2 = initialize(ints)
// 			res := addTwoNumbers(num1, num2)
// 			for {
// 				fmt.Printf("%d,", res.Val)
// 				res = res.Next
// 				if res == nil {
// 					fmt.Printf("\n")
// 					break
// 				}
// 			}
// 		}
// 	}
// }

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: l1.Val + l2.Val}

	for next1, next2, current := l1.Next, l2.Next, result; ; {
		carryBit := 0
		if current.Val >= 10 {
			current.Val = current.Val - 10
			carryBit = 1
		}

		if carryBit == 0 && next1 == nil && next2 == nil {
			break
		}

		val1 := 0
		if next1 != nil {
			val1 = next1.Val
			next1 = next1.Next
		}

		val2 := 0
		if next2 != nil {
			val2 = next2.Val
			next2 = next2.Next
		}

		val := carryBit + val1 + val2

		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}}

	//l1 := &ListNode{Val: 5}
	//l2 := &ListNode{Val: 5}

	fmt.Print(addTwoNumbers(l1, l2))
}

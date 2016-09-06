package main

import (
	"fmt"
)

// ListNode ...
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	next := head
	s := fmt.Sprintf("%d", next.Val)

	for next.Next != nil {
		next = next.Next
		s += fmt.Sprintf(" %d", next.Val)
	}

	return s
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	size := getSize(head)

	if size <= 2 {
		return
	}

	totalStep := size / 2

	prev := head
	target := getLastNode(head)
	last := moveAfter(target, prev)
	moveStep := 1

	for moveStep < totalStep {
		prev = target.Next
		target = last
		last = moveAfter(last, prev)
		moveStep++
	}
}

func getSize(head *ListNode) int {
	next := head
	size := 1

	for next.Next != nil {
		next = next.Next
		size++
	}

	return size
}

func getLastNode(head *ListNode) *ListNode {
	next := head

	for next.Next != nil {
		next = next.Next
	}

	return next
}

func moveAfter(target *ListNode, prev *ListNode) *ListNode {

	if target == prev {
		panic("target and prev cannot be same object")
	}

	next := prev
	for next.Next != nil {
		if next.Next == target {
			next.Next = target.Next
			break
		}

		next = next.Next
	}

	target.Next = prev.Next
	prev.Next = target

	return next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	// head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}

	fmt.Println(head)

	reorderList(head)

	fmt.Println(head)
}

package _2

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	target := &ListNode{}
	add(l1, l2, target, 0)
	return target
}

func add(l1 *ListNode, l2 *ListNode, target *ListNode, val int) *ListNode {
	if l1 == nil && l2 == nil {
		target.Val = val
		return target
	}
	i := val
	var l1N, l2N *ListNode
	if l1 != nil {
		i += l1.Val
		l1N = l1.Next
	}

	if l2 != nil {
		i += l2.Val
		l2N = l2.Next
	}

	//i := l1.Val + l2.Val + val
	if i >= 10 {
		i = i - 10
		val = 1
	} else {
		val = 0
	}
	target.Val = i
	if l1N == nil && l2N == nil {
		if val == 0 {
			return target
		}

	}
	target.Next = &ListNode{}
	add(l1N, l2N, target.Next, val)
	return nil
}

func getListNode(nums []int) *ListNode {
	l := &ListNode{}
	ll(l, nums)

	return l
}

func ll(node *ListNode, nums []int) *ListNode {
	if len(nums) == 0 {
		node = nil
		return node
	}
	node.Val = nums[0]

	if len(nums) == 1 {
		node.Next = nil
		return node
	}
	node.Next = &ListNode{}

	ll(node.Next, nums[1:])
	return node
}

func (l *ListNode) String() string {
	var result string
	current := l
	for current != nil {
		result += fmt.Sprintf("%d ", current.Val)
		current = current.Next
	}
	return result
}

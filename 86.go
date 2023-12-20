package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func partition(head *ListNode, x int) *ListNode {

	// 准备两个链表
	// 创建两个哑节点作为两个链表的头部
	leftDummy := &ListNode{}
	rightDummy := &ListNode{}
	// 创建两个指针用来构建新的链表
	left := leftDummy
	right := rightDummy

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}

		head = head.Next
	}

	right.Next = nil
	// 连接两个链表
	left.Next = rightDummy.Next

	return leftDummy.Next

}

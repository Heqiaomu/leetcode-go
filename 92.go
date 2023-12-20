package leetcode_go

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// 首先设定一个 dummy 节点
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	pre := dummy

	// 用数组的方式来存储所有节点
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	// 对 node 数组进行轮换  1,2,3,4,5  [left-1, right-1] 之间进行交换
	leftIdx := left - 1
	rightIdx := right - 1

	for leftIdx <= rightIdx {
		nodes[leftIdx], nodes[rightIdx] = nodes[rightIdx], nodes[leftIdx]
		leftIdx++
		rightIdx--
	}

	// 重连链表
	for _, node := range nodes {
		pre.Next = node
		pre = pre.Next
	}
	pre.Next = nil // 重要：确保链表尾部是 nil

	return dummy.Next

}

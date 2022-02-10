package main

// 反转链表

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

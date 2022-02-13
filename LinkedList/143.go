package main

func getMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse2(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTerm := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTerm
	}
	return prev
}

func merge(l1, l2 *ListNode) {
	var l1Term, l2Term *ListNode
	for l1 != nil && l2 != nil {
		l1Term = l1.Next
		l2Term = l2.Next
		l1.Next = l2
		l1 = l1Term

		l2.Next = l1
		l2 = l2Term
	}
}

func reorderList(head *ListNode) {
	mid := getMiddleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverse2(l2)
	merge(l1, l2)
}

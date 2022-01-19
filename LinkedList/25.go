package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, tail := dummy, dummy
	for prev.Next != nil {
		for i := 0; i < k && tail.Next != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		nextHead := tail.Next
		head = prev.Next
		tail.Next = nil
		prev.Next = reverse(head)
		head.Next = nextHead
		prev = head
		tail = prev
	}
	return dummy.Next
}

func reverse(node *ListNode) *ListNode {
	var pre *ListNode = nil
	var curr = node
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

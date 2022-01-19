package main

// 环形指针

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for fast.Next != nil && head.Next != nil {
				if fast == head{
					return head
				}
				fast = fast.Next
				head = head.Next
			}

		}
	}
	return nil
}

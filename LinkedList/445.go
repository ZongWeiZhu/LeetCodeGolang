package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	p1 := len(stack1) - 1
	p2 := len(stack2) - 1
	stack3 := []int{}
	pre := 0
	for p1 >= 0 || p2 >= 0 {
		sum := pre
		if p1 >= 0 {
			sum += stack1[p1]
		}
		if p2 >= 0 {
			sum += stack2[p2]
		}
		stack3 = append(stack3, sum%10)
		pre = sum / 10
		p1--
		p2--
	}

	if pre != 0 {
		stack3 = append(stack3, pre)
	}
	dummy := &ListNode{

	}
	slow := &ListNode{
		Val:  stack3[len(stack3)-1],
		Next: nil,
	}
	dummy.Next = slow
	for i := len(stack3) - 2; i >= 0; i-- {
		fast := &ListNode{
			Val:  stack3[i],
			Next: nil,
		}
		slow.Next = fast
		slow = fast
	}
	return dummy.Next
}

package main


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var ans = lists[0]
	for i:= 1 ;i <len(lists) ; i++{
		ans = merge2Lists(ans,lists[i])
	}
	return ans
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var res = &ListNode{}
	tail := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		}else{
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		 tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return res.Next
}

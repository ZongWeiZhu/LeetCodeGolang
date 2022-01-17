package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{}
	res := [][]int{}
	flags := true
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		n := len(q)
		data := make([]int, n)
		for i := 0; i < n; i++ {
			if flags == true {
				data[i] = q[i].Val
			} else {
				data[n-i-1] = q[i].Val
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		flags = !flags
		q = q[n:]
		res = append(res, data)
	}
	return res
}

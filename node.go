package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(x ...int) *ListNode {
	if len(x) == 0 {
		return nil
	}
	h := &ListNode{
		Val: x[0],
	}

	head := h

	for i := 0; i < len(x); i++ {
		h.Next = &ListNode{
			Val: x[i],
		}
		h = h.Next
	}

	return head
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type timberjack struct {
	rs []int
}

func (j *timberjack) walk(t *TreeNode) {

	if t.Left != nil {
		j.walk(t.Left)
	}

	j.rs = append(j.rs, t.Val)

	if t.Right != nil {
		j.walk(t.Right)
	}
}

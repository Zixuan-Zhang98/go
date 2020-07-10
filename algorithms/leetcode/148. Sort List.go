package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 1. base case
	if head == nil || head.Next == nil {
		return head
	}

	// 2. halving
	mid := middleNode(head)
	left, right := head, mid.Next
	mid.Next = nil

	// 3. sorting
	left = sortList(left)
	right = sortList(right)

	// 4. merging
	list := mergeLists(left, right)
	return list
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeLists(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	dummy := &ListNode{}
	pre := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}

	if left != nil {
		pre.Next = left
	} else {
		pre.Next = right
	}

	return dummy.Next
}

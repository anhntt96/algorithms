package leetcode

//https://leetcode.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head

	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{(carry + x + y) % 10, nil}
		curr = curr.Next
		carry = (carry + x + y) / 10
	}

	return head.Next
}

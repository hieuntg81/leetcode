package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addition := 0
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil || addition != 0 {
		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		cVal := n1 + n2 + addition
		addition = cVal / 10
		cur.Next = &ListNode{Val: cVal % 10, Next: nil}
		cur = cur.Next
	}
	return head.Next
}

//func main() {
//	l13 := &ListNode{
//		Val:  3,
//		Next: nil,
//	}
//	l12 := &ListNode{
//		Val:  4,
//		Next: l13,
//	}
//	l11 := &ListNode{
//		Val:  2,
//		Next: l12,
//	}
//
//	l23 := &ListNode{
//		Val:  4,
//		Next: nil,
//	}
//	l22 := &ListNode{
//		Val:  6,
//		Next: l23,
//	}
//	l21 := &ListNode{
//		Val:  5,
//		Next: l22,
//	}
//
//	x := addTwoNumbers(l11, l21)
//	println("Done", x)
//}

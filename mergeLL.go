package mergell

// ListNode is a simple linked list node structure.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists will take in two sorted linked list and merge them together.``
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{-1, nil}

	curr := res

Loop:
	for l1 != nil || l2 != nil {

		switch {
		case l1 == nil:
			curr.Next = l2
			break Loop
		case l2 == nil:
			curr.Next = l1
			break Loop
		default:
			if l1.Val < l2.Val {
				curr.Next = l1
				l1 = l1.Next
			} else {
				curr.Next = l2
				l2 = l2.Next
			}
			curr = curr.Next
		}

	}

	return res.Next
}

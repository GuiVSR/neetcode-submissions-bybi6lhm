/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	head = head.Next

	for head.Next != nil {
		temp := newHead
		nextHead := head.Next
		newHead = head
		head = nextHead

		newHead.Next = temp
	}

	temp := newHead
	newHead = head
	newHead.Next = temp

	return newHead
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil { return list2 }
	if list2 == nil { return list1 }

	var mergedHead *ListNode

	if list1.Val <= list2.Val {
		mergedHead = &ListNode{
			Val: list1.Val,
			Next: nil,
		}
		list1 = list1.Next

	} else {
		mergedHead = &ListNode{
			Val: list2.Val,
			Next: nil,
		}
		list2 = list2.Next
	}

	merged := mergedHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			merged.Next = &ListNode{
				Val: list1.Val,
				Next: nil,
			}
			list1 = list1.Next
			merged = merged.Next
			
		} else {
			merged.Next = &ListNode{
				Val: list2.Val,
				Next: nil,
			}
			list2 = list2.Next
			merged = merged.Next
		}
		
	}

	if list1 != nil {
		merged.Next = list1
	} else {
		merged.Next = list2
	}

	return mergedHead
}

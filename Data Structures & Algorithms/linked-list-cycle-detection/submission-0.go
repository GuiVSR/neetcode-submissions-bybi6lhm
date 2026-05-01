/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	 visitedMap := make(map[*ListNode]bool)

	for head != nil {
		_, alreadyVisited := visitedMap[head]
		if alreadyVisited {
			return true
		} else {
			visitedMap[head] = true
			head = head.Next
		}
	}
	return false   
}

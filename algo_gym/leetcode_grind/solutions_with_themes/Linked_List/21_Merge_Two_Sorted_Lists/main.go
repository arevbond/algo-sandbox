/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0, Next: nil}
    cur := dummy
    for list1 != nil || list2 != nil {
        if list1 == nil {
            cur.Next = list2
            break
        } else if list2 == nil {
            cur.Next = list1 
            break
        }

        if list1.Val <= list2.Val {
            cur.Next = &ListNode{Val: list1.Val, Next: nil}
            cur = cur.Next
            list1 = list1.Next
        } else {
            cur.Next = &ListNode{Val: list2.Val, Next: nil}
            cur = cur.Next
            list2 = list2.Next
        }
    }
    return dummy.Next   
}
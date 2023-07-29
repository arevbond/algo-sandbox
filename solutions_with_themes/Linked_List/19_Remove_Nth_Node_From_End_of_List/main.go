/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head}
    cur := dummy
    rigth := head 
    for n != 0 {
        n--
        rigth = rigth.Next
    }
    for rigth != nil {
        cur = cur.Next
        rigth = rigth.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}
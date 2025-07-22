package mergetwosortedlists

type ListNode struct {
    Val int
    Next *ListNode
}

/*


*/


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    } 

    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    dummy := &ListNode{}
    cur := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            cur.Next = &ListNode{Val: list1.Val}
            list1 = list1.Next
        } else {
            cur.Next = &ListNode{Val: list2.Val}
            list2 = list2.Next
        }
        cur = cur.Next
    }

    for list1 != nil {
        cur.Next = &ListNode{Val: list1.Val}
        list1 = list1.Next
        cur = cur.Next
    }

    for list2 != nil {
        cur.Next = &ListNode{Val: list2.Val}
        list2 = list2.Next
        cur = cur.Next
    }
    return dummy.Next
}

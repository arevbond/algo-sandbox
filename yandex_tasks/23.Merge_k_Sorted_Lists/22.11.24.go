package mergeksortedlists

import (
	"container/heap"
	"sort"
)

type ListNode struct {
    Val int
    Next *ListNode 
}

func mergeKLists(lists []*ListNode) *ListNode {
    mp := make(map[int]int)    
    for _, node := range lists {
        cur := node
        for cur != nil {
            mp[cur.Val]++
            cur = cur.Next
        }
    }

    vals := make([]int, 0, len(mp))

    for val := range mp {
        for mp[val] > 0 {
            vals = append(vals, val)
            mp[val]--
        }
    }

    sort.Ints(vals)

    dummy := &ListNode{}
    head := dummy

    for _, val := range vals {
        head.Next = &ListNode{Val: val}
        head = head.Next
    }

    return dummy.Next
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type HeapNode struct {
    Val int
    Indx int
    Node *ListNode
}

type Heap []*HeapNode

func (h Heap) Len() int { return len(h)}
func (h Heap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}
func (h Heap) Swap(i, j int) {
   h[i], h[j] = h[j], h[i] 
}

func (h *Heap) Push(x any) {
    *h = append(*h, x.(*HeapNode))
}

func (h *Heap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
    h := &Heap{}
    heap.Init(h)

    for i, l := range lists {
        if l == nil {
            continue
        }
        heap.Push(h, &HeapNode{Val: l.Val, Indx: i, Node: l})
    }

    dummy := &ListNode{}
    cur := dummy

    // O(N * log(L)) time ; N = len(lists)
    // O(N) space 
    for h.Len() > 0 {
        val := heap.Pop(h)

        heapNode := val.(*HeapNode)
        indx, node := heapNode.Indx, heapNode.Node

        cur.Next = &ListNode{Val: node.Val}
        cur = cur.Next

        if node.Next != nil {
            heap.Push(h, &HeapNode{Val: node.Next.Val, Indx: indx, Node: node.Next})
        }
    }
    return dummy.Next
 
}

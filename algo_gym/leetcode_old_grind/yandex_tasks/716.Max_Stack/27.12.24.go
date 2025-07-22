package main    
/*
Design a max stack that supports push, pop, top, peekMax and popMax.

    push(x) -- Push element x onto stack.
    pop() -- Remove the element on top of the stack and return it.
    top() -- Get the element on the top.
    peekMax() -- Retrieve the maximum element in the stack.
    popMax() -- Retrieve the maximum element in the stack, and remove it. If you find more than one maximum elements, only remove the top-most one.

Example 1:

MaxStack stack = new MaxStack();
stack.push(5);
stack.push(1);
stack.push(5);
stack.top(); -> 5
stack.popMax(); -> 5
stack.top(); -> 1
stack.peekMax(); -> 5
stack.pop(); -> 1
stack.top(); -> 5
*/

/*
    1. store all nums values in double linked list [2 3 5 1 6]
    2. store current maximum values in additional array maxNums [2 3 5 6] with maximum nodes
    3. when push element: check than it's new maximum value then append it to maxNums
    4. when pop element: compare last element in nums with last element in maxNums if they equal remove both
    5. popMax: 
*/

type Node struct {
    Val int
    Next, Prev *Node
}
 
type MaxStack struct {
    head *Node
    maxVals []*Node
}

func (s *MaxStack) push(x int) {
    var newNode *Node
    if s.head != nil {
        newNode = &Node{Val: x, Prev: s.head}
        s.head.Next = newNode
    } else {
        newNode = &Node{Val: x}
        s.head = newNode
    }

    if len(s.maxVals) == 0 {
        s.maxVals = append(s.maxVals, newNode)
    } else {
        last := s.maxVals[len(s.maxVals) - 1]
        if last.Val <= newNode.Val {
            s.maxVals = append(s.maxVals, newNode)
        }
    }
}

func (s *MaxStack) pop() (int, bool) {
    if s.head == nil {
        return -1, false
    }

    last := s.head
    curMax := s.maxVals[len(s.maxVals) - 1]
    if last.Val == curMax.Val {
        s.maxVals = s.maxVals[:len(s.maxVals)-1]
    }
    if s.head.Prev != nil {
        s.head.Prev.Next = nil
    }
    s.head = s.head.Prev
    return last.Val, true
}

func (s *MaxStack) top() (int, bool) {
    if s.head == nil {
        return -1, false
    }
    return s.head.Val, true
}

func (s *MaxStack) peekMax() (int, bool) {
    if s.head == nil {
        return -1, false
    }

    maxNode := s.maxVals[len(s.maxVals)-1]
    return maxNode.Val, true
}

func (s *MaxStack) popMax() (int, bool) {
    if s.head == nil {
        return -1, false
    }
    
    maxNode := s.maxVals[len(s.maxVals) - 1]
    if maxNode.Prev != nil {
        maxNode.Prev.Next = maxNode.Next
    } 
    if maxNode.Next != nil {
        maxNode.Next.Prev = maxNode.Prev
    }
    if s.head == maxNode {
        s.head = maxNode.Prev
    }
    return maxNode.Val, true
}


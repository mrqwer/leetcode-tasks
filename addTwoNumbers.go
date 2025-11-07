package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 2 -> 4 -> 3
    // 5 -> 6 -> 4
    // 342 + 465 = 807 => 7 -> 0 -> 8

    resultHead := &ListNode{} // new linked list
    currResult := resultHead  // get a link to the head of new linked list 
    curr1 := l1               // link to the head of first linked list
    curr2 := l2               // link to the head of second linked list
    var inMind int                 

    for curr1 != nil || curr2 != nil {  // traverse until end of long linked list
        val1, val2 := 0, 0    // initial values

        if curr1 != nil {     // get value and link from the first linked list
            val1 = curr1.Val
            curr1 = curr1.Next
        }
        if curr2 != nil {      // get value and link from the second linked list
            val2 = curr2.Val
            curr2 = curr2.Next
        }

        sum := val1 + val2 + inMind  // sum values of two linked list and carry value
        mod := sum % 10              // get remainder from sum value
        inMind = sum / 10            // get carry value from sum value

        node := &ListNode{Val: mod}  // new linked list
        currResult.Next = node       // assign new node to the next node of current node of new linked list
        currResult = currResult.Next // move forward
    }

    if inMind > 0 {
        currResult.Next = &ListNode{Val: inMind} // in case where carry value is left, then we create new node with inMind value
    }

    return resultHead.Next  // return second node as a head node since we start adding from the second node
}

func printList(head *ListNode){
    curr := head
    for curr != nil {
        fmt.Println(curr.Val)
        curr = curr.Next
    }
}

func main(){
    l1 := &ListNode{Val: 2}
    l11 := &ListNode{Val: 4}
    l12 := &ListNode{Val: 3}
    l1.Next = l11
    l1.Next.Next = l12

    l2 := &ListNode{Val: 5}
    l21 := &ListNode{Val: 6}
    l22 := &ListNode{Val: 4}
    l2.Next = l21
    l2.Next.Next = l22

    res := addTwoNumbers(l1, l2)    
    printList(res)
}

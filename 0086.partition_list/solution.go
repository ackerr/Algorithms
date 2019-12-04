package main


type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }

    low := &ListNode{Val: -1}
    ans := low
    high := &ListNode{Val: -1}
    target := high
    for head != nil{
        if head.Val < x{
            low.Next = head
            low = low.Next
        }else {
            high.Next = head
            high = high.Next
        }
        head = head.Next
    }
    high.Next = nil
    low.Next = target.Next
    return ans.Next
}
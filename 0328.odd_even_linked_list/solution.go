package main


type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil{
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

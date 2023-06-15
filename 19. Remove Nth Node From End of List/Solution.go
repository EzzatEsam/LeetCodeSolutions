package main
import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	oldHead := head
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	if length == n {
		if length == 1  {
			return nil
		} else {
			return oldHead.Next
		}
	}
	head = oldHead
	for i := 0; i < length - n - 1 ; i++ {
		head = head.Next
	}

	if head.Next != nil  {
		head.Next = head.Next.Next
	}
    return oldHead
}


func main(){  
	fmt.Println("hello world")
}
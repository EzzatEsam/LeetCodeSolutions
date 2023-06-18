func calculateListLength(head *ListNode) int {
    n :=0
    for head != nil {
        n++
        head = head.Next
    }
    return n
}
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }
    n := calculateListLength(head)
    fmt.Println(n)
    oldHead := head
    k = k % n
    if k ==0 || n == 1 {
        return head
    }
    for i := 0; i < n -k -1; i++ {
        head = head.Next
    }
    temp := head.Next
    head.Next = nil
    head = temp
    for head.Next != nil {
        head = head.Next
    }
    head.Next = oldHead
    return temp
}
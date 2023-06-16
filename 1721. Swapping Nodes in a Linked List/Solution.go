func swapNodes(head *ListNode, k int) *ListNode {
	oldHead := head
	var firstPre *ListNode
	length := 0
	for head != nil {
		length++
		if length == k -1 {
			firstPre = head
		}
		head = head.Next
	}
	if length == 1 {
		return oldHead
	} else if length == 2 {
		second := oldHead.Next
		second.Next = oldHead
		oldHead.Next = nil
		return second
	}
	if k == length {
		second := firstPre.Next
		second.Next = oldHead.Next
		firstPre.Next = oldHead
		oldHead.Next = nil
		return second
	}
	
	head = oldHead

	for i := 0; i < length - k - 1;  i++ {
		head = head.Next
	}
	
	secondPre := head   
	second := head.Next  
	temp := second.Next  
	
	if k == 1  {
		second.Next = oldHead.Next
		secondPre.Next = oldHead
		oldHead.Next = nil
		return second
	}



	first := firstPre.Next 
	if second.Next == first {
		secondPre.Next = first
		second.Next = first.Next
		first.Next = second
		return oldHead
	}

	if first.Next == second  {
		second.Next = first
	}else {
		second.Next = first.Next   
		secondPre.Next = first	
	}
	first.Next = temp
	firstPre.Next = second


  return oldHead
}




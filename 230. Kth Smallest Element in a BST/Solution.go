package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	idx := 0
	smallest := 0
	inOrderTraversal(root, k ,&idx, &smallest)
	return smallest
}

func inOrderTraversal( root * TreeNode ,k int ,idx * int , smallest *int )  {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left,k, idx, smallest)
	*idx ++
	if *idx == k {
		*smallest = root.Val
		return
	}

	inOrderTraversal(root.Right,  k,idx, smallest)
}
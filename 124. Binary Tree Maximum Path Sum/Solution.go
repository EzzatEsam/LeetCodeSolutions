
import (
	"math"
)

// type TreeNode struct {
// 	Val int
// 	Left *TreeNode     
// 	Right *TreeNode
// }



func maxPathSum(root *TreeNode) int {
    maxVal := math.MinInt
	

	NodeValue(root , &maxVal)
	return maxVal
}

func NodeValue(node *TreeNode , maxVal *int )  int {  
	if node == nil {
		return 0
	}
	connectionSum ,leftSum , rightSum := node.Val , 0 ,0
	
	leftSum = NodeValue(node.Left , maxVal)
	rightSum = NodeValue(node.Right , maxVal)

	connectionSum += max(leftSum , rightSum)
	fullSum := node.Val + rightSum + leftSum

	*maxVal = max( *maxVal , fullSum , connectionSum , node.Val)
	
	return max(connectionSum , node.Val)
}

func max(nums ...int) int {
    if len(nums) == 0 {
        return 0 // or return an error, depending on your use case
    }
    
    maxNum := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxNum {
            maxNum = nums[i]
        }
    }
    return maxNum
}
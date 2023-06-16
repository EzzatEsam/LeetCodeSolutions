var memo []int
func rob(nums []int) int {
    
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 { 
		return max(nums[0], nums[1])
	}
	memo = MakeFilledArray(len(nums), -1)

	return robIndex(nums, 0)

}

func robIndex(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	if memo[index] != -1 {
		return memo[index]
	}

	result := max(robIndex(nums, index + 2) + nums[index], robIndex(nums, index + 1))

	memo[index] = result

	return result
}


func max(a , b int) int {  
	if a > b {
		return a
	}
	return b
}

func MakeFilledArray(n , k int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		 arr[i] = k
	}
	return arr
}
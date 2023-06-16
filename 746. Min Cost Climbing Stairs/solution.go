var memo []int
func minCostClimbingStairs(cost []int) int {
	memo = MakeFilledArray(len(cost), -1)
	return goToIdx(cost, 0)
}

func goToIdx(cost []int, idx int) int {
	
	if idx >= len(cost) -1 {
		return 0
	}

	if memo[idx] != -1  {
		return memo[idx]
	}
	// if we take current
	sumIfWeTakeCurrent := cost[idx] + goToIdx(cost, idx+1)
	// if we take next
	sumIfWeTakeNext := cost[idx+1] + goToIdx(cost, idx+2)

	result :=  min(sumIfWeTakeCurrent, sumIfWeTakeNext)
	memo[idx] = result

	return result
}

func min(a , b int) int {  
	if a < b {
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
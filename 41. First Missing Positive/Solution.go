package main



func firstMissingPositive(nums []int) int {
	//hashList := make(map[int]bool )

	hashList := make([]bool, len(nums)+1)
	for _, num := range nums {
		if num > 0 && num <= len(nums) {
			hashList[num] = true
		}
	}

	for i := 1; i <= len(nums); i++ {
		// if _, ok := hashList[i]; !ok {
		// 	return i
		// }
		if ! hashList[i] {
			return i 
		}
	}
	return len(nums) + 1
}
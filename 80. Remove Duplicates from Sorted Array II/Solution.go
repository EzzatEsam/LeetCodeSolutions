func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
    idx := 0
	duplicate := false
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx]    {
			duplicate = false
			idx++
			nums[idx] = nums[i]
		} else if !duplicate {
				idx++
				nums[idx] = nums[i]
                duplicate  = true
	    }
	}
	return idx +1
}
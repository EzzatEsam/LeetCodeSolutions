package main

func trap(height []int) int {
	totalRain := 0
	left, right := 0, len(height)-1

	for right - left > 1  {
		println(left, right , totalRain)
		if height[left] < height[right] {
			println("left")
			idx := left
			for  idx < right {
				if height[idx] > height[left]  {
					
					break
				}
				totalRain += height[left] - height[idx]
				idx++	
			}
			left = idx
			
		} else {
			println("right")
			i := right
			for  i > left {
				if height[i] > height[right] {
					break
				}
				totalRain += height[right] - height[i]
				i--
			}
			right = i
			
		}
				
	}

    return totalRain
}

func min(a , b int) int {
	if a < b {
		return a
	}
	return b	
}


func main()  {
	testCase := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	println("result 1", trap(testCase))
	
	testCase2 := []int{4,2,3}
	println("result 2", trap(testCase2))
}





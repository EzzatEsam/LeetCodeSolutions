package main

func maxArea(height []int) int {
    maxArea := 0
    left := 0
    right := len(height) - 1
    for left < right {
        minHeight := height[left]
        if height[right] < minHeight {
            minHeight = height[right]
        }
        area := (right - left) * minHeight
        if area > maxArea {
            maxArea = area
        }
        for height[left] <= minHeight && left < right {
            left++
        }
        for height[right] <= minHeight && left < right {
            right--
        }
    }
    return maxArea
}

	

class Solution {
    public int maxArea(int[] height) {
        int maxArea = 0;
        int left = 0;
        int right = height.length - 1;
        while (left < right) {
            var minHeight = (height[left] < height[right])?height[left] : height[right];
            var area = (right - left) * minHeight;
            maxArea = (area > maxArea) ? area : maxArea;
            while (height[left] <= minHeight && left < right) {
                left++;
            }
            while (height[right] <= minHeight && left < right) {
                right--;
            }
        }
        return maxArea;       
    }
}
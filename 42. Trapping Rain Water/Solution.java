public class Solution {
    public int trap(int[] height) {
        int totalRain = 0;
        int left = 0;
        int right = 0;

        while (right < height.length) {
            if (right - left > 1) {
                while (right < height.length && height[right] < height[left]) {
                    right++;
                }

                if (right >= height.length) {
                    right = height.length - 1;
                    int minH = Math.min(height[left], height[right]);
                    for (int i = right - 1; i > left; i--) {
                        minH = Math.max(minH, height[i]);
                        totalRain += minH - height[i];
                    }
                    break;
                }

                int minH = Math.min(height[left], height[right]);
                System.out.println(left + " " + right + " " + minH);
                for (int i = left; i < right; i++) {
                    totalRain += minH - height[i];
                }
                left = right;
            } else if (height[right] >= height[left]) {
                left = right;
            }
            right++;
        }

        return totalRain;
    }

    public int min(int a, int b) {
        return Math.min(a, b);
    }

    public int max(int a, int b) {
        return Math.max(a, b);
    }


}

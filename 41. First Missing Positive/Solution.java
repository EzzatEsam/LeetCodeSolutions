
class Solution {
    public int firstMissingPositive(int[] nums) {
        var numsList = new boolean[nums.length + 1];

        for (var num : nums)  {
            if (num > 0 && num <= nums.length) {
                numsList[num] = true;
            }   
        }

        for (int i = 1; i < numsList.length; i++) {
            if (!numsList[i]) {
                return i;
            }
        }
        return nums.length +1; 
    }
}
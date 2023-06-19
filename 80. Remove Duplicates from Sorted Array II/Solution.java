class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length < 2) {
            return nums.length;
        }
        int idx = 0;
        boolean duplicate = false;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[idx]) {
                duplicate = false;
                idx++;
                nums[idx] = nums[i];
            } else if (!duplicate) {
                idx++;
                nums[idx] = nums[i];
                duplicate = true;
            }
        }
        return idx + 1;
    }
}
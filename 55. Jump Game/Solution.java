class Solution {
    boolean[] db;
    public boolean canJump(int[] nums) {
        db = new boolean[nums.length];
       return jumpWithIndex(nums, 0); 
    }

    boolean jumpWithIndex(int[] nums, int index) {
       System.out.println(index);
       System.out.println(nums[index]);
       if (index > nums.length - 1 || db[index]) {
           return false;
       }

       if (index == nums.length - 1) {
           return true;
       }
        for (int i = nums[index]; i >= 1; i--) {
            if (jumpWithIndex(nums, index + i )) {
                return true;
            } 
        }
        db[index] = true;
        return false;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.canJump(new int[]{2,3,1,1,4}));
    }
}
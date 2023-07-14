


class Solution {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode() {}
        TreeNode(int val) { this.val = val; }
        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }
    int idx = 0;
    int smallest = 0;
    
    public int kthSmallest(TreeNode root, int k) {
        inOrderTraversal(root, k);
        return smallest;
    }

    private void inOrderTraversal(TreeNode root , int k) {
        if (root == null) {
            return;
        }
        inOrderTraversal(root.left, k);
        idx++;
        if (idx == k) {
            smallest = root.val;
            return;
        }
        inOrderTraversal(root.right, k);   
    }
}
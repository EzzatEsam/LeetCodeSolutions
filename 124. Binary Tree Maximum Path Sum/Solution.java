/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int _maxVale;
    public int maxPathSum(TreeNode root) {
        _maxVale = Integer.MIN_VALUE;
        nodeValue(root);
        return _maxVale;
    }
    
    private int nodeValue(TreeNode node) {
        if (node == null) {
            return 0;
        }
        
        int connectionSum = node.val;
        int leftSum = nodeValue(node.left);
        int rightSum = nodeValue(node.right);
        
        connectionSum += Math.max(leftSum, rightSum);
        int fullSum = node.val + leftSum + rightSum;
        
        _maxVale = Math.max(_maxVale, Math.max(fullSum, Math.max(connectionSum, node.val)));
        
        return Math.max(connectionSum, node.val);
    }
}
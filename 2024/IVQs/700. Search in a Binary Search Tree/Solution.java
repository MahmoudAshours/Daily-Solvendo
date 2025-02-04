/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class Solution {
    public TreeNode searchBST(TreeNode root, int val) {
        try {
            if (val == root.val || root == null) return root;
            if (val > root.val) return searchBST(root.right, val);
            if (val < root.val) return searchBST(root.left, val);
            return null;
        } catch (Exception e) {
            return root;
        }
    }
}
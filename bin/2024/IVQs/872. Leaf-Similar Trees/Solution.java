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
    public ArrayList<Integer> l1 = new ArrayList();
    public ArrayList<Integer> l2 = new ArrayList();

    public boolean leafSimilar(TreeNode root1, TreeNode root2) {

        dfs(l1,root1);
        dfs(l2,root2);

        if(l1.size()!=l2.size()) return false;
        return l1.equals(l2) ? true : false;
    }


    private void dfs(ArrayList l , TreeNode root){
        if(root.left==null && root.right==null){
            l.add(root.val);
            return;
        }
        if(root.left!=null )
            dfs(l,root.left);

        if(root.right!=null )
            dfs(l,root.right);
    }
}
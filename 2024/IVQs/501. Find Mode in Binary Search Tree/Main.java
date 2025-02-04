public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        TreeNode node = new TreeNode(1, new TreeNode(2, new TreeNode(2), new TreeNode(3)), new TreeNode(10));
        sol.findMode(node);
    }
}

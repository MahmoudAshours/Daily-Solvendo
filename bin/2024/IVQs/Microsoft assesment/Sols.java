public class Sols {

    private boolean leftTree = true;
    private boolean rightTree = true;

    public boolean solution(TreeNode tree) {
        return checkBST(tree, (int) Double.NEGATIVE_INFINITY, (int) Double.POSITIVE_INFINITY);

    }

    private boolean checkBST(TreeNode root, int min, int max) {
        if (root.left == null && root.right == null)
            return true;

        if ((root.left != null && root.left.val > root.val) || (root.right != null && root.right.val < root.val)
                || (root.val < min && root.val > max)) {
            return false;
        }

        if (root.left != null) {
            max = root.val - 1;
            leftTree = checkBST(root.left, min, max);
        }
        if (root.right != null) {
            min = root.val - 1;
            rightTree = checkBST(root.right, min, max);
        }
        return (leftTree && rightTree);
    }
}

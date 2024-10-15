import javax.swing.tree.TreeNode;
import java.util.ArrayList;
import java.util.List;

class Solution {
    int maxlevel = 0;

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> list = new ArrayList<>();
        right(root, 1, list);
        return list;
    }

    void right(TreeNode root, int level, List<Integer> list) {
        if (root == null) {
            return;
        }
        if (maxlevel < level) {
            list.add(root.val);
            maxlevel = level;
        }
        right(root.right, level + 1, list);
        right(root.left, level + 1, list);

    }
}
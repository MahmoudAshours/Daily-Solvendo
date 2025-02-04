import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int[] findMode(TreeNode root) {
        HashMap<Integer, Integer> traverser = new HashMap();
        ArrayList<Integer> s = new ArrayList();
        traverseTree(traverser, root);

        int maxRepitition = Integer.MIN_VALUE;
        for (Map.Entry entry : traverser.entrySet()) {
            if (maxRepitition <= (int) entry.getValue()) {
                maxRepitition = (int) entry.getValue();
                s.add((int) entry.getKey());
            }
        }

        int[] returned = new int[s.size() - 1];
        int i = 0;
        for (Integer in : s) {
            returned[i] = in;
            i++;
        }
        return returned;
    }

    public HashMap<Integer, Integer> traverseTree(HashMap<Integer, Integer> traverser, TreeNode root) {
        if (root != null) {
            traverseTree(traverser, root.left);
            traverseTree(traverser, root.right);
            traverser.put(root.val, traverser.getOrDefault(root.val, 0) + 1);
        }
        return traverser;
    }
}

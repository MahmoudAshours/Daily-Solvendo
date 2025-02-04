import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;

public class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        HashSet<Integer> set = new HashSet<>();
        ArrayList<Integer> ans = new ArrayList<>();

        for (int j : nums1) {
            set.add(j);
        }

        for (int j : nums2) {
            if (set.contains(j)) {
                ans.add(j);
                set.remove(j);
            }
        }
        int[] arr = new int[ans.size()];

        for (int i = 0; i < ans.size(); i++) {
            arr[i] = ans.get(i);
        }
        return arr;
    }
}

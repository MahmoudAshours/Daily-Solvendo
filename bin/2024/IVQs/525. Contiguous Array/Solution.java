import java.util.HashMap;

public class Solution {
    public int findMaxLength(int[] nums) {
// 0 1 0 1 0 1
//
// 1 1 2 2 3

// 0 0 0 1 1 0
// 0 0 1 2 2

// 1000100111
// 1 1 1 2 2 2 3 2 3 4
        int maxLength = 0;
        int count = 0;
        HashMap<Integer, Integer> countMap = new HashMap<>();
        countMap.put(0, -1);
        for (int i = 0; i < nums.length; i++) {
            count += nums[i] == 1 ? 1 : -1;
            if (countMap.containsKey(count)) {
                maxLength = Math.max(maxLength, i - countMap.get(count));
            } else {
                countMap.put(count, i);
            }
        }

        return maxLength;
    }
}

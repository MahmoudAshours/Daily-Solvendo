import java.util.HashMap;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        /**
         {
         a
         b
         c
         a? exists
         }
         a b b c

         a c b b
         */

        HashMap mp = new HashMap();

        int l = 0;
        int r = 0;
        int count = 1;

        if (s.isEmpty()) return 0;
        while (r < s.length()) {
            while (mp.containsKey(s.charAt(r))) {
                mp.remove(s.charAt(l));
                l++;
            }

            count = Math.max(count, r - l + 1);
            mp.put(s.charAt(r), 1);
            r++;
        }
        return count;
    }
}
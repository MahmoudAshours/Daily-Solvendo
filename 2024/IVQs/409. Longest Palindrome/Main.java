import java.util.HashMap;

public class Main {
  public static void main(String[] args) {
    Solution solution = new Solution();
    solution.longestPalindrome("abccccdd");
  }
}

class Solution {
  public int longestPalindrome(String s) {
    HashMap<Character, Integer> map = new HashMap<Character, Integer>();
    for (int i = 0; i < s.length(); i++) {
      map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
    }
    return 1;
  }
}

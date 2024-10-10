import java.util.Arrays;

public class Main {
  public static void main(String[] args) {
    Solution solution = new Solution();
    System.out.println(solution.heightChecker(new int[] { 1, 1, 4, 2, 1, 3 }));
  }
}

class Solution {
  public int heightChecker(int[] heights) {
    int[] expected = Arrays.copyOf(heights, heights.length);
    int counter = 0;
    Arrays.sort(heights);
    for (int i = 0; i < heights.length; i++) {
      if (heights[i] != expected[i]) {
        counter++;
      }
    }
    return counter;
  }
}

import java.util.*;

public class Main {
  public static void main(String[] args) {
    Solution solution = new Solution();
    solution.commonChars(new String[] { "bella", "label", "roller" });
  }
}

/**
 * b e l l a => map(b:1,e:1,l:1,a:1)
 * l a b e l => map ( , ,l:2,a:2)
 */


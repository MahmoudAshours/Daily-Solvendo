import java.util.HashMap;

public class Main {
  public static void main(String[] args) {
    Solution solution = new Solution();
    solution.relativeSortArray(new int[] { 2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19 }, new int[] { 2, 1, 4, 3, 9, 6 });
  }
}

class Solution {
  public int[] relativeSortArray(int[] arr1, int[] arr2) {

    int[] answer = new int[arr1.length];
    HashMap<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < arr1.length; i++) {
      map.put(arr1[i], map.getOrDefault(arr1[i], 1) + 1);
    }
    
    return answer;
  }
}

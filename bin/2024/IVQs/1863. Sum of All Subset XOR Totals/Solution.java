class Solution {

  public static void main(String[] args) {
    System.out.println(subsetXORSum(new int[] { 5, 1, 6 }));
  }

  public static int subsetXORSum(int[] nums) {
    return helper(nums, 0, 0);
  }

  public static int helper(int nums[], int level, int currentXOR) {
    if (level == nums.length)
      return currentXOR;
    int include = helper(nums, level + 1, currentXOR ^ nums[level]);
    int exclude = helper(nums, level + 1, currentXOR);
    return include + exclude;
  }
}

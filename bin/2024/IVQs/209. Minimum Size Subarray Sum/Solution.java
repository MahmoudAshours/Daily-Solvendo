class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int n = nums.length;
        int left = 0;
        int right = 0;
        int sum = 0;
        int min_len = 10000000;
        while (left < n) {
            // right = left;
            while (right < n && sum < target) {
                sum += nums[right];
                right++;
            }
            if (sum >= target) {
                min_len = Math.min(min_len, right - left);
            }
            sum -= nums[left];
            left++;

        }
        return min_len == 10000000 ? 0 : min_len;
    }
}
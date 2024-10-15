class Solution {
    public int largestAltitude(int[] gain) {
        // -5 1 5 0 -7
        int sum = 0;
        int max = Integer.MIN_VALUE;
        for (int i : gain) {
            sum += i;
            System.out.println(sum);
            if (sum >= max) {
                max = sum;
            }
        }
        return max > 0 ? max : 0;
    }
}
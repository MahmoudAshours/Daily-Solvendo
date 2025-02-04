import java.lang.Math;

public class Solution {
    public int pivotInteger(int n) {
        if (n == 1) return n;
        int[] prefixSum = new int[n + 1];
        int sum = 1;
        for (int i = 2; i <= n; i++) {
            sum = sum + i;
            prefixSum[i] = sum;
            System.out.println(prefixSum[i]);
        }
        return isPerfectSquareByUsingSqrt(prefixSum[n]) ? (int) Math.sqrt(prefixSum[n]) : -1;
    }

    public static boolean isPerfectSquareByUsingSqrt(long n) {
        if (n <= 0) {
            return false;
        }
        double squareRoot = Math.sqrt(n);
        long tst = (long) (squareRoot + 0.5);
        return tst * tst == n;
    }
}

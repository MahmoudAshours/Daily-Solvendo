class Solution {
    public boolean isPowerOfThree(int n) {
        if (n == 1) return true;
        if (n % 3 != 0 || n == 0) return false;
        double x = n;
        while (true) {
            x /= 3;
            if (x == 1) return true;
            else if (x != (int) x) return false;
        }
    }
}
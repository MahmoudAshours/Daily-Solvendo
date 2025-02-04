public class Solution {
    public boolean isHappy(int n) {
        if (n == 1) return false;
        if (n * n < 10) return false;
        while (n!= 1) {
            String numString = n + "";
            int result = 0;
            for (int i = 0; i < numString.length(); i++) {
                result +=  (numString.charAt(i)- '0') * (numString.charAt(i)- '0');
            }
            n = result;
        }
        return true;
    }
}

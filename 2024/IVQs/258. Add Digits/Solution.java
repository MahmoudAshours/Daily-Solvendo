class Solution {
    public int addDigits(int num) {
        boolean isDoubleDigit = (num > 9);
        if (!isDoubleDigit) return num;
        int digit, sum = 0;
        while (num > 0) {
            digit = num % 10;
            sum = sum + digit;
            num = num / 10;
        }
        return addDigits(sum);
    }
}
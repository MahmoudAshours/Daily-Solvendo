class Solution {
    public String maximumOddBinaryNumber(String s) {
        String res = "";
        int onesCount = 0;
        for (char c : s.toCharArray()) {
            if (c == '1') onesCount++;
        }
        int zerosCount = s.length() - onesCount;

        for (int i = 0; i < onesCount - 1; i++) {
            res = "1" + res;
        }

        for (int i = 0; i < zerosCount; i++) {
            res = res + "0";
        }
        res = res + "1";

        return res;
    }
}
class Solution {
    public boolean isPalindrome(String s) {
        int start = 0, end = s.length() - 1;

        while (start <= end) {
            char currStart = s.charAt(start);
            char currEnd = s.charAt(end);
            if (!Character.isLetterOrDigit(currStart)) {
                start++;
            } else if (!Character.isLetterOrDigit(currEnd)) {
                end--;
            } else {
                if (Character.toLowerCase(currStart) != Character.toLowerCase(currEnd)) {
                    return false;
                }
                start++;
                end--;
            }
        }

        return true;

    }
}
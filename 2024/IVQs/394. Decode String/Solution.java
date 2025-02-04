class Solution {
    public String decodeString(String s) {
        return dec(s);
    }

    public String dec(String s) {

        int pos = 0;
        while (pos < s.length()) {

            int numL = pos;
            int numR = numL;
            while (numL < s.length() && !Character.isDigit(s.charAt(numL))) numL++;
            if (numL == s.length()) return s;

            numR = numL;
            while (numR < s.length() && s.charAt(numR) != '[') numR++;
            // RESULT: we found pos of '[' which is goes right after the last difit of the number


            // 3. Parse number, now we now start and end pos
            int num = Integer.parseInt(s.substring(numL, numR));


            int wordL = numR;
            int wordR = wordL;
            int cnt = 0;
            while (wordR < s.length()) {
                if (s.charAt(wordR) == '[') cnt++;
                if (s.charAt(wordR) == ']') cnt--;
                if (cnt == 0) break;
                wordR++;
            }


            String decoded = dec(s.substring(wordL + 1, wordR));
            decoded = decoded.repeat(num);


            String left = s.substring(0, numL);
            String right = s.substring(wordR + 1, s.length());
            s = left + decoded + right;

            pos = left.length() + decoded.length();
        }

        return s;
    }
}
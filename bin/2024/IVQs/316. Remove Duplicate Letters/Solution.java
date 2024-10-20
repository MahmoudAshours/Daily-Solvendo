import java.util.HashMap;
import java.util.Stack;

class Solution {
    public String removeDuplicateLetters(String s) {
        HashMap<Character, Integer> freqArr = new HashMap();
        Stack<Character> stack = new Stack();
        boolean[] visited = new boolean[26];
        for (int i = 0; i < s.length(); i++) {
            freqArr.put(s.charAt(i), i);
        }
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            int vis = c - 'a';
            if (!visited[vis]) {
                while (!stack.isEmpty() && stack.peek() > c && freqArr.get(stack.peek()) > i) {
                    visited[stack.pop() - 'a'] = false;
                }
                stack.push(c);
                visited[vis] = true;
            }
        }
        StringBuilder result = new StringBuilder();
        for (Character chars : stack) {
            result.append(chars);
        }
        return result.toString();
    }
}
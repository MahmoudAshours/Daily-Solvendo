import java.util.Arrays;
import java.util.HashMap;

public class Solution {
    public String customSortString(String order, String s) {
        HashMap<Character, Integer> hashMap = new HashMap<Character, Integer>();
        StringBuilder orderedString = new StringBuilder();
        for (char c : s.toCharArray()) {
            hashMap.put(c, hashMap.getOrDefault(c, 0) + 1);
        }
        for (char c : order.toCharArray()) {
            if (hashMap.containsKey(c)) {
                orderedString.append(String.valueOf(c).repeat(Math.max(0, hashMap.get(c))));
                hashMap.remove(c);
            }
        }
        for (char c : hashMap.keySet()) {
            orderedString.append(String.valueOf(c).repeat(Math.max(0, hashMap.get(c))));
        }
        return orderedString.toString();
    }
}

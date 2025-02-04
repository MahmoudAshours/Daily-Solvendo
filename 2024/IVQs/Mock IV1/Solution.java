import java.util.HashMap;

public class Solution {


    public boolean solve(String s) {

        StringBuilder reversed = new StringBuilder(s).reverse();

        StringBuilder newS = new StringBuilder(s);

        HashMap<String, Integer> map = new HashMap();

        if (newS.toString().equals(reversed.toString())) return true;

        for (int i = 0; i < s.length(); i++) {
            while (reversed.charAt(i) != newS.charAt(i)) {
                System.out.println(newS.toString());
                newS.append(newS.charAt(i));
                newS.deleteCharAt(i);
                if (map.containsKey(newS.toString())) {
                    return false;
                } else {
                    map.put(newS.toString(), 1);
                }
            }
            System.out.println(newS.toString());
        }
        if (reversed.toString().equals(newS.toString())) return true;
        System.out.println(newS.toString());
        return false;
    }
}


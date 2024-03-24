public class Solution{
  public int lengthOfLastWord(String s) {
        var arr = s.trim().split(" ");
        return arr[arr.length-1].length();
    }
}

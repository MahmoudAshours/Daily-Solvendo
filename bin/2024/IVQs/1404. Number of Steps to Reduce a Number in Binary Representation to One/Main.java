
public class Main {
  public static void main(String[] args) {
    Solution solution = new Solution();
    System.out.println(solution.numSteps("1111011110000011100000110001011011110010111001010111110001"));
  }
}
class Solution {
    public int numSteps(String s) {
        int cnt = 0;
        int cary = 0;

        for(int i = s.length()-1; i >= 1; i--){
            int num = s.charAt(i) -'0';
            if(num == 0 && cary == 0){
                cnt++;
            }else if(num == 1 && cary == 1){
                cnt++;
                cary = 1;
            }else{
                cnt += 2;
                cary = 1;
            }
        }
        if(cary == 1)cnt++;
        return cnt;
    }
}

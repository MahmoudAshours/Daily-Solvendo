import java.util.Arrays;

public class Solution {
    public int leastInterval(char[] tasks, int n) {
        int[] frequency = new int[26];
        for (char task : tasks) {
            frequency[task - 'A']++;
        }
        Arrays.sort(frequency);
        int chunk = frequency[25] - 1;
        int idle = chunk * n;
        for (int i = 24; i >= 0; i--) {
            idle -= Math.min(chunk, frequency[i]);
        }
        return idle < 0 ? tasks.length : tasks.length + idle;
    }
}

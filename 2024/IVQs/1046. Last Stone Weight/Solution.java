import java.util.Arrays;
import java.util.Collections;
import java.util.PriorityQueue;

class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        for (int stone : stones) {
            maxHeap.add(stone);
        }

        int x = 0;
        int y = 0;
        while (maxHeap.size() > 1) {
            y = maxHeap.poll();
            x = maxHeap.poll();
            if (y > x) {
                maxHeap.offer(y - x);
            }
        }

        //If the maxHeap is empty, the expression evaluates to 0
        //If the maxHeap is not empty, the expression evaluates to maxHeap.poll()
        return maxHeap.isEmpty() ? 0 : maxHeap.poll();
    }
}
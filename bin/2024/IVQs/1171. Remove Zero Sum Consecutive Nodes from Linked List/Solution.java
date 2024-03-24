import java.util.HashMap;
import java.util.Map;

class Solution {
    public ListNode removeZeroSumSublists(ListNode head) {
        ListNode dummy = new ListNode();
        dummy.next = head;
        int prefixSum = 0;
        Map<Integer, ListNode> prefixSumToNode = new HashMap<>();

        for (ListNode current = dummy; current != null; current = current.next) {
            prefixSum += current.val;
            // If this prefix sum has been seen before, it means the sublist sums to zero
            if (prefixSumToNode.containsKey(prefixSum)) {
                ListNode prev = prefixSumToNode.get(prefixSum);
                ListNode toRemove = prev.next;
                int p = prefixSum + (toRemove != null ? toRemove.val : 0);

                while (p != prefixSum) {
                    prefixSumToNode.remove(p);
                    toRemove = toRemove.next;
                    p += (toRemove != null ? toRemove.val : 0);
                }
                prev.next = current.next;
            } else {
                prefixSumToNode.put(prefixSum, current);
            }
        }
        return dummy.next;
    }
}
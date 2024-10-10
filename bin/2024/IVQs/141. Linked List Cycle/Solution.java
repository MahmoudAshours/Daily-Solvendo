import java.util.HashMap;

/**
 * Definition for singly-linked list.
 * class ListNode {
 * int val;
 * ListNode next;
 * ListNode(int x) {
 * val = x;
 * next = null;
 * }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {

        HashMap<ListNode, Integer> mp = new HashMap();
        boolean exist = true;
        mp.put(head, 0);
        if (head == null) return false;
        ListNode tempHead = head.next;
        while (tempHead != null) {
            if (!mp.containsKey(tempHead)) {
                mp.put(tempHead, 0);
                tempHead = tempHead.next;
            } else {
                return true;
            }
        }
        return false;
    }
}
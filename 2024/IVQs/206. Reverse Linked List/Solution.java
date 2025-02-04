public class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode current = head;
        ListNode prev = null;

        while (true) {
            ListNode temp_curr = current.next;
            current.next = prev;
            prev = current;
            try {
                if (temp_curr == null) {
                    break;
                }
            } catch (NullPointerException e) {
                break;
            }
            current = temp_curr;
        }
        return current;
    }
}

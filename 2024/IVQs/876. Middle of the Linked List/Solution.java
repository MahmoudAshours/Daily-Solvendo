public class Solution {
    public ListNode middleNode(ListNode head) {
        ListNode fastforwardNode = head;
        try {
            if (head.next == null) return head;
        } catch (NullPointerException e) {
            return head.next;
        }
        try {
            fastforwardNode = fastforwardNode.next.next;
            while (head.next != null) {
                fastforwardNode = fastforwardNode.next.next;
                head = head.next;
            }
        } catch (Exception e) {
            return head.next;
        }
        return fastforwardNode;
    }
}

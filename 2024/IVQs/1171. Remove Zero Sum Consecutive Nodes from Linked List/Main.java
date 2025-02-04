public class Main {
    public static void main(String[] args) {
        Solution solution = new Solution();
        ListNode node = new ListNode();
        node.val = 1;
        node.next = new ListNode(2, new ListNode(3, new ListNode(-3, new ListNode(-2))));
        node = solution.removeZeroSumSublists(node);
        try {
            while (node != null) {
                System.out.println(node.val);
                node = node.next;
            }
        } catch (NullPointerException ignored) {
        }
    }
}

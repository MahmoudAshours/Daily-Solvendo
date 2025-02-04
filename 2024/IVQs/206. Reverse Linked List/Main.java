public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        ListNode listNode = new ListNode(1, new ListNode(2));
        ListNode node = sol.reverseList(listNode);
        try {
            while (node != null) {
                System.out.println(node.val);
                node = node.next;
            }
        } catch (NullPointerException e) {

        }
    }
}

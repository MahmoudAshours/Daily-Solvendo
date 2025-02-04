public class Main {
    public static void main(String[] args) {


        Solution solution = new Solution();

        ListNode l1 = new ListNode(10, new ListNode(1, new ListNode(13, new ListNode(6, new ListNode(9, new ListNode(5))))));
        ListNode l2 = new ListNode(1000000, new ListNode(1000001, new ListNode(1000002)));
        ListNode l3 = solution.mergeInBetween(l1, 3, 4, l2);

        try {
            while (l3 != null) {
                System.out.println(l3.val);
                l3 = l3.next;
            }
        } catch (Exception e) {

        }
    }
}

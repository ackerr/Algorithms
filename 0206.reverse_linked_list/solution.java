public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode ans = null;
        while (head != null){
            ListNode temp = head.next;
            head.next = ans;
            ans = head;
            head = temp;
        }
        return ans;
    }
}
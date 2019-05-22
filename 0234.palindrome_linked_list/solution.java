public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class Solution {
    public boolean isPalindrome(ListNode head) {
        if (head == null) return true;

        ListNode fast = head, slow = head;
        ListNode pre = null;
        while (fast != null && fast.next != null){
            if (fast.next != null)
                fast = fast.next.next;
            else
                fast = fast.next;
            slow = slow.next;
        }

        while (slow != null){
            ListNode temp = slow.next;
            slow.next = pre;
            pre = slow;
            slow = temp;
        }

        while (pre != null){
            if (pre.val != head.val) return false;
            pre = pre.next;
            head = head.next;
        }
        return true;
    }
}
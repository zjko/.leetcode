/*
 * @lc app=leetcode.cn id=2 lang=java
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list. public class ListNode { int val; ListNode
 * next; ListNode(int x) { val = x; } }
 */

class Solution {

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode res, p;
        boolean carry = false;
        res = p = new ListNode(0);
        while (l1 != null && l2 != null) {
            int val = (l1.val + l2.val + (carry ? 1 : 0)) % 10;
            carry = (l1.val + l2.val + (carry ? 1 : 0)) > 9;

            // System.out.println(l1.val + " " + l2.val + " " + carry + " " + val);

            p.val = val;
            if (l1.next != null || l2.next != null) {
                p.next = new ListNode(0);
                p = p.next;
            }
            l1 = l1.next;
            l2 = l2.next;

        }

        ListNode t = l1 != null ? l1 : l2;

        while (t != null) {
            // System.out.println(t.val);
            int val = (t.val + (carry ? 1 : 0)) % 10;
            carry = (t.val + (carry ? 1 : 0)) > 9;
            p.val = val;
            if (t.next != null) {
                p.next = new ListNode(0);
                p = p.next;
            }
            t = t.next;
        }
        if(carry){
            p.next = new ListNode(1);
        }

        return res;
    }
}
// [1,2]\n[9,9,9]
// [9,1,6]\n[0]
// @lc code=end

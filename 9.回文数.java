/*
 * @lc app=leetcode.cn id=9 lang=java
 *
 * [9] 回文数
 */

// @lc code=start
class Solution {
    public boolean isPalindrome(int x) {
        if(x <0 || (x != 0 &&x % 10 == 0)){
            return false;
        }
        int a = 0;
        while(a<x){
            a = a * 10 + x%10;
            x/=10;
            //System.out.println(a + " " + x);
        }

        return a == x || a/10 == x;
    }
}

// 1223221
// @lc code=end


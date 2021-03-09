/*
 * @lc app=leetcode.cn id=7 lang=java
 *
 * [7] 整数反转
 */

// @lc code=start
class Solution {
    public int reverse(int x) {
        boolean negative = x<0;
        x = Math.abs(x);
        boolean zero = true;
        long res = 0L;
        while(x!=0){
            int a = x % 10;
            if(a != 0 || !zero) {
                zero = false;
                res *= 10L;
                res += a;
            }
            x /=10;
        }
        res *= (negative?-1L:1L);
        return res<-2147483648 || res>2147483647?0:(int)res;
    }
}
// @lc code=end


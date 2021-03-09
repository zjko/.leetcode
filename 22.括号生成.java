/*
 * @lc app=leetcode.cn id=22 lang=java
 *
 * [22] 括号生成
 */

// @lc code=start
class Solution {
    private static List<String> res;
    public static List<String> generateParenthesis(int n) {
        res = new LinkedList<>();
        func("(",n,0);
        return res;
    }

    /**
     *
     * @param c 已生成的字符串
     * @param level 括号的层级 ((( 表示当前层级为3  (() 表示当前层级为1
     * @param n 开括号， 有多少个开括号 就必须有多少个闭括号
     */
    private static void func(String c,int level,int n){
//        层级为1 表示到底了 不能再深入了
        if(level == 1){
            String newString = c;
            while (n-->-1){
                newString = newString.concat(")");
            }
            res.add(newString);
            return;
        }
//        每个括号后面有两种分支可能  分别为 ( 和 ) 其中(表示层级进一层，开括号多一个
        func(c.concat("("),level-1,n+1);
        if(n > -1){
            func(c.concat(")"), level ,n - 1);
        }
    }
}
// @lc code=end


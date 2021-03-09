import java.util.HashSet;

/*
 * @lc app=leetcode.cn id=3 lang=java
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
class Solution {
    public int lengthOfLongestSubstring(String s) {
        // System.out.println(s);
        int max = 0;
        int length = s.length();
        HashSet<Character> set = new HashSet<>(length);
        for(int i=0;i<length; i++){
            char c = s.charAt(i);
            // System.out.println(c);
            if(!set.contains(c)){
                set.add(c);
            }else{
                if(max < set.size()){
                    max = set.size();
                }
                set.clear();
            }
        }
        // System.out.println(set);
        return set.size();
    
    }
}
// "pwwkew"
// @lc code=end


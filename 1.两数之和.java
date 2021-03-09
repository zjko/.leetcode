/*
 * @lc app=leetcode.cn id=1 lang=java
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {

    // public static void main(String[] args) {
    //     twoSum(new int[]{3,2,4},6);
    // }

    public int[] twoSum(int[] nums, int target) {
        int size = nums.length;
        int[] res = new int[2];
        for(int x = 0; x<size-1;x++){
            for(int y=x+1;y<size;y++){
                //System.out.println(x+" "+y);
                if(nums[x]+nums[y] == target){
                    res[0] = x;
                    res[1] = y;
                    return res;
                }
            }
        }
        return null;
    }
}
// @lc code=end


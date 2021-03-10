/*
 * @lc app=leetcode.cn id=16 lang=java
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int len = nums.length;
        int a = 0, b, c;
        int res = 0;
        int answer = Integer.MAX_VALUE;
        while (a < len - 2) {
            b = a + 1;
            c = len - 1;
            while (b < c) {
                int sum = nums[a] + nums[b] + nums[c];
//                print(nums,a,b,c);
                if (sum == target) {
                    return sum;
                }
                int ans = Math.abs(sum - target);
//            System.out.println(ans +" " + answer);
                if(answer > ans){
                    answer = ans;
                    res = sum;
                }

                if (sum > target) {
                    c = prev(nums, c);
                } else {
                    b = next(nums, b);
                }
            }
            a = next(nums, a);
        }
        return res;
    }

    private int prev(int[] num, int a) {
        a--;
        while (a > 0 && num[a] == num[a + 1]) {
            a--;
        }
        return a;
    }

    private int next(int[] num, int a) {
        a++;
        int len = num.length - 1;
        while (a < len && num[a - 1] == num[a]) {
            a++;
        }
        return a;
    }

    private void print(int[] num, int a, int b, int c) {
        int sum = num[a] + num[b] + num[c];
        System.out.print(a + " " + b + " " + c + " " + sum + " ||");
        for (int i : num) {
            System.out.print(i + " ");
        }
        System.out.println();
    }
}
// @lc code=end
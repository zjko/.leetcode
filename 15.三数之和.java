import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashSet;

/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
    public static List<List<Integer>> threeSum(int[] nums) {
        // 结果集
        List<List<Integer>> res = new ArrayList<>();
        // step1 排序
        Arrays.sort(nums);
        // step2 计算 -nums[i] = num[n] + num[p]
        for (int i = 0; i < nums.length; i++) {
            // step3 找出 n 和 p 具体的值
            int n = i + 1, p = nums.length - 1;
            while (n < p) {
                switch (Integer.compare(nums[n] + nums[p], -nums[i])) {
                    case 1:
                        // 结果大于目标值 右指针左移
                        p--;
                        break;
                    case 0:
                        List<Integer> case1 = new ArrayList<>();
                        case1.add(nums[i]);
                        case1.add(nums[n]);
                        case1.add(nums[p]);
                        res.add(case1);
                        // 结果去重
                        while (n + 1 < nums.length && nums[n] == nums[n + 1]) {
                            n++;
                        }
                        while (p - 1 > -1 && nums[p] == nums[p - 1]) {
                            p--;
                        }
                        while (i + 1 < nums.length && nums[i] == nums[i + 1]) {
                            i++;
                        }
                        n++;
                        p--;
                        break;
                    case -1:
                        // 结果小于目标值 左指针右移
                        n++;
                        break;
                }
            }

        }
        return res;
    }

}
// @lc code=end

/*
 * @lc app=leetcode.cn id=12 lang=java
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
class Solution {
    private Map<Integer,String> word;
    {
        word = new HashMap<>();
        word.put(1,     "I");
        word.put(2,     "II");
        word.put(3,     "III");
        word.put(4,     "IV");
        word.put(5,     "V");
        word.put(6,     "VI");
        word.put(7,     "VII");
        word.put(8,     "VIII");
        word.put(9,     "IX");
        word.put(10,    "X");
        word.put(20,    "XX");
        word.put(30,    "XXX");
        word.put(40,    "XL");
        word.put(50,    "L");
        word.put(60,    "LX");
        word.put(70,    "LXX");
        word.put(80,    "LXXX");
        word.put(90,    "XC");
        word.put(100,   "C");
        word.put(200,   "CC");
        word.put(300,   "CCC");
        word.put(400,   "CD");
        word.put(500,   "D");
        word.put(600,   "DC");
        word.put(700,   "DCC");
        word.put(800,   "DCCC");
        word.put(900,   "CM");
        word.put(1000,  "M");
        word.put(2000,  "MM");
        word.put(3000,  "MMM");
    }
    private int[] list = new int[]{1,2,3,4,5,6,7,8,9,10,20,30,40,50,60,70,80,90,100,200,300,400,500,600,700,800,900,1000,2000,3000};

    public String intToRoman(int num) {

        StringBuilder res = new StringBuilder();
        while (num > 0) {
            int key = getKey(num);
            num -= list[key];
            res.append(word.get(list[key]));
            System.out.println(num + " " + list[key] +" "+ res.toString());
        }


        return res.toString();
    }


     //找到一个最接近且小于等于num的项
     private int getKey(int num){

        int l =0, r = 29;

        int mid = 0;
        while(l <= r){
            mid = (l+r) /2;
//            System.out.println(mid + " " + list[mid]);
            if(num == list[mid]){
                return mid;
            }else if(num < list[mid]){
                r = mid - 1;
            }else {
                l = mid + 1;
            }

        }
        if(list[mid] > num){
            mid--;
        }
        System.out.println(num + " " + list[mid]);
        return mid;
     }
}
// @lc code=end


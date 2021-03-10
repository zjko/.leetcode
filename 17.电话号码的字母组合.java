/*
 * @lc app=leetcode.cn id=17 lang=java
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
class Solution {
    private Map<String,String[]> map;
    {
        map = new HashMap<>();
        map.put("2",new String[]{"a","b","c"});
        map.put("3",new String[]{"d","e","f"});
        map.put("4",new String[]{"g","h","i"});
        map.put("5",new String[]{"j","k","l"});
        map.put("6",new String[]{"m","n","o"});
        map.put("7",new String[]{"p","q","r","s"});
        map.put("8",new String[]{"t","u","v"});
        map.put("9",new String[]{"w","x","y","z"});
    }

    private List<String> list = new LinkedList<>();

    public List<String> letterCombinations(String digits) {
        if("".equals(digits)){
            return new LinkedList();
        }
        getString(digits,"");
        return this.list;
    }
    
    private void getString(String digits,String str){
        if(digits.length() > 0){
            String[] strings = map.get(digits.substring(0, 1));
            for(String s:strings){
                getString(
                        digits.substring(1),str + s
                );
            }
        }else {
            list.add(str);
        }
    }
}
// @lc code=end


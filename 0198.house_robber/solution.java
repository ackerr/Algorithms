class Solution {
    public int rob(int[] nums){
        int last = 0, now = 0;
        for (int num: nums){
            int temp = now;
            now = Math.max(num + last, now);
            last = temp;
        }
        return now;
    }
}
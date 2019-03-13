class solution{
    public int majorityElement(int[] nums){
        int total = 0, ans = 0;
        for (int num: nums){
            if (total == 0)
                ans = num;
            if (ans == num)
                total++;
            else
                total--;
        }
        return ans;
    }

    public int otherSolution(int[] nums){
        Arrays.sort(nums);
        return nums[nums.length/2]
    }
}
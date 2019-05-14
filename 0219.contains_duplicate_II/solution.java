class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Map<Integer, Integer> storage = new HashMap<>();
        for (int i=0;i<nums.length;i++){
            if (storage.containsKey(nums[i])){
                if ((i - storage.get(nums[i])) <= k){
                    return true;
                }
            }
            storage.put(nums[i], i);
        }
        return false;
    }
}
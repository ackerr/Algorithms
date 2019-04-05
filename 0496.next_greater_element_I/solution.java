class Solution {
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Map<Integer, Integer> values = new HashMap<>();
        Stack<Integer> stack = new Stack<>();

        for (int num: nums2){
            while (!stack.isEmpty() && stack.peek() < num){
                values.put(stack.pop(), num);
            }
            stack.push(num);
        }

        int[] ans = new int[nums1.length];
        for (int i=0;i<nums1.length;i++){
            ans[i] = values.getOrDefault(nums1[i], -1);
        }
        return ans;
    }
}
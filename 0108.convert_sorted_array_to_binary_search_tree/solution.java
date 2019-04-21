 public class TreeNode {
     int val;
     TreeNode left;
     TreeNode right;
     TreeNode(int x) { val = x; }
 }


class Solution {
    public TreeNode sortedArrayToBST(int[] nums) {
        return sortedArray(nums, 0, nums.length-1);
    }
    
    public TreeNode sortedArray(int[] nums, int start, int end){
        if (start > end) return null;
        if (start == end) return new TreeNode(nums[start]);
        int mid = (start + end) / 2;
        TreeNode root = new TreeNode(nums[mid]);
        root.left = sortedArray(nums, start, mid-1);
        root.right = sortedArray(nums, mid+1, end);
        return root;
    }
}
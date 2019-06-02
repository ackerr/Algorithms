public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public int pathSum(TreeNode root, int sum) {
        if (root == null) return 0;
        return Sum(root, sum) + pathSum(root.left, sum) + pathSum(root.right, sum);
        
    }

    public int Sum(TreeNode root, int sum){
        if (root == null) return 0;
        int count = 0;
        if (root.val == sum){
            count++;
        }
        count += Sum(root.left, sum-root.val);
        count += Sum(root.right, sum-root.val);
        return count;
    }
}
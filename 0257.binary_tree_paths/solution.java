import java.util.ArrayList;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public List<String> binaryTreePaths(TreeNode root) {
        ArratList<String> paths = new ArrayList<>();
        String temp = "";
        deep(root, paths, temp);
        return paths;
    }

    public void deep(List<TreeNode> root, List<String> paths, String temp){
        if (root == null) return paths;
        temp += root.val;
        if (root.left == null && root.right == null){
            paths.add(temp);
        }
        deep(root.left, paths, temp+"->");
        deep(root.right, paths, temp+"->");
    }
}
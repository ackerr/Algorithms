import java.util.LinkedList;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        LinkedList<List<Integer>> ans = new LinkedList<List<Integer>>();
        addLevel(root, 0, ans);
        return ans;
    }

    public void addLevel(TreeNode node, int level, LinkedList<List<Integer>> list){
        if (node == null){
            return;
        }
        if (level+1 > list.size()){
            list.addFirst(new LinkedList<Integer>());
        }
        list.get(list.size()-1-level).add(node.val);
        addLevel(node.left, level+1, list);
        addLevel(node.right, level+1, list);
    }
}
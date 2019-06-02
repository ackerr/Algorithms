import java.util.List;

class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
}

class Solution {
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> ans = new ArrayList<>();
        if (root == null) return ans;
        Queue<Node> nodes = new LinkedList<>();
        nodes.add(root);
        while (!nodes.isEmpty()){
            int size = nodes.size();
            List<Integer> temp = new ArrayList<>();
            for (int i=0; i< size; i++){
                temp.add(nodes.peek().val);
                nodes.addAll(nodes.poll().children);
            }
            ans.add(temp);
        }
        return ans;
    }
}
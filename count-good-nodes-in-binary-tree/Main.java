public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    public int goodNodes(TreeNode root) {
        return dfs(root, root, -101);
    }

    private int dfs(TreeNode root, TreeNode node, int parentMax) {
        if (node == null) {
            return 0;
        }
        int result = 0;
        if (node.val >= parentMax) {
            result++;
        }
        int currentMax = Math.max(node.val, parentMax);
        result += dfs(root, node.left, currentMax);
        result += dfs(root, node.right, currentMax);
        return result;
    }
}

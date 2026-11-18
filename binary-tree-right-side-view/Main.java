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
    // public List<Integer> rightSideView(TreeNode root) {
    //     if (root == null) {
    //         return Collections.emptyList();
    //     }
    //     List<Integer> result = new ArrayList<>();
    //     Queue<TreeNode> queue = new ArrayDeque<>();
    //     queue.offer(root);
    //     while (!queue.isEmpty()) {
    //         int count = queue.size();
    //         for (int i = 0; i < count; i++) {
    //             TreeNode x = queue.poll();
    //             if (x.left != null) {
    //                 queue.offer(x.left);
    //             }
    //             if (x.right != null) {
    //                 queue.offer(x.right);
    //             }
    //             if (i == count - 1) {
    //                 result.add(x.val);
    //             }
    //         }
    //     }
    //     return result;
    // }

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        dfs(root, 0, result);
        return result;
    }

    private void dfs(TreeNode root, int depth, List<Integer> result) {
        if (root == null) {
            return;
        }
        if (depth == result.size()) {
            result.add(root.val);
        }
        dfs(root.right, depth + 1, result);
        dfs(root.left, depth + 1, result);
    }
}

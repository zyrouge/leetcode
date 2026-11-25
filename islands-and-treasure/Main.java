class Solution {
    public void islandsAndTreasure(int[][] grid) {
        Queue<int[]> queue = new LinkedList<>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 0) {
                    queue.offer(new int[] { i, j });
                }
            }
        }
        int[][] directions = {{ -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 }};
        while (!queue.isEmpty()) {
            int[] position = queue.poll();
            for (int[] x : directions) {
                int i = position[0] + x[0];
                int j = position[1] + x[1];
                if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != Integer.MAX_VALUE) {
                    continue;
                }
                grid[i][j] = grid[position[0]][position[1]] + 1;
                queue.offer(new int[] { i, j });
            }
        }
    }

    public void dfs(int[][] grid, int count, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == -1 || grid[i][j] != Integer.MAX_VALUE) {
            return;
        }
        grid[i][j] = count;
        dfs(grid, count + 1, i - 1, j);
        dfs(grid, count + 1, i + 1, j);
        dfs(grid, count + 1, i, j - 1);
        dfs(grid, count + 1, i, j + 1);
    }
}

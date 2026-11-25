class Solution {
    public int maxAreaOfIsland(int[][] grid) {
        int result = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 1) {
                    result = Math.max(result, dfs(grid, i, j));
                    // result = Math.max(result, bfs(grid, i, j));
                }
            }
        }
        return result;
    }

    public int dfs(int[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != 1) {
            return 0;
        }
        grid[i][j] = 0;
        return 1 + dfs(grid, i - 1, j) + dfs(grid, i + 1, j) + dfs(grid, i, j - 1) + dfs(grid, i, j + 1);
    }

    // public int bfs(int[][] grid, int row, int column) {
    //     int result = 0;
    //     Queue<int[]> queue = new LinkedList<>();
    //     queue.add(new int[] { row, column });
    //     while (!queue.isEmpty()) {
    //         int[] position = queue.poll();
    //         int i = position[0], j = position[1];
    //         if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != 1) {
    //             continue;
    //         }
    //         grid[i][j] = 0;
    //         result++;
    //         queue.offer(new int[] { i - 1, j });
    //         queue.offer(new int[] { i + 1, j });
    //         queue.offer(new int[] { i, j - 1 });
    //         queue.offer(new int[] { i, j + 1 });
    //     }
    //     return result;
    // }
}

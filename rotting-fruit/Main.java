class Solution {
    public int orangesRotting(int[][] grid) {
        int freshCount = 0;
        Queue<int[]> rotten = new LinkedList<>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                switch (grid[i][j]) {
                    case 1:
                        freshCount++;
                        break;
                    case 2:
                        rotten.offer(new int[] { i, j });                    
                }
            }
        }
        int minutes = 0;
        int[][] directions = new int[][] {{ -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 }};
        while (freshCount > 0 && !rotten.isEmpty()) {
            minutes++;
            int rottenCount = rotten.size();
            for (int p = 0; p < rottenCount; p++) {
                int[] position = rotten.poll();
                for (int[] x : directions) {
                    int i = position[0] + x[0];
                    int j = position[1] + x[1];
                    if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != 1) {
                        continue;
                    }
                    freshCount--;
                    grid[i][j] = 2;
                    rotten.offer(new int[] { i, j });
                }
            }
        }
        if (freshCount > 0) {
            return -1;
        }
        return minutes;
    }
}

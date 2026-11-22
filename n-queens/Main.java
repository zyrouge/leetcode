class Solution {
    // public List<List<String>> solveNQueens(int n) {
    //     List<List<String>> result = new ArrayList<>();
    //     boolean[][] board = new boolean[n][n];
    //     dfs(n, 0, board, result);
    //     return result;
    // }

    // private void dfs(int n, int i, boolean[][] board, List<List<String>> result) {
    //     if (i == n) {
    //         result.add(this.stringifyBoard(n, board));
    //         return;
    //     }
    //     for (int j = 0; j < n; j++) {
    //         if (canAcquire(n, board, i, j)) {
    //             board[i][j] = true;
    //             dfs(n, i + 1, board, result);
    //             board[i][j] = false;
    //         }
    //     }
    // }

    // private boolean canAcquire(int n, boolean[][] board, int x, int y) {
    //     for (int i = 0; i < n; i++) {
    //         if (board[x][i] || board[i][y] ||
    //             (x + i < n && y + i < n && board[x + i][y + i]) || (x - i >= 0 && y - i >= 0 && board[x - i][y - i]) ||
    //             (x + i < n && y - i >= 0 && board[x + i][y - i]) || (x - i >= 0 && y + i < n && board[x - i][y + i])) {
    //             return false;
    //         }
    //     }
    //     return true;
    // }

    // private List<String> stringifyBoard(int n, boolean[][] board) {
    //     List<String> result = new ArrayList<>();
    //     StringBuilder builder = new StringBuilder();
    //     for (int i = 0; i < n; i++) {
    //         for (int j = 0; j < n; j++) {
    //             builder.append(board[i][j] ? 'Q' : '.');
    //         }
    //         result.add(builder.toString());
    //         builder.setLength(0);
    //     }
    //     return result;
    // }

    public List<List<String>> solveNQueens(int n) {
        List<List<String>> result = new ArrayList<>();
        char[][] board = new char[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                board[i][j] = '.';
            }
        }
        dfs(n, 0, board, new ArrayList<>(), result);
        return result;
    }

    private void dfs(int n, int i, char[][] board, List<String> current, List<List<String>> result) {
        if (i == n) {
            result.add(new ArrayList<>(current));
            return;
        }
        for (int j = 0; j < n; j++) {
            if (canAcquire(n, board, i, j)) {
                board[i][j] = 'Q';
                current.add(new String(board[i]));
                dfs(n, i + 1, board, current, result);
                current.removeLast();
                board[i][j] = '.';
            }
        }
    }

    private boolean canAcquire(int n, char[][] board, int x, int y) {
        for (int i = 0; i < n; i++) {
            if (board[x][i] == 'Q' || board[i][y] == 'Q' ||
                (x + i < n && y + i < n && board[x + i][y + i] == 'Q') || (x - i >= 0 && y - i >= 0 && board[x - i][y - i] == 'Q') ||
                (x + i < n && y - i >= 0 && board[x + i][y - i] == 'Q') || (x - i >= 0 && y + i < n && board[x - i][y + i] == 'Q')) {
                return false;
            }
        }
        return true;
    }
}
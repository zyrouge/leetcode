class Solution {
    public void solve(char[][] board) {
        int rowCount = board.length;
        int columnCount = board[0].length;
        for (int i = 0; i < rowCount; i++) {
            if (board[i][0] == 'O') {
                dfs(board, i, 0);
            }
            if (board[i][columnCount - 1] == 'O') {
                dfs(board, i, columnCount - 1);
            }
        }
        for (int i = 0; i < columnCount; i++) {
            if (board[0][i] == 'O') {
                dfs(board, 0, i);
            }
            if (board[rowCount - 1][i] == 'O') {
                dfs(board, rowCount - 1, i);
            }
        }
        for (int i = 0; i < rowCount; i++) {
            for (int j = 0; j < columnCount; j++) {
                if (board[i][j] == 'O') {
                    board[i][j] = 'X';
                } else if (board[i][j] == 'T') {
                    board[i][j] = 'O';
                }
            }
        }
    }

    public void dfs(char[][] board, int i, int j) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] != 'O') {
            return;
        }
        board[i][j] = 'T';
        dfs(board, i - 1, j);
        dfs(board, i + 1, j);
        dfs(board, i, j - 1);
        dfs(board, i, j + 1);
    }
}
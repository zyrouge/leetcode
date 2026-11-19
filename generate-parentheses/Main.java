class Solution {
    public List<String> generateParenthesis(int n) {
        Set<String> results = new HashSet<>();
        dfs(n, "", 0, 0, results);
        return new ArrayList<>(results);
    }

    public void dfs(int n, String current, int opened, int closed, Set<String> results) {
        if (opened == n && closed == n) {
            results.add(current);
            return;
        }
        if (opened < n) {
            dfs(n, current + "(", opened + 1, closed, results);
        }
        if (closed < opened) {
            dfs(n, current + ")", opened, closed + 1, results);
        }
    }
}

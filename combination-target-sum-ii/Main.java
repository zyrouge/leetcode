class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        List<List<Integer>> results = new ArrayList<>();
        Arrays.sort(candidates);
        dfs(candidates, target, 0, new ArrayList<>(), 0, results);
        return results;
    }

    private void dfs(int[] candidates, int target, int i, List<Integer> current, int currentSum, List<List<Integer>> results) {
        if (currentSum == target) {
            results.add(new ArrayList<>(current));
        }
        if (currentSum >= target || i == candidates.length) {
            return;
        }
        current.add(candidates[i]);
        currentSum += candidates[i];
        dfs(candidates, target, i + 1, current, currentSum, results);
        current.removeLast();
        currentSum -= candidates[i];
        while (i + 1 < candidates.length && candidates[i] == candidates[i + 1]) {
            i++;
        }
        dfs(candidates, target, i + 1, current, currentSum, results);
    }
}

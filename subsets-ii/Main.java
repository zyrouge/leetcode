class Solution {
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> path = new ArrayList<>();
        dfs(nums, 0, path, result);
        return result;
    }

    private void dfs(int[] nums, int i, List<Integer> path, List<List<Integer>> result) {
        if (i == nums.length) {
            result.add(new ArrayList<>(path));
            return;
        }
        path.add(nums[i]);
        dfs(nums, i + 1, path, result);
        path.removeLast();
        while (i + 1 < nums.length && nums[i] == nums[i + 1]) {
            i++;
        }
        dfs(nums, i + 1, path, result);
    }
}

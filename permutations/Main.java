class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> results = new ArrayList<>();
        // dfs(nums, new LinkedHashSet<>(), results);
        dfs(nums, 0, results);
        return results;
    }

    // private void dfs(int[] nums, Set<Integer> current, List<List<Integer>> results) {
    //     if (current.size() == nums.length) {
    //         results.add(new ArrayList<>(current));
    //         return;
    //     }
    //     for (int x : nums) {
    //         if (current.contains(x)) {
    //             continue;
    //         }
    //         current.add(x);
    //         dfs(nums, current, results);
    //         current.remove(x);
    //     }
    // }

    private void dfs(int[] nums, int start, List<List<Integer>> results) {
        if (start == nums.length) {
            results.add(Arrays.stream(nums).boxed().collect(Collectors.toList()));
            return;
        }
        for (int i = start; i < nums.length; i++) {
            swap(nums, start, i);
            dfs(nums, start + 1, results);
            swap(nums, i, start);
        }
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}

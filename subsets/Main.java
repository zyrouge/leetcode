class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        result.add(Collections.emptyList());
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            int rn = result.size();
            for (int j = 0; j < rn; j++) {
                List<Integer> y = new ArrayList<>(result.get(j));
                y.add(x);
                result.add(y);
            }
        }
        return result;
    }
}

class Solution {
    // public int findKthLargest(int[] nums, int k) {
    //     PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Comparator.reverseOrder());
    //     for (int x : nums) {
    //         maxHeap.offer(x);
    //     }
    //     for (int i = 0; i < k - 1; i++) {
    //         maxHeap.poll();
    //     }
    //     return maxHeap.peek();
    // }

    // public int findKthLargest(int[] nums, int k) {
    //     PriorityQueue<Integer> maxHeap = new PriorityQueue<>();
    //     for (int x : nums) {
    //         maxHeap.offer(x);
    //         if (maxHeap.size() > k) {
    //             maxHeap.poll();
    //         }
    //     }
    //     return maxHeap.peek();
    // }

    public int findKthLargest(int[] nums, int k) {
        return quickSelect(nums, nums.length - k, 0, nums.length - 1);
    }

    private int quickSelect(int[] nums, int k, int left, int right) {
        int swapTo = left;
        for (int i = left; i < right; i++) {
            if (nums[i] <= nums[right]) {
                swap(nums, i, swapTo);
                swapTo++;
            }
        }
        swap(nums, right, swapTo);
        if (k == swapTo) {
            return nums[swapTo];
        }
        if (k < swapTo) {
            return quickSelect(nums, k, left, swapTo - 1);
        }
        return quickSelect(nums, k, swapTo + 1, right);
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
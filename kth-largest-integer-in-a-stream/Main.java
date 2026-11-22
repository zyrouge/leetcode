class KthLargest {
    // private int k;
    // private List<Integer> nums;

    // public KthLargest(int k, int[] nums) {
    //     this.k = k;
    //     this.nums = new ArrayList<>();
    //     for (int x : nums) {
    //         this.nums.add(x);
    //     }
    // }
    
    // public int add(int val) {
    //     this.nums.add(val);
    //     Collections.sort(this.nums);
    //     return this.nums.get(this.nums.size() - this.k);
    // }

    private int k;
    private PriorityQueue<Integer> heap;

    public KthLargest(int k, int[] nums) {
        this.k = k;
        this.heap = new PriorityQueue<>();
        for (int x : nums) {
            this.add(x);
        }
    }
    
    public int add(int val) {
        heap.offer(val);
        if (heap.size() > this.k) {
            heap.poll();
        }
        return heap.peek();
    }
}

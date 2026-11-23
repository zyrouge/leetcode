class Solution {
    public int[][] kClosest(int[][] points, int k) {
        PriorityQueue<int[]> minHeap = new PriorityQueue<>(Comparator.comparingInt(x -> x[0]));
        for (int[] x : points) {
            int distance = (x[0] * x[0]) + (x[1] * x[1]);
            minHeap.offer(new int[] { distance, x[0], x[1] });
        }
        int[][] result = new int[k][2];
        for (int i = 0; i < k; i++) {
            int[] x = minHeap.poll();
            result[i][0] = x[1];
            result[i][1] = x[2];
        }
        return result;
    }
}
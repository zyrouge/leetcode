class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        for (int x : stones) {
            maxHeap.add(x);
        }
        while (maxHeap.size() > 1) {
            int y = maxHeap.poll();
            int x = maxHeap.poll();
            if (x == y) {
                continue;
            }
            maxHeap.add(y - x);
        }
        return !maxHeap.isEmpty() ? maxHeap.peek() : 0;
    }
}
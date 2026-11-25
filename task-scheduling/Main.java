class Solution {
    public int leastInterval(char[] tasks, int n) {
        int[] counts = new int[26];
        for (char x : tasks) {
            counts[x - 'A']++;
        }
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        for (int x : counts) {
            if (x > 0) {
                maxHeap.offer(x);
            }
        }
        int result = 0;
        Queue<int[]> queue = new LinkedList<>();
        int i = 25;
        while (!maxHeap.isEmpty() || !queue.isEmpty()) {
            result++;
            if (!maxHeap.isEmpty()) {
                int current = maxHeap.poll();
                if (current > 1) {
                    queue.add(new int[] { current - 1, result + n });
                }
            }
            if (!queue.isEmpty() && queue.peek()[1] == result) {
                maxHeap.offer(queue.poll()[0]);
            }
        }
        return result;
    }
}

class Solution {
    public int carFleet(int target, int[] position, int[] speed) {
        Integer[] ordered = new Integer[position.length];
        for (int i = 0; i < position.length; i++) {
            ordered[i] = i;
        }
        // for (int i = 0; i < position.length; i++) {
        //     for (int j = i + 1; j < position.length; j++) {
        //         if (position[ordered[j]] < position[ordered[i]]) {
        //             int temp = ordered[i];
        //             ordered[i] = ordered[j];
        //             ordered[j] = temp;
        //         }
        //     }
        // }
        Arrays.sort(ordered, (a, b) -> Integer.compare(position[a], position[b]));
        int result = 0;
        double previousTime = 0;
        for (int i = position.length - 1; i >= 0; i--) {
            int x = ordered[i];
            double currentTime = (double) (target - position[x]) / speed[x];
            if (currentTime > previousTime) {
                result++;
                previousTime = currentTime;
            }
        }
        return result;
    }
}

class Solution {
    // public int[] plusOne(int[] digits) {
    //     Queue<Integer> queue = new LinkedList<>();
    //     int carry = 1;
    //     for (int i = digits.length - 1; i >= 0; i--) {
    //         int current = digits[i] + carry;
    //         queue.offer(current % 10);
    //         carry = current == 10 ? 1 : 0;
    //     }
    //     if (carry != 0) {
    //         queue.offer(carry);
    //     }
    //     int[] result = new int[queue.size()];
    //     for (int i = result.length - 1; i >= 0; i--) {
    //         result[i] = queue.poll();
    //     }
    //     return result;
    // }

    public int[] plusOne(int[] digits) {
        int n = digits.length;
        for (int i = n - 1; i >= 0; i--) {
            if (digits[i] < 9) {
                digits[i]++;
                return digits;
            }
            digits[i] = 0;
        }
        int[] result = new int[n + 1];
        result[0] = 1;
        return result;
    }
}

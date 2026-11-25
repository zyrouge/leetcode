class Solution {
    public boolean isHappy(int n) {
        Set<Integer> visited = new HashSet<>();
        while (!visited.contains(n)) {
            visited.add(n);
            int current = 0;
            while (n > 0) {
                int digit = n % 10;
                current += digit * digit;
                n /= 10;
            }
            n = current;
            if (n == 1) {
                return true;
            }
        }
        return false;
    }
}

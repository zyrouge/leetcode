class Solution {
    public int evalRPN(String[] tokens) {
        Deque<Integer> stack = new ArrayDeque<>();
        for (String x : tokens) {
            if (x.equals("+")) {
                stack.push(stack.pop() + stack.pop());
            } else if (x.equals("-")) {
                int second = stack.pop();
                int first = stack.pop();
                stack.push(first - second);
            } else if (x.equals("*")) {
                stack.push(stack.pop() * stack.pop());
            } else if (x.equals("/")) {
                int second = stack.pop();
                int first = stack.pop();
                stack.push(first / second);
            } else {
                stack.push(Integer.valueOf(x));
            }
        }
        return stack.peek();
    }
}

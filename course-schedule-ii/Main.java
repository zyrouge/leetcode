class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        Map<Integer, List<Integer>> nodes = new HashMap<>();
        for (int[] x : prerequisites) {
            List<Integer> edges = nodes.get(x[0]);
            if (edges == null) {
                edges = new ArrayList<>();
                nodes.put(x[0], edges);
            }
            edges.add(x[1]);
        }
        int[] visited = new int[numCourses];
        int[] order = new int[numCourses];
        int orderIndex = 0;
        for (int i = 0; i < numCourses; i++) {
            orderIndex = dfs(nodes, i, visited, order, orderIndex);
            if (orderIndex == -1) {
                return new int[0];
            }
        }
        return order;
    }

    public int dfs(Map<Integer, List<Integer>> nodes, int value, int[] visited, int[] order, int orderIndex) {
        if (visited[value] == 1) {
            return -1;
        }
        if (visited[value] == 2) {
            return orderIndex;
        }
        List<Integer> edges = nodes.get(value);
        if (edges != null) {
            visited[value] = 1;
            for (int x : edges) {
                orderIndex = dfs(nodes, x, visited, order, orderIndex);
                if (orderIndex == -1) {
                    return -1;
                }
            }
        }
        visited[value] = 2;
        order[orderIndex] = value;
        return orderIndex + 1;
    }
}

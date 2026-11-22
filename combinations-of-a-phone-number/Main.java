class Solution {
    private String[] digitCharacters = {"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        if (digits.isEmpty()) {
            return Collections.emptyList();
        }
        List<String> result = new ArrayList<>();
        dfs(digits, 0, new StringBuilder(), result);
        return result;
    }

    private void dfs(String digits, int i, StringBuilder current, List<String> result) {
        if (i == digits.length()) {
            result.add(current.toString());
            return;
        }
        char[] possibilities = this.digitCharacters[digits.charAt(i) - '2'].toCharArray();
        for (char x : possibilities) {
            current.append(x);
            dfs(digits, i + 1, current, result);
            current.deleteCharAt(current.length() - 1);
        }
    }
}
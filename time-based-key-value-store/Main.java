private class TimeMap {
    public record Entry(String value, int timestamp) {}

    private Map<String, List<Entry>> entries;

    public TimeMap() {
        this.entries = new HashMap<>();
    }
    
    public void set(String key, String value, int timestamp) {
        List<Entry> values = this.entries.get(key);
        if (values == null) {
            values = new ArrayList<>();
            this.entries.put(key, values);
        }
        values.add(new Entry(value, timestamp));
    }
    
    public String get(String key, int timestamp) {
        List<Entry> values = this.entries.get(key);
        if (values == null || values.isEmpty()) {
            return "";
        }
        int n = values.size();
        int i = 0, j = n - 1;
        while (i <= j) {
            int mid = (i + j) / 2;
            if (values.get(mid).timestamp <= timestamp) {
                if (mid == n - 1 || values.get(mid + 1).timestamp > timestamp) {
                    return values.get(mid).value;
                }
                i = mid + 1;
            } else {
                j = mid - 1;
            }
        }
        return "";
    }
}

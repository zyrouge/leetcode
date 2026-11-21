class LRUCache {
    // private int capacity;
    // private LinkedHashMap<Integer, Integer> entries;

    // public LRUCache(int capacity) {
    //     this.capacity = capacity;
    //     this.entries = new LinkedHashMap<>();
    // }
    
    // public int get(int key) {
    //     Integer value = this.entries.remove(key);
    //     if (value == null) {
    //         return -1;
    //     }
    //     this.entries.putLast(key, value);
    //     return value;
    // }
    
    // public void put(int key, int value) {
    //     this.entries.putLast(key, value);
    //     if (this.entries.size() > capacity) {
    //         this.entries.pollFirstEntry();
    //     }
    // }

    private class Node {
        public int key;
        public int value;
        public Node previous;
        public Node next;

        public Node(int key, int value) {
            this.key = key;
            this.value = value;
        }
    }

    private int capacity;
    private Node head;
    private Node tail;
    private Map<Integer, Node> map;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.head = null;
        this.tail = null;
        this.map = new HashMap<>();
    }
    
    public int get(int key) {
        Node node = this.map.get(key);
        if (node == null) {
            return -1;
        }
        this.unlink(node);
        this.link(node);
        return node.value;
    }
    
    public void put(int key, int value) {
        Node node = this.map.get(key);
        if (node != null) {
            this.unlink(node);
        }
        node = new Node(key, value);
        this.map.put(key, node);
        this.link(node);
        if (this.map.size() > capacity && this.head != null) {
            this.map.remove(this.head.key);
            this.head = this.head.next;
            if (this.head != null) {
                this.head.previous = null;
            }
        }
    }

    private void unlink(Node node) {
        if (node.previous != null) {
            node.previous.next = node.next;
        }
        if (node.next != null) {
            node.next.previous = node.previous;
        }
        if (node == this.head) {
            this.head = this.head.next;
            if (this.head != null) {
                this.head.previous = null;
            }
        }
        if (node == this.tail) {
            this.tail = this.tail.previous;
            if (this.tail != null) {
                this.tail.next = null;
            }
        }
    }

    private void link(Node node) {
        if (this.head == null) {
            this.head = node;
        }
        if (this.tail != null) {
            this.tail.next = node;
            node.previous = this.tail;
        }
        this.tail = node;
    }
}

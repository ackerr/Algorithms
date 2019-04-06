class MyLinkedList {
    class Node {
        int val;
        Node next;

        Node(int val) {
            this.val = val;
        }
    }

    private Node head;
    private int size;

    public MyLinkedList() {
        head = null;
        size = 0;
    }
    
    public int get(int index) {
        if (index >= size){
            return -1;
        }
        Node cur = head;
        while (index-- > 0){
            cur = cur.next;
        }
        return cur.val;
    }
    
    public void addAtHead(int val) {
        Node n = new Node(val);
        if (head != null){
            n.next = head;
        }
        head = n;
        size++;
    }

    public void addAtTail(int val) {
        Node n = new Node(val);
        if (head == null){
            head = n;
        }else{
            Node cur = head;
            while (cur.next != null)
                cur = cur.next;
            cur.next = n;
        size++;
        }
    }
    
    public void addAtIndex(int index, int val) {
        if (index > size){
            return;
        }
        if (index == 0){
            addAtHead(val);
        }else{
            Node n = new Node(val);
            Node cur = head;
            while (--index > 0){
                cur = cur.next;
            }
            n.next = cur.next;
            cur.next = n;
            size++;
        }
    }
    
    public void deleteAtIndex(int index) {
        if (index < size){
            Node cur = head;
            while (--index > 0) {
                cur = cur.next;
            }
            cur.next = cur.next.next;
            size--;
        }
    }
}
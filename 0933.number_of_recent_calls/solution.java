class RecentCounter {
    private Queue<Integer> q;

    public RecentCounter() {
        q = new LinkedList<>();
    }

    public int ping(int t) {
        while (!q.isEmpty() && t - 3000 > q.peek()){
            q.poll();
        }
        q.add(t);
        return q.size();
    }
}
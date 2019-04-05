class MyQueue {
    private Stack<Integer> stack;
    private Stack<Integer> stackTemp;

    public MyQueue() {
        stack = new Stack<>();
        stackTemp = new Stack<>();
    }

    public void push(int x) {
        stack.push(x);
    }

    public int pop() {
        while (!this.empty()){
            stackTemp.push(stack.pop());
        }
        int temp = stackTemp.pop();
        while (!stackTemp.isEmpty()){
            stack.push(stackTemp.pop());
        }
        return temp;
    }

    public int peek() {
        while (!this.empty()){
            stackTemp.push(stack.pop());
        }
        int temp = stackTemp.peek();
        while (!stackTemp.isEmpty()){
            stack.push(stackTemp.pop());
        }
        return temp;
    }

    public boolean empty() {
        return stack.isEmpty();
    }
}
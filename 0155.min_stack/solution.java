import java.util.*;

class MinStack {
    /**
     * use a stack
     */

    private int min = Integer.MAX_VALUE;
    private Stack<Integer> stack = new Stack<>();

    public void push(int x) {
        if (min >= x) {
            stack.push(min);
            min = x;
        }
        stack.push(x);
    }

    public void pop() {
        if (stack.pop() == min) {
            min = stack.pop();
        }
    }

    public int top() {
        return stack.peek();
    }

    public int getMin() {
        return min;
    }

}

class OtherSolution {
    /**
     * use a arraylist
     */

    private int min = Integer.MAX_VALUE;
    private ArrayList<Integer> stack = new ArrayList<>();

    public void push(int x) {
        if (min >= x) {
            stack.add(min);
            min = x;
        }
        stack.add(x);
    }

    public void pop() {
        if (stack.remove(stack.size()-1) == min) {
            min = stack.remove(stack.size()-1);
        }
    }

    public int top() {
        return stack.get(stack.size()-1);
    }

    public int getMin() {
        return min;
    }

}

class solution{
    public static void main(String[] args) {
        MinStack stack = new MinStack();
//        OtherSolution stack = new OtherSolution();
        stack.push(3);
        stack.push(2);
        stack.pop();
        stack.push(4);
        System.out.println(stack.top());
        System.out.println(stack.getMin());
    }
}

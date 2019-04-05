import java.util.ArrayList;

class MyStack{
    private ArrayList<Integer> stack;

    public MyStack(){
        stack = new ArrayList<>();
    }

    public void push(int x){
        stack.add(x);
    }

    public int pop(){
        if (!this.empty()){
            return stack.remove(stack.size() - 1);
        }
        return -1;
    }

    public int top(){
        if (!this.empty()){
            return stack.get(stack.size() - 1);
        }
        return -1;
    }

    public boolean empty(){
        return stack.isEmpty();
    }
}
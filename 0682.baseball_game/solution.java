class Solution {
    public int calPoints(String[] ops) {
        Stack<Integer> stack = new Stack<>();
        int score = 0;
        for (String op:ops){
            if (op.equals("C")){
                score -= stack.pop();
            }else if (op.equals("D")){
                stack.push(stack.peek() * 2);
                score += stack.peek();
            }else if (op.equals("+")){
                int pre = stack.pop();
                int value = pre + stack.peek();
                score += value;
                stack.push(pre);
                stack.push(value);
            }else{
                stack.push(Integer.valueOf(op));
                score += stack.peek();
            }
        }
        return score;
    }


    public int otherSolution(String[] ops) {
        int[] nums = new int[ops.length];
        int i=0;
        for(String op:ops){
            if (op.equals("+")){
                nums[i] = nums[i-1] + nums[i-2];
                i++;
            }else if (op.equals("C")){
                nums[i-1] = 0;
                i--;
            }else if (op.equals("D")){
                nums[i] = nums[i-1]*2;
                i++;
            }else{
                nums[i] = Integer.valueOf(op);
                i++;
            }
        }
        int sum=0;
        for (int num: nums) sum+=num;
        return sum;
    }

    public
}
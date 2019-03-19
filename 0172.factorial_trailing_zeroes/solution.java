class Solution {
    public int trailingZeroes(int n) {
        int total = 0;
        while (n != 0){
            n /= 5;
            total += n;
        }
        return total;
    }
}
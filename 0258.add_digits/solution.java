class Solution {
    public int addDigits(int num) {
        if (num < 10) return num;
        num %= 9;
        if (num == 0) return 9;
        return num;
    }
}
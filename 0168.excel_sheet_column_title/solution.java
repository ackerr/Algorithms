class Solution {
    public String convertToTitle(int n) {
        StringBuffer ans = new StringBuffer();
        while (n>0){
            int mod = n % 26;
            n = n / 26;
            if (mod == 0){
                mod = 26;
                n--;
            }
            ans.append((char) mod+64);
        }
        return ans.reverse().toString();
    }
}
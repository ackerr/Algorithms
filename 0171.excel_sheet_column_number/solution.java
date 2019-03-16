class solution {
    public int titleToNumber(String s) {
        int ans = 0;
        for (int i=0;i<s.length();i++) {
            ans = ans * 26 + (int) s.charAt(i) - 64;
        }
        return ans;
    }
}
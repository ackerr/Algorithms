class Solution{
    public boolean isIsomorphic(String s, String t){
        int [] m1 = new int[256];
        int [] m2 = new int[256];
        // 存两个字符串对应索引的值
        for (int i=0;i<s.length();i++){
            if (m1[s.charAt(i)] != m2[t.charAt(i)]){
                return false;
            }
            m1[s.charAt(i)] = m2[s.chatAt(i)] = i+1;
        }
        return true;
    }
}
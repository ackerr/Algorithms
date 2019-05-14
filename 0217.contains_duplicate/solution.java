import java.util.Set;

class Solution {
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> set = new HashSet<Integer>();
		for(int i : nums)
		    if(!set.add(i))  // had same value
			    return true; 
		return false;
    }
}
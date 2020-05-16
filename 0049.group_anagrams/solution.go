package leetcode0049

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		k := strArray(str)
		s, ok := m[k]
		if !ok {
			s = make([]string, 0)
		}
		s = append(s, str)
		m[k] = s
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func strArray(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
		res[v-'a']++
	}
	return res
}

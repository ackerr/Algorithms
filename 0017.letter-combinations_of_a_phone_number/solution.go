package leetcode0017

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ans := []string{""}
	for i := 0; i < len(digits); i++ {
		if digits[i] == '1' {
			continue
		}
		var temp []string
		for j := 0; j < len(ans); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ans[j]+m[digits[i]][k])
			}
		}
		ans = temp
	}
	return ans
}

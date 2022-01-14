package main

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	println(longestCommonPrefix([]string{"a"}))
	println(longestCommonPrefix([]string{"aa", "a"}))
	println(longestCommonPrefix([]string{"aa", "a", "aaa"}))
	println(longestCommonPrefix([]string{"aa", "a", "aaa", "aaaa"}))
	println(longestCommonPrefix([]string{"aa", "a", "aaa", "aaaa", "aaaaa"}))
}

func longestCommonPrefix(strs []string) string {
	pre := ""

	if len(strs) == 1 {
		return strs[0]
	}

	s := strs[0]
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != s[i] {
				return pre
			}
		}
		pre = pre + string(s[i])

	}

	return pre
}
